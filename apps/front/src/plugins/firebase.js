import { initializeApp } from "firebase/app";
// import { getAnalytics } from "firebase/analytics";

const firebaseConfig = {
  apiKey: "AIzaSyBXg5nEwBKWVaRO-xYbC74iw5i1OyUDnfY",
  authDomain: "anan-project.firebaseapp.com",
  databaseURL: "https://anan-project.firebaseio.com",
  projectId: "anan-project",
  storageBucket: "anan-project.appspot.com",
  messagingSenderId: "474794270878",
  appId: "1:474794270878:web:ff0a48e2c095d2722e5b8f",
  measurementId: "G-J3JEBQSB29"
};

const app = initializeApp(firebaseConfig);

export default app
