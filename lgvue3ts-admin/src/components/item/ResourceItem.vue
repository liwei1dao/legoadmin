<template>
  <v-card draggable="true" @dragstart="dragStart">
    <v-img
      :src="resource.cover"
      @mousemove="showGif($event, formatData.source)"
      @mouseout="showGif($event, formatData.cover)"
    ></v-img>
    <div>{{ resource.name }}</div>
  </v-card>
</template>
<script setup lang="ts">
import { computed } from "vue";
import { useTrackState } from "@/stores/trackState";
const props = defineProps({
  resource: {
    type: Object as () => IResource,
    required: true,
  },
});
const formatData = computed(() => {
  let { time, frameCount } = props.resource;
  if (props.resource.rtype === "video" && !time) {
    time = parseInt(`${(frameCount / 30) * 1000}`);
  }
  return {
    ...props.resource,
    time,
  };
});
const store = useTrackState();
function dragStart(event: DragEvent) {
  event.stopPropagation();
  const dragInfo = {
    type: props.resource.rtype,
    ...formatData.value,
  };
  store.dragData.dataInfo = JSON.stringify(dragInfo);
  store.dragData.dragType = props.resource.rtype;
  store.dragData.dragPoint.x = event.offsetX;
  store.dragData.dragPoint.y = event.offsetY;
  store.selectTrackItem.line = -1;
  store.selectTrackItem.index = -1;
}
function showGif(event: MouseEvent, imageSource: ImageTractItem["type"]) {
  if (["image"].includes(props.resource.rtype) && event.target) {
    (event.target as HTMLImageElement).src = imageSource;
  }
}
</script>
<style lang=""></style>
