import React from 'react';
import PropTypes from 'prop-types';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Fab from '@mui/material/Fab';
import Icon from '@mui/material/Icon';

function FAB() {
  return (
    <Fab
      color="primary"
      aria-label="add"
      sx={{
        transform: 'translate(0, -50%)',
      }}
    >
      <Icon baseClassName="material-symbols-outlined">add</Icon>
    </Fab>
  );
}

function BottomBar({ my, total }) {
  return (
    <Paper
      sx={{
        position: 'fixed',
        bottom: 0,
        left: 0,
        right: 0,
      }}
    >
      <Stack direction="row" sx={{ justifyContent: 'space-between' }}>
        <Box className="mx-4 my-2">
          <Typography variant="body-2" component="p" sx={{ textTransform: 'uppercase' }}>My Total</Typography>
          <Typography variant="subtitle-2" component="p" sx={{ fontWeight: 'bold' }}>{`₹${my.toFixed(2)}`}</Typography>
        </Box>
        <FAB />
        <Box className="mx-4 my-2">
          <Typography variant="body-2" component="p" sx={{ textTransform: 'uppercase', textAlign: 'right' }}>Total Expenses</Typography>
          <Typography variant="subtitle-2" component="p" sx={{ fontWeight: 'bold', textAlign: 'right' }}>{`₹${total.toFixed(2)}`}</Typography>
        </Box>
      </Stack>
    </Paper>
  );
}

BottomBar.propTypes = {
  my: PropTypes.number.isRequired,
  total: PropTypes.number.isRequired,
};

export default BottomBar;
