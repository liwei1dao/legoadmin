<template>
  <v-card variant="outlined" height="60">
    <v-row class="pl-3" align="center" no-gutters style="height: 60px">
      <v-col cols="auto"> <strong>ID:</strong>{{ product._id }} </v-col>
      <v-spacer></v-spacer>
      <v-col cols="auto">
        <v-btn class="mx-2" fab dark color="brown" @click="editor">
          <v-icon dark> mdi-pencil </v-icon>
          编辑
        </v-btn>
      </v-col>
    </v-row>
  </v-card>
</template>

<script>
import { useRouter } from "vue-router";

export default {
  name: "ProductItem",
  props: {
    product: {
      type: Object,
      default() {
        return {
          titile: "测试产品",
        };
      },
    },
  },
  setup() {
    const router = useRouter();
    return {
      router,
    };
  },
  methods: {
    editor() {
      console.log("编辑");
      this.router.push({
        path: "/add",
        query: {
          _id: this.product._id,
        },
      });
    },
  },
};
</script>

<style></style>
