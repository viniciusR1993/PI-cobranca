// ;(function() {
//     console.log("Hello Word!")
//   })()

//captura o formulario
const formLogin = document.getElementById("login");

//Observa o aperta de botação
formLogin.onsubmit = (e) => {
    e.preventDefault() //Evita que a pagina refresh
    const username = formLogin.elements["username"].value //captura o valor do forms
    const password = formLogin.elements["password"].value
    const url = "http://localhost:8000/login"
    const options = {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },  
        mode:'no-cors',
        body: JSON.stringify({
            "Name":username,
            "Password":password
        })
    }

    fetch(url, options)
    .then(response => {
        if (!response.ok) {
            console.log(response)
            throw new Error('Falha na requisição')
        }
        return response.json();
    })
    .then(data => {
        data == 2 ? alert("Email ou Senha Errado!"): window.location.href = "http://localhost:8000/pginicial"
    })
    .catch(err => {
        alert(err)
    })
}