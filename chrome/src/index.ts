import { initializeApp, setLogLevel as firebaseSetLogLevel } from 'firebase/app';
import { getFirestore, doc, getDoc } from 'firebase/firestore/lite';


const firebaseConfig = {
  authDomain: "click-to-call-d2769.firebaseapp.com",
  projectId: "click-to-call-d2769",
  storageBucket: "click-to-call-d2769.appspot.com",
  messagingSenderId: "241315818421",
  appId: "1:241315818421:web:3923fcd304ec18796007fc",
  measurementId: "G-PSHP0W4DNB"
};
const app = initializeApp(firebaseConfig);
const db = getFirestore(app);


async function getUser(db: any, id: string) {
  const docRef = doc(db, "users", id);
  const docSnap = await getDoc(docRef);

  if (docSnap.exists()) {
    var data = docSnap.data()
    console.log("Document data:", data);
    return data
  } else {
    // docSnap.data() will be undefined in this case
    console.log("No such document!");
  }
  
}

console.log(getUser(db, "UGcFEEx4kUzS7Y90NogY"))


chrome.contextMenus.onClicked.addListener(onClick);

function onClick(info: any) {
  var linkUrl = info.linkUrl;
  var phone = linkUrl.substring(4)
  console.log("Call ", phone)
};

chrome.runtime.onInstalled.addListener(function () {
  chrome.contextMenus.create({
      title: "Click TO Call",
      contexts: ["link"],
      id: "click-to-call"
    });
  }
);