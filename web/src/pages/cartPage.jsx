import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Grid, Typography, Button } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    padding: theme.spacing(2),
    backgroundColor: '#fff',
  },
  itemContainer: {
    display: 'flex',
    alignItems: 'center',
    marginBottom: theme.spacing(2),
  },
  itemImage: {
    width: 64,
    height: 64,
    objectFit: 'cover',
    marginRight: theme.spacing(2),
  },
  itemName: {
    fontWeight: 'bold',
    marginBottom: theme.spacing(1),
  },
  itemPrice: {
    color: '#999',
  },
  quantityInput: {
    width: 50,
    marginLeft: theme.spacing(2),
    marginRight: theme.spacing(2),
  },
  checkoutButton: {
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

export default function CartPage() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={4}>
        <Grid item xs={12}>
          <Typography variant="h4" gutterBottom>Your Cart</Typography>
        </Grid>
        <Grid item xs={12}>
          <div className={classes.itemContainer}>
            <img src="https://via.placeholder.com/64" alt="item" className={classes.itemImage} />
            <div>
              <Typography className={classes.itemName}>Pepperoni Pizza</Typography>
              <Typography className={classes.itemPrice}>$12.99</Typography>
              <input type="number" defaultValue={1} className={classes.quantityInput} />
              <Button variant="outlined">Remove</Button>
            </div>
          </div>
          <div className={classes.itemContainer}>
            <img src="https://via.placeholder.com/64" alt="item" className={classes.itemImage} />
            <div>
              <Typography className={classes.itemName}>Cheeseburger</Typography>
              <Typography className={classes.itemPrice}>$8.99</Typography>
              <input type="number" defaultValue={2} className={classes.quantityInput} />
              <Button variant="outlined">Remove</Button>
            </div>
          </div>
          <div className={classes.itemContainer}>
            <img src="https://via.placeholder.com/64" alt="item" className={classes.itemImage} />
            <div>
              <Typography className={classes.itemName}>Spaghetti Bolognese</Typography>
              <Typography className={classes.itemPrice}>$11.99</Typography>
              <input type="number" defaultValue={1} className={classes.quantityInput} />
              <Button variant="outlined">Remove</Button>
            </div>
          </div>
          <Button variant="contained" className={classes.checkoutButton}>Checkout</Button>
        </Grid>
      </Grid>
    </div>
  );
}
