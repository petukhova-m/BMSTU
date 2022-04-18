'use strict'



const baseURL = "http://localhost:3000/api";


class Api{

    async auth(username, password) {
        let data = "failed";
        console.log(98);
        //alert(123);
        await fetch(`${baseURL}/auth`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
              },
            body: JSON.stringify({
                "username": username,
                "password": password
            })
        }).then(response => {
             data = response.json();
             //return data;
        }).catch(err => {console.log(err)})
        return data;
      }


    async register(username, password){
        let data = "failed";
        console.log("FIR");
        await fetch(`${baseURL}/registration`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
              },
            body: JSON.stringify({
                "username": username,
                "password": password
            })
        }).then(response => {
             data = response.json();
        }).catch(err => {console.log(err)})


        return data;
      }

    async send(name, to, message){

        fetch(`${baseURL}/addMessage`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
              },
            body: JSON.stringify({
                "from": name,
                "to": to,
                "cont": message
            })
        })
        //return await data.json();
    }


    async getAllUsers(){
        let data = "failed";
        data = await fetch(`${baseURL}/getAllUsers`,{
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
              },
        })
       return await data.json();
    }

    async getUserId(username, password) {
        let data = "failed";
        //alert(123);
        data = await fetch(`${baseURL}/getUserId`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
              },
            body: JSON.stringify({
                "username": username,
                "password": password
            })
        })
        return await data.json()
    }
    async changeName(id, name) {
        let data = "failed";
        //alert(123);
        await fetch(`${baseURL}/changeName`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
              },
            body: JSON.stringify({
                "id": id,
                "name": name
            })
        }).catch(err => {console.log(err)})
        return data;
      }

    async changePassword(id, password) {
        let data = "failed";
        //alert(123);
        await fetch(`${baseURL}/changePassword`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
              },
            body: JSON.stringify({
                "id": id,
                "password": password
            })
        }).catch(err => {console.log(err)})
        return data;
      }


    async getMessages(login, password){
        let data = "failed";
        data = await fetch(`${baseURL}/getMessages`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
              },
            body: JSON.stringify({
                "username": login,
                "password": password
            })
        })
        return await data.json();}

    async addFriend(myid, id){
        await fetch(`${baseURL}/addFriend`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
              },
            body: JSON.stringify({
                "myId": myid,
                "friendId": id
            })
        })
      }

      async getFriends(id) {
        let data = "failed";
        //alert(123);
        data = await fetch(`${baseURL}/getFriends`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
              },
            body: JSON.stringify({
                "id": id,
            })
        })
        return await data.json()
    }
}


export default new Api();
