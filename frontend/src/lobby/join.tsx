import { useEffect, useState } from 'react';
import { APIListGames, ListGameResponse } from '../api/lobby';

const Join = () => {
    const [games, setGames] = useState<ListGameResponse[]>([]);

    useEffect(() => {
        const interval = setInterval(async () => {
            const activeGames = await APIListGames();
            setGames(activeGames);
        }, 1000);

        return () => clearInterval(interval);
    }, []);

    return (
        <div>
            <h2>Join a Game</h2>
            {(games.length == 0) && <label>No active games on the server</label>}
            <ul>
                {games.sort((a, b) => (a.id > b.id ? 1 : -1)).map((game, index) => (
                    <li key={index} style={{marginBottom:"3px"}}>
                        <button style={{padding:"5px 10px", marginInline:"5px"}}>Join</button>
                        <span>GameID = {game.id}, {game.currentPlayers}/{game.maxPlayers} Players</span>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default Join;
