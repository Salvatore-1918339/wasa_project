<script>


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
      localStorage.removeItem('nickname')
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
  },
}
</script>



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