<script>
export default {	

	props:['modal_id','likes'],

	methods:{
		goToProfile(profileId){
			this.$router.replace("/Users/"+profileId);
		}
	},

}
</script>


<template>
	<div class="modal fade my-modal-disp-none" :id="modal_id" tabindex="-1" aria-hidden="true">
		<div class="modal-dialog modal-dialog-centered">
			<div class="modal-content">
				<div class="modal-header d-flex align-items-center">
					<h1 class="modal-title fs-5 mb-0" :id="modal_id">Mi piace</h1>
					<button type="button" class="btn-close ml-auto" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body p-4 " >
					<div class="d-flex flex-wrap justify-content-center">
						<LikeComponent v-for="(user) in likes"
							:key="user.user_id"
							:identifier="user.user_id" 
							:nickname="user.nickname"
							@clickedUser="goToProfile"
							style="height:40px;"/>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
<style>

.my-modal-disp-none{
	display: none;
}

</style>