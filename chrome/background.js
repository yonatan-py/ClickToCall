chrome.contextMenus.onClicked.addListener(genericOnClick);

// A generic onclick callback function.
function genericOnClick(info) {
  var linkUrl = info.linkUrl;
  var phone = linkUrl.substring(4)
  console.log("Call ", phone)
}

chrome.runtime.onInstalled.addListener(function () {
  // Create one test item for each context type.
  chrome.contextMenus.create({
      title: "Click TO Call",
      contexts: ["link"],
      id: "click-to-call"
    });
  }
);