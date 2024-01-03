chrome.runtime.onInstalled.addListener(async () => {
  const data = await chrome.storage.local.get(["clickToCall.secret", "clickToCall.userID"])
  
  console.log(data)
  var secret = data["clickToCall.secret"]
  var userID = data["clickToCall.userID"] 
  
  function onClick(info: any) {
    var linkUrl = info.linkUrl;
    var number = linkUrl.substring(4)
    var data = {
      number: number,
      userId: userID,
      secret: secret,
    }
    fetch('http://localhost:8080/call', {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
          "Content-type": "application/json"
      }
    }).then(res => {
      console.log("call sent to server")
      console.log(res.json())
    })
  };
  if (secret && userID) {
    chrome.contextMenus.create({
      title: 'Call',
      id: 'clickToCall',
      contexts: ["all", "link"],
    });
    chrome.contextMenus.onClicked.addListener(onClick);
  }
})