<template>
  <div ref="iconList">
    <span
      class="w-full h-5 sticky top-0 left-0 right-0 z-20 dark:bg-gray-800 bg-gray-50"
    ></span>
  </div>
</template>
<script setup lang="ts">
import { ref, watch } from "vue";
const props = defineProps({
  listData: {
    type: Array,
    default() {
      return [];
    },
  },
  offsetTop: {
    type: Number,
    default: 0,
  },
});

const iconList = ref();
watch(
  () => props.offsetTop,
  () => {
    iconList.value.scrollTop = props.offsetTop;
  }
);
</script>
<style lang=""></style>
