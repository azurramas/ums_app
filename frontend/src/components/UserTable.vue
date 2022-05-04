<template>
  <v-col>
    <v-card flat rounded="lg">
      <v-card-title>
        <v-row>
          <v-col lg="6">
            <v-text-field
              v-model="search"
              append-icon="mdi-magnify"
              label="Search"
              single-line
              hide-details
            >
            </v-text-field>
          </v-col>
        </v-row>
      </v-card-title>
      <v-data-table :search="search" :headers="headers" :items="users"
        ><template v-slot:[`item.actions`]="{ item }">
          <v-icon color="warning" class="mr-2" @click="toggleEditUser(item.id)">
            mdi-pencil
          </v-icon>
          <v-icon color="primary" class="mr-2" @click="editItem(item)">
            mdi-lock
          </v-icon>

          <v-icon color="error" @click="promptDeleteUser(item)">
            mdi-delete
          </v-icon>
        </template>
      </v-data-table>
      <v-dialog width="600px" v-model="deleteUserDialog">
        <v-card>
          <v-card-title class="primDark--text"
            >Are you sure you want to delete the following user:
          </v-card-title>
          <v-card-text>
            First name: <strong>{{ userToDelete.firstName }}</strong> <br />
            Last name: <strong> {{ userToDelete.lastName }}</strong>
            <br />
            Username: <strong> {{ userToDelete.username }}</strong> <br />
            Email: <strong> {{ userToDelete.email }}</strong>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primDark" @click="deleteUserDialog = false" text
              >Close</v-btn
            >
            <v-btn text color="error" @click="deleteUser()"
              ><v-icon color="error" left> mdi-delete </v-icon>Delete</v-btn
            >
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-card>
  </v-col>
</template>

<script>
export default {
  created() {
    this.$store.dispatch("getUsers");
  },
  computed: {
    users() {
      return this.$store.getters.getUsers;
    },
  },
  data() {
    return {
      search: "",
      headers: [
        {
          text: "First Name",
          value: "firstName",
        },
        {
          text: "Last Name",
          value: "lastName",
        },
        {
          text: "Username",
          value: "username",
        },
        { text: "Email", value: "email" },
        { text: "Status", value: "status" },
        { text: "Actions", value: "actions", sortable: false, align: "center" },
      ],
      deleteUserDialog: false,
      userToDelete: {},
    };
  },
  methods: {
    promptDeleteUser(item) {
      this.deleteUserDialog = true;
      this.userToDelete = item;
    },
    deleteUser() {
      this.$store.dispatch("deleteUser", this.userToDelete.id).then(() => {
        this.deleteUserDialog = false;
        this.$store.dispatch("getUsers");
      });
    },
    toggleEditUser(id) {
      this.$router.push({ name: "Edit User", params: { id: id } });
    },
  },
};
</script>
