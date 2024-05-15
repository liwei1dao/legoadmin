<template>
    <div ref="frameContainer" class="absolute top-0 left-0 right-0 h-full">
    <canvas
        v-bind="canvasAttr"
        ref="canvas"
    >
    </canvas>
  </div>
</template>

<script>
export default {
    name:"VideoFrame"
}
</script>

<style>

</style>