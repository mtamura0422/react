'use strict';

var fs        = require('fs');
var path      = require('path');
var Sequelize = require('sequelize');
var basename  = path.basename(__filename);
var env       = process.env.NODE_ENV || 'development';
//var config    = require(__dirname + '/../config/config.json')[env];
var db        = {};

process.env.EX_MYSQL_USER = "to"
process.env.EX_MYSQL_PASS = "Ac7PPsLG6H"
process.env.EX_MYSQL_DATABASE = "to_development"
process.env.EX_MYSQL_SERVER = "localhost"

var sequelize = new Sequelize({
  username: process.env.EX_MYSQL_USER,
  password: process.env.EX_MYSQL_PASS,
  database: process.env.EX_MYSQL_DATABASE,
  host: process.env.EX_MYSQL_SERVER,
  dialect: "mysql",
  operatorsAliases: false
});


fs
  .readdirSync(__dirname)
  .filter(file => {
    return (file.indexOf('.') !== 0) && (file !== basename) && (file.slice(-3) === '.js');
  })
  .forEach(file => {
    var model = sequelize['import'](path.join(__dirname, file));
    db[model.name] = model;
  });

Object.keys(db).forEach(modelName => {
  if (db[modelName].associate) {
    db[modelName].associate(db);
  }
});

db.sequelize = sequelize;
db.Sequelize = Sequelize;

module.exports = db;
