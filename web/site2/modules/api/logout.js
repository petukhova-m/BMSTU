"use strict";
const express = require("express");
const router = express.Router();


router.post('/', (req, res)=>{
    res.send("123");
});


module.exports = router;
