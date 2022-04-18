'use strict'


import api from './api.js';



function getCookie(name) {
  let matches = document.cookie.match(new RegExp(
    "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
  ));
  return matches ? decodeURIComponent(matches[1]) : "";
}

window.onload = ()=>{
  if(getCookie('session') !== ""){
    document.location.href = "http://localhost:3000/chat";
};
}


document.getElementById("login").value=getCookie("login");
document.getElementById("password").value=getCookie("password");
if (getCookie("login")!=""&& getCookie("password")) {
  api.auth(getCookie("login"), getCookie("password")).then((res)=> {
    if(res.isKnownUser=== true)
        {
          document.location.href="http://localhost:3000/chat";
        }
  })
}

login_button.onclick = () => {
    let name = document.getElementById("login").value;
    let password = document.getElementById("password").value;
    api.auth(name, password)
    .then((res)=>{
        if(res.isKnownUser=== true)
            {
              document.cookie=`login=${name}`;
              document.cookie=`password=${password}`;
              document.cookie =`session=${true}`;
              document.location.href="http://localhost:3000/chat";
            }
        else{
            document.getElementById("incorrect").hidden = false;
        }})};


iAmNew.onclick = () => {document.location.href="http://localhost:3000/reg"};
