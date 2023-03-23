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
    const value = await redisClient.get('hello');
    console.log(value);
})()

const RateLimiter = async function(req, res, next) {
    const counter = await redisClient.get('counter');
    if (counter > 10) {
        console.log("Rate access limit");
        res.send("Rate access limit, try again later")
        return;
    } else {
        console.log("Allow access");
        await redisClient.set('counter', counter + 1);
        console.log(counter);
    }
    next();
}

module.exports = RateLimiter;