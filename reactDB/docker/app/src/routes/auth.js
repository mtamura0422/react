var express = require('express');
var router = express.Router();

var passport = require('passport');
var logger = require('../logger');
var util = require('util')


/* GET home page. */
router.get('/', function(req, res, next) {

  logger.debug.info('url:'+ decodeURI(req.url));
  //ログイン時のエラーメッセージ表示
  var errorMsg = '';
//  var error = req.flash().error;
//  if(error){
//    var errorMsg = error[0];
//  }

  res.render('_auth/login', { title: 'Express', msg: errorMsg });
});


//ログイン処理
//badRequestMessage は入力値が空の場合にflashにセットされるメッセージ。
//つまり、入力値が空かどうかはpassport側でチェックしてくれているので、それ用のバリデーションは不要。
router.post(
  '/login',
  passport.authenticate( 'local', { successRedirect: '/', failureRedirect: '/_auth/', failureFlash: false, badRequestMessage: '入力値が空です' }),
  function(req, res, next) {
    //認証成功した場合のコールバック
    //成功時のリダイレクトは successRedirect で指定しているので、ここですることは特にないかもしれない。
  }
);



//ログアウト
router.get('/logout', function(req, res, next) {
  req.logout();
  res.redirect('/');
});

module.exports = router;