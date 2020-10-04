import React, { useState, useEffect, Fragment } from "react";
import "./App.css";
import {
  Grid,
  TextField,
  Button,
  makeStyles,
  createStyles,
  Theme,
  ThemeProvider,
  CssBaseline,
} from "@material-ui/core";
import { w3cwebsocket as W3CWebSocket } from "websocket";
import { create, add, startGame, turn } from "./utils/api";
import { Card } from "./Card/component";
import { Table } from "./Table/component";
import { theme } from "./theme";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: { flexGrow: 2 },
    actionContainer: {
      paddingLeft: "25%",
      paddingRight: "25%",
      paddingTop: "10%",
    },
    button: {
      height: theme.spacing(7),
    },
  })
);

const App = (): JSX.Element => {
  const classes = useStyles();

  const [gameID, setGameID] = useState<string>("");
  const [enteredID, setEnteredID] = useState<string>("");
  const [game, setGame] = useState<any>({});
  const [playerName, setPlayerName] = useState<string>("");
  const [joined, setJoined] = useState<boolean>(false);
  const [creator, setCreator] = useState<boolean>(false);

  const createGame = async () => {
    const gameID = await create("5Lives");

    setGameID(gameID);
    setCreator(true);
  };

  const start = async () => {
    await startGame("5Lives", gameID);
  };

  const changeEnteredID = (event: any) => {
    setEnteredID(event.target.value);
  };

  const join = () => {
    setGameID(enteredID);
  };

  const changePlayerName = (event: any) => {
    setPlayerName(event.target.value);
  };

  const addPlayer = async () => {
    await add("5Lives", gameID, playerName);
    setJoined(true);
  };

  const cardClicked = async (cardDescriptor: string) => {
    await turn("5Lives", gameID, cardDescriptor);
  };

  useEffect(() => {
    (async () => {
      if (gameID !== "") {
        const client = new W3CWebSocket(
          `ws://localhost:8080/games/5Lives/instances/${gameID}`
        );

        client.onmessage = (message) => {
          if (typeof message.data === "string") {
            console.log(JSON.parse(message.data));
            setGame(JSON.parse(message.data));
          }
        };
      }
    })();
  }, [gameID]);

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Grid
        className={classes.root}
        container
        spacing={0}
        alignItems="center"
        justify="center"
      >
        <Grid spacing={5} container>
          <Grid item xs={false} md={4}></Grid>
          <Grid item xs={12} md={4} container spacing={5}>
            {gameID === "" ? (
              <Fragment>
                <Grid item md={3}></Grid>
                <Grid container justify="center">
                  <h1>The 5 Lives Game</h1>
                </Grid>
                <Grid item md={3}></Grid>
                <Grid container spacing={5}>
                  <Grid item md={8}>
                    <TextField
                      value={enteredID}
                      onChange={changeEnteredID}
                      fullWidth
                      variant="outlined"
                      id="join"
                      helperText="Enter a code to join an existing game"
                    />
                  </Grid>
                  <Grid item md={4}>
                    <Button
                      className={classes.button}
                      fullWidth
                      color="primary"
                      variant="contained"
                      onClick={join}
                    >
                      Join
                    </Button>
                  </Grid>
                </Grid>
                <Grid item md={12}>
                  <Button
                    className={classes.button}
                    fullWidth
                    color="secondary"
                    variant="contained"
                    onClick={createGame}
                  >
                    Or Create a game!
                  </Button>
                </Grid>
              </Fragment>
            ) : (
              <Fragment>
                <Grid item md={12}>
                  <h1>The 5 Lives Game</h1>
                  <h2>Code: {game.id}</h2>
                  {playerName && <h2>Player: {playerName}</h2>}
                </Grid>
                {game.game_finished && (
                  <Grid item md={12}>
                    {game.players.map((player: any, index: number) => (
                      <Fragment key={index}>
                        {player.lives > 0 && (
                          <Fragment>
                            <h2>The winner is: {player.name}</h2>
                            <Button
                              className={classes.button}
                              fullWidth
                              color="secondary"
                              variant="contained"
                              onClick={start}
                            >
                              Play again
                            </Button>
                          </Fragment>
                        )}
                      </Fragment>
                    ))}
                  </Grid>
                )}
                {game.players && (
                  <Grid item md={12}>
                    {game.players.map((player: any, index: number) => (
                      <div key={index}>
                        {player.cards && (
                          <div>
                            {player.name === playerName &&
                              game.last_player_name && (
                                <Fragment>
                                  <h2>{game.last_player_name} just played: </h2>
                                  <Card
                                    descriptor={game.previous_card.descriptor}
                                    height={200}
                                    width={150}
                                  />
                                </Fragment>
                              )}
                            {player.name === playerName && <h2>Your Cards</h2>}
                            {player.name === playerName && player.is_turn && (
                              <Fragment>
                                <h2>It's your turn!</h2>
                                <h3>Click on a card to play</h3>
                              </Fragment>
                            )}
                            {player.cards.map((card: any, subindex: number) => (
                              <Fragment key={subindex}>
                                {player.name === playerName && (
                                  <Fragment>
                                    {player.is_turn ? (
                                      <Fragment>
                                        <Button
                                          onClick={() => {
                                            cardClicked(card.descriptor);
                                          }}
                                        >
                                          <Card
                                            key={subindex}
                                            descriptor={card.descriptor}
                                            height={200}
                                            width={150}
                                          />
                                        </Button>
                                      </Fragment>
                                    ) : (
                                      <Card
                                        key={subindex}
                                        descriptor={card.descriptor}
                                        height={200}
                                        width={150}
                                      />
                                    )}
                                  </Fragment>
                                )}
                              </Fragment>
                            ))}
                          </div>
                        )}
                      </div>
                    ))}
                    <h2>Table</h2>
                    <Table players={game.players} />
                  </Grid>
                )}
                {!joined && (
                  <Grid container spacing={2}>
                    <Grid item md={8}>
                      <TextField
                        value={playerName}
                        onChange={changePlayerName}
                        fullWidth
                        variant="outlined"
                        id="player"
                      />
                    </Grid>
                    <Grid item md={4}>
                      <Button
                        className={classes.button}
                        fullWidth
                        color="secondary"
                        variant="contained"
                        onClick={addPlayer}
                      >
                        Set player name!
                      </Button>
                    </Grid>
                  </Grid>
                )}
                {game.players && !game.started && creator && (
                  <Grid item md={12}>
                    <Button
                      className={classes.button}
                      fullWidth
                      color="secondary"
                      variant="contained"
                      onClick={start}
                    >
                      Start game!
                    </Button>
                  </Grid>
                )}
              </Fragment>
            )}
          </Grid>
          <Grid item xs={false} md={4}></Grid>
        </Grid>
      </Grid>
    </ThemeProvider>
  );
};

export default App;
