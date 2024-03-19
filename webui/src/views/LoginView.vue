<script>
/**

Ha alcuni dati, tra cui un messaggio di errore, un identificatore di utente e un flag per l'abilitazione o meno del pulsante di login. 
C'è anche un metodo login che viene chiamato quando l'utente fa clic sul pulsante di login.

La funzione login effettua una richiesta HTTP POST a "/session" con l'identificatore dell'utente.
Se la richiesta riesce, il token dell'utente viene memorizzato nel local storage e l'utente viene reindirizzato alla pagina "home". 
Se c'è un errore nella richiesta, il messaggio di errore viene memorizzato nella proprietà errormsg.

Quando il componente viene montato, viene controllato se esiste un token nell'local storage.
Se esiste, l'utente viene reindirizzato alla pagina "home". 
Questo è per garantire che l'utente non possa tornare alla schermata di login se è già loggato.

 */
export default {
	data: function() {
		return {
			errormsg: null,
			nickname: "",
			disabled: true,
		}
	},
	methods: {
		async login() {
			// this.loading = true;
			this.errormsg = null;
			try {
				// Login (POST): "/session"
				let response = await this.$axios.post("/session", {
                    nickname: this.nickname.trim()
                });

				localStorage.setItem('token',response.data.user_id);
				localStorage.setItem('nickname', this.nickname)
				this.$axios.defaults.headers.common['Authorization']= 'Bearer ' + response.data.user_id
				this.$router.replace("/home")
				this.$emit('updatedLoggedChild',true)
				
			} catch (e) {
				this.errormsg = e.toString();
			}
			// this.loading = false;
		},
	},
	mounted(){
		if (localStorage.getItem('token')){
			this.$router.replace("/home")
		}
	},
	
}
</script>

<!--
Questo codice è un modello HTML/CSS/JS che rappresenta una pagina di login. 
La pagina di login consiste di una semplice form che accetta un input di testo "Username" e ha un pulsante "Accedi" che, 
una volta premuto, esegue un'azione "login". La pagina di login include anche un logo e un titolo per l'applicazione. 
La forma viene stilizzata utilizzando CSS per creare un'interfaccia utente.
-->

<template>
	
	<div class="container-fluid h-100 w-100 d-flex justify-content-center m-0 p-0 login">
	  <div class="login-card p-5 text-center">
		<img src="../assets/images/InstaSnapText1.png" ass="mb-4" alt="WasaPhoto logo" style="width: 320px; height: 100px;"><img>
		<form @submit.prevent="login">
		  <div class="form-group">
			<input type="text" class="form-control" v-model="nickname" placeholder="Nickname">
		  </div>
		  <button class="btn btn-primary" :disabled="nickname.length < 4 || nickname.length > 16">Accedi</button>
		</form>
	  </div>
	</div>
  </template>
  
  <style>

  .text-primary {
	color: rgba(0,0,200,255) !important;
  }
  
  .login {
  display: flex;
  align-items: center;
  justify-content: center;
  background-image: url("../assets/images/sfondo.png");
  background-size: cover;
  background-position: center;
  height: 100vh;
  width: 100%;
  position: absolute;
  top: 0;
  left: 0;
}

.login-card {
	background-color: rgba(241, 241, 241, 0.35);
	box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
  	border-radius: 25px;
  	display: flex;
  	flex-direction: column;
  	align-items: center;
  	justify-content: center;
  	padding: 2rem;
  	text-align: center;
	width: 350px;
	height: 400px;
}

.login-card img {
	width: auto;
  	height: auto;
  	text-align: center;
  	margin-bottom: 10px;
}

.btn-primary {
	margin-top: 30px;
	margin-left: auto;
	margin-right: auto;
	display: block;
	background-color: #0000c8 !important;
	color: white;
	border-radius: 50px;
	padding: 15px 25px;
	font-weight: bold;
	font-size: 18px;
	box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
	transition: all 0.3s ease;
  }
  
.btn-primary:hover {
	transform: translateY(-2px);
	box-shadow: 0 8px 10px rgba(0, 0, 0, 0.1);
  }
  
.btn-primary:active {
	transform: translateY(0);
	box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  }

.form-control{
	width: 300px !important;
}

  
  </style>