import React from 'react';
import TopBar from './TopBar/index.jsx';
import BottomBar from './BottomBar/index.jsx';
import Main from './Main/index.jsx';

export default function App() {
  const participants = [
    'Alex',
    'Brian',
    'Julia',
    'Thomas',
  ];
  return (
    <>
      <TopBar title="City trip" participants={participants} />
      <Main />
      <BottomBar my={39.25} total={162.00} />
    </>
  );
}
