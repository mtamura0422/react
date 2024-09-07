var log4js = require('log4js');
var logger = exports = module.exports = {};
log4js.configure({
  appenders: { debug: { type: 'file', filename: 'logs/debug.log' } },
  categories: { default: { appenders: ['debug'], level: 'debug' } }
});


logger.debug = log4js.getLogger('debug');