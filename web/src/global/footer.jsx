import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Typography } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    padding: theme.spacing(2),
    backgroundColor: '#f5f5f5',
  },
  footerText: {
    color: '#555',
  },
}));

export default function Footer() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Typography variant="body1" align="center" className={classes.footerText}>
        Food Ordering Platform &copy; {new Date().getFullYear()}
      </Typography>
    </div>
  );
}
