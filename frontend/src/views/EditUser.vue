a
<template>
  <v-card flat rounded="lg">
    <v-card-title class="primDark--text"> Edit User </v-card-title>
    <v-card-text>
      <v-form ref="userForm" class="mt-6">
        <v-row>
          <v-col cols="12" sm="6" lg="4">
            <v-text-field
              class="rounded-lg"
              outlined
              v-model="user.username"
              label="Username"
              :rules="rules"
            ></v-text-field>
          </v-col>
          <v-col cols="12" sm="6" lg="4">
            <v-text-field
              disabled
              class="rounded-lg"
              outlined
              v-model="user.password"
              :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
              :type="showPassword ? 'text' : 'password'"
              label="Password"
              @click:append="showPassword = !showPassword"
            >
            </v-text-field>
          </v-col>
          <v-col cols="12" sm="6" lg="4">
            <v-text-field
              :rules="rules"
              class="rounded-lg"
              outlined
              v-model="user.firstName"
              label="First Name"
            ></v-text-field>
          </v-col>
          <v-col cols="12" sm="6" lg="4">
            <v-text-field
              class="rounded-lg"
              :rules="rules"
              outlined
              v-model="user.lastName"
              label="Last Name"
            ></v-text-field>
          </v-col>
          <v-col cols="12" sm="6" lg="4">
            <v-text-field
              :rules="rules"
              class="rounded-lg"
              outlined
              v-model="user.email"
              label="Email"
            ></v-text-field>
          </v-col>
          <v-col cols="12" sm="6" lg="4">
            <v-select
              class="rounded-lg"
              :rules="rules"
              outlined
              v-model="user.status"
              label="Status"
              :items="statuses"
            ></v-select>
          </v-col>
        </v-row>
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
      <v-btn
        text
        @click="saveChanges()"
        class="rounded-lg"
        dark
        elevation="0"
        color="primary"
      >
        <v-icon left>mdi-check</v-icon> Save Changes</v-btn
      >
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      showPassword: false,
      rules: [(v) => !!v || "This value is required"],
      statuses: ["Not active", "Active", "Pending", "Suspended"],
    };
  },
  created() {
    this.$store.dispatch("getUser", this.$route.params.id);
  },
  computed: {
    user() {
      return this.$store.getters.getUser;
    },
  },
  methods: {
    saveChanges() {
      if (this.$refs.userForm.validate()) {
        this.$store.dispatch("editUser", this.user);
      }
    },
  },
};
</script>

<style></style>
