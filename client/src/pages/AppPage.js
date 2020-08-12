import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Button from "@material-ui/core/Button";
import IconButton from "@material-ui/core/IconButton";
import Typography from "@material-ui/core/Typography";
import AccountCircle from "@material-ui/icons/AccountCircle";
import TextField from "@material-ui/core/TextField";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogContentText from "@material-ui/core/DialogContentText";
import DialogTitle from "@material-ui/core/DialogTitle";
import RefreshIcon from "@material-ui/icons/Refresh";

import Domains from "../components/Domains";
import useSWR from "swr";
import fetcher from "../api/fetcher";
import { Avatar } from "@material-ui/core";

const useStyles = makeStyles((theme) => ({
  root: {
    margin: "16px",
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
  heading: {
    fontSize: theme.typography.pxToRem(20),
    fontWeight: theme.typography.fontWeightBold,
  },
}));

export default function AppPage() {
  const classes = useStyles();
  const [open, setOpen] = React.useState(false);
  const [loggedIn, setLoggedIn] = React.useState(false);
  const [otp, setOtp] = React.useState(null);
  // const {loading, data: tokenEndpoints, error} = useSWR("/api/token", fetcher);
  const [identityEndpoint, setIdentityEndpoint] = React.useState(null);

  React.useEffect(() => {
    fetch("/api/auth/login").then((resp) => {
      if (resp.status === 200) {
        setLoggedIn(true);
      } else {
        setLoggedIn(false);
      }
    });
  }, []);

  React.useEffect(() => {
    fetch("/api/token")
      .then((r) => r.json())
      .then(({ passcode_endpoint }) => {
        setIdentityEndpoint(passcode_endpoint);
      });
  }, []);

  const { loading: userLoading, data: user, error: userError } = useSWR(
    "/api/user",
    fetcher
  );
  const { loading: prefLoading, data: userPref, error: prefError } = useSWR(
    user ? `/api/user/${user.iam_id}/preference` : null,
    fetcher
  );

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleLogout = React.useCallback(async () => {
    const { status } = await fetch("/api/auth/logout", {
      method: "POST",
    });
    if (status === 200) {
      setLoggedIn(false);
    }
  }, []);

  const handleSubmitOTP = React.useCallback(async (otp) => {
    const { status } = await fetch("/api/auth/authenticate", {
      method: "POST",
      body: JSON.stringify({
        otp: otp,
      }),
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (status === 200) {
      setLoggedIn(true);
      setOpen(false);
      setOtp("");
      return;
    }
    setLoggedIn(false);
  }, []);

  return (
    <div>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            IBM Developer DNS
          </Typography>

          {loggedIn ? (
            <>
              <Button color="inherit" onClick={handleLogout}>
                Logout
              </Button>
              <Avatar src={userPref && userPref.photo} />
            </>
          ) : (
            <>
              <Button
                href={identityEndpoint}
                rel="noopener noreferer"
                target="_blank"
                onClick={handleClickOpen}
                color="inherit"
              >
                Login
              </Button>

              <IconButton
                aria-label="account of current user"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                color="inherit"
              >
                <AccountCircle />
              </IconButton>
            </>
          )}
        </Toolbar>
      </AppBar>
      <Dialog
        maxWidth="md"
        open={open}
        onClose={() => handleSubmitOTP(otp)}
        aria-labelledby="form-dialog-title"
      >
        <DialogTitle id="form-dialog-title">Subscribe</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Paste the OTP from iam identity.
          </DialogContentText>
          <TextField
            autoFocus
            margin="dense"
            id="name"
            label="OTP"
            value={otp}
            onChange={(e) => setOtp(e.target.value)}
            type="email"
            fullWidth
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => handleSubmitOTP(otp)} color="primary">
            Submit
          </Button>
        </DialogActions>
      </Dialog>
      <Domains />
      {loggedIn && user && user.email.includes("ibm.com") ? (
        <div>You are cool to add new records</div>
      ) : (
        <div>No new record for you</div>
      )}
    </div>
  );
}
