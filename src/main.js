import Vue from "vue";
import App from "./App.vue";
import Vuetify from "vuetify";
import "vuetify/dist/vuetify.css";
import router from "./router";
import store from "./store";

Vue.use(Vuetify);

Vue.config.productionTip = false;

const vuetify = new Vuetify({
  // Vuetifyの設定をここに追加することもできます。
});

new Vue({
  vuetify,
  router,
  store,
  // 他のオプションも追加できます。
  render: (h) => h(App),
}).$mount("#app");
