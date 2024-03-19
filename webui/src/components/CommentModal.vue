<script>

/**
Questo codice rappresenta un componente Vue.js che gestisce i commenti su una foto.
Il componente ha diverse proprietà e metodi che gli permettono di visualizzare, aggiungere e eliminare commenti.

La proprietà "modal_id" rappresenta l'identificatore del modal che verrà visualizzato quando si clicca su un commento.
La proprietà "comments_list" rappresenta la lista dei commenti associati alla foto.
La proprietà "photo_owner" rappresenta l'identificatore dell'utente che ha caricato la foto.
La proprietà "photo_id" rappresenta l'identificatore univoco della foto.

Il componente ha diversi metodi, tra cui:

1) "addComment", viene chiamato quando un utente inserisce un commento.
	 Il commento viene inviato come richiesta POST al server all'indirizzo "/users/:id/photos/:photo_id/comments", 
	 dove ":id" viene sostituito con l'identificatore dell'utente proprietario della foto e ":photo_id" viene sostituito con
	 l'identificatore univoco della foto. 
	 La richiesta include il contenuto del commento e l'identificatore dell'utente che ha inserito il commento.

2) Il metodo "eliminateCommentToParent" viene chiamato quando un utente elimina un commento. 
	Emette un evento "eliminateComment" per informare il componente padre che il commento deve essere eliminato.

3) Il metodo "addCommentToParent" viene chiamato quando un nuovo commento viene inserito. 
	Emette un evento "addComment" per informare il componente padre che un nuovo commento è stato inserito e deve essere
	visualizzato nell'elenco dei commenti.
 */

export default {	
	data(){
		return{
			commentValue:"",
		}
	},
	props:['modal_id','comments_list','photo_owner','photo_id'],

	methods: {
		async addComment(){
			try{
				// Comment post: /users/:id/photos/:photo_id/comments
				let response = await this.$axios.post("/Users/"+ this.photo_owner +"/Photos/"+this.photo_id+"/comments",{
					comment_string: this.commentValue
				},{
					headers:{
						'Content-Type': 'application/json'
					}
				})

				this.$emit('addComment',{
					comment_id: response.data.comment_id, 
					photo_id: this.photo_id, 
					owner: response.data.owner,
					comment_string: this.commentValue,
					timestamp: response.data.timestamp,
					},	
				)
				this.commentValue = ""
				
			}catch(e){
				console.log(e.toString())
			}
		},

		eliminateCommentToParent(value){
			this.$emit('eliminateComment',value)
		},

		addCommentToParent(newCommentJSON){
			this.$emit('addComment',newCommentJSON)
		},
	},
}
</script>

<!--Il codice HTML crea una finestra di dialogo modale con un'intestazione e un corpo che contiene una lista di commenti
	visualizzati utilizzando la direttiva Vue v-for. Il corpo del modale ha anche una sezione di piede con un'area di testo
	in cui gli utenti possono inserire un commento e un pulsante per inviare il commento.

Il codice CSS definisce lo stile per la finestra di dialogo modale, incluso il modo in cui i commenti sono visualizzati
all'interno del modale.
Il codice JavaScript basato su Vue.js implementa la logica per visualizzare e gestire i commenti.
Utilizza le proprietà e i metodi del componente Vue per recuperare i commenti da una fonte di dati,
gestire l'aggiunta di nuovi commenti e eliminare i commenti esistenti.
 -->

<template>
	<div class="modal fade my-modal-disp-none" :id="modal_id" tabindex="-1" aria-hidden="true">
	  <div class="modal-dialog modal-dialog-center modal-dialog modal-dialog-scrollable ">
		<div class="modal-content">
		  <div class="modal-header">
			<h1 class="modal-title fs-5" :id="modal_id">Commenti</h1>
			<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
		  </div>
  
		  <div class="modal-body" >
			<PhotoComment v-for="(comm) in comments_list" 
						  :key="comm.comment_identifier" 
						  :author="comm.owner.user_id" 
						  :username="comm.owner.nickname"
						  :comment_id="comm.comment_identifier"
						  :photo_id="comm.photo_id"
						  :content="comm.comment_string"
						  :photo_owner="photo_owner"
						  @eliminateComment="eliminateCommentToParent"
			/>
		  </div>
		  <div class="modal-footer">
    <textarea class="form-control" id="exampleFormControlTextarea1" 
     placeholder="Scrivi un commento" rows="1" maxLength="300" v-model="commentValue"></textarea>
    <button type="button" class="btn btn-primary" 
     @click.prevent="addComment" 
     :disabled="commentValue.length < 1 || commentValue.length > 300">
     Invia
    </button>
</div>
		</div>
	  </div>
	</div>
  </template>
  
  <style> 

.modal-footer {
    display: flex;
    align-items: center;
}

.form-control {
    width: 100%;
    display: inline-block;
    margin-left: 20px;
	margin-right: 20px;
}


  </style>
  
