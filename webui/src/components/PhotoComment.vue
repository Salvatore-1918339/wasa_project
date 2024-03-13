<script>
/**
  Questo codice è un componente Vue.js che rappresenta un singolo commento in un'applicazione web. 
  Il componente fornisce funzionalità per visualizzare il contenuto del commento, l'autore e l'ora della pubblicazione, 
  nonché per eliminare un commento specifico.

  La funzione principale è "deleteComment", che utilizza Axios per inviare una richiesta HTTP DELETE al server 
  per eliminare un commento specifico tramite l'ID del commento e l'ID del proprietario della foto.
  Se la richiesta ha successo, viene emesso un evento "eliminateComment" con l'ID del commento eliminato come parametro, 
  che può essere gestito da un componente padre per rimuovere il commento visualizzato nella UI.

  La funzione "mounted" viene eseguita all'avvio del componente e imposta il valore "user" sul token di autenticazione memorizzato
  nel localStorage del browser dell'utente.
 */
export default {
    data(){
        return {
            user: "",
        }
    },
	props: ['content','author','photo_owner','comment_id','photo_id','username'],

    methods:{
        async deleteComment(){
            try{
                // Delete comment: "/users/:id/photos/:photo_id/comments/:comment_id"
                await this.$axios.delete("/users/"+this.photo_owner+"/photos/"+this.photo_id+"/comments/"+this.comment_id)

                this.$emit('eliminateComment',this.comment_id)

            }catch (e){
                console.log(e.toString())
            }
        },
    },

    mounted(){
        this.user = localStorage.getItem('token')
    }

}
</script>

<!--
Questo codice è un componente Vue.js per un commento. 
Il componente include una porzione di codice HTML che definisce come dovrebbe essere visualizzato un commento,
e una porzione di codice CSS che definisce il look e lo stile del componente.

Viene mostrato il nome dell'autore e il contenuto del commento. Inoltre, c'è un pulsante che viene visualizzato solo se 
l'utente corrente è l'autore del commento o il proprietario della foto.
Questo pulsante, se cliccato, triggera l'evento "deleteComment" che elimina il commento.

Il codice CSS definisce lo stile del pulsante "deleteComment", ad esempio il colore e la dimensione del testo durante 
il passaggio del mouse.
-->

<template>
	<div class="container-fluid">

        <hr>
        <div class="row">
            <div class="col-10">
                <h5>{{username}} @{{author}}</h5>
            </div>

            <div class="col-2">
                <button v-if="user === author || user === photo_owner" class="btn my-btn-comm" @click="deleteComment">
                    <i class="fa-regular fa-trash-can my-trash-icon"></i>
                </button>
            </div>

        </div>

        <div class="row">
            <div class="col-12">
                {{content}}
            </div>

        </div>
        <hr>
    </div>
</template>

<style>
.my-btn-comm{
    border: none;
}
.my-btn-comm:hover{
    border: none;
    color: var(--color-red-danger);
    transform: scale(1.1);
}

</style>
