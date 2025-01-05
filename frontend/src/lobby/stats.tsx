import { useGlobalState } from "../hooks/globalState";
import { PlayerColors, ThreePlayerColors } from "../colors";

const Stats = () => {
    const { gameState, lobbyState, ws, setAuditLog } = useGlobalState();

    const getPlayerColor = (playerID: number) => {
        if (gameState.players.length == 3) {
            return ThreePlayerColors[playerID + 1];
        }
        return PlayerColors[playerID + 1];
    }

    const getPlayerTurnColor = (playerID: number) => {
        if (gameState.players.length == 3) {
            return ThreePlayerColors[playerID];
        }
        return PlayerColors[playerID];
    }

    const handleSkipTurn = () => {
        if (!ws) return;

        const request = JSON.stringify({
            type: "player",
            action: "move",
            player_id: lobbyState.playerID,
        });

        setAuditLog((prevAuditLog) => [...prevAuditLog, `TX ${request}`]);
        ws.send(request);
    };

    return (
        <div>
            <h2>Game Stats ({lobbyState.gameVariant})</h2>
            <p>All Players {gameState.players.map((player) => player).join(", ")}</p>
            <hr/>
            <div style={{ display: "flex", flexDirection: "row", alignItems: "center" }}>
                <div style={{ marginRight: "10px" }}>
                    My Player
                </div>
                <div
                    className={"board-tile-circle board-tile-circle-active"}
                    style={{
                        border: `4px solid ${getPlayerColor(gameState.color)}`,
                        color: `${getPlayerColor(gameState.color)}`,
                        backgroundColor: `${getPlayerColor(gameState.color)}`,
                    }}>
                    {gameState.color}
                </div>
            </div>
            <hr />
            <div style={{ display: "flex", flexDirection: "row", alignItems: "center" }}>
                <div style={{ marginRight: "10px" }}>
                    Player to Move
                </div>
                <div
                    className={"board-tile-circle board-tile-circle-active"}
                    style={{
                        border: `4px solid ${getPlayerTurnColor(gameState.turn % gameState.players.length + 1)}`,
                        color: `${getPlayerTurnColor(gameState.turn % gameState.players.length + 1)}`,
                        backgroundColor: `${getPlayerTurnColor(gameState.turn % gameState.players.length + 1)}`,
                    }}>
                </div>
            </div>
            <hr />
            <div style={{ display: "flex", flexDirection: "row", alignItems: "center" }}>
                <button className={"wide"} onClick={handleSkipTurn}>Skip Turn</button>
            </div>
        </div>
    );
}

export default Stats;