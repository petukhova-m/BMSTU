
'use strict'

const express = require("express");
const router = express.Router();
const bodyParser = require('body-parser');
let ObjectID= require('mongodb').ObjectID;

router.post('/', async (req, res) => {
    let id = req.body.id;
    console.log(req.body.id)
    const collection = req.app.locals.collectionUsers;
    console.log(321);
    collection.findOne({_id : ObjectID(id)}, async(err, obj) => {
        if(err) return console.log(err);
        console.log(obj.id);
        let arrOfFriends = new Array();
        const abc = async()=> {
        for (let i=0; i<obj.friends.length; i++) {
          console.log(i);
          let friendId=obj.friends[i]
          let friend = await collection.findOne({_id : ObjectID(friendId)})
          delete friend.password
          delete friend.username
          arrOfFriends.push(friend);
         }
       }
      await abc()
      res.json(arrOfFriends)
    });
})

module.exports = router;
