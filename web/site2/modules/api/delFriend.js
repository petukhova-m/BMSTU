'use strict'

const express = require("express");
const router = express.Router();
const bodyParser = require('body-parser');
let ObjectID= require('mongodb').ObjectID;

router.post('/', (req, res) => {
    let myId = req.body.myId;
    let friendId =req.body.friendId;
    const collection = req.app.locals.collectionUsers;
    collection.updateOne({_id : ObjectID(myId)}, {$pull : {friends : friendId}}, (err) => {
        if(err) return console.log(err);
        res.sendStatus(200);
    });
    collection.updateOne({_id : ObjectID(friendId)}, {$pull : {friends : myId}}, (err) => {
        if(err) return console.log(err);
        res.sendStatus(200);
    });

})

module.exports = router;
