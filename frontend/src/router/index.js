import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: () =>
      import(/* webpackChunkName: "home" */ "../views/UserList.vue"),
  },
  {
    path: "/create-user",
    name: "Create User",
    component: () =>
      import(/* webpackChunkName: "createuser" */ "../views/CreateUser.vue"),
  },
  {
    path: "/edit-user/:id",
    name: "Edit User",
    component: () =>
      import(/* webpackChunkName: "edituser" */ "../components/EditUser.vue"),
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
