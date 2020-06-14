import React from "react";
import { Grid } from "@material-ui/core";
import { Card } from "../Card/component";

export type Props = {
  players: any;
};

export const Table = ({ players }: Props): JSX.Element => {
  return (
    <Grid container spacing={2}>
      {/* <img  src={TableIMG} alt="table" /> */}
      <Grid item md={3}>
        {players.length > 0 && (
          <div>
            <p>{players[0].name}</p>
            <p>Lives: {players[0].lives}</p>
            {players[0].is_turn && <p>Currently Playing</p>}
            {players[0].last_card_played && (
              <Card
                descriptor={players[0].last_card_played.descriptor}
                height={80}
                width={60}
              />
            )}
          </div>
        )}
      </Grid>
      <Grid item md={3}>
        {players.length > 1 && (
          <div>
            <p>{players[1].name}</p>
            <p>Lives: {players[1].lives}</p>
            {players[1].is_turn && <p>Currently Playing</p>}
            {players[1].last_card_played && (
              <Card
                descriptor={players[1].last_card_played.descriptor}
                height={80}
                width={60}
              />
            )}
          </div>
        )}
      </Grid>
      <Grid item md={3}>
        {players.length > 2 && (
          <div>
            <p>{players[2].name}</p>
            <p>Lives: {players[2].lives}</p>
            {players[2].is_turn && <p>Currently Playing</p>}
            {players[2].last_card_played && (
              <Card
                descriptor={players[2].last_card_played.descriptor}
                height={80}
                width={60}
              />
            )}
          </div>
        )}
      </Grid>
      <Grid item md={3}>
        {players.length > 3 && (
          <div>
            <p>{players[3].name}</p>
            <p>Lives: {players[3].lives}</p>
            {players[3].is_turn && <p>Currently Playing</p>}
            {players[3].last_card_played && (
              <Card
                descriptor={players[3].last_card_played.descriptor}
                height={80}
                width={60}
              />
            )}
          </div>
        )}
      </Grid>
      <Grid item md={3}>
        {players.length > 9 && (
          <div>
            <p>{players[9].name}</p>
            <p>Lives: {players[9].lives}</p>
            {players[9].is_turn && <p>Currently Playing</p>}
            {players[9].last_card_played && (
              <Card
                descriptor={players[9].last_card_played.descriptor}
                height={80}
                width={60}
              />
            )}
          </div>
        )}
      </Grid>
      <Grid item md={3}></Grid>
      <Grid item md={3}></Grid>
      <Grid item md={3}>
        {players.length > 4 && (
          <div>
            <p>{players[4].name}</p>
            <p>Lives: {players[4].lives}</p>
            {players[4].is_turn && <p>Currently Playing</p>}
            {players[4].last_card_played && (
              <Card
                descriptor={players[4].last_card_played.descriptor}
                height={80}
                width={60}
              />
            )}
          </div>
        )}
      </Grid>
      <Grid item md={3}>
        {players.length > 8 && (
          <div>
            <p>{players[8].name}</p>
            <p>Lives: {players[8].lives}</p>
            {players[8].is_turn && <p>Currently Playing</p>}
            {players[8].last_card_played && (
              <Card
                descriptor={players[8].last_card_played.descriptor}
                height={80}
                width={60}
              />
            )}
          </div>
        )}
      </Grid>
      <Grid item md={3}>
        {players.length > 7 && (
          <div>
            <p>{players[7].name}</p>
            <p>Lives: {players[7].lives}</p>
            {players[7].is_turn && <p>Currently Playing</p>}
            {players[7].last_card_played && (
              <Card
                descriptor={players[7].last_card_played.descriptor}
                height={80}
                width={60}
              />
            )}
          </div>
        )}
      </Grid>
      <Grid item md={3}>
        {players.length > 6 && (
          <div>
            <p>{players[6].name}</p>
            <p>Lives: {players[6].lives}</p>
            {players[6].is_turn && <p>Currently Playing</p>}
            {players[6].last_card_played && (
              <Card
                descriptor={players[6].last_card_played.descriptor}
                height={80}
                width={60}
              />
            )}
          </div>
        )}
      </Grid>
      <Grid item md={3}>
        {players.length > 5 && (
          <div>
            <p>{players[5].name}</p>
            <p>Lives: {players[5].lives}</p>
            {players[5].is_turn && <p>Currently Playing</p>}
            {players[5].last_card_played && (
              <Card
                descriptor={players[5].last_card_played.descriptor}
                height={80}
                width={60}
              />
            )}
          </div>
        )}
      </Grid>
    </Grid>
  );
};
