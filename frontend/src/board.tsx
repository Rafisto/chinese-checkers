import { useEffect, useState } from "react";
import { useGlobalState } from "./hooks/globalState";
import { LogicBoard } from "./logic/board";
import { FindMoves, PerformMove } from "./logic/state";
import AuditLog from "./components/auditLog";

const colors: Record<string, string> = {
    '-1': 'transparent',
    '0': 'gray',
    '1': 'red',
    '2': 'blue',
    '3': 'lime',
    '4': 'cyan',
    '5': 'magenta',
    '6': 'yellow'
}

const playerColors: Record<string, string> = {
    '0': 'transparent',
    '1': 'red',
    '2': 'blue',
    '3': 'lime',
    '4': 'cyan',
    '5': 'magenta',
    '6': 'yellow'
}

interface BoardTileProps {
    value: number;
    state: number;
    selectPiece: () => void;
    tryMovePiece: () => void;
    selected: boolean;
    available?: boolean;
}

const BoardTile: React.FC<BoardTileProps> = ({ value, state, selectPiece, tryMovePiece, selected, available }: BoardTileProps) => {
    if (value === -1) {
        return <div className={"board-tile"}></div>;
    }
    if (state === 0) {
        return (
            <div className={"board-tile"}>
                <div
                    onClick={() => tryMovePiece()}
                    className={"board-tile-circle"}
                    style={{
                        border: `2px solid ${(available) ? 'white' : colors[value]}`,
                        backgroundColor: 'transparent',
                    }}>
                </div>
            </div>
        );
    }
    return (
        <div className={"board-tile"}>
            <div
                onClick={() => selectPiece()}
                className={"board-tile-circle board-tile-circle-active"}
                style={{
                    border: `4px solid ${(selected) ? 'white' : playerColors[value]}`,
                    color: playerColors[state],
                    backgroundColor: playerColors[state],
                }}>
            </div>
        </div>
    );
}

type Point = {
    row: number;
    col: number;
}

const Board = () => {
    const { boardState, setBoardState, playerID, ws, setAuditLog } = useGlobalState();
    const [selectedTile, setSelectedTile] = useState<Point | null>(null);
    const [availableMoves, setAvailableMoves] = useState<Point[]>([]);

    useEffect(() => {
        if (selectedTile) {
            const passMap = Array.from({ length: boardState.length }, () => Array(boardState[0].length).fill(false));
            const moves = FindMoves(boardState, { row: selectedTile.row, col: selectedTile.col }, passMap, []);
            setAvailableMoves(moves);
        }
    }, [selectedTile]);

    useEffect(() => {
        if (ws) {
            ws.onmessage = (event) => {
                const data = JSON.parse(event.data);
                const message = JSON.parse(data.message);
                console.log(message);
                setAuditLog((prevAuditLog: string[]) => [...prevAuditLog, `Received message: ${data}`]);
                if (message.type === "player" && message.action === "move" && message.player_id !== playerID) {
                    setBoardState((bs) => PerformMove(bs, message.start, message.end));
                }
            }
        }
    }, [ws]);

    const handleTryMovePiece = (end: Point) => {
        if (selectedTile == undefined) {
            return;
        }
        if (availableMoves.length === 0) {
            return;
        }
        if (!availableMoves.some((move) => move.row === end.row && move.col === end.col)) {
            setSelectedTile(null);
            setAvailableMoves([]);
            return;
        }
        setBoardState(PerformMove(boardState, { row: selectedTile.row, col: selectedTile.col }, end));
        sendNewMove(playerID, { row: selectedTile.row, col: selectedTile.col }, end);
        setSelectedTile(null);
        setAvailableMoves([]);
    }

    const sendNewMove = (playerID: number, start: Point, end: Point) => {
        if (!ws) {
            return;
        }

        ws.send(JSON.stringify(
            {
                "type": "player",
                "action": "move",
                "player_id": playerID,
                "start": {
                    "row": start.row,
                    "col": start.col,
                },
                "end": {
                    "row": end.row,
                    "col": end.col,
                }
            }
        ));
    }

    return (
        <div className="board">
            <AuditLog />
            {LogicBoard.map((row: number[], rowIndex: number) => (
                <div key={rowIndex} className="board-row">
                    {row.map((tile, colIndex) => (
                        <BoardTile
                            key={colIndex}
                            value={tile}
                            state={boardState[rowIndex][colIndex]}
                            selectPiece={() => setSelectedTile({ row: rowIndex, col: colIndex })}
                            tryMovePiece={() => handleTryMovePiece(
                                { row: rowIndex, col: colIndex })}
                            selected={selectedTile?.row === rowIndex && selectedTile?.col === colIndex}
                            available={(availableMoves.length > 0) && availableMoves.some((move) => move.row === rowIndex && move.col === colIndex)}
                        />
                    ))}
                </div>
            ))}
        </div>
    )
}

export default Board;