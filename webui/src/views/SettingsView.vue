<script>
/**
 * Questo codice rappresenta un componente Vue.js che modifica lo username di un utente.
 * Il componente ha una proprietà chiamata "username" che rappresenta lo username da modificare. 
 * Il componente ha anche un metodo chiamato "modifyusername" che effettua una richiesta PUT all'indirizzo "/users/:id" passando
 * come parametro il nuovo username. 
 * Se la richiesta ha successo, il metodo resetta il valore della proprietà "username".
 * In caso di errore, viene salvato un messaggio di errore nella proprietà "errormsg".
 */
 export default {
	data: function () {
		return {
			errormsg: null,
			err_nickname_exist:false,
			old_nickname: "",
			temp_nickname:"",
			nickname: "",
			changed: false,
		}
	},

	methods:{

		async loadNickname(){
			if (this.$route.params.id === undefined){
				return
			}
			try{
				// Get user profile info: /Users/:id
				let response = await this.$axios.get("/Users/"+this.$route.params.id);
				this.old_nickname = response.data.nickname

			}catch(e){
				this.currentIsBanned = true
			}
		},

		async modifyNickname(){
			try{
				this.temp_nickname = this.nickname
				// Nickname put: /Users/:id
				let resp = await this.$axios.put("/Users/"+this.$route.params.id,{
					nickname: this.nickname,
				})
				this.err_nickname_exist = false
				this.changed=true
				this.nickname=""
			}catch (e){
				switch(e.response.status){
					case 400: 
						this.err_nickname_exist = true;
						break;
					default: 
						this.errormsg = e.toString();
						break;
				}
			}
		},
	},

	async mounted(){
    	this.$axios.defaults.headers.common['Authorization']= 'Bearer ' + localStorage.getItem('token')
		await this.loadNickname()
	},

}
</script>

<!--
Questo codice è un modello Vue.js che implementa una pagina "Impostazioni utente".
La pagina visualizza l'ID dell'utente corrente, che viene passato come parametro all'URL. 
Il codice consente all'utente di modificare il proprio username, che è la parte dell'identificatore dell'utente prima della @.

Il codice utilizza un input di testo v-model per acquisire il nuovo username dell'utente e un pulsante per inviare 
la richiesta di modifica. 
La funzione modifyusername viene eseguita quando viene fatto clic sul pulsante. 
L'anteprima del nuovo identificatore dell'utente viene visualizzata sotto l'input di testo.
Il codice utilizza anche un componente ErrorMsg per visualizzare eventuali messaggi di errore durante la modifica dell'username.

Infine, il codice utilizza stili CSS per formattare la pagina e rendere il testo più leggibile.
-->

<template>
	<div class="container-fluid p-5" style="padding-top: 150px;">
	  <div class="row">
		<div class="col d-flex justify-content-center mb-4" style="padding-top: 100px;">
		  <h1 class="text-primary">Impostazioni</h1>
		</div>
	  </div>
  
	  <div class="row justify-content-center">
			<div v-if="!err_nickname_exist">
				<div v-if="!changed" class="col-12 d-flex justify-content-center mb-3" >
					<div class="font-large">
						<p class="ml-2">Il tuo nickname è: </p>
					</div>
					<div class="font-large">
						<p class="ml-2 text-success">&nbsp{{old_nickname}}</p>
					</div>
				</div>
				<div v-else class="col-12 d-flex justify-content-center mb-3"  style="background-color: #2b961f; border-radius: 8px; width: auto;">
					<div class="font-large">
						<p class="ml-2" style="margin-left: 5px; color: white;">Hai cambiato correttamente il tuo Nickname: </p>
					</div>
					<div class="font-large">
						<p style="font-weight: bolder; color: white;">&nbsp{{temp_nickname}}</p>
					</div>
				</div>
			</div>
			<div v-else>
				<div class="col-12 d-flex justify-content-center mb-3"  style="background-color: red; border-radius: 8px; width: auto;">
					<div class="font-large">
						<p class="ml-2" style="margin-left: 5px; color: white;">Il Nickname scelto esiste già : </p>
					</div>
					<div class="font-large">
						<p style="font-weight: bolder; color: white;">&nbsp{{temp_nickname}}</p>
					</div>
				</div>
			</div>
		</div>
  
		<div class="row">
			<div class="col d-flex justify-content-center">

				<div class="input-group mb-3 w-70">
					<input
					type="text"
					class="form-control w-25"
					placeholder="Inserisci il tuo nuovo Nickname"
					maxlength="25"
					minlength="5"
					v-model="nickname"
					/>
					<div class="input-group-append">
					<button
						class="btn btn-primary"
						@click="modifyNickname"
						:disabled="nickname === null || nickname.length > 25 || nickname.length < 5 || nickname.trim().length === 0">
						Modifica
					</button>
					</div>
				</div>
			</div>
	  </div>
    
	  <div class="row">
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	  </div>
	</div>
  </template>
  
  <style>

  body {
    font-family: "Times New Roman", Times, serif;
  }
  .font-large {
  font-size: 1.5em;
  font-family: 'Times New Roman', Times, serif;
}
.input-group-append {
  
  margin-top: -25px;
}

</style>

