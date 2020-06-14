import { createMuiTheme } from "@material-ui/core/styles";

export const theme = createMuiTheme({
  palette: {
    type: "dark",
    primary: {
      main: "rgba(255, 255, 255, 0.87)",
    },
    secondary: {
      main: "#4275e3",
    },
    background: {
      paper: "#1abfc4",
      default: "#1abfc4",
    },
    text: {
      primary: "rgba(255, 255, 255, 0.87)",
      secondary: "rgba(255, 255, 255, 0.87)",
    },
  },
});
