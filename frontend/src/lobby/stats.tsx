import { useState } from "react";
import { handleSkipTurn } from "../api/websocket";
import { useGlobalState } from "../hooks/useGlobalState";
import { getPlayerColor, getPlayerTurnColor } from "../logic/colors";
import { APISaveGame } from "../api/save";

const Stats = () => {
    const { serverAddress, gameState, lobbyState, ws, auditLog, setAuditLog } = useGlobalState();
    const [saveName, setSaveName] = useState<string>("");

    const handleSaveGame = async () => {
        try {
            await APISaveGame(serverAddress, lobbyState.gameID, saveName);
            setAuditLog([...auditLog, `Saved game ${lobbyState.gameID} as ${saveName}`]);
        }
        catch (error) {
            console.error(error);
            setAuditLog([...auditLog, `Failed to saved game ${saveName}`]);
        }
    }

    return (
        <div>
            <h2>Game Stats ({lobbyState.gameVariant})</h2>
            <p>All Players {gameState.players.map((player) => player).join(", ")}</p>
            <hr />
            <div style={{ display: "flex", flexDirection: "row", alignItems: "center" }}>
                <div style={{ marginRight: "10px" }}>
                    My Player
                </div>
                <div
                    className={"board-tile-circle board-tile-circle-active"}
                    style={{
                        border: `4px solid ${getPlayerColor(gameState.color, gameState.players.length)}`,
                        color: `${getPlayerColor(gameState.color, gameState.players.length)}`,
                        backgroundColor: `${getPlayerColor(gameState.color, gameState.players.length)}`,
                    }}>
                    {gameState.color}
                </div>
            </div>
            <p>Enter game name</p>
            <input type="text" value={saveName} onChange={(e) => setSaveName(e.currentTarget.value)}></input>
            <button onClick={() => handleSaveGame()}>Save as '{saveName}'</button>
            <hr />
            {!gameState.ended ?
                <>
                    <div style={{ display: "flex", flexDirection: "row", alignItems: "center" }}>
                        <div style={{ marginRight: "10px" }}>
                            Player to Move
                        </div>
                        <div
                            className={"board-tile-circle board-tile-circle-active"}
                            style={{
                                border: `4px solid ${getPlayerTurnColor(gameState.turn % gameState.players.length + 1, gameState.players.length)}`,
                                color: `${getPlayerTurnColor(gameState.turn % gameState.players.length + 1, gameState.players.length)}`,
                                backgroundColor: `${getPlayerTurnColor(gameState.turn % gameState.players.length + 1, gameState.players.length)}`,
                            }}>
                        </div>
                    </div>
                    <hr />
                    <div style={{ display: "flex", flexDirection: "row", alignItems: "center" }}>
                        <button className={"wide"} onClick={() => handleSkipTurn(ws, setAuditLog, lobbyState.playerID)}>Skip Turn</button>
                    </div>
                </>
                :
                <>
                    <div style={{ display: "flex", flexDirection: "row", alignItems: "center" }}>
                        <div style={{ marginRight: "10px" }}>
                            Game Over. Player
                        </div>
                        <div
                            className={"board-tile-circle board-tile-circle-active"}
                            style={{
                                border: `4px solid ${getPlayerTurnColor(gameState.turn % gameState.players.length + 1, gameState.players.length)}`,
                                color: `${getPlayerTurnColor(gameState.turn % gameState.players.length + 1, gameState.players.length)}`,
                                backgroundColor: `${getPlayerTurnColor(gameState.turn % gameState.players.length + 1, gameState.players.length)}`,
                            }}>
                        </div>
                        <div style={{ marginLeft: "10px" }}>
                            Wins!
                        </div>
                    </div>
                </>
            }
        </div>
    );
}

export default Stats;