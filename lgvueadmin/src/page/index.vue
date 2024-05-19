<template>
    <v-container fluid>

    </v-container>
</template>

<script>
import { inject } from 'vue';
import {useStore} from "vuex"
export default {
    name:"IndexPage",
    components:{

    },
    data: () => ({
        products:[]
    }),
    setup() {
        const websocket = inject('$websocket');
        return {
            websocket,
        }
    },
    mounted () {
        this.updateView()
    },
    methods: {
        updateView(){
            let {commit} = useStore();
            //获取产品列表
            this.websocket.send("factory","inquireInfo",{},(resp)=>{
                console.log("收到查询消息:%v",resp)
                this.products = resp.products
                commit("product/adds", resp.products)  //commit 提交数据
            })
        },
        addnewproduct(){

        }
    }
}
</script>

<style>

</style>