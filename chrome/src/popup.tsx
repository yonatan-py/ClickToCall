import React, { useState, useEffect } from 'react';
import { createRoot } from 'react-dom/client';

import SetCode from './setCode';
import LoggedIn from './loggedIn';


function App() {
  const [secret, setSecret] = useState();
  const [userID, setUserID] = useState();
  function updateCredentials() {
      chrome.storage.local.get(["clickToCall.secret", "clickToCall.userID"], (data) => {
          console.log(data)
          setSecret(data["clickToCall.secret"])
          setUserID(data["clickToCall.userID"])
      })
  }
  function isLoggedin() {
      return secret && userID
  }
  return (
    isLoggedin() ?
    <LoggedIn updateCredentials={updateCredentials}/>:
    <SetCode updateCredentials={updateCredentials} isLoggedin={isLoggedin} />
  );
}

const container = document.getElementById('root');
const root = createRoot(container); 
root.render(<App/>);
