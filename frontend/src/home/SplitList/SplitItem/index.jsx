import React from 'react';
import PropTypes from 'prop-types';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';

function Item({ title, desc }) {
  return (
    <Button
      component="a"
      sx={{
        display: 'block',
      }}
      color="inherit"
    >
      <Typography variant="h5" sx={{ textTransform: 'none' }}>{ title }</Typography>
      <Typography variant="body1" sx={{ textTransform: 'none' }}>{ desc }</Typography>
    </Button>
  );
}

Item.propTypes = {
  title: PropTypes.string.isRequired,
  desc: PropTypes.string.isRequired,
};

export default Item;
