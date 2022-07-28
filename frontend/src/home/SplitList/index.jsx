import React from 'react';
import PropTypes from 'prop-types';
import Stack from '@mui/material/Stack';
import Divider from '@mui/material/Divider';
import Typography from '@mui/material/Typography';
import Item from './SplitItem/index.jsx';

function List({ data }) {
  let listItems = data.map((splitup) => (
    <Item key={splitup.id} title={splitup.title} desc={splitup.desc} />
  ));
  if (listItems.length === 0) {
    listItems = <Typography variant="h6" align="center">No splitcounts yet</Typography>;
  }
  return (
    <Stack divider={<Divider />} className="bg-white rounded">{listItems}</Stack>
  );
}

List.propTypes = {
  data: PropTypes.arrayOf(PropTypes.shape({
    id: PropTypes.string,
    title: PropTypes.string,
    desc: PropTypes.string,
  })).isRequired,
};

export default List;
