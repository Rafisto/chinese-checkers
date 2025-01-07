import { useEffect, useState } from 'react';
import { useGlobalState } from '../hooks/useGlobalState';
import { APIJoinGame, APIListGames, ListGameResponse } from '../api/lobby';

type JoinProps = {
    joined: boolean;
    setJoined: (joined: boolean) => void;
}

const Join = ({ joined, setJoined }: JoinProps) => {
    const { serverAddress, lobbyState, setLobbyState, auditLog, setAuditLog } = useGlobalState();
    const [games, setGames] = useState<ListGameResponse[]>([]);

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

    return (
        <div>
            <h2>Join a Game</h2>
            {(games.length == 0) && <label>No active games on the server</label>}
            <ul>
                {games.sort((a, b) => (a.id > b.id ? 1 : -1)).map((game, index) => (
                    <li key={index} style={{ marginBottom: "3px" }}>
                        <button onClick={() => handleJoinGame(game.id, game.variant)} style={{ padding: "5px 10px", marginInline: "5px" }}>Join</button>
                        <span>GameID = {game.id} ({game.variant}), {game.currentPlayers}/{game.maxPlayers} Players</span>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default Join;
