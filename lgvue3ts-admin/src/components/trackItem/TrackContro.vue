<template>
  <v-toolbar height="30">
    <v-row class="mt-3" dense>
      <v-col cols="auto" v-for="item of icons" :key="item.title">
        <v-btn icon @click="handlerIcon(item)">
          <v-icon>{{ item.icon }}</v-icon>
        </v-btn></v-col
      >
      <v-spacer></v-spacer>
      <v-col cols="3">
        <v-slider
          v-model="modelValue"
          color="indigo"
          track-color="grey"
          :min="sliderProps.min"
          :max="sliderProps.max"
          :step="sliderProps.step"
        >
          <template v-slot:prepend>
            <v-btn
              size="small"
              variant="text"
              icon="mdi-minus"
              color="indigo"
              @click="changeScale(-10)"
            ></v-btn>
          </template>

          <template v-slot:append>
            <v-btn
              size="small"
              variant="text"
              icon="mdi-plus"
              color="indigo"
              @click="changeScale(10)"
            ></v-btn>
          </template>
        </v-slider>
      </v-col>
    </v-row>
  </v-toolbar>
</template>
<script setup lang="ts">
import { usePageState } from "@/stores/pageState";
import { useTrackState } from "@/stores/trackState";
import { reactive, computed } from "vue";
const props = defineProps({
  modelValue: {
    type: Number,
    default: 30,
  },
});
const emit = defineEmits({
  "update:modelValue": (val) => {
    return val !== null;
  },
});
const modelValue = computed({
  get() {
    return props.modelValue;
  },
  set(value) {
    emit("update:modelValue", value);
  },
});
const store = usePageState();
const trackStore = useTrackState();
const statePoint = computed(() => store._stepInfo.statePoint);
const stateLength = computed(() => store._stepInfo.stateLength);
const sliderProps = reactive({
  showTooltip: false,
  size: "small",
  step: 10,
  max: 100,
  min: 0,
});
function changeScale(val: number) {
  let newVal = modelValue.value + val;
  if (newVal > sliderProps.max) newVal = sliderProps.max;
  if (newVal < sliderProps.min) newVal = sliderProps.min;
  modelValue.value = newVal;
}
const icons = computed(() => [
  {
    title: "撤销",
    disable: stateLength.value === 0 || statePoint.value === 0,
    icon: "mdi-skip-backward",
  },
  {
    title: "前进",
    disable: statePoint.value === stateLength.value,
    icon: "mdi-skip-forward",
  },
  {
    title: "分割",
    disable: true,
    icon: "mdi-box-cutter",
  },
  {
    title: "删除",
    disable:
      trackStore.selectTrackItem.line === -1 &&
      trackStore.selectTrackItem.index === -1,
    icon: "mdi-trash-can",
  },
]);

function handlerIcon(item: Record<string, any>) {
  const { icon: type, disable } = item;
  if (disable) {
    return;
  }
  if (type === "DeleteIcon") {
    if (
      trackStore.selectTrackItem.line !== -1 &&
      trackStore.selectTrackItem.index !== -1
    ) {
      trackStore.removeTrack(
        trackStore.selectTrackItem.line,
        trackStore.selectTrackItem.index
      );
      trackStore.selectTrackItem.line = -1;
      trackStore.selectTrackItem.index = -1;
    }
  } else if (type === "UndoIcon") {
    store._undo();
  } else if (type === "RedoIcon") {
    store._redo();
  }
}
</script>
<style lang=""></style>
