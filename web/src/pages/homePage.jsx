import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Grid, TextField, Button } from '@material-ui/core';


const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    padding: theme.spacing(2),
    backgroundColor: '#fff',
  },
  searchContainer: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    marginBottom: theme.spacing(4),
  },
  searchInput: {
    backgroundColor: '#f9f9f9',
    borderRadius: theme.shape.borderRadius,
    marginRight: theme.spacing(2),
  },
  popularRestaurants: {
    marginBottom: theme.spacing(2),
  },
  popularTitle: {
    fontWeight: 'bold',
    marginBottom: theme.spacing(2),
  },
  popularItem: {
    display: 'flex',
    alignItems: 'center',
    marginBottom: theme.spacing(2),
    cursor: 'pointer',
  },
  popularImage: {
    width: 64,
    height: 64,
    marginRight: theme.spacing(2),
  },
  popularName: {
    fontWeight: 'bold',
  },
  ctaButton: {
    backgroundColor: '#d42a2a',
    color: '#fff',
    padding: theme.spacing(2),
    borderRadius: theme.shape.borderRadius,
    fontWeight: 'bold',
    '&:hover': {
      backgroundColor: '#c92121',
    },
  },
}));

export default function HomePage() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={4}>
        <Grid item xs={12} className={classes.searchContainer}>
          <TextField
            variant="outlined"
            placeholder="Search for restaurants or dishes"
            className={classes.searchInput}
          />
          <Button variant="contained" className={classes.ctaButton}>Search</Button>
        </Grid>
        <Grid item xs={12} className={classes.popularRestaurants}>
          <h2 className={classes.popularTitle}>Popular Restaurants</h2>
          <div className={classes.popularItem}>
            <img src="https://via.placeholder.com/64" alt="restaurant" className={classes.popularImage} />
            <div>
              <h3 className={classes.popularName}>Pizza Hut</h3>
              <p>Italian, Pizza, Fast Food</p>
            </div>
          </div>
          <div className={classes.popularItem}>
            <img src="https://via.placeholder.com/64" alt="restaurant" className={classes.popularImage} />
            <div>
              <h3 className={classes.popularName}>McDonald's</h3>
              <p>Burgers, Fast Food, American</p>
            </div>
          </div>
          <div className={classes.popularItem}>
            <img src="https://via.placeholder.com/64" alt="restaurant" className={classes.popularImage} />
            <div>
              <h3 className={classes.popularName}>Subway</h3>
              <p>Subs, Sandwiches, Fast Food</p>
            </div>
          </div>
        </Grid>
      </Grid>
    </div>
  );
}
