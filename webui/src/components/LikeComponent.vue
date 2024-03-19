<script>

     export default {
        props: ['identifier','nickname'],
    
        methods:{
    
            clickedUser(){
                this.$emit('clickedUser',this.identifier)
            }
        },
    }
</script>



<template>

    <div class="row mb-2 mt-2" >
        <div class="col d-flex justify-content-center" style="margin-top: 1px; margin-bottom: 1px; height: 30px;">
            <div class="like-card card">
                <div style="position: relative; width: auto;" >
                    <h5 class="User-nickname" @click="clickedUser" style="padding-top: 6px;font-weight: bold;font-size: 1.8em; margin-left: -230px;">{{nickname}}</h5>
                </div>
            </div>
        </div>
    </div>
    
</template>

<style>


.user{
    color: var(--color-text);
}

.User-nickname{
    transition: all 400ms ease-in-out;
}

.User-nickname:hover{
    transition: 0.3s ease;
    transform: translateY(-5px);
    color: 	#dc143c;
    cursor: pointer;
}
.like-card{
    display: flex; 
    align-items: center; 
    width: 320px;
    border: none;
}


</style>
