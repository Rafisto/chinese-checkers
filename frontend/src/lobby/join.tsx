import { useEffect, useState } from 'react';
import { useGlobalState } from '../hooks/globalState';
import { APIJoinGame, APIListGames, ListGameResponse } from '../api/lobby';

type JoinProps = {
    setJoined: (joined: boolean) => void;
}

const Join = ({setJoined}: JoinProps) => {
    const { serverAddress, playerName, setPlayerID, setGameID, auditLog, setAuditLog } = useGlobalState();
    const [games, setGames] = useState<ListGameResponse[]>([]);

    useEffect(() => {
        const interval = setInterval(async () => {
            const activeGames = await APIListGames(serverAddress);
            setGames(activeGames);
        }, 1000);

        return () => clearInterval(interval);
    }, []);

    const handleJoinGame = async (gameID: number) => {
        try {
            const resp = await APIJoinGame(serverAddress, gameID, playerName);
            setPlayerID(resp.id);
            setGameID(gameID);
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
                        <button onClick={()=>handleJoinGame(game.id)} style={{ padding: "5px 10px", marginInline: "5px" }}>Join</button>
                        <span>GameID = {game.id}, {game.currentPlayers}/{game.maxPlayers} Players</span>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default Join;
