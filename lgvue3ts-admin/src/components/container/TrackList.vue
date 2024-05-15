<template>
  <div class="flex flex-1 flex-row w-full">
    <TrackListIcon :listData="showTrackList" :offsetTop="startY" />
    <div
      class="flex-1 overflow-x-scroll overflow-y-auto flex-col shrink-0 grow relative"
      ref="trackList"
      @scroll="handleScroll"
      @wheel="handleWheel"
      @click="setSelectTract($event, -1, -1)"
    >
      <TimeLine
        :start="startX"
        :scale="trackScale"
        :step="defaultFps"
        @selectFrame="handlerSelectFrame"
      ></TimeLine>
    </div>
  </div>
</template>
<script setup lang="ts">
import TrackListIcon from "@/components/trackItem/TrackListIcon.vue";
import TimeLine from "@/components/trackItem/TimeLine.vue";
import { getGridPixel, getSelectFrame } from "@/utils/canvasUtil";
import { formatTime, isVideo, getJsonParse } from "@/utils/common";
import { ref, computed } from "vue";
import { debounce } from "lodash-es";
import { useTrackState } from "@/stores/trackState";
import { usePlayerState } from "@/stores/playerState";

const store = useTrackState();
const playerStore = usePlayerState();
const trackList = ref();
const trackListContainer = ref();
const offsetLine = {
  left: 10, // 容器 margin, 为了显示拖拽手柄
  right: 200,
};
const startX = ref(0 - offsetLine.left); // 与容器padding对齐
const startY = ref(0); // 左侧icons对齐
const trackScale = computed(() => store.trackScale);
const trackStyle = computed(() => {
  return {
    width:
      getGridPixel(trackScale.value, playerStore.frameCount) + offsetLine.right,
  };
});
const defaultFps = ref(30); // 帧率
let mainIndex = ref(0); // main 行下标
const showTrackList = computed(() => {
  return store.trackList.map((line, lineIndex) => {
    line.main && (mainIndex.value = lineIndex);
    const newList = line.list.map((item) => {
      const { time } = item as VideoTractItem;
      return {
        ...item,
        showWidth: `${getGridPixel(trackScale.value, item.end - item.start)}px`,
        showLeft: `${getGridPixel(trackScale.value, item.start)}px`,
        time: isVideo(line.type) ? `${formatTime(time || 0).str}` : "",
      };
    });
    return {
      ...line,
      list: newList,
    };
  });
});
function handleScroll() {
  const { scrollLeft, scrollTop } = trackList.value;
  startX.value = scrollLeft - offsetLine.left;
  startY.value = scrollTop;
}
function handlerSelectFrame(frame: number) {
  const playFrame = frame - 1;
  const startFrane =
    playFrame < 0
      ? 0
      : playFrame > playerStore.frameCount
      ? playerStore.frameCount
      : playFrame;
  playerStore.playStartFrame = startFrane;
  playerStore.playAudioFrame = startFrane;
}
let maxDelta = 0;
const setScale = debounce(() => {
  store.trackScale -= maxDelta > 0 ? 10 : -10;
  maxDelta = 0;
}, 100);
const handleWheel = (event: WheelEvent) => {
  if (event.ctrlKey || event.metaKey) {
    event.preventDefault();
    maxDelta || (maxDelta = event.deltaY);
    setScale();
  }
};
function setSelectTract(event: Event, line: number, index: number) {
  event.preventDefault();
  event.stopPropagation();
  store.selectTrackItem.line = line;
  store.selectTrackItem.index = index;
}
</script>
<style lang=""></style>
