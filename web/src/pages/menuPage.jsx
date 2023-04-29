import React, { useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Button, Card, CardActions, CardContent, CardMedia, Typography } from '@material-ui/core';

const useStyles = makeStyles({
  root: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    padding: '2rem',
  },
  title: {
    marginBottom: '2rem',
  },
  card: {
    display: 'flex',
    marginBottom: '1rem',
    width: '80%',
  },
  media: {
    width: '25%',
    paddingTop: '25%',
  },
  details: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'flex-start',
    marginLeft: '1rem',
  },
});

const MenuPage = () => {
  const classes = useStyles();

  // example menu items data
  const [menuItems] = useState([
    {
      id: 1,
      name: 'Burger',
      description: 'A classic burger with lettuce, tomato, and cheese',
      price: 9.99,
      image: 'https://source.unsplash.com/collection/1163637/400x400',
    },
    {
      id: 2,
      name: 'Pizza',
      description: 'A delicious pizza with cheese, pepperoni, and mushrooms',
      price: 12.99,
      image: 'https://source.unsplash.com/collection/186703/400x400',
    },
    {
      id: 3,
      name: 'Taco',
      description: 'A spicy taco with beef, lettuce, and salsa',
      price: 7.99,
      image: 'https://source.unsplash.com/collection/364929/400x400',
    },
  ]);

  return (
    <div className={classes.root}>
      <Typography variant="h4" className={classes.title}>
        Menu
      </Typography>
      {menuItems.map((menuItem) => (
        <Card key={menuItem.id} className={classes.card}>
          <CardMedia className={classes.media} image={menuItem.image} title={menuItem.name} />
          <div className={classes.details}>
            <CardContent>
              <Typography variant="h5" component="h2">
                {menuItem.name}
              </Typography>
              <Typography variant="body2" color="textSecondary" component="p">
                {menuItem.description}
              </Typography>
              <Typography variant="h6" component="p">
                ${menuItem.price}
              </Typography>
            </CardContent>
            <CardActions>
              <Button size="small" color="primary">
                Add to Cart
              </Button>
            </CardActions>
          </div>
        </Card>
      ))}
    </div>
  );
};

export default MenuPage;
