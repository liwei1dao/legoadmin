import { App } from 'vue'
const weburl = import.meta.env.WebSocketUrl
//MySocekt
class MySocket implements ISocket {
    url:string                                      //链接地址
    socket: WebSocket | null;                       //socket对象
    isConnected: boolean;                           //是否链接
    isReonnect = true                               //是否重连
    reconnect_count = 3                             //最大重连次数
    reconnect_current = 1                           //当前链接次数
    reconnect_interval = 3000                       //重连间隔
    heartbeatTimer:NodeJS.Timeout|null;             //心跳计时器
    hearbeat_interval = 30000                       //心跳间隔
    messageHandle:Map<string,Function>;             //常驻消息监听
    tempmessageHandle:Map<string,Function>;         //临时消息监听

    constructor(url:string) {
        this.url = url
        this.socket = null;
        this.isConnected = false
        this.heartbeatTimer = null
        this.messageHandle = new Map<string,Function>()
        this.tempmessageHandle = new Map<string,Function>()
        this.init()
    }

    init(): void{
        if (this.socket) {
            return
        }
        this.socket = new WebSocket(this.url)
        this.socket.onmessage =  (event)=> {
            console.log('[接收消息]:'+event.data)
            const msgpackage = JSON.parse(event.data) as IMessagePackage
            // 处理收到的数据
            for (const message of msgpackage.messages) {
                let msgName = message.mtype+"."+message.stype
                if (msgName != "system.heart"){
                    let callback = this.tempmessageHandle.get(msgName)
                    if (callback){
                        callback(message.data)
                        this.tempmessageHandle.delete(msgName)
                        break
                    }else{
                        let callback = this.messageHandle.get(msgName)
                        if (callback){
                            callback(message.data)
                            break
                        }else{
                            console.log("no register handele:"+msgName)
                        }
                    }
                }
                return
            }
        }
        
        // 连接成功
        this.socket.onopen =  ()=> {
            console.log('连接成功')
            this.isConnected = true
            // 开启心跳
            this.heartbeat()
        }
        // 关闭连接
        this.socket.onclose =  (e)=> {
            console.log('连接已断开')
            console.log('connection closed (' + e.code + ')')
            clearInterval(this.hearbeat_interval)
            this.isConnected = false

            // 需要重新连接
            if (this.isReonnect) {
            this.heartbeatTimer = setTimeout(() => {
                // 超过重连次数
                if (this.reconnect_current > this.reconnect_count) {
                    if (this.heartbeatTimer){
                        clearTimeout(this.heartbeatTimer)
                    }
                return
                }

                // 记录重连次数
                this.reconnect_current++
                this.reconnect()
            }, this.reconnect_interval)
            }
        }
        // 连接发生错误
        this.socket.onerror =  ()=> {
            console.log('Webthis连接发生错误')
        }
    }
    registerHandle(msgpath:string,callback:Function):void{
        this.messageHandle.set(msgpath,callback)
        return
    }
    unregisterHandle(msgpath:string):void{
        this.messageHandle.delete(msgpath)
        return
    }

    send(mtype:string,stype:string,data:any,callback:Function|null):void{
        // 开启状态直接发送
        if (this.socket?.readyState === this.socket?.OPEN) {
            let msgpackage: IMessagePackage = {
                messages: [
                    {
                        mtype: mtype,
                        stype: stype,
                        data: data,
                    }
                ],
            }
            console.log('[发送消息]:'+JSON.stringify(msgpackage))
            this.socket?.send(JSON.stringify(msgpackage))
            if (callback){
            let msgPath = mtype+"."+stype
            this.tempmessageHandle.set(msgPath,callback)
            }
        }else if (this.socket?.readyState === this.socket?.CONNECTING) {
            setTimeout(()=> {
            this.send(mtype,stype,data,callback )
            }, 1000)
        // 未开启，则等待1s后重新调用
        } else {
            this.init()
            setTimeout( ()=> {
            this.send(mtype,stype,data,callback)
            }, 1000)
      }
    }

    /**
     * 心跳
     */
    heartbeat()  {
        console.log('this', 'ping')
        if (this.heartbeatTimer) {
            clearInterval(this.heartbeatTimer)
        }
        this.heartbeatTimer = setInterval(() => {
            this.send("system","heart",null,null)
        }, this.hearbeat_interval)
    }
    /**
     * 重新连接
     */
    reconnect() {
        console.log('发起重新连接', this.reconnect_current)
        if (this.socket && this.isConnected) {
          this.socket.close()
        }
        this.init()
      }
}

const WebSocketPlugin = {
    install(app: App) {
        let mysocket = new MySocket(weburl)
        // 将WebSocket实例添加到Vue的全局属性中
        app.config.globalProperties.$socket = mysocket
      }
}
export default WebSocketPlugin