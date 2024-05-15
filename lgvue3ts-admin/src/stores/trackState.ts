
import { defineStore } from "pinia";
import { computed,reactive, ref } from 'vue'
import { getJsonParse } from '@/utils/common';
import { useTrackAttrState } from '@/stores/trackAttribute';

export const useTrackState = defineStore('trackState', () => {
    const attrStore = useTrackAttrState();
   
    const dragData = reactive({ // 拖拽数据
        dataInfo: '',
        dragType: '',
        dragPoint: {
          x: 0,
          y: 0
        }
    });
    const moveTrackData = reactive({ // 行内移动
        lineIndex: -1,
        itemIndex: -1
    });
    // 轨道放大比例
    const trackScale = ref(parseInt(localStorage.trackS || '60'));
    const trackList = reactive<TrackLineItem[]>(localStorage.trackList ? getJsonParse(localStorage.trackList) : []);
    
    // 选中元素坐标
    const selectTrackItem = reactive({
        line: -1,
        index: -1
    });

    // 选中元素
    const selectResource = computed(() => {
        if (selectTrackItem.line === -1) {
            return null;
        }
        return trackList[selectTrackItem.line]?.list[selectTrackItem.index] || null;
    });
    // 删除元素
    function removeTrack(lineIndex: number, itemIndex: number, removeAttr = true) {
        const removeItem = trackList[lineIndex].list[itemIndex];
        trackList[lineIndex].list.splice(itemIndex, 1);
        if (trackList[lineIndex].list.length === 0 && !trackList[lineIndex].main) {
        trackList.splice(lineIndex, 1);
        }
        if (trackList.length === 1 && trackList[0].list.length === 0) {
        trackList.splice(0, 1);
        }
        removeAttr && attrStore.deleteTrack(removeItem.id);
    }
    return {
        dragData,
        selectTrackItem,
        selectResource,
        trackScale,
        trackList,
        removeTrack
    }
})