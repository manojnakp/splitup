import React, { useState } from 'react';
import PropTypes from 'prop-types';
import SwipeableViews from 'react-swipeable-views';
import AppBar from '@mui/material/AppBar';
import Tab from '@mui/material/Tab';
import Tabs from '@mui/material/Tabs';

function ExpensesPanel({ value }) {
  return (
    <div role="tabpanel" hidden={value !== 0} id="tabpanel-expenses" aria-labelledby="tab-expenses">
      {value === 0 && (
        <div>Content</div>
      )}
    </div>
  );
}

ExpensesPanel.propTypes = {
  value: PropTypes.oneOf([0, 1]).isRequired,
};

function BalancesPanel({ value }) {
  return (
    <div role="tabpanel" hidden={value !== 1} id="tabpanel-balances" aria-labelledby="tab-balances">
      {value === 1 && (
        <div>Content</div>
      )}
    </div>
  );
}

BalancesPanel.propTypes = {
  value: PropTypes.oneOf([0, 1]).isRequired,
};

function Main() {
  // const mapping = ['expenses', 'balances'];
  const [value, setValue] = useState(0);
  const handleChange = (event, newValue) => {
    setValue(newValue);
  };
  const handleChangeIndex = (index) => {
    setValue(index);
  };
  return (
    <main>
      <AppBar component="div" position="static">
        <Tabs
          value={value}
          onChange={handleChange}
          indicatorColor="secondary"
          textColor="inherit"
          variant="fullWidth"
          aria-label="splitup tabs"
        >
          <Tab label="expenses" id="tab-expenses" aria-controls="tabpanel-expenses" />
          <Tab label="balances" id="tab-balances" aria-controls="tabpanel-balances" />
        </Tabs>
      </AppBar>
      <SwipeableViews index={value} onChangeIndex={handleChangeIndex}>
        <ExpensesPanel value={value} />
        <BalancesPanel value={value} />
      </SwipeableViews>
    </main>
  );
}

export default Main;
