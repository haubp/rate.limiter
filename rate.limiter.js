const redis = require('redis');
const fs = require("fs");

let configRawStr = fs.readFileSync(__dirname + "/config.json");
let config = JSON.parse(configRawStr);

const redisClient = redis.createClient({
    socket: {
        host: config["server"],
        port: config["port"]
    }
});

redisClient.on('error', err => {
    console.log('Error ' + err);
});

(async function() {
    await redisClient.connect()
    await redisClient.set('hello', 'hello world');
    const value = await redisClient.get('hello');
    console.log(value);
    await redisClient.disconnect();
})()

const RateLimiter = function(req, res, next) {
    next();
}

module.exports = RateLimiter;