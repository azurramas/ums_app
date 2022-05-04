import axios from "axios";
import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    apiUrl: "http://localhost:3010",
    snackbar: {
      status: false,
      color: "",
      desc: "",
    },
    users: [],
    user: {},
    permissions: [],
    userPermissions: [],
  },
  mutations: {
    snackbar(state, payload) {
      state.snackbar = payload;
    },
    setUsers(state, payload) {
      state.users = payload;
    },
    setUser(state, payload) {
      state.user = payload;
    },
    setPerms(state, payload) {
      state.permissions = payload;
    },
    setUserPerms(state, payload) {
      state.userPermissions = payload;
    },
  },
  actions: {
    getUsers(context) {
      return new axios.get(this.state.apiUrl + "/users")
        .then((res) => {
          context.commit("setUsers", res.data);
        })
        .catch((err) => {
          let snackbar = {
            status: true,
            color: "error",
            desc: err.response.data.error + ": " + err.response.data.desc,
          };
          context.commit("snackbar", snackbar);
        });
    },
    addUser(context, payload) {
      return new axios.post(this.state.apiUrl + "/user", payload)
        .then(() => {
          let snackbar = {
            status: true,
            color: "success",
            desc: "Successfully added user.",
          };
          context.commit("snackbar", snackbar);
        })
        .catch((err) => {
          let snackbar = {
            status: true,
            color: "error",
            desc: err.response.data.error + ": " + err.response.data.desc,
          };
          context.commit("snackbar", snackbar);
        });
    },
    editUser(context, payload) {
      return new axios.put(this.state.apiUrl + "/user", payload)
        .then(() => {
          let snackbar = {
            status: true,
            color: "success",
            desc: "Successfully saved changes.",
          };
          context.commit("snackbar", snackbar);
        })
        .catch((err) => {
          let snackbar = {
            status: true,
            color: "error",
            desc: err.response.data.error + ": " + err.response.data.desc,
          };
          context.commit("snackbar", snackbar);
        });
    },
    getUser(context, payload) {
      return new axios.get(this.state.apiUrl + "/user/" + payload)
        .then((res) => {
          context.commit("setUser", res.data);
        })
        .catch((err) => {
          let snackbar = {
            status: true,
            color: "error",
            desc: err.response.data.error + ": " + err.response.data.desc,
          };
          context.commit("snackbar", snackbar);
        });
    },
    deleteUser(context, payload) {
      return new axios.delete(this.state.apiUrl + "/user/" + payload)
        .then(() => {
          let snackbar = {
            status: true,
            color: "success",
            desc: "Successfully deleted user.",
          };
          context.commit("snackbar", snackbar);
        })
        .catch((err) => {
          let snackbar = {
            status: true,
            color: "error",
            desc: err.response.data.error + ": " + err.response.data.desc,
          };
          context.commit("snackbar", snackbar);
        });
    },
    getPermissions(context) {
      return new axios.get(this.state.apiUrl + "/perms")
        .then((res) => {
          context.commit("setPerms", res.data);
        })
        .catch((err) => {
          let snackbar = {
            status: true,
            color: "error",
            desc: err.response.data.error + ": " + err.response.data.desc,
          };
          context.commit("snackbar", snackbar);
        });
    },
    addUserPerm(context, payload) {
      return new axios.post(
        this.state.apiUrl + "/perms/" + payload.userID + "/" + payload.permID
      )
        .then(() => {
          let snackbar = {
            status: true,
            color: "success",
            desc: "Successfully updated permissions.",
          };

          context.commit("snackbar", snackbar);
        })
        .catch((err) => {
          let snackbar = {
            status: true,
            color: "error",
            desc: err.response.data.error + ": " + err.response.data.desc,
          };
          context.commit("snackbar", snackbar);
        });
    },
    deleteUserPerm(context, payload) {
      return new axios.delete(
        this.state.apiUrl + "/perms/" + payload.userID + "/" + payload.permID
      )
        .then(() => {
          let snackbar = {
            status: true,
            color: "success",
            desc: "Successfully updated permissions.",
          };

          context.commit("snackbar", snackbar);
        })
        .catch((err) => {
          let snackbar = {
            status: true,
            color: "error",
            desc: err.response.data.error + ": " + err.response.data.desc,
          };
          context.commit("snackbar", snackbar);
        });
    },
    getUserPermissions(context, payload) {
      return new axios.get(this.state.apiUrl + "/perms/" + payload)
        .then((res) => {
          context.commit("setUserPerms", res.data);
        })
        .catch((err) => {
          let snackbar = {
            status: true,
            color: "error",
            desc: err.response.data.error + ": " + err.response.data.desc,
          };
          context.commit("snackbar", snackbar);
        });
    },
  },
  getters: {
    getUsers(state) {
      return state.users;
    },
    getUser(state) {
      return state.user;
    },
    getSnackbar(state) {
      return state.snackbar;
    },
    getPerms(state) {
      return state.permissions;
    },
    getUserPerms(state) {
      return state.userPermissions;
    },
  },
});
