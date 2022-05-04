import Vue from "vue";
import Vuetify from "vuetify/lib/framework";

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: "#2377FF",
        warning: "#FF8D0A",
        primDark: "#16313B",
        primLight: "#F2F5F8",
      },
    },
  },
});
