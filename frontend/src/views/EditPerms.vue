a
<template>
  <v-card flat rounded="lg">
    <v-card-title class="primDark--text">
      Edit Permissions for User
    </v-card-title>
    <v-card-text>
      <v-form ref="userForm" class="mt-6">
        <v-row v-if="loadReady">
          <v-col
            v-for="perm in permissions"
            :key="perm.id"
            cols="12"
            sm="6"
            md="4"
          >
            <v-switch
              v-model="perm.status"
              :label="getPermLabel(perm)"
              hide-details
              @change="updatePerm(perm)"
            ></v-switch></v-col
        ></v-row>
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer> </v-spacer>
      <v-btn
        text
        @click="$router.push('/')"
        class="rounded-lg mr-2 primDark--text"
        elevation="0"
      >
        Go to Home Page</v-btn
      >
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      loadReady: false,
    };
  },
  created() {
    this.$store.dispatch("getPermissions").then(() => {
      this.$store
        .dispatch("getUserPermissions", this.$route.params.id)
        .then(() => {
          this.initPerms();
        });
    });
  },
  computed: {
    permissions() {
      return this.$store.getters.getPerms || [];
    },
    userPermissions() {
      return this.$store.getters.getUserPerms || [];
    },
  },
  methods: {
    initPerms() {
      this.permissions.forEach((p) => {
        this.userPermissions.forEach((up) => {
          if (p.id === up.permID) {
            p.status = true;
          }
        });
      });
      this.loadReady = true;
    },
    updatePerm(perm) {
      let payload = {
        userID: this.$route.params.id,
        permID: perm.id,
      };

      if (perm.status) {
        this.$store.dispatch("addUserPerm", payload);
      } else {
        this.$store.dispatch("deleteUserPerm", payload);
      }
    },
    getPermLabel(perm) {
      return perm.code + ": " + perm.desc;
    },
  },
};
</script>

<style></style>
