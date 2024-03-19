<script>
/**
 Questo codice rappresenta la visualizzazione della pagina "home".
 Il metodo loadStream è una funzione asincrona che carica i dati dalla pagina "home" utilizzando la richiesta HTTP GET.
 La richiesta viene effettuata a "/users/:id/home", dove :id viene sostituito con il token dell'utente recuperato dal local storage.
 Se la richiesta riesce, i dati restituiti vengono memorizzati nella proprietà photos.
 Se c'è un errore nella richiesta, il messaggio di errore viene memorizzato nella proprietà errormsg.
 Infine, quando il componente è montato, viene chiamato il metodo loadStream per caricare i dati.
 
 */
export default {
	data: function () {
		return {
			errormsg: null,
			photos: [],
		}
	},

	methods: {
		
		async loadStream() {
			try {
				this.errormsg = null
				// Home get: "/Users/:id/home"
				let response = await this.$axios.get("/Users/" + localStorage.getItem('token') + "/home")
				if (response.data != null){
					this.photos = response.data
				}
			} catch (e) {
				console.log(e.toString())
				this.errormsg = e.toString()
			}
		}
	},

	async mounted() {
		await this.loadStream()
	}

}
</script>

<!--
Il codice utilizza un componente chiamato "Photo" per visualizzare ogni foto, e visualizza un messaggio di errore utilizzando 
un componente "ErrorMsg" se c'è un errore.

La struttura del componente è costituita da una div contenitore, che contiene una riga. Se non ci sono foto da visualizzare, 
verrà visualizzato un messaggio che invita l'utente a iniziare a seguire qualcuno per vedere le foto.
In caso contrario, viene visualizzata una foto utilizzando il componente "Photo" per ogni elemento nell'array "photos".

Il CSS viene utilizzato per stilizzare il messaggio che viene visualizzato quando non ci sono foto da visualizzare.
-->

<template>
	<div class="home container-fluid" >
	  <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	  <div class="row" >
		<div class="align-items-center justify-content-center" v-if="photos.length === 0">
			<h1 class="text-black mt-3">Inizia a seguire qualcuno per vedere le foto!</h1>
		</div>

		<Photo 
			v-for="(photo) in photos"
			:key="photo.photo_id" 
			:owner="photo.owner.user_id" 
			:nickname="photo.owner.nickname"
			:photo_id="photo.photo_id" 
			:comments="photo.comments" 
			:likes="photo.likes" 
			:upload_date="photo.timestamp" 
			:isOwner="sameUser" 
			@removePhoto="removePhotoFromList"
        />
		</div>
	</div>
</template>

<style>
.home {
	padding-top: 100px;
}
.text-black.mt-3 {
  position: relative;
  top: 300px;
  font-size: 28px;
}

</style>