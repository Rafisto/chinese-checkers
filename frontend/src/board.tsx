import { useEffect, useState } from "react";
import { useGlobalState } from "./hooks/globalState";
import { FindMoves, PerformMove } from "./logic/state";
import AuditLog from "./components/auditLog";
import BoardTile from "./boardtile";

// Type definitions
type Point = {
    row: number;
    col: number;
};

const Board = () => {
    // Global and local state
    const { gameState, setGameState, lobbyState, ws, setAuditLog } = useGlobalState();
    const [selectedTile, setSelectedTile] = useState<Point | null>(null);
    const [availableMoves, setAvailableMoves] = useState<Point[]>([]);

    // Fetch available moves when a tile is selected
    useEffect(() => {
        if (!selectedTile) {
            setAvailableMoves([]);
            return;
        }
        const passMap = Array.from({ length: gameState.state.length }, () =>
            Array(gameState.state[0].length).fill(false)
        );
        const moves = FindMoves(gameState.state, selectedTile, passMap, []);
        setAvailableMoves(moves);
    }, [selectedTile, gameState.state]);

    // Handle WebSocket messages
    useEffect(() => {
        if (ws) {
            ws.onopen = () => {
                console.log("WebSocket connection established.");
                sendBoardRequest();
                sendPawnsRequest();
                sendStateRequest();
            };

            ws.onmessage = (event) => {
                const data = JSON.parse(event.data);
                const message = JSON.parse(data.message);
                console.log(message);
                setAuditLog((prevAuditLog) => [...prevAuditLog, `Received message: ${data}`]);
                if (message.type === "server" && message.board !== undefined) {
                    if (message.board == null || message.board.length === 0) {
                        console.log("Received empty board state.");
                    }
                    setGameState((prevGameState) => ({ ...prevGameState, board: message.board }));
                }

                if (message.type === "server" && message.pawns !== undefined) {
                    if (message.pawns == null || message.pawns.length === 0) {
                        console.log("Received empty pawns state.");
                    }
                    setGameState((prevGameState) => ({ ...prevGameState, state: message.pawns }));
                }

                if (message.type === "server" && message.action === "state") {
                    setGameState((prevGameState) => ({ ...prevGameState, players: message.players, turn: message.turn, current: message.current, color: message.color }));
                };

                if (message.type === "server" && message.action === "move" && message.player_id !== lobbyState.playerID) {
                    setGameState((prevGameState) => ({ ...prevGameState, state: PerformMove(prevGameState.state, message.start, message.end) }));
                }
            }
        };
    }, [ws]);

    // Send state request to server
    const sendStateRequest = () => {
        if (!ws || ws.readyState !== WebSocket.OPEN) {
            console.warn("WebSocket is not ready. Cannot send state request.");
            return;
        }

        ws.send(
            JSON.stringify({
                type: "player",
                action: "state",
            })
        );
    }

    // Send board request to server
    const sendBoardRequest = () => {
        if (!ws || ws.readyState !== WebSocket.OPEN) {
            console.warn("WebSocket is not ready. Cannot send board request.");
            return;
        }

        ws.send(
            JSON.stringify({
                type: "player",
                action: "board",
            })
        );
    };

    // Send pawn request to server
    const sendPawnsRequest = () => {
        if (!ws) return;

        ws.send(
            JSON.stringify({
                type: "player",
                action: "pawns",
            })
        );
    };

    // Send move request to server
    const sendNewMove = (playerID: number, start: Point, end: Point) => {
        if (!ws) return;

        ws.send(
            JSON.stringify({
                type: "player",
                action: "move",
                player_id: playerID,
                start,
                end,
            })
        );
    };

    // Handle move attempt
    const handleTryMovePiece = (end: Point) => {
        if (!selectedTile || !availableMoves.some((move) => move.row === end.row && move.col === end.col)) {
            setSelectedTile(null);
            setAvailableMoves([]);
            return;
        }

        // Perform local move
        // setGameState({ ...gameState, state: PerformMove(gameState.state, selectedTile, end) });

        // Notify server of the move
        sendNewMove(lobbyState.playerID, selectedTile, end);
        setSelectedTile(null);
        setAvailableMoves([]);
    };

    // Render board UI
    return (
        <div className="board">
            <AuditLog />
            {gameState.board.map((row: number[], rowIndex: number) => (
                <div key={rowIndex} className="board-row">
                    {row.map((tile, colIndex) => {
                        const isSelected = selectedTile?.row === rowIndex && selectedTile?.col === colIndex;
                        const isAvailable = availableMoves.some((move) => move.row === rowIndex && move.col === colIndex);

                        return (
                            <BoardTile
                                key={colIndex}
                                value={tile}
                                state={gameState.state[rowIndex][colIndex]}
                                selectPiece={() => setSelectedTile({ row: rowIndex, col: colIndex })}
                                tryMovePiece={() => handleTryMovePiece({ row: rowIndex, col: colIndex })}
                                selected={isSelected}
                                available={isAvailable}
                            />
                        );
                    })}
                </div>
            ))}
        </div>
    );
};

export default Board;
