<template>
  <v-container>
    <v-card>
      <v-row>
        <v-col cols="12" v-for="(procesline, i) in proceslines" :key="i">
          <ProceslineItem :procesline="procesline"></ProceslineItem>
        </v-col>
      </v-row>
    </v-card>
  </v-container>
</template>
<script setup lang="ts">
import ProceslineItem from "@/components/item/ProceslineItem.vue";
import { computed } from "vue";
import { useModelState } from "@/stores/modelState";
import { usePageState } from "@/stores/pageState";
const app = usePageState();
app.pageTitle = "主页";
const model = useModelState();
const proceslines = computed(() => model.proceslines);
</script>

<style></style>
