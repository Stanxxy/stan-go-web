import React, { useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { TextField, Button, Typography, Link, IconButton, Box } from '@material-ui/core';
import { Visibility, VisibilityOff } from '@material-ui/icons';
import { GoogleLogin } from "react-google-login";
import GTranslateIcon from '@material-ui/icons/GTranslate';
import FacebookIcon from '@material-ui/icons/Facebook';
import AppleIcon from '@material-ui/icons/Apple';

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    padding: theme.spacing(2),
    '& .MuiTextField-root': {
      margin: theme.spacing(1),
      width: '300px',
    },
    '& .MuiButton-root': {
      margin: theme.spacing(2),
      width: '300px',
    },
  },
  socialLogin: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    // width: '300px',
    margin: theme.spacing(2),
  },
  loginButton: {
    width: '100%',
  },
  socialButton: {
    margin: theme.spacing(1, 0),
    padding: theme.spacing(2),
    textAlign: 'center',
  },
  socialIcon: {
    marginRight: theme.spacing(1),
  }
}));

function Login() {
  const classes = useStyles();
  const [showPassword, setShowPassword] = useState(false);
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleUsernameChange = (event) => {
    setUsername(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    // handle login logic
  };


  const handleGoogleLogin = () => {
    const provider = new GoogleLogin({
        clientId: "YOUR_CLIENT_ID",
        scope: "https://www.googleapis.com/auth/userinfo.profile",
        onLogin: () => {
          setIsLoggedIn(true);
        },
        onLogout: () => {
          setIsLoggedIn(false);
        },
      });
  
    provider.login();
    console.log("clicked google login");
  };
  
  const handleFacebookLogin = () => {
    // Handle Facebook login logic
    console.log("clicked fb login");
  };
  
  const handleAppleLogin = () => {
    // Handle Apple login logic
    console.log("clicked apple login");
  };
  
  const SocialLoginButtons = () => {
    const classes = useStyles();
  }

  const handleShowPassword = () => {
    setShowPassword(!showPassword);
  };

  return (
    <div className={classes.root}>
      <Typography variant="h4" gutterBottom>
        Login
      </Typography>
      <form onSubmit={handleSubmit}>
        <TextField
          label="Username or email"
          variant="outlined"
          value={username}
          onChange={handleUsernameChange}
        />
        <TextField
          label="Password"
          variant="outlined"
          type={showPassword ? 'text' : 'password'}
          value={password}
          onChange={handlePasswordChange}
          InputProps={{
            endAdornment: (
              <IconButton onClick={handleShowPassword}>
                {showPassword ? <Visibility /> : <VisibilityOff />}
              </IconButton>
            ),
          }}
        />
        <Button
          variant="contained"
          color="primary"
          type="submit"
          className={classes.loginButton}
        >
          Sign In
        </Button>
        <Box className={classes.socialLogin}>
        <Button
            className={classes.socialButton}
            variant="contained"
            color="primary"
            startIcon={<GTranslateIcon className={classes.socialIcon} />}
            onClick={handleGoogleLogin}
        >
            Login with Google
        </Button>
        <Button
            className={classes.socialButton}
            variant="contained"
            color="primary"
            startIcon={<FacebookIcon className={classes.socialIcon} />}
            onClick={handleFacebookLogin}
        >
            Login with Facebook
        </Button>
        <Button
            className={classes.socialButton}
            variant="contained"
            color="primary"
            startIcon={<AppleIcon className={classes.socialIcon} />}
            onClick={handleAppleLogin}
        >
            Login with Apple
        </Button>
        </Box>
        <Typography variant="body1">
          Don't have an account? <Link href="/signup">Sign up</Link>
        </Typography>
        <Typography variant="body1">
          Forgot your password? <Link href="/forgot-password">Reset Password</Link>
        </Typography>
        <Typography variant="body1">
          Login with phone number? <Link href="/login-with-phone">Sign in with phone number</Link>
        </Typography>
      </form>
    </div>
  );
}

export default Login;
