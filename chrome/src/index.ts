import { initializeApp, setLogLevel as firebaseSetLogLevel } from 'firebase/app';
import { getFirestore, doc, getDoc, setDoc } from 'firebase/firestore/lite';


async function main() {
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
  const userId = "UGcFEEx4kUzS7Y90NogY";

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

  
  
  async function onClick(info: any) {
    debugger
    var user = doc(db, "users", userId)
    console.log("user:", user)
    var linkUrl = info.linkUrl;
    var phone = linkUrl.substring(4)
    setDoc(user, {call: phone}, {merge: true});
    
  };

  chrome.contextMenus.onClicked.addListener(onClick);

  chrome.runtime.onInstalled.addListener(function () {
    console.log('creating menu context')
    chrome.contextMenus.create({
        title: "Click TO Call",
        contexts: ["link"],
        id: "click-to-call"
      });
    }
  );
}
main()