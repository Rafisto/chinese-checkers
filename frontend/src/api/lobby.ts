import axios from 'axios';

const api_url = "http://localhost:8080"

type CreateGameResponse = {
    message: string,
}

const APICreateGame = async (playerCount: number) => {
    const response = await axios.post(`${api_url}/games`, {
        playerNum: playerCount,
    })

    await new Promise(resolve => setTimeout(resolve, 1000));

    if (response.status !== 201) {
        throw new Error("Unable to create the game.");
    }

    else {
        return response.data as CreateGameResponse;
    }
};

type ListGameResponse = {
    id: number,
    currentPlayers: number,
    maxPlayers: number,
}

const APIListGames = async () => {
    const response = await axios.get(`${api_url}/games`);

    if (response.status !== 200) {
        throw new Error("Unable to list games.");
    }

    else {
        return response.data as ListGameResponse[];
    }
}

export { APICreateGame, APIListGames };
export type { CreateGameResponse, ListGameResponse };