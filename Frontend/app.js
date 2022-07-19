async function getUsers() {
var add=document.getElementById("email").value;
let url = `http://localhost:8080/?email=${add}`;
fetch(url,{
    method:'POST',
    headers:{
        'Content-Type':'application/json',
    },
    body:{},
})
.then( response => response.json() )
.then( data => {
        // Do something with response.
        console.log(data);
        if (data =="Email Verified"){
            document.getElementById("msg").style.color="green";
        }
        else 
        {
            document.getElementById("msg").style.color="red";
        }
        document.getElementById("msg").innerHTML=data;
    } );
    
}