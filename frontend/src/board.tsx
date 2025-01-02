import { LogicBoard } from "./logic/board";

const colors: Record<string, string> = {
    '-1': 'transparent',
    '0': 'gray',
    '1': 'red',
    '2': 'blue',
    '3': 'green',
    '4': 'cyan',
    '5': 'magenta',
    '6': 'yellow'
}

interface BoardTileProps {
    value: number;
}

const BoardTile: React.FC<BoardTileProps> = ({ value }: BoardTileProps) => {
    return (
        <div className={"board-tile"}>
            <div 
            className={"board-tile-circle"} 
            style={{ 
                border: `4px solid ${colors[value]}`,
                backgroundColor: 'transparent',
             }}>
            </div>
        </div>
    );

}

const Board = () => {
    return (
        <div className="board">
            {LogicBoard.map((row: number[], rowIndex: number) => (
                <div key={rowIndex} className="board-row">
                    {row.map((tile, colIndex) => (
                        <BoardTile key={colIndex} value={tile} />
                    ))}
                </div>
            ))}
        </div>
    )
}

export default Board;