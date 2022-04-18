
'use strict'

const express = require("express");
const router = express.Router();
const bodyParser = require('body-parser');
let ObjectID= require('mongodb').ObjectID;

router.post('/', (req, res) => {
    let from=req.body.from
    let to=req.body.to
    let date=req.body.date
    const collectionChat = req.app.locals.collectionChat;
    collectionChat.update({_id : ObjectID(from)}, {$pull: {messages: {date: date}}}, (err)=> {
      if (err)  {
        console.log(err)
        return res.sendStatus(500)
      }
    })
    collectionChat.update({_id : ObjectID(to)}, {$pull: {messages: {date: date}}}, (err)=> {
      if (err) {
        console.log(err)
        return res.sendStatus(500)
      }
    })
    return res.sendStatus(200)
})

module.exports = router;
