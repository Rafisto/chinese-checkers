import { useGlobalState } from "../hooks/globalState";
import { PlayerColors } from "../colors";

const Stats = () => {
    const { gameState } = useGlobalState();

    return (
        <div>
            <h2>Game Stats</h2>
            <p>All Players {gameState.players.map((player) => player).join(", ")}</p>
            <div style={{ display: "flex", flexDirection: "row", alignItems: "center" }}>
                <div style={{ marginRight: "10px" }}>
                    My Player
                </div>
                <div
                    className={"board-tile-circle board-tile-circle-active"}
                    style={{
                        border: `4px solid ${PlayerColors[gameState.color + 1]}`,
                        color: `${PlayerColors[gameState.color + 1]}`,
                        backgroundColor: `${PlayerColors[gameState.color + 1]}`,
                    }}>
                    {gameState.color}
                </div>
            </div>
            <div style={{ display: "flex", flexDirection: "row", alignItems: "center" }}>
                <div style={{ marginRight: "10px" }}>
                    Player to Move
                </div>
                <div
                    className={"board-tile-circle board-tile-circle-active"}
                    style={{
                        border: `4px solid ${PlayerColors[gameState.turn % gameState.players.length + 1]}`,
                        color: `${PlayerColors[gameState.turn % gameState.players.length + 1]}`,
                        backgroundColor: `${PlayerColors[gameState.turn % gameState.players.length + 1]}`,
                    }}>
                </div>
            </div>
        </div>
    );
}

export default Stats;