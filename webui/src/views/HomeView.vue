<script>

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