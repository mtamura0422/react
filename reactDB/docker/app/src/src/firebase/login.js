import React, { Component } from 'react'
import firebase from './firebase'
import firebaseui from 'firebaseui'
import 'firebaseui/dist/firebaseui.css';


class Login extends React.Component {

  constructor(props) {

    super(props);

    this.state = {
      loading: true,
      authenticated: false,
      currentUser: null
    };

  }


  componentWillMount() {

    let ui = firebaseui.auth.AuthUI.getInstance();
    if (!ui) {
      ui = new firebaseui.auth.AuthUI(firebase.auth());
    }

    const uiConfig = {
      callbacks: {
        signInSuccessWithAuthResult: (authResult, redirectUrl) => {
          return true;
        },
        uiShown: () => {
          document.getElementById('loader').style.display = 'none';
        }
      },
      signInFlow: 'popup',
      signInSuccessUrl: 'http://localhost:3001',
      signInOptions: [
        firebase.auth.GoogleAuthProvider.PROVIDER_ID,
        firebase.auth.FacebookAuthProvider.PROVIDER_ID,
      ],
      tosUrl: 'http;//localhost:3001'
    };


    firebase.auth().onAuthStateChanged(user => {

console.log("user"+user);



      if (user) {

        this.setState({
          authenticated: true,
          user: user,
          loading: false
        });
      } else {
        this.setState({
          authenticated: false,
          user: null,
          loading: false
        });

        ui.start('#firebaseui-auth-container', uiConfig);
      }
    })


  }

/*
  componentDidMount() {
    let ui = firebaseui.auth.AuthUI.getInstance();
    if (!ui) {
      ui = new firebaseui.auth.AuthUI(firebase.auth());
    }

    const uiConfig = {
      callbacks: {
        signInSuccessWithAuthResult: (authResult, redirectUrl) => {
          return true;
        },
        uiShown: () => {
          document.getElementById('loader').style.display = 'none';
        }
      },
      signInFlow: 'popup',
      signInSuccessUrl: 'http://localhost:3001',
      signInOptions: [
        firebase.auth.GoogleAuthProvider.PROVIDER_ID,
        firebase.auth.FacebookAuthProvider.PROVIDER_ID,
      ],
      tosUrl: 'http;//localhost:3001'
    };
    if (this.state.user == null) {
      ui.start('#firebaseui-auth-container', uiConfig);
    }
  }

*/
  render() {


    var name, email, photoUrl, uid, emailVerified, topText;

    if (this.state.loading) {

      topText = <div>

                  <div id="loader">Now Loading...</div>
                </div>
    } else {

      const { component : Component} = this.props


      if (this.state.user != null) {
        name = this.state.user.displayName;
        email = this.state.user.email;
        photoUrl = this.state.user.photoURL;
        emailVerified = this.state.user.emailVerified;
        uid = this.state.user.uid;

        console.log("email:" + email);

        topText = <div>
                    <p>uid: {uid}</p>
                    <p>: {name}</p>
                    <p>eusermail: {email}</p>
                    <p><button onClick={ () => firebase.auth().signOut() }>sign out</button></p>
                    <Component {...this.props} user={this.state.user} />
                  </div>
      }



    }

    return (
      <div>
        <div id="firebaseui-auth-container" />
        {topText}
      </div>
    );


  /*
    const { authenticated, loading } = this.state;
    if (loading) return <p>loading..</p>;

    var userInfo = null

    if (this.state.user) {
      userInfo = <div>
                  <p>UID : {this.state.user.uid}</p>
                  <p>email : {this.state.user.email}</p>
                  <button onClick={ () => firebase.auth().signOut() }>sign out</button>
                </div>
    } else {

      const provider = new firebase.auth.GoogleAuthProvider()
      userInfo = <div>
                  <button onClick={ () => firebase.auth().signInWithRedirect(provider) }>Login</button>
                </div>
    }

    return (

      <div>
        <h2>test</h2>
         <p className="App-intro">
          {userInfo}
        </p>
      </div>

    )
*/
  }
}


export default Login