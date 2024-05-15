<template>
  <div class="flex flex-col" :style="trackHeight">
    <TrackContro v-model="store.trackScale"></TrackContro>
    <TrackList />
  </div>
</template>
<script setup lang="ts">
import { computed, reactive } from "vue";
import TrackContro from "@/components/trackItem/TrackContro.vue";
import TrackList from "@/components/container/TrackList.vue";
import { useTrackState } from "@/stores/trackState";
import { usePageState } from "@/stores/pageState";
const page = usePageState();
const store = useTrackState();
const trackHeight = computed(() => ({
  height: `${page.trackHeight}px`,
}));
const limitSize = reactive({
  minHeight: 200,
  maxHeight: document.body.getBoundingClientRect().height - 200,
});
</script>
<style lang=""></style>
