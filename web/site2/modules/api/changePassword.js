
'use strict'

const express = require("express");
const router = express.Router();
const bodyParser = require('body-parser');
let ObjectID= require('mongodb').ObjectID;

router.post('/', (req, res) => {
    let id=req.body.id
    let newPassword=req.body.password
    if (newPassword==undefined) {
      console.log("error in api/changePassword.js")
    }
    const collectionUsers = req.app.locals.collectionUsers;
    collectionUsers.updateOne({_id : ObjectID(id)}, {$set: {password : newPassword}}, (err)=> {
      if (err) {
        console.log(err)
        return res.sendStatus(500)
      }
      return res.sendStatus(200)
    })
})

module.exports = router;
