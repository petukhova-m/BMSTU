"use strict"

const express = require("express");
const router = express.Router();
const bodyParser = require('body-parser');

router.post('/', (req,res)=>{
    const collection = req.app.locals.collectionUsers;
    collection.findOne({ username: req.body.username, password: req.body.password}, (err, doc) =>{
      if (err) {
        console.log(err);
        return res.sendStatus(500);
      }
      else {
        if (doc) {
          console.log(`user ${JSON.stringify(req.body.username)} найден` );
          res.json({isKnownUser: true});
          //document.location.href="http://localhost:3000/chat";
        }
        else {
          console.log(`user ${JSON.stringify(req.body.username)} не найден` );
          //res.redirect("public/html/incorrect.html");
          res.json({isKnownUser: false});
        }
      }
    });
  });

module.exports = router;
