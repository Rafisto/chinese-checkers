import { useEffect, useState } from 'react';
import { useGlobalState } from '../hooks/useGlobalState';
import { APIJoinGame, APIAddBot, APIListGames, ListGameResponse } from '../api/lobby';
import { APILoadGame } from '../api/save';

type JoinProps = {
    joined: boolean;
    setJoined: (joined: boolean) => void;
}

const Join = ({ joined, setJoined }: JoinProps) => {
    const { serverAddress, lobbyState, setLobbyState, auditLog, setAuditLog } = useGlobalState();
    const [games, setGames] = useState<ListGameResponse[]>([]);
    const [loadName, setLoadName] = useState<string>("");
    const [successLoad, setSuccessLoad] = useState<boolean>(false);

    useEffect(() => {
        // polling via REST
        const interval = setInterval(async () => {
            if (!joined) {
                const activeGames = await APIListGames(serverAddress);
                setGames(activeGames);
            }
        }, 1000);

        return () => clearInterval(interval);
    }, [joined, serverAddress]);

    const handleJoinGame = async (gameID: number, gameVariant: string) => {
        try {
            const resp = await APIJoinGame(serverAddress, gameID, lobbyState.playerName);
            setLobbyState({ ...lobbyState, playerID: resp.id, gameID: gameID, gameVariant: gameVariant });
            setJoined(true);
            setAuditLog([...auditLog, `Joined game ${gameID}, my playerID is ${resp.id}`]);
        }
        catch (error) {
            console.error(error);
            setJoined(false);
            setAuditLog([...auditLog, `Failed to join game ${gameID}`]);
        }
    }

    const handleAddBot = async(gameID: number) => {
        try {
            await APIAddBot(serverAddress, gameID);
            setAuditLog([...auditLog, `Added bot to game ${gameID}`])
        }
        catch (error) {
            setAuditLog([...auditLog, `Failed to add a bot to game ${gameID}, ${error}`]);
        }
    }

    const handleLoadGame = async () => {
        try {
            setSuccessLoad(false);
            await APILoadGame(serverAddress, loadName);
            setAuditLog([...auditLog, `Loaded game ${loadName}`]);
            setSuccessLoad(true);
        }
        catch (error) {
            console.error(error);
            setAuditLog([...auditLog, `Failed to load game ${loadName}`]);
            setSuccessLoad(false);
        }
    }

    return (
        <div>
            <h2>Load a Game</h2>
            <p>Enter game name</p>
            <input type="text" value={loadName} onChange={(e) => setLoadName(e.currentTarget.value)}></input>
            <button onClick={() => handleLoadGame()}>Load</button>
            {successLoad ? <p color="lime">Game loaded successfully!</p> : null}
            <h2>Join a Game</h2>
            {(games.length == 0) && <label>No active games on the server</label>}
            <ul>
                {games.sort((a, b) => (a.id > b.id ? 1 : -1)).map((game, index) => (
                    <li key={index} style={{ marginBottom: "3px" }}>
                        <button onClick={() => handleAddBot(game.id)}  style={{ padding: "5px 10px", marginInline: "5px" }}>Add Bot</button>
                        <button onClick={() => handleJoinGame(game.id, game.variant)} style={{ padding: "5px 10px", marginInline: "5px" }}>Join</button>
                        <span>GameID = {game.id} ({game.variant}), {game.currentPlayers}/{game.maxPlayers} Players</span>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default Join;
