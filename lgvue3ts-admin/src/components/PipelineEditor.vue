<template>
  <v-card>
    <v-toolbar :title="selectPipeline.name"></v-toolbar>
    <v-card-text>
      <component
        :is="componentMap.get(selectPipeline.name)"
        :pipeline="selectPipeline"
      ></component>
    </v-card-text>
  </v-card>
</template>
<script setup lang="ts">
import CollectorPipelineEditor from "@/components/pipeline/CollectorPipelineEditor.vue";
import { inject, reactive, computed, toRefs } from "vue";
import { useModelState } from "@/stores/modelState";
const model = useModelState();
const selectPipeline = computed(() => {
  return model.selectPipeline;
});
const componentMap = new Map([["采集器:西瓜视频", CollectorPipelineEditor]]);
</script>
<style lang=""></style>
