export default {
    //采集器
    collectors:[
        {
            name:"采集器:西瓜视频",
            describe:"西瓜视频采集器!采集单个视频"
        }
    ],
    //发布器
    publishers:[
        {
            name:"发布器:YouTuBe",
            describe:"自动发布视频到YouTube平台上"
        }
    ],
    //改造器
    modifiers:[
        {
            name:"裁剪",
            color: 'deep-purple',
            icon: 'mdi-content-cut',
            describe:"裁剪视频,保留裁剪部分作为视频主题!"
        },
        {
            name:"添加片头",
            color: 'indigo',
            icon: 'mdi-book-variant',
            describe:"裁剪视频,保留裁剪部分作为视频主题!"
        },
        {
            name:"替换音乐",
            color: 'lime',
            icon: 'mdi-music-circle-outline',
            describe:"裁剪视频,保留裁剪部分作为视频主题!"
        }
    ]
};

