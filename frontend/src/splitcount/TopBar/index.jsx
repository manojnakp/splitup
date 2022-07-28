import React from 'react';
import PropTypes from 'prop-types';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';
import Icon from '@mui/material/Icon';

function Title({ title, participants }) {
  const n = participants.length;
  let plist;
  if (n === 0) {
    plist = '-';
  } else {
    [plist] = participants;
  }
  if (n > 1) {
    plist = plist.concat(`, ${participants[1]}`);
  }
  if (n > 2) {
    plist = plist.concat(', ...');
  }
  return (
    <Box sx={{ flexGrow: 1 }} className="px-4">
      <Typography variant="h5" component="h1">{title}</Typography>
      <Typography variant="subtitle2" component="p">{plist}</Typography>
    </Box>
  );
}

Title.propTypes = {
  title: PropTypes.string.isRequired,
  participants: PropTypes.arrayOf(PropTypes.string).isRequired,
};

function TopBar({ title, participants }) {
  return (
    <AppBar position="static">
      <Toolbar>
        <IconButton color="inherit">
          <Icon baseClassName="material-symbols-outlined">arrow_back</Icon>
        </IconButton>
        <Title title={title} participants={participants} />
        <IconButton color="inherit">
          <Icon baseClassName="material-symbols-outlined">more_vert</Icon>
        </IconButton>
      </Toolbar>
    </AppBar>
  );
}

TopBar.propTypes = {
  title: PropTypes.string.isRequired,
  participants: PropTypes.arrayOf(PropTypes.string).isRequired,
};

export default TopBar;
