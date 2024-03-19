<script>
/**
Questo codice è un componente Vue.js che implementa una funzionalità di ricerca utente. 
La logica principale è definita nella sezione "methods".

Il componente ha una proprietà chiamata "searchValue", che è un valore passato da un componente padre. 
Questa proprietà viene osservata utilizzando la sezione "watch", e ogni volta che viene modificata, viene eseguita la funzione 
"loadSearchedUsers".

La funzione "loadSearchedUsers" utilizza l'axios per inviare una richiesta GET all'URL "/users" con il parametro "id" impostato 
sulla "searchValue". Se la richiesta viene eseguita correttamente, i dati della risposta vengono assegnati alla proprietà "users". 
In caso contrario, viene assegnato un messaggio di errore alla proprietà "errormsg".

Quando il componente viene montato, viene verificato se l'utente è autenticato utilizzando "localStorage.getItem('token')". 
Se non lo è, l'utente viene reindirizzato alla pagina di login. Alla fine, viene chiamata la funzione "loadSearchedUsers".

Infine, la funzione "goToProfile" viene utilizzata per reindirizzare l'utente alla pagina del profilo utente specifico 
passando l'ID del profilo come parametro.

 */
 export default {
	data: function() {
		return {
			users: [],
			errormsg: null,
		}
	},

	props:['searchValue'],

	watch:{
		searchValue: function(){
			this.loadSearchedUsers()
		},
	},

	methods:{
		async loadSearchedUsers(){
			this.errormsg = null;
			if (
				this.searchValue === undefined ||
				this.searchValue === "" || 
				this.searchValue.includes("?") ||
				this.searchValue.includes("_")){
				this.users = []
				return 
			}
			try {
				// Search user (PUT):  "/Users"
				let response = await this.$axios.get("/Users",{
						params: {
							user_query_id: this.searchValue,
					},
				});
				this.users = response.data

			} catch (e) {
				this.errormsg = e.toString();
			}
		},

		goToProfile(profileId){
			this.$router.replace("/Users/"+profileId)
		}
	},

	async mounted(){
		// Check if the user is logged
		if (!localStorage.getItem('token')){
			this.$router.replace("/login")
		}
		await this.loadSearchedUsers()
		
	},
}
</script>

<!--
Il componente visualizza un elenco di utenti tramite una struttura di dati chiamata "users". 
Ogni utente viene rappresentato tramite un componente "UserMiniCard". Se l'elenco "users" è vuoto, viene visualizzato 
il testo "Non ho trovato alcun utente". Inoltre, se viene impostato un messaggio di errore, verrà visualizzato tramite 
il componente "ErrorMsg". La proprietà "key" viene utilizzata per identificare in modo univoco ogni componente "UserMiniCard"
nell'elenco. La proprietà "identifier" viene utilizzata per impostare l'ID univoco dell'utente corrispondente, mentre la proprietà
"username" viene utilizzata per impostare il nome utente. 

L'evento "clickedUser" viene gestito per visualizzare il profilo dell'utente selezionato.
-->

<template>
	<div class="container-fluid h-100 " style="padding-top: 100px;" >
		<UserMiniCard v-for="(user) in users" 
		:key="user.user_id"
		:identifier="user.user_id" 
		:nickname="user.nickname" 
		@clickedUser="goToProfile"/>

		<p v-if="users.length == 0" class="no-result-text d-flex justify-content-center"> Non ho trovato alcun utente.</p>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>

.no-result-text{
	color: white;
	font-style: italic;
}
</style>

