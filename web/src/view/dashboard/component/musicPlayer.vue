<template>
    <div>
        <div style="width: 100%">
            <APlayer :audio="audio"></APlayer>
        </div>
    </div>
</template>

<script>
    import { APlayer } from '@moefe/vue-aplayer';
    export default {
        name: "musicPlayer",
        components:{
            APlayer
        },
        data(){
            return {
                audio :[
                ]
            }
        }
    }
</script>

<style scoped>

</style>
