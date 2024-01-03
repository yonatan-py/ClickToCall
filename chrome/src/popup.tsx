import React, { useState, useEffect } from 'react';
import { createRoot } from 'react-dom/client';

import SetCode from './setCode';


function App() {
  return (
    <div>
      <SetCode/>
    </div>
  );
}

const container = document.getElementById('root');
const root = createRoot(container); 
root.render(<App/>);
