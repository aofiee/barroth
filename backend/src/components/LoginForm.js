import React, { Component } from 'react';
import firebase from 'firebase/app';
import 'firebase/auth';
import swal from 'sweetalert';
import firebaseConfig from '../config/firebase';
import '../styles/LoginForm/google_button.css';

if (!firebase.apps.length) {
  firebase.initializeApp(firebaseConfig);
} else {
  firebase.app();
}

export class LoginForm extends Component {
  constructor(props) {
    super(props);
    this.title = 'Sign In';
    this.desc = 'with Firebase Authentication';
    this.FacebookSignInAction = this.FacebookSignInAction.bind(this);
    this.AppleSignInAction = this.AppleSignInAction.bind(this);
    this.GoogleSigninAction = this.GoogleSigninAction.bind(this);
  }

  RegisterUser = async (currentUser) => {
    const res = await fetch(`http://${process.env.REACT_APP_HOST}:${process.env.REACT_APP_PORT}/register`, {
      mode: 'cors',
      cache: 'no-cache',
      credentials: 'same-origin',
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        email: currentUser.email,
        uuid: currentUser.uid,
        provider: currentUser.providerId,
        display_name: currentUser.displayName,
        photo_url: currentUser.photoURL,
      }),
    });
    const body = await res.json();
    if (body.error === null) {
      swal({
        title: 'Good job!',
        text: body.msg,
        icon: 'success',
        button: 'ok',
      });
    } else {
      swal({
        title: body.msg,
        text: body.error,
        icon: 'error',
        button: 'ok',
      });
    }
  }

  GoogleSigninAction = () => {
    const provider = new firebase.auth.GoogleAuthProvider();
    firebase.auth().languageCode = 'th';
    provider.setCustomParameters({
      login_hint: 'user@example.com',
    });
    firebase.auth()
      .signInWithPopup(provider)
      .then((result) => {
        // const { credential } = result;
        // const token = credential.accessToken;
        const { user } = result;
        console.log('user', user);
        this.RegisterUser(user);
      }).catch((error) => {
        const errorCode = error.code;
        // const errorMessage = error.message;
        // const { email } = error;
        // const { credential } = error;
        if (errorCode === 'auth/account-exists-with-different-credential') {
          firebase.auth().currentUser.linkWithPopup(provider).then((result) => {
            // const c = result.credential;
            const u = result.user;
            this.RegisterUser(u);
          }).catch((e) => {
            console.log('error.code', e.code);
          });
        }
      });
  }

  FacebookSignInAction = () => {
    const provider = new firebase.auth.FacebookAuthProvider();
    provider.addScope('user_birthday');
    provider.addScope('public_profile');
    provider.addScope('email');
    firebase.auth().languageCode = 'th';
    provider.setCustomParameters({
      display: 'popup',
    });
    firebase
      .auth()
      .signInWithPopup(provider)
      .then((result) => {
        const { user } = result;
        console.log('user', user);
        this.RegisterUser(user);
      })
      .catch((error) => {
        const errorCode = error.code;
        if (errorCode === 'auth/account-exists-with-different-credential') {
          firebase.auth().currentUser.linkWithPopup(provider).then((result) => {
            const u = result.user;
            console.log('user', u);
            this.RegisterUser(u);
          }).catch((e) => {
            console.log('error.code', e.code);
          });
        }
      });
  }

  AppleSignInAction = () => {
    const provider = new firebase.auth.OAuthProvider('apple.com');
    provider.addScope('email');
    provider.addScope('name');
    provider.setCustomParameters({
      locale: 'th',
    });
    firebase
      .auth()
      .signInWithPopup(provider)
      .then((result) => {
        const { user } = result;
        console.log('user', user);
        this.RegisterUser(user);
      })
      .catch((error) => {
        const errorCode = error.code;
        if (errorCode === 'auth/account-exists-with-different-credential') {
          firebase.auth().currentUser.linkWithPopup(provider).then((result) => {
            const u = result.user;
            console.log('user', u);
            this.RegisterUser(u);
          }).catch((e) => {
            console.log('error.code', e.code);
          });
        }
      });
  }

  render() {
    return (
      <div className="container">
        <h1>{this.title}</h1>
        <h1>{this.desc}</h1>

        <div className="col">
          <button type="button" className="loginBtn loginBtn--apple" onClick={this.AppleSignInAction}>
            Sign In with Apple
          </button>
          <br />
          <button type="button" className="loginBtn loginBtn--facebook" onClick={this.FacebookSignInAction}>
            Sign In with Facebook
          </button>
          <br />
          <button type="button" className="loginBtn loginBtn--google" onClick={this.GoogleSigninAction}>
            Sign In with Google
          </button>
          <br />
        </div>
        {/* <button type="button" onClick={() => firebase.auth().signOut()}>Sign-out</button> */}
      </div>
    );
  }
}
export default LoginForm;
