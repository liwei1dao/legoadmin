interface ISocket {
    socket: WebSocket | null;
    isConnected: boolean;
    messageHandle:Map<string,Function>;
    tempmessageHandle:Map<string,Function>;
    init(url:string): void;
    registerHandle(msgpath:string,callback:Function):void
    unregisterHandle(msgpath:string):void
    send(mtype:string,stype:string,data:any,callback:Function|null):void
}


interface IUserMessage {
    mtype:string
    stype:string
    data:any
}
interface IMessagePackage {
    messages:IUserMessage[]
}