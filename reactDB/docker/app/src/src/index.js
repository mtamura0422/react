//var React = require('react');
//var ReactDOM = require('react-dom');
//var rrd = require('react-router-dom');
//var BrowserRouter = rrd.BrowserRouter;
//var Route = rrd.Route;
//var Link = rrd.Link;
//var MasterList = require('./master/list');

import React from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter, Switch, Route, Link, Redirect} from 'react-router-dom';
import Login from './firebase/login';
import MasterList from './master/list';



const session = require('express-session');



class MasterApp extends React.Component {

  render() {
    return (
      <BrowserRouter>

        <div>
          <Login exact path="/master" component={MasterList} />
        </div>

      </BrowserRouter>
    );
  }
}

class Home extends React.Component {
  render() {
    return (
      <div>
        <h2>Home</h2>
      </div>
    );
  }
}

/*

// DOMのレンダリング処理
//   see. https://reactjs.org/docs/react-dom.html#render
ReactDOM.render(
  <MasterApp />,            // Appをレンダリングする
  document.getElementById('root')  // id=root要素に対してレンダリングする
);

*/


ReactDOM.render(
  <MasterApp />,            // Appをレンダリングする
  document.getElementById('root')  // id=root要素に対してレンダリングする
);