<script>
/**

 */
export default {
    data(){
        return {
            user: "",
        }
    },

	props: ['content','author','photo_owner','comment_identifier','photo_id','username'],
    
    // Funzioni di controllo Booleani
    computed:{
        is_comment_author(){
            return this.author === +this.user
        },
        is_photo_author() {
            return this.photo_owner === +this.user
        }
    },

    methods:{
        async deleteComment(){
            try{
                // Delete comment: "/users/:id/photos/:photo_id/comments/:comment_identifier"
                await this.$axios.delete("/Users/"+this.photo_owner+"/Photos/"+this.photo_id+"/comments/"+this.comment_identifier)
                this.$emit('eliminateComment',this.comment_identifier)

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

-->

<template>
	<div class="container-fluid" style="padding-bottom: 5px;">

        
        <div class="comment-card">
            <div class="row">
                <div class="col-10">
                    <h5 style="padding-left: 10px; font-weight: bold;" >{{username}}</h5>
                </div>

            <div v-if="is_comment_author || is_photo_author" class="col-2">
                <img  src="../assets/images/trash.png" alt="Delete Comment" @click="deleteComment">
            </div>

            </div>
            <div class="row">
                <div class="col-12" style="padding-left: 20px;">
                    {{content}}
                </div>

            </div>
        </div>
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

.comment-card {
    border-radius: 5px;
    border: 1px solid black;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.7);}
</style>
