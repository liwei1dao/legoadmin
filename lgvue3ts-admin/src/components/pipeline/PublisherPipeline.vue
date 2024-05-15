<template>
  <BasePipeline
    :pipeline="pipeline"
    color="indigo-lighten-2"
    icon="mdi-upload-multiple"
  >
    <v-card>
      <v-card-title class="bg-indigo-lighten-2">
        <h2 class="font-weight-light">{{ pipeline?.name }}</h2>
      </v-card-title>
    </v-card>
  </BasePipeline>
</template>
<script setup lang="ts">
import BasePipeline from "@/components/pipeline/BasePipeline.vue";
defineProps({
  pipeline: {
    type: Object as () => Pipeline,
    default: {},
  },
});
</script>
<style lang=""></style>
