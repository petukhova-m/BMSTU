"use strict";
const express = require("express");
const router = express.Router();


router.post('/', (req, res)=>{
    let text = req.body.message;

});


module.exports = router;