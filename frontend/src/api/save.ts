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
    throw new Error("Unable to create the bot");
  } else {
    return;
  }
};

const APILoadGame = async (api_url: string, gameName: string) => {
  const response = await axios.get(`${api_url}/load/${gameName}`)

  if (response.status !== 201) {
    throw new Error("Unable to create the bot");
  } else {
    return;
  }
};

export { APISaveGame, APILoadGame };
