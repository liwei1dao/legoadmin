<template>
    <v-container fluid>
        <v-card 
            flat
            class="py-12"
            min-height="530"
            >
            <v-row justify="center">
                <v-col v-for="(product,i) in products" :key="i">
                    <pitem :product="product"></pitem>
                </v-col>
            </v-row>
        </v-card>
    </v-container>
</template>

<script>
import { inject } from 'vue';
import {useStore} from "vuex"
import pitem from '@/components/pitem.vue'
export default {
    name:"IndexPage",
    components:{
        pitem,
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