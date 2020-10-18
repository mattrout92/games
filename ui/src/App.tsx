import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Home from "./Pages/Home";
import ChaseTheAce from "./Pages/ChaseTheAce";

import "./App.css";
import { ThemeProvider, CssBaseline } from "@material-ui/core";
import { theme } from "./theme";
import FiveLives from "./Pages/5Lives";

const App = (): JSX.Element => {
  return (
    <ThemeProvider theme={theme}>
      <Router>
        <Switch>
          <Route exact path="/" component={Home} />
          <Route exact path="/five-lives" component={FiveLives} />
          <Route exact path="/chase-the-ace" component={ChaseTheAce} />
        </Switch>
        <CssBaseline />
      </Router>
    </ThemeProvider>
  );
};

export default App;
