
'use strict'

const express = require("express");
const router = express.Router();
const bodyParser = require('body-parser');


router.get('/', (req, res) => {
    const collection = req.app.locals.collectionUsers;
    collection.find().toArray(function(err, users){

        if(err) return console.log(err);
        for (let i=0; i<users.length; i++) {
          delete users[i].password
          delete users[i].username
        }
        res.json(users)
    });
})

module.exports = router;
