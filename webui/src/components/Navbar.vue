<script>

/**
   Questo script rappresenta un componente Vue.js che viene utilizzato per la barra di navigazione dell'applicazione.
   La barra di navigazione contiene funzionalità come la ricerca, il logout, la visualizzazione del profilo personale, 
   e la navigazione verso la home page.

   Il componente utilizza una serie di metodi che gestiscono i click sui pulsanti e sulle icone presenti nella barra di navigazione. 
   Ad esempio, il metodo logout rimuove il token dalla memoria locale del browser, 
   il metodo goBackHome emette un evento che richiede di visualizzare la home page, 
   il metodo searchFunc emette un evento che richiede la ricerca di elementi sul sito, 
   il metodo myProfile emette un evento che richiede la visualizzazione del profilo personale, e così via.

  Inoltre, il componente utilizza anche alcune proprietà di dati come textVar, che memorizza il testo inserito nella barra di ricerca,
  e iconProfile, che memorizza lo stato dell'icona del profilo (attivo o inattivo). 
  Questi dati sono utilizzati per gestire l'interfaccia utente e renderla dinamica in base alle azioni dell'utente.
 */
export default {
  data(){
    return{
      textVar: "",
      iconProfile: "fa-regular",
    }
  },
  methods:{
    logout(){
      localStorage.removeItem('token')
      this.$emit('logoutNavbar',false)
    },
    goBackHome(){
      this.$emit('requestUpdateView',"/home")
    },
    searchFunc(){
      this.$emit('searchNavbar',this.textVar)
      this.textVar=""
    },
    myProfile(){
      this.$emit('requestUpdateView',"/Users/"+localStorage.getItem('token'))
    },
    profileIconInactive(){
      this.iconProfile = "fa-regular"
    },
    profileIconActive(){
      this.iconProfile = "fa-solid"
    },
  },
}
</script>

<!--
Questo codice definisce un componente di navigazione per un'applicazione web. Il componente visualizza una barra di navigazione
con un pulsante "Home" a sinistra, una casella di testo di ricerca al centro e due pulsanti "Profilo" e "Esci" a destra.
Il componente utilizza anche lo stile CSS per definire l'aspetto dei componenti.
-->

<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-custom">
   
    <div class="logo" style="margin-left: 10px;" @click="goBackHome" ><img src="../assets/images/InstaSnapText_ridimen.png" alt="Home"></div>

    <form class="form-inline" @submit.prevent="searchFunc" style="width: 40%;  display: flex;  justify-content: center;">
      <div class="bar">
        <input class="form-control mr-sm-2 w-100" v-model="textVar" type="search" placeholder="Cerca utenti" aria-label="Search">
        <button class="btn btn-light ms-2" type="submit">Cerca</button>
      </div>
    </form>
    <div class="d-flex" style="right: 0; top: 0">
      <button @click="myProfile" class="button-profilo" type="button">
        Profilo
      </button>
      <button @click="logout" class="button-esci" type="button">
        Esci
      </button>
    </div>
  </nav>

</template>

<style>

.logo {
  transition: all 400ms ease-in-out;
}
.logo:hover{
    filter: drop-shadow(20px 10px 4px #37305f);
    transition: 0.3s ease;
    transform: translateY(-5px);

  }

.navbar{
  width: 100%;
  height: 85px;
  background-color: #53498f;

  display: flex;
  justify-content: space-between;
  align-items: center;
  
  padding: 0px 50px;
  position: absolute;
}

.text-white {
  margin-left: 615px;
  margin-top: -50px;
  font-family: 'Javanese Text';
  font-style: italic;
  font-size: 60px;
}

.button-profilo {
  margin-left: 0px;
  margin-right: 15px;
  border: none;
  font-size: 22px;
  font-family: sans-serif;
  color: #fff;

  border-radius: 5px;
  background-color: #53498f !important;

  transition: 0.3s ease;
}

.button-profilo:hover {
  color: #ccc
}


.button-esci {
  font-family: sans-serif;
  font-size: 22px;
  border-radius: 5px;
  color: #fff;

  border: none;
  background-color: #53498f !important;

  transition: 0.3s ease;
}

.button-esci:hover{
  color: #ccc
}

.form-control{
  font-size: 22px;

}

.btn-light{
  font-size: 22px;
}

.bar {
  display: flex;
  align-items: center;
}


</style>