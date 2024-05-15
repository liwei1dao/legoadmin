<template>
  <v-app>
    <v-app-bar color="deep-purple accent-4" dense dark>
      <v-toolbar-title to="/index">AITools</v-toolbar-title>
      <v-spacer></v-spacer>
      <div class="flex flex-row-reverse pr-10 items-center">
        <v-switch v-model="store.isDark" hide-details></v-switch>
      </div>
    </v-app-bar>
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>
<script setup lang="ts">
import { ref, computed } from "vue";
import { usePageState } from "@/stores/pageState";
const store = usePageState();
</script>
<style></style>
