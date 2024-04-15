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
        <div class="col d-flex justify-content-center" style="margin-top: 5px; margin-bottom: 5px; height: 50px;border: 1px;">
            <div @click="clickedUser" class="user-mini-card card" style="display: flex; align-items: center; ">
                <div class="left-banner">
                    <img src="../assets/images/user.png" alt="User" height="50"  style="position:absolute;right:-23px; padding-bottom: 4px;">
                </div>
                <div style="position: relative; width: auto;" >
                    <h5  class="" style="padding-top: 6px; font-weight: bold;font-size: 1.8em; position: center; left: -60px;">{{nickname}}</h5>
                </div>
            </div>
        </div>
    </div>
    
</template>

<style>


.user{
    color: var(--color-text);
}

.left-banner {
    position: absolute;
    padding-bottom: 4px;
    left: 0;
    width: 65px;
    height: 100%;
    background-color: #ffcc00;
    border-right: 2px solid black;
}

.user-mini-card{
    border: 2px solid black;
    color: black;
    width: 320px;
    transition: all 400ms ease-in-out;
}

.user-mini-card:hover{
    cursor: pointer;
    transition: 0.3s ease;
    transform: translateY(-5px);
    filter: drop-shadow(5px 5px 0px black);
}
</style>
