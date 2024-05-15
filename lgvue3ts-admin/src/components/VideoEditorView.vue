<template>
  <v-layout class="fill-height">
    <v-app-bar color="grey-lighten-4" density="compact" elevation="0">
      <v-app-bar-title>编辑面板</v-app-bar-title>
      <template v-slot:append>
        <v-btn icon="mdi-dots-vertical"></v-btn>
      </template>
    </v-app-bar>
    <v-main>
      <v-container fluid>
        <TrackContainer></TrackContainer>
      </v-container>
    </v-main>
  </v-layout>
</template>
<script setup lang="ts">
import TrackContainer from "@/components/container/TrackContainer.vue";
</script>
<style lang=""></style>
