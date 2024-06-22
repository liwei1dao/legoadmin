<template>
    <v-container fluid>
        <button @click="serializeData">Serialize Data</button>
        <button @click="deserializeData">Deserialize Data</button>
    </v-container>
</template>

<script>
const { AgentMessage } = require('@/pb/js/comm_pb');
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
        serializeData() {
            const message = new AgentMessage();
            message.setMsgname('test/liwei1dao');
            this.serializedData = message.serializeBinary();
            console.log('Serialized data:', this.serializedData);
        },
        deserializeData() {
            if (this.serializedData) {
                const message1 = AgentMessage.deserializeBinary(this.serializedData);
                console.log('Deserialized data:', message1.toObject());
            } 
        },
    }
}
</script>

<style>

</style>