import { BoardColors, PlayerColors } from "./colors";

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
                        border: `2px solid ${(available) ? 'white' : BoardColors[value]}`,
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
                    border: `4px solid ${(selected) ? 'white' : PlayerColors[value]}`,
                    color: PlayerColors[state],
                    backgroundColor: PlayerColors[state],
                }}>
            </div>
        </div>
    );
}

export default BoardTile;