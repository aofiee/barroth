import React, { Component } from 'react';
import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth';
import firebase from 'firebase';
import firebaseConfig from '../config/firebase';

firebase.initializeApp(firebaseConfig);
export class LoginForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      isLogin: false,
    };
    this.title = 'FirebaseUI';
    this.desc = 'with Firebase Authentication';
  }

  componentDidMount() {
    this.unregisterAuthObserver = firebase
      .auth()
      .onAuthStateChanged((user) => this.setState({ isLogin: !!user }));
  }

  componentWillUnmount() {
    this.unregisterAuthObserver();
  }

  render() {
    const uiConfig = {
      signInFlow: 'popup',
      signInOptions: [
        firebase.auth.GoogleAuthProvider.PROVIDER_ID,
        firebase.auth.FacebookAuthProvider.PROVIDER_ID,
        'apple.com',
      ],
      callbacks: {
        signInSuccess: () => false,
      },
    };
    const { isLogin } = this.state;
    if (!isLogin) {
      return (
        <div className="container">
          <h1>{this.title}</h1>
          <h1>{this.desc}</h1>
          <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={firebase.auth()} />
        </div>
      );
    }
    return (
      <div className="container">
        <h1>{this.title}</h1>
        <h1>{this.desc}</h1>
        <img alt="profile" className="p" src={firebase.auth().currentUser.photoURL} />
        <br />
        {firebase.auth().currentUser.displayName}
        <br />
        {firebase.auth().currentUser.email}
        <br />
        {firebase.auth().currentUser.phoneNumber}
        <br />
        {firebase.auth().currentUser.providerId}
        <br />
        {firebase.auth().currentUser.uid}
        <br />
        <button type="button" onClick={() => firebase.auth().signOut()}>Sign-out</button>
      </div>
    );
  }
}
export default LoginForm;
