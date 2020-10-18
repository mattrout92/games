import React from "react";
import {
  Grid,
  Button,
  makeStyles,
  createStyles,
  Theme,
} from "@material-ui/core";
import { Link } from "react-router-dom";

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

const Home = (): JSX.Element => {
  const classes = useStyles();

  return (
    <Grid container justify="center">
      <Grid item xs={6}>
        <Button
          className={classes.button}
          fullWidth
          color="primary"
          variant="contained"
          component={Link}
          to="/five-lives"
        >
          Five Lives
        </Button>
      </Grid>
      <Grid item xs={6}>
        <Button
          className={classes.button}
          fullWidth
          color="secondary"
          variant="contained"
          component={Link}
          to="/chase-the-ace"
        >
          Chase The Ace
        </Button>
      </Grid>
    </Grid>
  );
};

export default Home;
