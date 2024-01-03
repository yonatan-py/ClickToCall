import React, { useState, useEffect } from 'react';


const sendCodeToServer = async (code: string, onDone: () => void) => {
    const response = await fetch("http://localhost:8080/code", {
        method: "POST",
        body: JSON.stringify({ code }),
        headers: {"Content-type": "application/json"}
    })
    const data = await response.json()
    chrome.storage.local.set({
        "clickToCall.userID": data.userid
    })
    chrome.storage.local.set({
        "clickToCall.secret": data.secret
    })
    onDone()
}

const SetCode = () => {
    
    
    const [secret, setSecret] = useState();
    const [userID, setUserID] = useState();
    const [code, setCode] = useState();
    function updateCredentials() {
        chrome.storage.local.get(["clickToCall.secret", "clickToCall.userID"], (data) => {
            console.log(data)
            setSecret(data["clickToCall.secret"])
            setUserID(data["clickToCall.userID"])
        })
    }
    updateCredentials()
    
    const onCodeUpdate = async (event: any) => {
        const code = event.target.value
        setCode(code)
        // TODO: check if code should be longer
        if (code.length === 3) {
            sendCodeToServer(code, updateCredentials)
        }
    }
    return (secret && userID ?
            <div>Logged in!</div>:
            <input type="text" value={code} onChange={onCodeUpdate}/>)
}

export default SetCode;