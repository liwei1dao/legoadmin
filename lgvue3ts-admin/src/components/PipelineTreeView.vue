<template>
  <v-card>
    <v-toolbar title="自动化管道面板"></v-toolbar>
    <v-card-text>
      <v-timeline align="start">
        <template v-for="(pipeline, i) of procesline?.pipelines" :key="i">
          <component
            :is="componentMap.get(pipeline.name)"
            :pipeline="pipeline"
          ></component>
        </template>
      </v-timeline>
    </v-card-text>
  </v-card>
</template>
<script setup lang="ts">
import CollectorPipeline from "@/components/pipeline/CollectorPipeline.vue";
import PublisherPipeline from "@/components/pipeline/PublisherPipeline.vue";

defineProps({
  procesline: {
    type: Object as () => Procesline | null,
    default: {},
  },
});

const componentMap = new Map([
  ["采集器:西瓜视频", CollectorPipeline],
  ["发布器:Youtbes", PublisherPipeline],
]);
</script>
<style lang=""></style>
