
'use strict'

const express = require("express");
const router = express.Router();
const bodyParser = require('body-parser');
let ObjectID= require('mongodb').ObjectID;

router.post('/', (req, res) => {
    let date=new Date()
    let message = {
      from: req.body.from,
      to: req.body.to,
      date: date,
      cont: req.body.cont
    }
    console.log("cont = " + message.cont)
    const collectionChat = req.app.locals.collectionChat;
    collectionChat.updateOne({_id : ObjectID(message.from)}, {$push: {messages: message}}, (err)=> {
      if (err) {
        console.log(err)
        return res.sendStatus(500)
      }
    })
    collectionChat.updateOne({_id : ObjectID(message.to)}, {$push: {messages: message}}, (err)=> {
      if (err) {
        console.log(err)
        return res.sendStatus(500)
      }
    })
    return res.sendStatus(200)
})

module.exports = router;
