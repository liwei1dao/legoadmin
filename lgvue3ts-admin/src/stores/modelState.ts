// 1.在src目录下创建 store 文件
// 2.新增 counter.ts 测试仓库,代码如下

// counter.ts
import { defineStore } from "pinia";
import {  inject,reactive ,computed,ref} from 'vue'

//defineStore 接受两个参数
//参数1：仓库的id（字符串）
//参数2：options（对象）
export const useModelState =  defineStore("model",()=> {

    const proceslines = reactive([
        {
            name:"测试管道",
            pipelines:[
                {
                    name:"采集器:西瓜视频",
                    state:0,
                },
                {
                    name:"发布器:Youtbes",
                    state:0,
                }
            ]
        }
    ]) as Procesline[]

    // 选中元素坐标
    const selectProcesline = reactive<Procesline>(proceslines[0]);
    // 选中元素坐标
    const selectPipeline = reactive<Pipeline>(proceslines[0].pipelines[0]);

    const  getList =() =>{
        const socket = inject<ISocket>('socket')
        socket?.send("","",{},null)
    }
    const  getprocesline =(pname:string):Procesline|null =>{
        for (const pipeline of proceslines) {
            if (pipeline.name == pname){
                return pipeline
            }
        }
        return null
    }

    return {
        proceslines,
        selectProcesline,
        selectPipeline,
        getList,
        getprocesline,
    }
})