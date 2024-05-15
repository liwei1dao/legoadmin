/**
 *  时间格式化
 * */
export function formatTime(time: number) {
    let second = Math.ceil(time / 1000);
    const s = second % 60;
    second = Math.floor(second / 60);
    const m = second % 60;
    second = Math.floor(second / 60);
    const h = second % 60;
    return {
        s,
        m,
        h,
        str: `${h === 0 ? '' : `${h < 10 ? '0' : ''}${h}:`}${m < 10 ? '0' : ''}${m}:${s < 10 ? '0' : ''}${s}`
    };
}


// 封装json格式化, 避免error
export function getJsonParse(jsonStr: string): any {
    let res = '';
    try {
        res = JSON.parse(jsonStr);
    } catch (e) {
        console.log(e);
        res = '';
    }
    return res;
}

export function isVideo(type: string) {
    return type === 'video';
}