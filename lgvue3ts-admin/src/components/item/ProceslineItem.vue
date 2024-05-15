<template>
  <v-card variant="outlined" height="60">
    <v-row class="pl-3" align="center" no-gutters style="height: 60px">
      <v-col cols="auto"> <strong>Name:</strong>{{ procesline.name }} </v-col>
      <v-spacer></v-spacer>
      <v-col cols="auto">
        <v-btn class="mx-2" fab dark color="brown" @click="goeditor">
          <v-icon dark> mdi-pencil </v-icon>
          编辑
        </v-btn>
      </v-col>
    </v-row>
  </v-card>
</template>
<script setup lang="ts">
import { useRouter } from "vue-router";
import { useModelState } from "@/stores/modelState";
const props = defineProps({
  procesline: {
    type: Object as () => Procesline,
    required: true,
  },
});

const router = useRouter();
const model = useModelState();
const goeditor = () => {
  model.selectProcesline = props.procesline;
  router.push({
    path: "/editor",
  });
};
</script>
<style lang=""></style>
