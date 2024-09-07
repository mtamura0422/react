var express = require('express');
var router = express.Router();

var logger = require('../logger');
var util = require('util')


var env    = process.env.NODE_ENV || 'development';
var redis = require('redis');
var redisClient = redis.createClient({port:process.env.EX_REDIS_PORT, host:process.env.EX_REDIS_SERVER});
var DateFormat = require('dateformat');


/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('_admin/index', { title: '管理画面' });
});



/* GET home page. */
router.get('/user/input', function(req, res, next) {
  res.render('_admin/user/input', { title: 'ユーザー作成' });
});


/* GET home page. */
router.post('/user/confirm', function(req, res, next) {
    logger.debug.info('req:'+ util.inspect(req));



  var redisKey = "ex:user:"+req.body.username;


  // redisから編集状況を取得
  redisClient.get(redisKey, function(err, reply) {

    if(err) return console.log(err);
    var data = JSON.parse(reply);

    if (data == null) { data = {}}

    delete data[req.params.table_name];
    redisClient.set(redisKey, JSON.stringify(data), function(){
      res.json({ status: 1 });
    });

  });


  res.render('_admin/user/confirm', { title: 'ユーザー作成' });
});


/* GET home page. */
router.post('/user/create', function(req, res, next) {

  logger.debug.info('req:'+ util.inspect(req));

  res.render('_admin/user/confirm', { title: 'ユーザー作成' });
});



module.exports = router;