<template>
  <router-view />
</template>

<script>
import { inject } from 'vue';
export default {
  name: 'App',
  components: {
  },
  data: () => ({
    //
  }),
  setup() {
    const websocket = inject('$websocket');
    websocket.init()
  },
}
</script>
