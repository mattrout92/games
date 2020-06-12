import React, { useState, useEffect, Fragment } from "react";
import "./App.css";
import {
  Grid,
  TextField,
  Button,
  makeStyles,
  createStyles,
  Theme,
  IconButton,
} from "@material-ui/core";
import { w3cwebsocket as W3CWebSocket } from "websocket";
import { create, add, startGame, turn } from "./utils/api";
import { Card } from "./Card/component";
import { Table } from "./Table/component";

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

  const createGame = async () => {
    const gameID = await create("5Lives");

    setGameID(gameID);
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
    <Grid
      className={classes.root}
      container
      spacing={0}
      alignItems="center"
      justify="center"
    >
      <Grid spacing={5} container className={classes.actionContainer}>
        {gameID === "" ? (
          <Fragment>
            <Grid item xs={12}>
              <h1>The 5 Lives Game</h1>
            </Grid>
            <Grid item xs={8}>
              <TextField
                value={enteredID}
                onChange={changeEnteredID}
                fullWidth
                variant="outlined"
                id="join"
              />
            </Grid>
            <Grid item xs={4}>
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
            <Grid item xs={12}>
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
            <Grid item xs={12}>
              <h1>The 5 Lives Game</h1>
              <h2>Code: {game.id}</h2>
              {playerName && <h2>Player: {playerName}</h2>}
            </Grid>
            {game.players && (
              <Grid item xs={12}>
                {game.players.map((player: any, index: number) => (
                  <div>
                    {player.cards && (
                      <div>
                        {player.name === playerName && <h2>Your Cards</h2>}
                        {player.cards.map((card: any, subindex: number) => (
                          <Fragment key={subindex}>
                            {player.name === playerName && (
                              <Fragment>
                                {player.is_turn ? (
                                  <IconButton
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
                                  </IconButton>
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
              <Grid item xs={12}>
                <Grid item xs={8}>
                  <TextField
                    value={playerName}
                    onChange={changePlayerName}
                    fullWidth
                    variant="outlined"
                    id="player"
                  />
                </Grid>
                <Grid item xs={4}>
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
            {game.players && !game.started && (
              <Grid item xs={12}>
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
    </Grid>
  );
};

export default App;
