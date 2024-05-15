<template>
  <v-container class="d-flex justify-center fill-height" fluid>
    <div
      class="flex flex-row w-full h-full border-2 border-gray-700 divide-x-4 divide-dashed"
    >
      <div class="basis-2/5">
        <PipelineTreeView :procesline="procesline" />
      </div>
      <div class="basis-3/5">
        <PipelineEditor />
      </div>
    </div>
  </v-container>
</template>
<script setup lang="ts">
import PipelineTreeView from "@/components/PipelineTreeView.vue";
import PipelineEditor from "@/components/PipelineEditor.vue";

import { ref, computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useModelState } from "@/stores/modelState";
import { usePageState } from "@/stores/pageState";
const app = usePageState();
const model = useModelState();
app.pageTitle = model.selectProcesline.name;
const procesline = computed(() => {
  return model.selectProcesline;
});
if (procesline == null) {
  const router = useRouter();
  router.push("/index");
}
</script>
<style></style>
