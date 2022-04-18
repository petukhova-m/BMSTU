'use strict'


import api from './api.js';



registration_button.onclick = async () =>  {
    let login = document.getElementById("login").value
    let password = document.getElementById("password").value;
    let password2 = document.getElementById("password_2").value;
    if (password === password2 && login.length != 0 ){
      console.log(0);
        api.register(login, password)
        //.then(()=> {
          document.location.href="http://localhost:3000/suc";
        //}) 
            //alert("Okey");
            
    }else{
       alert("Пароль не совпадает");
    }
}
