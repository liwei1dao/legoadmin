<template>
  <v-app>
    <v-app-bar       
      color="deep-purple accent-4"
      dense
      dark>
      <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>AITools</v-toolbar-title>
      <v-spacer></v-spacer>
    </v-app-bar>
    <v-navigation-drawer
        v-model="drawer"
        :location="$vuetify.display.mobile ? 'bottom' : undefined"
        temporary
      >
        <v-list
          :items="items"
        ></v-list>
      </v-navigation-drawer>
    <v-main>
        <router-view></router-view>   
    </v-main>
    <v-footer
      class="bg-indigo-lighten-1 text-center d-flex flex-column"
    >
      <v-container>
        <div>
          <v-btn
            v-for="icon in icons"
            :key="icon"
            class="mx-4"
            :icon="icon"
            variant="text"
          ></v-btn>
        </div>

        <div class="pt-0">
          欢迎来到我们的网站！感谢您的光临！我们致力于为您提供优质的服务和丰富的内容。如果您有任何问题或建议，请随时与我们联系。祝您在我们的网站上度过愉快的时光！
        </div>

        <v-divider></v-divider>

        <div class="text-center">
          {{ new Date().getFullYear() }} — <strong>AiTool</strong>
        </div>
      </v-container>
    </v-footer>
  </v-app>
</template>

<script>
export default {
 name: 'DefLayout',
 data: () => { return {
  drawer:false,
 }},
}
</script>

<style>

</style>