import { useState } from "react";
import { useGlobalState } from "./hooks/globalState";
import { LogicBoard } from "./logic/board";

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
    onClick: () => void;
}

const BoardTile: React.FC<BoardTileProps> = ({ value, state, onClick }: BoardTileProps) => {
    if (value === -1) {
        return <div className={"board-tile"}></div>;
    }
    if (state === 0) {
        return (
            <div className={"board-tile"}>
                <div
                    className={"board-tile-circle"}
                    style={{
                        border: `4px solid ${colors[value]}`,
                        backgroundColor: "transparent",
                    }}>
                </div>
            </div>
        );
    }

    return (
        <div className={"board-tile"}>
            <div
                onClick={() => onClick()}
                className={"board-tile-circle board-tile-circle-active"}
                style={{
                    border: `4px solid ${colors[value]}`,
                    backgroundColor: playerColors[state],
                }}>
            </div>
        </div>
    );

}

type Tile = {
    row: number;
    col: number;
}

const Board = () => {
    const { boardState } = useGlobalState();
    const [selectedTile, setSelectedTile] = useState<Tile | null>(null);

    return (
        <div className="board">
            {LogicBoard.map((row: number[], rowIndex: number) => (
                <div key={rowIndex} className="board-row">
                    {row.map((tile, colIndex) => (
                        <BoardTile
                            key={colIndex}
                            value={tile}
                            state={boardState[rowIndex][colIndex]}
                            onClick={() => setSelectedTile({ row: rowIndex, col: colIndex })} />
                    ))}
                </div>
            ))}
        </div>
    )
}

export default Board;