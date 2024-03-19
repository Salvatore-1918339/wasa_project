<script>
export default {	

	props:['modal_id','likes'],

	methods:{
		goToProfile(profileId){
			this.$router.replace("/Users/"+profileId);
		}
	},

}
</script>

<!--
Questo codice definisce un modello per una finestra di dialogo di un'applicazione web.
Il modello viene utilizzato per mostrare i "Mi piace" di una foto o di un contenuto. 
La finestra di dialogo modale viene visualizzata su un'interfaccia utente, sovrapponendo il contenuto originale della pagina.

Il modello è costituito da un div con classe "modal fade my-modal-disp-none", che viene utilizzato come finestra di dialogo modale. 
Il contenuto del modello è contenuto all'interno di un div con classe "modal-dialog modal-dialog-center", 
che viene centrato sulla pagina.
Il contenuto del modello consiste in un'intestazione con un titolo "Mi piace" e un pulsante di chiusura, 
e un corpo che contiene un elenco di elementi "UserMiniCard", che mostrano informazioni sugli utenti che hanno espresso un "Mi piace" 
per il contenuto. 

Il codice CSS associato al modello include la classe "my-modal-disp-none" per nascondere il modello quando non è in uso.-->

<template>
	<div class="modal fade my-modal-disp-none" :id="modal_id" tabindex="-1" aria-hidden="true">
		<div class="modal-dialog modal-dialog-centered">
			<div class="modal-content">
				<div class="modal-header d-flex align-items-center">
					<h1 class="modal-title fs-5 mb-0" :id="modal_id">Mi piace</h1>
					<button type="button" class="btn-close ml-auto" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body p-4 ">
					<div class="d-flex flex-wrap justify-content-center">
						<UserMiniCard v-for="(user) in likes" 
						:key="user.user_id"
						:identifier="user.user_id" 
						:nickname="user.nickname"
						@clickedUser="goToProfile"/>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
<style>

.my-modal-disp-none{
	display: none;
}

</style>