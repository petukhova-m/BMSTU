'use strict'

const express = require('express');
const bodyParser = require("body-parser");
const cookieParser = require("cookie-parser");
const MongoClient= require('mongodb').MongoClient;
const session = require("express-session");
const path = require('path');
const index = require('./modules/router/index')
const http = require('http')
const app = express();
const WebSocket = require( "ws");
const server = http.createServer(app);
const webSocketServer = new WebSocket.Server({ server });

app.use(express.static('public'));



let peers =[];




app.use((req, res, next) => {
    res.setHeader("Access-Control-Allow-Origin", "*"); // update to match the domain you will make the request from
    res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
    res.header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE");
    next();
  });

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: true}));
app.use('/', index);

const mongoClient = new MongoClient("mongodb://localhost:27017/", { useUnifiedTopology: true, useNewUrlParser: true });

let dbClient;

webSocketServer.on('connection', ws => {
  console.log('connect');
  peers.push(ws);
  ws.on('message', message => {
    let peersHelp = peers.filter((lenght) => lenght != ws);
    console.log('------', message);
    const body = JSON.parse(message);
    let password = body.password
    let username = body.login
    console.log("pass = " + password)
    console.log("user = " + username)
    const collectionUsers = app.locals.collectionUsers;
    const collectionChat = app.locals.collectionChat;
    collectionUsers.findOne({username: username, password: password}, (err,doc) => {
    if (err) console.log(err)
    console.log("found in Users " + doc)
    let messages
    collectionChat.findOne({ _id: doc._id }, (err, doc) => {
      if (err) console.log(err)
      console.log(doc)
      messages=doc.messages
      console.log("found in Chat " + messages)
      messages.sort((a,b) => {
        //console.log(a)
        //console.log(b)
        return (a.date < b.date)
      })
      console.log('send');
      ws.send(JSON.stringify({messages: messages, help: true}));
      
    })
  });
  })
})

mongoClient.connect( (err, client) => {
  if (err) {
    console.log(`app.js : ошибка при попытке вызова connect к db\n`);
    return console.log(err);
  }

  dbClient = client;

  app.locals.collectionUsers = client.db("DB").collection("users");

  app.locals.collectionChat = client.db("DB").collection("chat");

  server.listen(3000, ()=> {
    console.log(`api запущен\n`);
  });
})


module.exports = app;
