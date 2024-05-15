<template>
    <v-container>
        <v-card flat class="py-12">
            <v-row justify="center">
                <v-col sm="12" lg="6">
                    <v-text-field 
                        v-model="producturl"
                        label="NewProduct" 
                        variant="outlined"
                        :rules="[rules.required]"
                        :error-messages="errorMessages"
                        append-inner-icon="mdi-note-plus"
                        @click:append-inner="addnewproduct"
                        >
                    </v-text-field>
                </v-col>
            </v-row>
        </v-card>
        <v-divider></v-divider>
        <v-card 
            flat
            min-height="560"
            class="pb-12"
            >
            <template v-if="product==null">
                <v-row 
                    align="center"
                    no-gutters
                    justify="center"
                    style="height: 550px;"
                    >
                    <v-col cols="auto">
                        没有可编辑的产品,请先添加产品!
                    </v-col>
                </v-row>
            </template>
            <template v-else>
                <v-card-text>
                    <v-row dense>
                        <v-col cols="12">
                            <strong>ID:&nbsp;&nbsp;</strong>{{product._id}}
                        </v-col>
                        <v-col cols="12">
                            <strong>Url:&nbsp;&nbsp;</strong>{{product.sourceurl}}
                        </v-col>
                        <v-col cols="12" v-if="product.title">
                            <strong>Title:&nbsp;&nbsp;</strong>{{product.title}}
                        </v-col>
                    </v-row>
                </v-card-text>
                <v-card-subtitle class="pt-12">编辑面板</v-card-subtitle>
                <v-divider></v-divider>
                <v-row justify="center">
                    <v-col cols="auto">
                        <v-card-subtitle class="pt-12">生产</v-card-subtitle>
                        <v-divider></v-divider>
                        <v-timeline class="mt-5" align="start">
                            <v-timeline-item
                                :dot-color="collectorColor"
                                icon="mdi-s_ider-thread"
                                fill-dot
                            >
                            <v-card min-width="400">
                                <v-card-title :class="`text-h6 bg-${collectorColor}`">
                                    <v-row>
                                        <v-col cols="auto">{{currcollector != null ? currcollector.name:"采集器:还未配置!"}}</v-col>
                                        <v-spacer></v-spacer>
                                        <v-col cols="auto" class="font-weight-black text-subtitle-1" align-self="center">{{pipelineState(currcollector)}}</v-col>
                                        <v-col cols="auto" v-if="currcollector != null" class="text-h7">
                                            <v-avatar 
                                                color="black"
                                                size="30">
                                                <v-avatar :color="currcollector.state == 0 ? 'white':'green'" size="18"></v-avatar>
                                            </v-avatar>
                                        </v-col>
                                    </v-row>
                                </v-card-title>
                                <v-card-text class="pt-2 bg-white text--primary">
                                    <template v-if="currcollector != null">
                                        <template v-if="currcollector.state == 2">
                                            <v-progress-linear
                                                indeterminate
                                                color="green"
                                            ></v-progress-linear>
                                        </template>
                                        <template v-if="currcollector.state == 10">
                                            <v-card flat class="pa-0">
                                                <v-card-title>
                                                    <strong>Title:&nbsp;&nbsp;</strong>{{product.title}}
                                                </v-card-title>
                                                <v-card-text>
                                                    <video-player
                                                        style="width: 100%;height: 100%;"
                                                        class="vjs-custom-skin"
                                                        controls
                                                        :options="collectorVidoe"
                                                    />
                                                </v-card-text>
                                            </v-card>
                                        </template>
                                    </template>
                                    <template v-else></template>
                                </v-card-text>
                                <v-divider></v-divider>
                                <v-card-actions>
                                    <v-spacer></v-spacer>
                                    <v-btn
                                        v-show="currcollector == null"
                                        color="red-lighten-2"
                                        variant="outlined"
                                    >
                                        配置
                                    </v-btn>
                                    <v-btn
                                        v-show="currcollector != null"
                                        color="red-lighten-2"
                                        variant="outlined"
                                        @click="collectorproduct"
                                    >
                                        {{currcollector.state == 10?'重新采集':'采集'}}
                                    </v-btn>
                                </v-card-actions>
                            </v-card>
                            </v-timeline-item>

                            <v-timeline-item
                            v-for="(modifier, i) in currmodifier"
                                :key="i"
                                :dot-color="modifierColor(modifier)"
                                :icon="modifier.icon"
                                fill-dot
                            >
                            <v-card>
                                <v-card-title :class="['text-h6', `bg-${modifierColor(modifier)}`]">
                                {{modifier.name}}
                                </v-card-title>
                                <v-card-text class="bg-white text--primary">
                                {{modifier.describe}}
                                </v-card-text>
                                <v-divider></v-divider>
                                <v-card-actions>
                                    <v-spacer></v-spacer>
                                    <v-btn
                                        :color="modifierColor(modifier)"
                                        variant="outlined"
                                    >
                                        执行
                                    </v-btn>
                                </v-card-actions>
                            </v-card>
                            </v-timeline-item>
                            
                            <v-timeline-item
                                :dot-color="publisherColor"
                                icon="mdi-transfer-up"
                                fill-dot
                            >
                            <v-card min-width="400">
                                <v-card-title :class="`text-h6 bg-${publisherColor}`">
                                    <v-row>
                                        <v-col cols="auto">{{currpublisher != null ? currpublisher.name:"发布器:还未配置!"}}</v-col>
                                        <v-spacer></v-spacer>
                                        <v-col cols="auto" v-if="currpublisher != null" class="text-h7">
                                            <v-avatar 
                                                color="grey darken-4"
                                                size="30">
                                                <v-avatar :color="currpublisher.state == 0 ? 'white':'green'" size="18"></v-avatar>
                                            </v-avatar>
                                        </v-col>
                                    </v-row>
                                </v-card-title>
                                <v-card-text class="pt-2 bg-white text--primary">
                                   {{currpublisher != null? currpublisher.describe:""}}
                                </v-card-text>
                                <v-divider></v-divider>
                                <v-card-actions>
                                    <v-spacer></v-spacer>
                                    <v-btn
                                        v-show="currpublisher == null"
                                        :color="publisherColor"
                                        variant="outlined"
                                    >
                                        配置
                                    </v-btn>
                                    <v-btn
                                        v-show="currpublisher != ''"
                                        :color="publisherColor"
                                        variant="outlined"
                                    >
                                        发布
                                    </v-btn>
                                </v-card-actions>
                            </v-card>
                            </v-timeline-item>
                        </v-timeline>
                    </v-col>
                    <v-spacer></v-spacer>
                    <v-divider class="mt-12" vertical></v-divider>
                    <v-col cols="auto">
                        <v-card-subtitle class="pt-12">采集器</v-card-subtitle>
                        <v-divider></v-divider>
                        <v-card-text>
                            <v-row dense>
                                <v-col cols="12">
                                    <v-chip v-if="currcollector != null" 
                                        color="blue"  
                                    >{{currcollector.name}}</v-chip>
                                </v-col>
                                <v-col cols="12">
                                    <v-btn block color="blue" @click="dialog = true">{{currcollector == null? '添加采集器':'更换采集器'}}</v-btn>
                                </v-col>
                            </v-row>
                        </v-card-text>
                        <v-card-subtitle class="pt-12">改造器</v-card-subtitle>
                        <v-divider></v-divider>
                        <v-card-text>
                            <v-row dense>
                                <v-col cols="12">
                                    <v-row>
                                        <v-col cols="auto" v-for="(modifier,i) in currmodifier" :key="i">
                                            <v-chip 
                                                color="blue" 
                                                close
                                                @click:close="removemodifier(modifier)"
                                            >{{modifier.name}}</v-chip>
                                        </v-col>
                                    </v-row>
                                </v-col>
                                <v-col cols="12">
                                    <v-btn block color="blue" @click="dialog1 = true">添加修改器</v-btn>
                                </v-col>
                            </v-row>
                        </v-card-text>

                        <v-card-subtitle class="pt-12">发布器</v-card-subtitle>
                        <v-divider></v-divider>
                        <v-card-text>
                            <v-row dense>
                                <v-col cols="12">
                                    <v-chip v-if="currpublisher != null" 
                                        color="blue"
                                    >{{currpublisher.name}}</v-chip>
                                </v-col>
                                <v-col cols="12">
                                    <v-btn block color="blue" @click="dialog2 = true">{{currpublisher == null? '添加发布器':'更换发布器'}}</v-btn>
                                </v-col>
                            </v-row>
                        </v-card-text>

                        <v-card-actions>
                            <v-btn v-show="change" @click="modifyproduct">修改</v-btn>
                        </v-card-actions>

                    </v-col>
                </v-row>
            </template>
        </v-card>
    </v-container>
    <v-dialog
        v-model="dialog"
        fullscreen
        hide-overlay
        transition="dialog-bottom-transition"
        scrollable
      >
        <v-card tile>
          <v-toolbar
            flat
            dark
            color="primary"
          >
            <v-btn
              icon
              dark
              @click="dialog = false"
            >
              <v-icon>mdi-close</v-icon>
            </v-btn>
            <v-toolbar-title>采集器选择</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>
          <v-card-text>
            <v-row>
                <v-col v-for="(collector,i) in config.collectors" :key="i">
                    <v-btn :color="currcollector != null && currcollector.name == collector.name ?'blue':'grey' " @click="modifycollector(collector)">{{collector.name}}</v-btn>
                </v-col>
            </v-row>
          </v-card-text>
        </v-card>
    </v-dialog>
    <v-dialog
        v-model="dialog1"
        fullscreen
        hide-overlay
        transition="dialog-bottom-transition"
        scrollable
      >
        <v-card tile>
          <v-toolbar
            flat
            dark
            color="primary"
          >
            <v-btn
              icon
              dark
              @click="dialog1 = false"
            >
              <v-icon>mdi-close</v-icon>
            </v-btn>
            <v-toolbar-title>修改器选择</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>
          <v-card-text>
            <v-row>
                <v-col cols="auto" v-for="(modifier,i) in config.modifiers" :key="i">
                    <v-btn color="grey" @click="modifymodifier(modifier)">{{modifier.name}}</v-btn>
                </v-col>
            </v-row>
          </v-card-text>
        </v-card>
    </v-dialog>
    <v-dialog
        v-model="dialog2"
        fullscreen
        hide-overlay
        transition="dialog-bottom-transition"
        scrollable
      >
        <v-card tile>
          <v-toolbar
            flat
            dark
            color="primary"
          >
            <v-btn
              icon
              dark
              @click="dialog2 = false"
            >
              <v-icon>mdi-close</v-icon>
            </v-btn>
            <v-toolbar-title>发布器选择</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>
          <v-card-text>
            <v-row>
                <v-col v-for="(publisher,i) in config.publishers" :key="i">
                    <v-btn :color="currpublisher != null && currpublisher.name == publisher.name ?'blue':'grey' " @click="modifypublisher(publisher)">{{publisher.name}}</v-btn>
                </v-col>
            </v-row>
          </v-card-text>
        </v-card>
    </v-dialog>
</template>

<script>
import { inject,onMounted,onUnmounted,ref } from 'vue';
import { useStore } from "vuex"
import { useRoute } from "vue-router"
export default {
    name:"AddPage",
    props:{
    },
    data: () => ({
        rules: {
          required: value => !!value || 'Required.',
          counter: value => value.length <= 20 || 'Max 20 characters',
          email: value => {
            const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
            return pattern.test(value) || 'Invalid e-mail.'
          },
        },
        errorMessages:"",
        producturl:"",
        dialog:false,
        dialog1:false,
        dialog2:false,
        change:false,       //产品配置变化
    }),
    computed: {    // 使用对象展开运算符将 getter 混入 computed 对象中

        collectorColor(){
            if (this.currcollector == null){
                return "grey-lighten-2"
            }
            if (this.currcollector.state == 0){
                return "red-lighten-5"
            }
            if (this.currcollector.state == 10){
                return "red-lighten-2"
            }
            return "grey-lighten-2"
        },
        publisherColor(){
            if (this.currpublisher == null){
                return "grey-lighten-2"
            }
            if (this.currpublisher.state == 0){
                return "orange-lighten-5"
            }
            if (this.currpublisher.state == 10){
                return "orange-lighten-2"
            }
            return "grey-lighten-2"
        },
        modifierColor(){
            return function(modifier){
                if (modifier.state == 0){
                    return `${modifier.color}-lighten-5`
                }
                if (modifier.state == 2){
                    return `${modifier.color}-lighten-2`
                }
                return "grey-lighten-2"
            }
        },
        pipelineState(){
            return function(pipeline){
                if (pipeline == null){
                    return "未配置"
                }
                if (pipeline.state == 0){
                    return "未执行"
                }
                if (pipeline.state == 1){
                    return "等待执行"
                }
                if (pipeline.state == 2){
                    return "正在执行"
                }
                if (pipeline.state == 10){
                    return "已执行"
                }
                return "状态未知"
            }
        },
        collectorVidoe(){
            return {
                playbackRates: [0.7, 1.0, 1.5, 2.0], //播放速度
                autoplay: false, //如果true,浏览器准备好时开始回放。
                muted: false, // 默认情况下将会消除任何音频。
                loop: false, // 导致视频一结束就重新开始。
                preload: 'auto', // 建议浏览器在<video>加载元素后是否应该开始下载视频数据。auto浏览器选择最佳行为,立即开始加载视频（如果浏览器支持）
                language: 'zh-CN',
                aspectRatio: '16:9', // 将播放器置于流畅模式，并在计算播放器的动态大小时使用该值。值应该代表一个比例 - 用冒号分隔的两个数字（例如"16:9"或"4:3"）
                fluid: true, // 当true时，Video.js player将拥有流体大小。换句话说，它将按比例缩放以适应其容器。
                sources: [
                {
                    type: 'video/mp4',
                    src: this.getstaticfile(this.product.videopath),
                },
                ],
                poster: this.product.coverurl, //你的封面地址
                notSupportedMessage: '此视频暂无法播放，请稍后再试', //允许覆盖Video.js无法播放媒体源时显示的默认信息。
                controlBar: {
                    timeDivider: true,
                    durationDisplay: true,
                    remainingTimeDisplay: false,
                    fullscreenToggle: true, //全屏按钮
                },
            }
        }
    },
    setup() {
        let route = useRoute(); 
        let url = route.query.url?route.query.url:""
        let store = useStore();
        const products = store.getters['product/products']
        // 调用mutations
        const addproduct = () => {
            store.commit('product/add')
        }
        let product = ref(null) 
        let currcollector = ref(null)
        let currmodifier = ref([])
        let currpublisher = ref(null)
        if (products[route.query._id]){
            product.value = Object.assign({}, products[route.query._id]); 
            if (product.value['setting'] != null){
                currcollector = product.value['setting']['collector']
                currmodifier = product.value['setting']['modifiers']
                currpublisher = product.value['setting']['publisher']
            }
        }
        onMounted(()=>{ 
            console.log("注册监听函数"); 
            websocket.registermessageHandle("factory.productchange",(resp)=>{
                if (resp.product._id == product.value._id) {
                    product.value = resp.product
                    if (product.value['setting'] != null){
                        currcollector = product.value['setting']['collector']
                        currmodifier = product.value['setting']['modifiers']
                        currpublisher = product.value['setting']['publisher']
                    }
                }
                store.commit('product/add',product.value)
            },)})

        onUnmounted(()=>{
            console.log("注销监听函数"); 
            websocket.unregistermessageHandle("factory.productchange")
        })
        let websocket = inject('$websocket');
        let config = inject('$config');
        return {
            config,
            websocket,
            url,
            product,
            currcollector,
            currmodifier,
            currpublisher,
            addproduct,
        }
    },       
    // onMounted () {
    //     console.log("添加长期监听")
    //     //注册监听函数
    //     this.websocket.registermessageHandle("factory.productchange",this.pushproductchange)
    // },
    // onUnmounted(){
    //     //注销监听函数
    //     this.websocket.unregistermessageHandle("factory.productchange")
    //     console.log("移除长期监听")
    // },
    methods: {
        addnewproduct(){
            if(this.producturl == ""){
                return
            }
            console.log("添加新的产品")
            //添加产品
            this.websocket.send("factory","addnewproduct",{url:this.producturl},(resp)=>{
                this.product = resp.product
            })
        },
        modifyproduct(){
            if(this.product == null){
                return
            }
            //修改产品
            this.websocket.send("factory","modifyproduct",{id:this.product._id,setting:this.product.setting},(resp)=>{
                    this.product = resp.product
            })
        },
        collectorproduct(){
            if(this.product == null || this.currcollector == null){
                return
            }
            //修改产品
            this.websocket.send("factory","collection",{id:this.product._id},(resp)=>{
                    this.product = resp.product
            })
        },
        //推送变化产品
        pushproductchange(resp){
            let product = resp.product
            if (resp.product._id == this.product._id) {
                this.product = resp.product
            }
            this.addproduct(product)
            if (product['setting'] != null){
                this.currcollector = product['setting']['collector']
                this.currmodifier = product['setting']['modifiers']
                this.currpublisher = product['setting']['publisher']
            }
        },

        //修改采集器
        modifycollector(collector){
            console.log("collector:%s",collector)
            if (this.product.setting == undefined)  {
                this.product.setting = {
                    collector:"",
                    modifiers:[],
                }
            }
            collector.state = 0
            this.product.setting["collector"] = collector
            this.currcollector = collector
            this.dialog = false
            this.change = true
        },
        //修改修改器
        modifymodifier(modifier){
            console.log("modifier:%s",modifier)
            if (this.product.setting == undefined)  {
                this.product.setting = {
                    collector:"",
                    modifiers:[],
                }
            }
            modifier.state = 0
            this.product.setting.modifiers.push(modifier)
            this.currmodifier.push(modifier)
            this.dialog1 = false
            this.change = true
        },

        removemodifier(modifier){
            console.log("modifier:%s",modifier)

        },

        //修改采集器
        modifypublisher(publisher){
            console.log("publisher:%s",publisher)
            if (this.product.setting == undefined)  {
                this.product.setting = {
                    collector:null,
                    modifiers:[],
                    publisher:null,
                }
            }
            publisher.state = 0
            this.product.setting["publisher"] = publisher
            this.currpublisher = publisher
            this.dialog2 = false
            this.change = true
        },
        //获取静态资源
        getstaticfile(fliepath){
            return `http://127.0.0.1:9269/${fliepath.replace("./download",'static')}` 
        },
        //刷新执行流
        refreshPipeline(){
            if (this.product.setting == null){
                this.pipeline = []
                return
            }

        }
    }
}
</script>

<style>

</style>