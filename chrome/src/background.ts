import { apiUrl } from "./config";


const main = async () => {
  console.log("installed")
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
    // TODO: deal with errors
    fetch(`${apiUrl}/call`, {
      method: "POST",
      body: JSON.stringify(data),
      headers: {"Content-type": "application/json"}
    })
  };
  // TODO: this is somewhat tricky...
  if (secret && userID) {
    chrome.contextMenus.create({
      title: 'Call',
      id: 'clickToCall',
      contexts: ["all", "link"],
    });
    chrome.contextMenus.onClicked.addListener(onClick);
  }
}

main()