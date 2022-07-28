import React from 'react';
import Fab from '@mui/material/Fab';
import Icon from '@mui/material/Icon';

function FAB() {
  return (
    <Fab
      color="primary"
      aria-label="add"
      sx={{
        position: 'fixed',
        bottom: '1em',
        right: '1em',
      }}
    >
      <Icon baseClassName="material-symbols-outlined">add</Icon>
    </Fab>
  );
}

export default FAB;
