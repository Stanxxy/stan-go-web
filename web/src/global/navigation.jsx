import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { AppBar, Toolbar, Typography } from '@material-ui/core';
import { Link, useLocation } from 'react-router-dom';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  title: {
    flexGrow: 1,
  },
  link: {
    color: '#fff',
    textDecoration: 'none',
    marginRight: theme.spacing(2),
    '&:hover': {
      textDecoration: 'underline',
    },
  },
  activeLink: {
    fontWeight: 'bold',
    textDecoration: 'underline',
  },
}));

export default function Navigation() {
  const classes = useStyles();
  const location = useLocation();

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            Food Ordering Platform
          </Typography>
          <Link to="/" className={`${classes.link} ${location.pathname === '/' && classes.activeLink}`}>
            Home
          </Link>
          <Link to="/menu" className={`${classes.link} ${location.pathname === '/menu' && classes.activeLink}`}>
            Menu
          </Link>
          <Link to="/cart" className={`${classes.link} ${location.pathname === '/cart' && classes.activeLink}`}>
            Cart
          </Link>
          <Link to="/discover" className={`${classes.link} ${location.pathname === '/discover' && classes.activeLink}`}>
            Discover
          </Link>
        </Toolbar>
      </AppBar>
    </div>
  );
}
