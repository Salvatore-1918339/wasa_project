<script>

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
				let response = await this.$axios.get("/Users/"+this.$route.params.id);
				this.old_nickname = response.data.nickname

			}catch(e){
				this.currentIsBanned = true
			}
		},

		async modifyNickname(){
			try{
				this.temp_nickname = this.nickname
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

