import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Grid, Typography, TextField, Button } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    padding: theme.spacing(2),
    backgroundColor: '#fff',
  },
  sectionTitle: {
    fontWeight: 'bold',
    marginBottom: theme.spacing(2),
  },
  formContainer: {
    marginBottom: theme.spacing(4),
  },
  formInput: {
    marginBottom: theme.spacing(2),
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

export default function CheckoutPage() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={4}>
        <Grid item xs={12}>
          <Typography variant="h4" gutterBottom>Checkout</Typography>
        </Grid>
        <Grid item xs={12}>
          <div className={classes.formContainer}>
            <Typography className={classes.sectionTitle}>Delivery Information</Typography>
            <TextField label="Full Name" variant="outlined" fullWidth className={classes.formInput} />
            <TextField label="Email" variant="outlined" fullWidth className={classes.formInput} />
            <TextField label="Phone Number" variant="outlined" fullWidth className={classes.formInput} />
            <TextField label="Address" variant="outlined" fullWidth className={classes.formInput} />
            <TextField label="City" variant="outlined" fullWidth className={classes.formInput} />
            <TextField label="State" variant="outlined" fullWidth className={classes.formInput} />
            <TextField label="Zip Code" variant="outlined" fullWidth className={classes.formInput} />
          </div>
          <div className={classes.formContainer}>
            <Typography className={classes.sectionTitle}>Payment Information</Typography>
            <TextField label="Cardholder Name" variant="outlined" fullWidth className={classes.formInput} />
            <TextField label="Card Number" variant="outlined" fullWidth className={classes.formInput} />
            <TextField label="Expiration Date" variant="outlined" fullWidth className={classes.formInput} />
            <TextField label="CVV" variant="outlined" fullWidth className={classes.formInput} />
          </div>
          <Button variant="contained" className={classes.checkoutButton}>Place Order</Button>
        </Grid>
      </Grid>
    </div>
  );
}
