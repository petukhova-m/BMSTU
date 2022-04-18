
'use strict'

import api from './api.js';

const wsSend = (data) => {
  if(!wsConnection.readyState){
    setTimeout(() => {
      wsSend(data);
    }, 100);
  } else {
    wsConnection.send(data);
  }
};

function getCookie(name) {
    let matches = document.cookie.match(new RegExp(
      "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
    ));
    return matches ? decodeURIComponent(matches[1]) : "";
  }



let users =[]

let messages = []

let pickedUser = 0

let currentUser = 0

document.getElementById("help").onclick = () =>{
  if (document.getElementById('search_friends').style.display == "block"){
    document.getElementById('search_friends').style.display = "none"
  }else{
    document.getElementById('search_friends').style.display = "block"
  }
}


document.getElementById("menu_1").onclick = () =>{
  if (document.getElementById('settings_menu').style.display == "block"){
    document.getElementById('settings_menu').style.display = "none"
  }else{
    document.getElementById('settings_menu').style.display = "block"
  }
}

submit_change.onclick = ()=>{
  let password1 = document.getElementById("change_1").value;
  let password2 = document.getElementById("change_2").value;
  console.log(password1)
  if(password1 === password2){
    api.changePassword(currentUser, password1);
    console.log("changed");
  }
}

submit_change_name.onclick = ()=>{
 let name = document.getElementById("change_name").value
 api.changeName(currentUser, name)
}


document.getElementById('logout').onclick = () => {
  document.cookie=`login=`;
  document.cookie=`password=`;
  document.cookie =`session=`;
  document.location.href = "http://localhost:3000";
}


const changeChat = ()=>{

}

const chatOnload =()=>{

}

let direct_id = "";

let currentUserName = ""

let friends = []

const url = 'ws://localhost:3000';
const connection = new WebSocket(url);

let helpId = null;

connection.onopen = () => {
  console.log('connect');
  const msg = {
    login: getCookie('login'),
    password: getCookie("password")
  }
  connection.send(JSON.stringify(msg));
}

connection.onmessage = e => {
  try {
    //console.log(e);
    console.log('Получил');
    let obj = JSON.parse(e.data);
    console.log(obj)
    messages = obj.messages;
    if(obj.help){
      if (direct_id){
        reloadfriends();
        generateChat(direct_id);
      } else {
        reloadfriends();
    }}else{
      const msg = {
        login: getCookie('login'),
        password: getCookie("password")
      }
      connection.send(JSON.stringify(msg));
    }
  
  }catch(e) {
    console.log(e);
  }
}

const reloadfriends = async ()=>{

  document.getElementById('friends_list').innerHTML = '';
  await friends.forEach(a => {
    let user = document.createElement('div');
    user.id = "user_in_list";
    let nameOfuser = document.createElement('p');
    nameOfuser.innerText = a.name;
    nameOfuser.id = a._id;
    helpId = a._id;
    user.onclick =() =>{
      generateChat(a._id)
    }
    user.appendChild(nameOfuser);
    document.getElementById("friends_list").appendChild(user);
    if (a._id == currentUser){
      document.getElementById("logged_as_1").innerText = "Logged as " + a.name;
    }
  })
}



const generateChat = async (id) =>{
  document.getElementById("textarea").innerHTML = ""
  let dialog =[]
  direct_id= id;
  messages.forEach(a => {
    if ((a.to === id && a.from === currentUser)|| (a.from === id && a.to === currentUser)){
      dialog.push(a)
    }
  })
  dialog.forEach( a=>{
    let divOneMessage =document.createElement('div');
    if(currentUser === a.to){
      divOneMessage.id = 'div_one_Message_left';
    } else{
      divOneMessage.id = 'div_one_Message_right';
    }
    console.log(a)
    let text = document.createElement('p');
    text.innerText = a.cont;
    let data = document.createElement('p');
    let x = new Date(a.date);
    console.log(a.date)
    data.innerText = x.getHours() + ":" + x.getMinutes();
    divOneMessage.appendChild(text);
    divOneMessage.appendChild(data);
    document.getElementById("textarea").appendChild(divOneMessage);
  })
}



send_message.onclick =()=>{
    const text = document.getElementById('text').value;
    document.getElementById('text').value = "";
    api.send(currentUser, direct_id, text);
    const msg = {
      login: getCookie('login'),
      password: getCookie("password")
    }
    connection.send(JSON.stringify(msg));
}




window.onload = async ()=>{
    if(getCookie('session') === ""){
        document.location.href = "http://localhost:3000";
    }else{
    currentUser = await api.getUserId(getCookie('login'), getCookie('password'));
    friends = await api.getFriends(currentUser);
    let i = 0;
    console.log(friends)


    // Добавляю список людей на добавления в друзья
    api.getAllUsers().then(a =>  a.forEach(elem => {

      if (elem._id !== currentUser){
          let divMessage = document.createElement('div');
          divMessage.className = 'messages';
          let divUserName = document.createElement('div');
          divUserName.className = 'userName';
          let divDate = document.createElement('div');
          divDate.className = 'date_time';
          let divText = document.createElement('div');
          divText.className = 'text_of_message';
          divText.innerText = elem.name
          let divicon = document.createElement('div');
          divicon.className = 'icon';
          let buttonAddFriend = document.createElement('button');
          buttonAddFriend.className = 'buttonAddFriend';
          buttonAddFriend.innerText = 'Add';
          buttonAddFriend.name = i;
          buttonAddFriend.onclick = () =>{
            api.getFriends(currentUser).then(z =>{if(!z.includes(currentUser)){api.addFriend(currentUser, elem._id)}});
          }

          divMessage.appendChild(divicon);
          divMessage.appendChild(divText);
          divMessage.appendChild(divDate);
          divMessage.appendChild(divUserName);
          divMessage.appendChild(buttonAddFriend);
          document.getElementById("Friends").appendChild(divMessage);
          }
      reloadfriends()
        }))};
  //messages = await api.getMessages(getCookie('login'), getCookie("password"));
  document.getElementById("settings_menu_name").innerText = getCookie("login");
}

