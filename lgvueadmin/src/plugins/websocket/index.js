
class MyWebSocket {
    websock = null
    ws_url = "ws://localhost:9567"
    this_open = false
    hearbeat_timer = null
    hearbeat_interval = 30000
    is_reonnect = true
    reconnect_count = 3
    reconnect_current = 1
    reconnect_timer = null
    reconnect_interval = 3000
    messageHandle = {}
    tempmessageHandle = {}
    
    /**
     * 初始化连接
     */
    init(){
      if(typeof(WebSocket) === 'undefined') {
        console.log('浏览器不支持websocket')
        return null
      }

      // 已经创建过连接不再重复创建
      if (this.websock) {
        return this.websock
      }

      this.websock = new WebSocket(this.ws_url)
      this.websock.onmessage =  (e)=> {
        this.receive(e)
      }

      // 关闭连接
      this.websock.onclose =  (e)=> {
        console.log('连接已断开')
        console.log('connection closed (' + e.code + ')')
        clearInterval(this.hearbeat_interval)
        this.this_open = false

        // 需要重新连接
        if (this.is_reonnect) {
          this.reconnect_timer = setTimeout(() => {
            // 超过重连次数
            if (this.reconnect_current > this.reconnect_count) {
              clearTimeout(this.reconnect_timer)
              return
            }

            // 记录重连次数
            this.reconnect_current++
            this.reconnect()
          }, this.reconnect_interval)
        }
      }

      // 连接成功
      this.websock.onopen =  ()=> {
        console.log('连接成功')
        this.this_open = true
        this.is_reonnect = true
        // 开启心跳
        this.heartbeat()
      }

      // 连接发生错误
      this.websock.onerror =  ()=> {
        console.log('Webthis连接发生错误')
      }
    }

    registermessageHandle(msgpath,callback){
      this.messageHandle[msgpath] = callback
    }
    unregistermessageHandle(msgpath){
      this.messageHandle[msgpath] = null
    }
    /**
     * 发送消息
     * @param {*} data 发送数据
     * @param {*} callback 发送后的自定义回调函数
     */
    send(mtype,stype,data,callback = null)  {
      // 开启状态直接发送
      if (this.websock.readyState === this.websock.OPEN) {
        let msgpackage = {
          messages: [
            {
              mtype: mtype,
              stype: stype,
              data: data,
            }
          ],
        }
        console.log('[发送消息]:'+JSON.stringify(msgpackage))
        this.websock.send(JSON.stringify(msgpackage))
        if (callback){
          let msgName = mtype+"."+stype
          this.tempmessageHandle[msgName] = callback
        }
        // 正在开启状态，则等待1s后重新调用
      } else if (this.websock.readyState === this.websock.CONNECTING) {
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
     * 接收消息
     * @param {*} message 接收到的消息
     */
    receive(message)  {
      console.log('[接收消息]:'+message.data)
      var msgpackage = JSON.parse(message.data)  
      if (msgpackage == undefined) {
        console.log("收到服务器空内容")
        return false
      }
      for (const message of msgpackage.messages) {
        let msgName = message.mtype+"."+message.stype
        if (msgName != "system.heart"){
          let callback = this.tempmessageHandle[msgName]
          if (callback != undefined){
            callback(message.data)
            this.tempmessageHandle[msgName] = null
            break
          }else{
            let callback = this.messageHandle[msgName]
            if (callback != undefined){
              callback(message.data)
              this.tempmessageHandle[msgName] = null
              break
            }else{
              console.log("no register handele:"+msgName)
            }
          }
        }
      }
    }

    /**
     * 心跳
     */
    heartbeat()  {
      console.log('this', 'ping')
      if (this.hearbeat_timer) {
        clearInterval(this.hearbeat_timer)
      }
      this.hearbeat_timer = setInterval(() => {
        this.send("system","heart",null,null)
      }, this.hearbeat_interval)
    }
      
    /**
     * 主动关闭连接
     */
    close() {
      console.log('主动断开连接')
      clearInterval(this.hearbeat_timer)
      this.is_reonnect = false
      this.websock.close()
    }
      
    /**
     * 重新连接
     */
    reconnect() {
      console.log('发起重新连接', this.reconnect_current)
  
      if (this.websock && this.this_open) {
        this.websock.close()
      }
  
      this.init()
    }
}

// this主要对象
const newtwebsocket = ()=>{
    return new MyWebSocket();
} 
export default newtwebsocket()