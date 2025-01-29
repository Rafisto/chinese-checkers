import axios from "axios";

const APISaveGame = async (
  api_url: string,
  gameID: number,
  gameName: string
) => {
  const response = await axios.post(`${api_url}/games/${gameID}/save`, {
    name: gameName,
  });

  if (response.status !== 201) {
    throw new Error("Unable to save the game");
  } else {
    return;
  }
};

const APILoadGame = async (api_url: string, gameName: string) => {
  const response = await axios.get(`${api_url}/load/${gameName}`)

  if (response.status !== 200) {
    throw new Error("Unable to load the game");
  } else {
    return;
  }
};

export { APISaveGame, APILoadGame };
