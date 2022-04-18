const express = require("express");
const router = express.Router();

const auth = require("../api/auth");

const registration = require("../api/registration");

const send = require("../api/send");

const logout = require("../api/logout");



const user = require("../api/user");

const getAllUsers = require("../api/getAllUsers");

const getMessages = require("../api/getMessages");

const addMessage = require("../api/addMessage");

const delMessage = require("../api/delMessage");

const changeName = require("../api/changeName");

const changeIcon = require("../api/changeIcon");

const changePassword = require("../api/changePassword")

const getFriends = require("../api/getFriends")

const addFriend = require("../api/addFriend")

const getUserId = require("../api/getUserId")

const delFriend = require("../api/delFriend")

router.get('/', ()=>{})

router.use('/auth', auth);

router.use('/registration', registration);


router.use('/logout', logout);

router.use('/user', user);

router.use('/getAllUsers', getAllUsers);

router.use('/getMessages', getMessages);

router.use('/addMessage', addMessage);

router.use('/delMessage', delMessage);

router.use('/changeName', changeName);

router.use('/changePassword', changePassword);

router.use('/changeIcon', changeIcon);

router.use('/getFriends', getFriends);

router.use('/addFriend', addFriend);

router.use('/getUserId', getUserId);

router.use('/delFriend', delFriend);

module.exports = router;
