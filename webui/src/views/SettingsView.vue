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
			nickname: "",
		}
	},

	methods:{
		async modifyNickname(){
			try{
				// Nickname put: /users/:id
				let resp = await this.$axios.put("/users/"+this.$route.params.id,{
					nickname: this.nickname,
				})

				this.nickname=""
			}catch (e){
				this.errormsg = e.toString();
			}
		},
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
	<div class="container-fluid p-5">
	  <div class="row">
		<div class="col d-flex justify-content-center mb-4">
		  <h1 class="text-primary">{{ this.$route.params.id }}&nbspImpostazioni</h1>
		</div>
	  </div>
  
	  <div class="row">
		<div class="col-12 d-flex justify-content-center mb-3">
			<div class="font-large">
		  <p class="ml-2">Un utente ha questa struttura: </p>
		  </div>
		  <div class="font-large">
		  <p class="ml-2 text-success">&nbspusername</p>
		  </div>
		  
		  <div class="font-large">
		  <p class="ml-2">@id.</p>
		</div>
		</div>
		<div class="col-12 d-flex justify-content-center mb-3">
			<div class="font-large">
		  <p>È possibile modificare solo la parte prima della @</p>
		  </div>
		  <div class="font-large">
		  <p class="text-success">&nbsp(username)&nbsp</p>
		  </div>
		  <div class="font-large">
		  <p>e non quella dopo (l'id).</p>
		  </div>
		</div>
		<div class="col-12 d-flex justify-content-center mb-3">
		</div>
	  </div>
  
	  <div class="row">
		<div class="col d-flex justify-content-center">
		  <div class="input-group mb-3 w-25">
			<input
			  type="text"
			  class="form-control w-25"
			  placeholder="Inserisci il tuo nuovo username"
			  maxlength="16"
			  minlength="3"
			  v-model="nickname"
			/>
			<div class="input-group-append">
			  <button
				class="btn btn-primary"
				@click="modifyNickname"
				:disabled="nickname === null || nickname.length > 16 || nickname.length < 3 || nickname.trim().length === 0"
			  >
				Modifica
			  </button>
			</div>
		  </div>
		</div>
	  </div>
  
	  <div class="row" v-if="nickname.trim().length > 0">
		<div class="col d-flex justify-content-center mb-3">
		  <p>Preview: {{ nickname }} @{{ this.$route.params.id }}</p>
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

