<template>
  <v-app>
    <v-app-bar flat app color="primary" dark>
      <v-toolbar-title>
        <strong>User Management System</strong>
      </v-toolbar-title>
    </v-app-bar>
    <v-snackbar
      multi-line
      bottom
      right
      :color="snackbar.color"
      v-model="snackbar.status"
      timeout="3500"
    >
      <strong>{{ snackbar.desc }}</strong>
    </v-snackbar>
    <v-main class="primLight">
      <v-container class="pt-8"> <router-view></router-view></v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: "App",

  computed: {
    snackbar() {
      return this.$store.getters.getSnackbar;
    },
  },
};
</script>
