var express = require('express');
var path = require('path');
var cookieParser = require('cookie-parser');
var morgan = require('morgan');
var bodyParser = require('body-parser');
var logger = require('./logger');
var util = require('util')



//import {createError} from 'http-errors'
//import express from 'express'
//import path from 'path'
//import cookieParser from 'cookie-parser'
//import morgan from 'morgan'
//import bodyParser from 'body-parser'
//import logger from './logger'
//import util from 'util'


//var methodOverride = require('method-override');



var api = require('./routes/api');
var auth = require('./routes/auth');
var admin = require('./routes/admin');



var app = express();

var port = process.env.PORT || 3001


//passportの初期化
var passport = require('passport')
var LocalStrategy = require('passport-local').Strategy;
app.use(require('express-session')({
    secret: 'fjaiofjfiafkldsfkadjkafk', //セッションのハッシュ文字列。任意に変更すること。
    resave: false,
    saveUninitialized: false
}));
app.use(passport.initialize());
app.use(passport.session());

//認証ロジック
passport.use(
  new LocalStrategy(
    function(username, password, done) {
      //routerのpassport.authenticate()が呼ばれたらここの処理が走る。
logger.debug.info('username:'+ username);
logger.debug.info('password:'+ password);


      // DBから取得する
      if(username == 'myname' && password == 'mypassword'){
        //ログイン成功
        //今回はここの判定式をハードコーディングにする
            return done(null, username);
      }

      //ログイン失敗
      //messageはログイン失敗時のフラッシュメッセージ。
      //各routerの req.flash() で取得できる。
          return done(null, false, {message:'ID or Passwordが間違っています。'});
    }
  )
);


//認証した際のオブジェクトをシリアライズしてセッションに保存する。
passport.serializeUser(function(username, done) {
  console.log('serializeUser');
  done(null, username);
});


//認証時にシリアライズしてセッションに保存したオブジェクトをデシリアライズする。
//デシリアライズしたオブジェクトは各routerの req.user で参照できる。
passport.deserializeUser(function(username, done) {
  console.log('deserializeUser');
  done(null, {name:username, msg:'my message'});
});



// ログインチェック
// auth以下は次の処理へ。
//app.get('*', function(req, res, next) {


// logger.debug.info('url:'+ decodeURI(req.url));

//  if (decodeURI(req.url).match(/^\/_auth*/) || req.isAuthenticated()) {  // 認証済 or ログイン画面
//    return next();
//  }
//  else {  // 認証されていない
//    res.redirect('/_auth/');  // ログイン画面に遷移
//  }
//});




//---------ここまでpassport-----------



app.use(morgan('dev'));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(cookieParser());



app.use('/_api/', api);
app.use('/_auth/', auth);
app.use('/_admin/', admin);



// override HTTP method to using CRUD
//app.use(methodOverride('_method'));


// allow CORS
//   see. https://developer.mozilla.org/ja/docs/Web/HTTP/HTTP_access_control
app.use(function(req, res, next) {
  res.header("Access-Control-Allow-Origin", "*");
  res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
  next();
});




app.set('view engine', 'pug')
app.set('views', path.join(__dirname, './views'));

// correspond Express and React
app.use(express.static(path.join(__dirname, 'build')));


//app.get('/*', function(req, res) {
//  res.sendFile(path.join(__dirname, 'build', 'index.html'));
//});






// catch 404 and forward to error handler
app.use(function(req, res, next) {
  var err = new Error('Not Found');
  err.status = 404;
  next(err);
});



// error handler
app.use(function(err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render('error');
});


//app.get('*', (req, res) => {
//  return handle(req, res);
//})

app.listen(port, (err) => {
  if (err) throw err
  console.log(`> Ready on http://localhost:${port}`)
})



module.exports = app;
