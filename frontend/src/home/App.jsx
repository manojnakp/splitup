import React from 'react';
import TopBar from './TopBar/index.jsx';
import FAB from './FAB/index.jsx';
import SplitList from './SplitList/index.jsx';

export default function App() {
  const list = [
    {
      id: 'splitup1',
      title: 'Name',
      desc: 'Long description of the splitup',
    },
    {
      id: 'splitup2',
      title: 'City trip',
      desc: 'This is a sample splitup',
    },
  ];
  return (
    <>
      <TopBar />
      <SplitList data={list} />
      <FAB />
    </>
  );
}
