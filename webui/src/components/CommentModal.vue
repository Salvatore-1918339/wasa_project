<script>
/*
	COMMENT-MODAL

	props : Proprietà passate dal 'padre' 

	Metodi:
		addComment()
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
				let response = await this.$axios.post("/Users/"+ this.photo_owner +"/Photos/"+this.photo_id+"/comments",{
					comment_string: this.commentValue
				},{
					headers:{
						'Content-Type': 'application/json'
					}
				})

				console.log("Commento:",response.data)

				this.$emit('addComment',{
					comment_identifier: response.data.comment_identifier, 
					photo_id: this.photo_id, 
					owner: response.data.owner,
					comment_string: this.commentValue,
					timestamp: response.data.timestamp,
					},	
				)
				console.log("Comment list:",this.comments_list)
				this.commentValue = ""
				
			}catch(e){
				console.log(e.toString())
			}
		},

		eliminateCommentToParent(value){
			console.log("NOnno elimina:",value)
			this.$emit('eliminateComment',value)
		},

		addCommentToParent(newCommentJSON){
			this.$emit('addComment',newCommentJSON)
		},
	},
}
</script>

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
						  :comment_identifier="comm.comment_identifier"
						  :photo_id="comm.photo_id"
						  :content="comm.comment_string"
						  :photo_owner="photo_owner"
						  @eliminateComment="eliminateCommentToParent"
			/>
		  </div>
		  <div class="modal-footer">
    <textarea class="form-control" :id="photo_id" 
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
  
