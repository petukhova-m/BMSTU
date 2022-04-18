
'use strict'

const express = require("express");
const router = express.Router();
const bodyParser = require('body-parser');
let ObjectID= require('mongodb').ObjectID;

router.post('/', (req, res) => {
    let id=req.body.id
    let newName=req.body.name
    if (newName==undefined) {
      console.log("error in api/changeName.js")
    }
    const collectionUsers = req.app.locals.collectionUsers;
    collectionUsers.updateOne({_id : ObjectID(id)}, {$set: {name : newName}}, (err)=> {
      if (err) {
        console.log(err)
        return res.sendStatus(500)
      }
      return res.sendStatus(200)
    })
})

module.exports = router;
