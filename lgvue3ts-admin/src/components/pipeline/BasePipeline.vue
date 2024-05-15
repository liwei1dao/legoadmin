<template>
  <v-timeline-item :dot-color="stateColor" size="small" fill-dot>
    <template v-slot:icon>
      <div
        class="rounded-full"
        :class="
          selectPipeline.name == pipeline.name ? 'border-solid border-4' : ''
        "
      >
        <v-btn
          variant="text"
          size="small"
          :color="color"
          :icon="icon"
          @click="selectPipelineFunc"
        ></v-btn>
      </div>
    </template>
    <slot></slot>
  </v-timeline-item>
</template>
<script setup lang="ts">
import { inject, reactive, computed, toRefs } from "vue";
import { useModelState } from "@/stores/modelState";
const props = defineProps({
  color: {
    type: String,
    default: "red-lighten-2",
  },
  pipeline: {
    type: Object as () => Pipeline,
    default: {},
  },
  icon: {
    type: String,
    default: "mdi-spider-thread",
  },
});
const model = useModelState();
const selectPipeline = computed(() => {
  return model.selectPipeline;
});
const stateColor = computed(() => {
  if (props.pipeline?.state == 0) {
    return "grey-lighten-2";
  } else {
    return props.color;
  }
});

function selectPipelineFunc() {
  model.selectPipeline = props.pipeline;
}
</script>
<style lang=""></style>
