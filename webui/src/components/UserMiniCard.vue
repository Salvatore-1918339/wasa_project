<script>
/**
  Questo codice rappresenta un componente Vue.js che fornisce una funzionalità di navigazione all'interno dell'applicazione.
  Il componente ha una proprietà di input chiamata "identifier" e "username".

  Nel metodo "clickedUser", il componente emette un evento "clickedUser" passando "identifier" come payload. 
  Questo evento potrebbe essere utilizzato da un componente padre per gestire la navigazione alla pagina del profilo dell'utente
  con l'identificativo specificato.
 */
 export default {
	props: ['identifier','nickname'],

    methods:{

        clickedUser(){
            this.$emit('clickedUser',this.identifier)
        }
    },
}
</script>

<!-- 
Questo codice presenta una componente Vue.js che mostra una "scheda mini utente". 
La scheda è costituita da un div con una classe CSS che lo formatta come una "card". La scheda mostra il nome utente e 
un identificatore. Se si clicca sul nome utente, viene attivata un'azione tramite la funzione "clickedUser". 
La funzione è legata all'evento "click" del nome utente.

Inoltre, il codice include anche una sezione di stile CSS che formatta l'aspetto della scheda. 
Questi stili includono la larghezza, il colore del testo, il bordo, e l'effetto al passaggio del mouse sul nome utente.
-->

<template>
    
    <div class="row mb-2 mt-2">
        <div class="col d-flex justify-content-center">
            <div class="user-mini-card card bg-transparent border-start">
                <div class="card-body">
                    <h5 @click="clickedUser" class="user-mini-card-title d-flex justify-content-center ">{{nickname }} @{{identifier }}</h5>
                </div>
            </div>
        </div>
    </div>
    
</template>

<style>

.user{
    color: var(--color-text);
}

.user-mini-card{
    color: black;
    width: 25%;
    border: none;
    border-color: black;
}

.user-mini-card-title:hover{
    cursor: pointer;
    font-weight: bold;
}
</style>
