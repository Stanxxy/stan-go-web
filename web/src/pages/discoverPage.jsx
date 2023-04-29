import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';
import { GoogleMap, useLoadScript, Marker } from '@react-google-maps/api';
import { getSearchResults } from '../utils/api';

const GOOGLE_MAPS_API_KEY = process.env.REACT_APP_GOOGLE_MAPS_API_KEY;

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
  }));

const libraries = ['places'];
const mapContainerStyle = {
  width: '100%',
  height: '500px',
};

const DiscoverPage = () => {
  const classes = useStyles();
  const navigate = useNavigate();
  const [searchTerm, setSearchTerm] = useState('');
  const [mapLoaded, setMapLoaded] = useState(false);
  const [map, setMap] = useState(null);
  // There should be a function to get center and remake center
  const [center, setCenter] = useState({
        lat: 40.73061,
        lng: -73.935242,
    });
  const [markers, setMarkers] = useState([]);
  const { isLoaded, loadError } = useLoadScript({
    googleMapsApiKey: process.env.REACT_APP_GOOGLE_MAPS_API_KEY,
    libraries,
  });

  const handleMarkerClick = (result) => {
    navigate.push(`/menu/${result.id}`);
  };

  useEffect(() => {
    if (isLoaded && !mapLoaded) {
      setMapLoaded(true);
      const mapInstance = new window.google.maps.Map(document.getElementById('map'), {
        center,
        zoom: 10,
      });
      setMap(mapInstance);
    }
  }, [isLoaded, mapLoaded]);

  const handleSearch = async () => {
    try {
        // dummy searchResults
      const searchResults = [
        {id:1, name: 'Restaurant A', location: { lat: 40.741895, lng: -73.989308 } },
        {id:2, name: 'Restaurant B', location: { lat: 40.738129, lng: -73.992236 } },
        {id:3, name: 'Restaurant C', location: { lat: 40.749825, lng: -73.991961 } },
      ];
      
      // await getSearchResults(searchTerm);

      // Clear existing markers
      setMarkers([]);

      // Add markers for search results
      searchResults.forEach((result) => {
        const marker = new window.google.maps.Marker({
          position: result.location,
          map,
          title: result.name,
          onClick: () => handleMarkerClick(result)
        });
        setMarkers((prevMarkers) => [...prevMarkers, marker]);
      });
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div className={classes.root}>
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
        <TextField
          label="Search by location or food provider"
          variant="outlined"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          style={{ marginRight: '10px' }}
        />
        <Button variant="contained" color="primary" onClick={handleSearch}>
          Search
        </Button>
      </div>
      <div id="map" style={mapContainerStyle}></div>
      <script src={`https://maps.googleapis.com/maps/api/js?key=${GOOGLE_MAPS_API_KEY}&libraries=places`}></script>
    </div>
  );
};

export default DiscoverPage;
