import React from 'react';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';
import Icon from '@mui/material/Icon';

function TopBar() {
  return (
    <AppBar position="static">
      <Toolbar sx={{
        justifyContent: 'space-between',
      }}
      >
        <Typography variant="h4" component="div">Splitup</Typography>
        <IconButton color="inherit">
          <Icon baseClassName="material-symbols-outlined">more_vert</Icon>
        </IconButton>
      </Toolbar>
    </AppBar>
  );
}

export default TopBar;
