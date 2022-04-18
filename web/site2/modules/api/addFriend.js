
'use strict'

const express = require("express");
const router = express.Router();
const bodyParser = require('body-parser');
let ObjectID= require('mongodb').ObjectID;

router.post('/', async (req, res) => {
    let myId = req.body.myId;
    let friendId =req.body.friendId;
    const collection = req.app.locals.collectionUsers;

    let obj = await collection.findOne({_id : ObjectID(myId), friends : friendId})
    if (obj==undefined) {
      await collection.updateOne({_id : ObjectID(myId)}, {$push : {friends : friendId}})
    }
    obj = await collection.findOne({_id : ObjectID(friendId), friends : myId})
    if (obj==undefined) {
      await collection.updateOne({_id : ObjectID(friendId)}, {$push : {friends : myId}})
    }
    return res.sendStatus(200)
})

module.exports = router;
