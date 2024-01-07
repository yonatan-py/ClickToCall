import React, { useState, useEffect } from 'react';

import { apiUrl } from './config';


const sendCodeToServer = async (code: string, onDone: () => void) => {
    const response = await fetch(`${apiUrl}/code`, {
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

type SetCodeProps = {
    updateCredentials: () => void,
    isLoggedin: () => boolean
}

function SetCode({isLoggedin, updateCredentials}: SetCodeProps) {
    updateCredentials()
    const [code, setCode] = useState("");
    
    
    const onCodeUpdate = async (newCode: string) => {
        // TODO: code should be longer
        if (newCode.length === 3 && !isLoggedin()) {
            sendCodeToServer(newCode, updateCredentials)
        }
    }
    useEffect(() => {
        console.log("code changed")
        onCodeUpdate(code)
    }, [code])
    return <input type="text" value={code} onChange={e => setCode(e.target.value)}/>
}

export default SetCode;