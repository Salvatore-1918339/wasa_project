<script>

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
		if (!localStorage.getItem('token')){
			this.$router.replace("/login")
		}
		await this.loadSearchedUsers()
		
	},
}
</script>



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

