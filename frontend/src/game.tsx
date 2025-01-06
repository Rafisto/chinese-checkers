import { useEffect, useState } from "react";
import { FindMoves } from "./logic/state";
import { Point } from "./logic/types";
import { handleWebSocketMessages, sendNewMove } from "./api/websocket";
import { useGlobalState } from "./hooks/useGlobalState";
import AuditLog from "./components/auditLog";
import Tile from "./tile";

const Game = () => {
    const { gameState, setGameState, lobbyState, ws, setAuditLog } = useGlobalState();

    const [selectedTile, setSelectedTile] = useState<Point | null>(null);
    const [availableMoves, setAvailableMoves] = useState<Point[]>([]);

    // Handle WebSocket messages
    useEffect(() => {
        if (ws) {
            handleWebSocketMessages(ws, setAuditLog, setGameState);
        }
    }, [ws]);
    

    // Fetch available moves when a tile is selected
    useEffect(() => {
        if (!selectedTile) {
            setAvailableMoves([]);
            return;
        }
        const passMap = Array.from({ length: gameState.state.length }, () =>
            Array(gameState.state[0].length).fill(false)
        );

        let moves: Point[] = [];

        if (lobbyState.gameVariant == "classic") {
            moves = FindMoves(gameState.board, gameState.state, selectedTile, passMap, [], true);
        }
        else {
            moves = FindMoves(gameState.board, gameState.state, selectedTile, passMap, [], false);
        }

        setAvailableMoves(moves);

    }, [selectedTile, gameState.state]);

    // Handle try select piece (only if owned by the player)
    const handleTrySelectPiece = (row: number, col: number) => {
        // this is not working properly for player id renumeration in 3-player game
        // if (gameState.state[row][col] !== gameState.color + 1) {
        //     console.log(`Cannot select piece that does not belong to the player. (${row}, ${col}) is not ${gameState.color + 1}`);
        //     return;
        // }

        setSelectedTile({ row, col });
    }

    // Handle move attempt
    const handleTryMovePiece = (end: Point) => {
        if (!selectedTile || !availableMoves.some((move) => move.row === end.row && move.col === end.col)) {
            setSelectedTile(null);
            setAvailableMoves([]);
            return;
        }

        // Perform local move (optimistic update)
        // setGameState({ ...gameState, state: PerformMove(gameState.state, selectedTile, end) });

        // Notify server of the move (await server acknowledgement)
        sendNewMove(ws, setAuditLog, lobbyState.playerID, selectedTile, end);
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
                            <Tile
                                key={colIndex}
                                value={tile}
                                state={gameState.state[rowIndex][colIndex]}
                                selectPiece={() => handleTrySelectPiece(rowIndex, colIndex)}
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

export default Game;
