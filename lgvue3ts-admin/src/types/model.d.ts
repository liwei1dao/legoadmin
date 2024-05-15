
interface Procesline{
  name:string
  pipelines:Pipeline[]

}
type PipelineType = "collector" | "modifiers" | "publisher";// 采集器 修改器 发布器
type PipelineState = 0 | 1 | 2 | 3 ;// 0未配置 1未执行 2执行中 3已执行

//管道
interface Pipeline{
  name:string
  state:PipelineState                        
  setting:Map<string,Object|Number|String>   //配置项
  output:Map<string,Number|String>           //输出项
}


type ResourceType = 'video' | 'audio' | 'text' | 'image' | 'effect' | 'transition' | 'filter';

//资源
interface IResource {
  name:string,
  cover:string,
  path:string,
  rtype:ResourceType
  fps: number,
  frameCount: number
  time:number
}

interface BaseTractItem {
  id: string,
  type: ResourceType,
  name: string,
  start: number,
  end: number,
  frameCount: number,
  offsetL: number, // 音视频左侧裁切
  offsetR: number, // 音视频右侧裁切
}


interface AudioTractItem extends BaseTractItem{
  time: number,
  format: string,
  source: string
  cover: string
}

interface VideoTractItem extends BaseTractItem{
  time: number,
  format: string,
  source: string,
  cover: string,
  width: number,
  height: number,
  fps: number
}

interface TextTractItem extends BaseTractItem{
  cover: string,
  templateId: number
}

interface ImageTractItem extends BaseTractItem{
  source: string,
  format: string,
  width: number,
  height: number,
  sourceFrame: number,
  cover: string
}

interface EffectTractItem extends BaseTractItem{
  templateId: number,
  cover: string
}

interface TransitionTractItem extends BaseTractItem{
  templateId: number,
  cover: string
}

interface FilterTractItem extends BaseTractItem{
  templateId: number,
  cover: string
}
type TrackItem = VideoTractItem | AudioTractItem | TextTractItem | ImageTractItem | EffectTractItem | TransitionTractItem | FilterTractItem;

interface TrackLineItem {
  type: TrackItem['type'],
  main?: boolean,
  list: TrackItem[]
}