<script>

export default {
	data(){
		return{
			photoURL: "",
			liked: false,
			allComments: [],
			allLikes: [],
			date: new Date()
		}
	},

	props: ['owner','nickname','likes','comments',"upload_date","photo_id","isOwner"], 

	methods:{
		loadPhoto(){
			// GET - URL Photo
			this.photoURL = __API_URL__+ "/Users/"+this.owner+"/Photos/"+this.photo_id 

			// GESTIONE DELLA DATA -- VALORE AGGIORNATO 
			this.date = new Date(this.upload_date);
            var giorno = this.date.getDate();
            var diff = new Date() - this.date;
            if (giorno !=  new Date().getDate()) {
                var giorniPassati = Math.floor(diff/(1000*60*60*24)) 
                this.date = giorniPassati + " giorni fa";
            } else {

                var orePassate = Math.floor(diff/(1000*60*60));
                
                if (orePassate > 0)
                    this.date = orePassate + " ore fa"
                else {
                    var minPassati = Math.floor(diff%(1000*60*60)/(1000*60));           
                    this.date = minPassati + " min fa"
                }
            }
			
		},

		async deletePhoto(){
			// Elimina la photo
			try{
				await this.$axios.delete("/Users/"+this.owner+"/Photos/"+this.photo_id)
				this.$emit("removePhoto",this.photo_id)
			}catch(e){
				//
			}
		},

		photoOwnerClick: function(){
			this.$router.replace("/Users/"+this.owner)
		},

		async toggleLike() {
			const token = localStorage.getItem('token')
			try{
				if (!this.liked){

					await this.$axios.put("/Users/"+ this.owner +"/Photos/"+this.photo_id+"/Likes")
					this.allLikes.push({
						user_id: token,
						nickname: localStorage.getItem('nickname')
					})
				}else{
					await this.$axios.delete("/Users/"+ this.owner  +"/Photos/"+this.photo_id+"/Likes/"+ token)
					this.allLikes.pop()
				}

				this.liked = !this.liked;
			}catch(e){
				//
			}
      		
    	},

		removeCommentFromList(value){
			this.allComments = this.allComments.filter(item=> item.comment_identifier !== value)
		},

		addCommentToList(comment){
			this.allComments.push(comment)
		},

		
	},
	
	async mounted(){
		await this.loadPhoto()

		if (this.likes != null){
			this.allLikes = this.likes
		}

		if (this.likes != null){
			this.liked = this.allLikes.some(obj => obj.user_id === +localStorage.getItem('token')) // + serve a convertire il valore in int
		}
		if (this.comments != null){
			this.allComments = this.comments
		}
		
	},

}
</script>



<template>
	<div class="container-fluid mt-3 mb-5">
	  <LikeModal :modal_id="'like_modal' + photo_id" :likes="allLikes"/>
	  <CommentModal 
		:modal_id="'comment_modal' + photo_id"
		:comments_list="allComments" 
		:photo_owner="owner" 
		:photo_id="photo_id"
		@eliminateComment="removeCommentFromList"
		@addComment="addCommentToList"/>
  
		<div class="d-flex flex-direction-row justify-content-center">
		<div class="card my-card">
			<div class="Title-section"> 
				<button class="Owner-Title my-trnsp-btn" @click="photoOwnerClick" style="display: flex;	justify-content: space-between;	height: 55px;">
				  <i>{{nickname}}</i>
				</button>
				<div v-if="isOwner" class="trash d-flex justify-content-start">
					<button class="my-trnsp-btn my-dlt-btn me-2" @click="deletePhoto" style="float: right; padding-top: 8px;">
					<i class="fa-solid fa-trash w-100 h-100" style="color: red;"></i>
					</button>
				</div>
			</div>
		  <div class="d-flex justify-content-center foto" style="padding-left: 15px; padding-right: 15px;">
			<img :src="photoURL" class="card-img-top img-fluid" >
		  </div>
		  <div class="card-body">
			<div class="container">
				<div class="d-flex flex-column align-items-center">

					<div class="d-flex justify-content-start">
						
						<button class="btn btn-primary mr-3 button-like d-flex align-items-center">
                                <i @click="toggleLike" :class="'me-1 my-heart-color w-100 h-100 fa '+(liked ? 'fa-heart' : 'fa-heart-o') "></i>
                                <i data-bs-toggle="modal" :data-bs-target="'#like_modal'+photo_id" class="my-comment-color ">
                                    {{allLikes.length}}
                                </i>
                        </button>

                        <button class="btn btn-primary mr-3 button-commento d-flex align-items-center" 
						data-bs-toggle="modal" :data-bs-target="'#comment_modal'+photo_id">

            	        <i class="my-comment-color fa-regular fa-comment me-1"></i>
                        <i class="my-comment-color-2"> {{allComments != null ? allComments.length : 0}}</i>

                        </button>
					</div>
				
			  </div>
			  <div class="d-flex flex-row justify-content-start align-items-center upload" style="color: #808080;">
				<p>Caricata il {{date}}</p>
			  </div>
			</div>
		  </div>
		</div>
	  </div>
	</div>
  </template>
  

<style>

.trash:hover {
	filter: drop-shadow(0px 0px 2px crimson);
}

.Title-section {
	display: flex;
	justify-content: space-between;
	height: 55px;
}

.Owner-Title{
	font-family: Verdana, sans-serif;
	font-weight: bold; 
    color: #333; 
	padding-left: 15px;
	font-size: 16px;
    margin-bottom: 15px;
	margin-top: 10px;
	width: 150px;
}

.my-card{
	box-shadow: 0 0 10px rgba(0, 0, 0, 0.7);
	width: auto;
	width: 400px;
	border-color: black;
	border-width: thin;
	display: inline-block;
}

.my-heart-color{
	color: grey;
}
.my-heart-color:hover{
	color: red;
}

.my-comment-color {
	color: grey;
}
.my-comment-color:hover{
	color: black;
}

.my-comment-color-2{
	color:grey
}

.my-dlt-btn{
	font-size: 19px;
}

.card-img-top {
  height: auto;
  width: auto;
}

.upload{
	font-family: 'Times New Roman', Times, serif;
	font-style: "Times New Roman";
	font-size: 18px;
}

.button-like{
	margin-left: 0px;
	margin-right: 15px;
	margin-top: 30px;
	margin-bottom: 25px;
	padding: 8px 16px;
	font-size: 22px;
	color: #0000c8;
	background-color: white !important;
	font-family: 'Times New Roman', Times, serif;
	font-style: "Times New Roman";
	font-size: 20px;
}

.button-commento{
	margin-left: 0px;
	margin-right: 15px;
	margin-top: 30px;
	margin-bottom: 25px;
	padding: 8px 16px;
	font-size: 22px;
	color: #0000c8;
	background-color: white !important;
	font-family: 'Times New Roman', Times, serif;
	font-style: "Times New Roman";
	font-size: 20px;
}

.my-trnsp-btn{
	font-family: 'Times New Roman', Times, serif;
	font-style: "Times New Roman";
	font-size: 22px;
}


</style>
