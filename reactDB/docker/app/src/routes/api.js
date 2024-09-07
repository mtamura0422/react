var express = require('express');
var router = express.Router();


var db = require('../models/db');
var env    = process.env.NODE_ENV || 'development';
//var config = require(__dirname + '/../config/config.json')[env];
var redis = require('redis');
var redisClient = redis.createClient({port:process.env.EX_REDIS_PORT, host:process.env.EX_REDIS_SERVER});
var DateFormat = require('dateformat');


const redisKey = "ex:editor";


router.get('/master/', function(req, res, next) {
  console.log('controller index');

  db.sequelize.getQueryInterface().showAllSchemas().then((tableObj) => {

    var tables = [];
    var masters = {}

    tableObj.forEach(function(val, key) {

      var table_name = val["Tables_in_"+process.env.EX_MYSQL_DATABASE];

      if (table_name.match(/^master_/)) {

        tables.push(table_name);

        console.log("val.Tables_in_to_development", val.Tables_in_to_development);
        console.log("table_name", table_name);
      }
    });

    // redisから編集状況を取得
    redisClient.get(redisKey, function(err, reply) {

      if(err) return console.log(err);
      var redisData = JSON.parse(reply);
      console.log(redisData);

      tables.map(function(v) {
        if (redisData && redisData[v]) {
          masters[v] = redisData[v];
          masters[v]["is_edit"] = true
        } else {
          masters[v] = {is_edit: false}
        }
      });

      console.log(masters);
      res.json({ tables: masters })

    });
  })
  .catch((err) => {
    console.log('showAllSchemas ERROR',err);
  });
});



router.get('/master/edit/:table_name/:username', function(req, res, next) {

  console.log('controller edit');
  console.log('table name', req.params.table_name);
  console.log('table name', req.params.username);

  var status = 1;
  var now = new Date();
  edit_time = DateFormat(now, "yyyy/mm/dd HH:MM:ss");



  // 編集者情報
  editInfo = {
    editor: req.params.username,
    edit_time: edit_time,
  }

  // redisから編集状況を取得
  redisClient.get(redisKey, function(err, reply) {

    if(err) return console.log(err);
    var data = JSON.parse(reply);

    if (data == null) { data = {}}

    data[req.params.table_name] = editInfo;
    redisClient.set(redisKey, JSON.stringify(data), function(){
      res.json({ status: status });
    });

  });
});


router.get('/master/delete/:table_name', function(req, res, next) {
  var tables = []
  console.log('controller delete');
  console.log('table name', req.params.table_name);

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
});


module.exports = router;