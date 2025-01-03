import axios from 'axios';

type BasicResponse = {
    message: string,
}

const APICreateGame = async (api_url: string, playerCount: number) => {
    console.log("Creating game with player count: ", playerCount);

    console.log("Querying: ", `${api_url}/games`);
    const response = await axios.post(`${api_url}/games`, {
        playerNum: playerCount,
    })

    if (response.status !== 201) {
        console.log(response.data);
        throw new Error("Unable to create the game.");
    }
    else {
        console.log(response.data);
        return response.data as BasicResponse;
    }
};

type ListGameResponse = {
    id: number,
    currentPlayers: number,
    maxPlayers: number,
}

const APIListGames = async (api_url: string) => {
    const response = await axios.get(`${api_url}/games`);

    if (response.status !== 200) {
        throw new Error("Unable to list games.");
    }
    else {
        return response.data as ListGameResponse[];
    }
}

type JoinGameResponse = {
    message: string,
    id: number,
}


const APIJoinGame = async (api_url: string, gameID: number, playerName: string) => {
    const response = await axios.post(`${api_url}/games/${gameID}/join`, {
        username: playerName
    })

    if (response.status !== 201) {
        throw new Error("Unable to join the game.");
    }
    else {
        return response.data as JoinGameResponse;
    }
}

export { APICreateGame, APIListGames, APIJoinGame };
export type { BasicResponse, ListGameResponse };