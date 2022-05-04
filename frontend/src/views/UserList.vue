<template>
  <div>
    <v-row>
      <v-col> <h2 class="primDark--text">User List</h2></v-col>
      <v-col class="d-flex justify-end">
        <!-- <v-btn
          @click="$router.push('/create-user')"
          class="rounded-lg mr-4"
          dark
          large
          elevation="0"
          color="primDark"
        >
          <v-icon left>mdi-account-plus</v-icon> Add a User (new page)</v-btn
        > -->
        <v-btn
          @click="dialog = true"
          class="rounded-lg"
          dark
          large
          elevation="0"
          color="primDark"
        >
          <v-icon left>mdi-account-plus</v-icon> Add a User</v-btn
        ></v-col
      >
    </v-row>
    <v-row>
      <UserTable />
    </v-row>
    <v-dialog persistent class="rounded-lg" v-model="dialog" width="1200px">
      <CreateUser @closeDialog="dialog = false" />
    </v-dialog>
  </div>
</template>

<script>
import UserTable from "../components/UserTable.vue";
import CreateUser from "./CreateUser.vue";
export default {
  components: {
    UserTable,
    CreateUser,
  },
  data() {
    return {
      dialog: false,
    };
  },
};
</script>
