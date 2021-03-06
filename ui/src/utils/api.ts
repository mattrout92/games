import Axios from "axios";

const apiURL = "https://api.mattgames.co.uk";

export const create = async (gameName: string): Promise<string> => {
  const response = await Axios.post(`${apiURL}/games/${gameName}`);

  if (response.data) {
    return response.data.id;
  }

  return "";
};

export const add = async (
  gameName: string,
  gameID: string,
  playerName: string
): Promise<boolean> => {
  const response = await Axios.post(
    `${apiURL}/games/${gameName}/instances/${gameID}/players/${playerName}`
  );

  if (response.data) {
    return true;
  }

  return false;
};

export const startGame = async (
  gameName: string,
  gameID: string
): Promise<boolean> => {
  const response = await Axios.post(
    `${apiURL}/games/${gameName}/instances/${gameID}`
  );

  if (response.data) {
    return true;
  }

  return false;
};

export const turn = async (
  gameName: string,
  gameID: string,
  turn: string,
  stick: boolean
): Promise<boolean> => {
  const response = await Axios.post(
    `${apiURL}/games/${gameName}/instances/${gameID}/turn/${turn}?stick=${stick}`
  );

  if (response.data) {
    return true;
  }

  return false;
};

export const nextRound = async (
  gameName: string,
  gameID: string
): Promise<boolean> => {
  const response = await Axios.post(
    `${apiURL}/games/${gameName}/instances/${gameID}/next-round`
  );

  if (response.data) {
    return true;
  }

  return false;
};
