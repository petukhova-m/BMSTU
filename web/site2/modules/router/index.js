"use strict";
const express = require("express");
const router = express.Router();
const routerApi = require('./routerApi');
const bodyParser = require("body-parser");

router.use(express.static('public'));
router.use(bodyParser.json());

router.get('/', (req, res, next) => res.sendfile("public/html/index.html"));

router.get('/reg', (req, res, next) => res.sendfile("public/html/register.html"));

router.get('/chat', (req, res, next) => res.sendfile("public/html/chat.html"));

router.get('/suc', (req, res, next) => res.sendfile("public/html/success.html"));

router.use('/api', routerApi);

module.exports = router;
