"use strict"
const express = require("express");
const router = express.Router();
const bodyParser = require('body-parser');


router.post('/',async (req, res)=> {
  console.log("SEC");
  let user = {
    name: "DefaultName",
    icon: -1,
    username: req.body.username,
    password: req.body.password,
    friends : []
  };
  console.log(123);
  const collection = req.app.locals.collectionUsers;
  const collectionChat= req.app.locals.collectionChat;
  let userf = await collection.findOne({username : user.username})
  console.log(userf)
  if (userf!=undefined) {
    console.log(500)
    return res.sendStatus(500)
  }
  collection.insertOne(user, (err, result) => {
    if (err) {
      console.log(err);
      return res.sendStatus(500);
    }
    console.log(`user ${user.username} добавлен в БД`);

  let id
  collection.findOne({username: user.username}, (err, doc) => {
    if (err) {
      console.log(err);
      return res.sendStatus(500);
    }
    id=doc._id
    console.log(id);
    let userChat = {
      _id: id,
      login: user.username,
      messages :[]
    }
    collectionChat.insertOne(userChat, (err, result)=> {
      if (err) {
        console.log(err);
        return res.sendStatus(500);
      }
      console.log(`user ${user.username} добавлен в БД chat`);
      return res.sendStatus(200);
    })
  })
})

})
module.exports = router;
