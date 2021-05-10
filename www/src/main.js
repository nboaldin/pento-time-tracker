import Vue from "vue";
import App from "./App.vue";
import vuetify from "./plugins/vuetify";
import axios from "axios";
import VueAxios from "vue-axios";

Vue.config.productionTip = false;

Vue.use(VueAxios, axios);
axios.defaults.baseURL = "http://localhost:8080/api";
axios({
  withCredentials: true,
  // credentials: 'same-origin',
  headers: { "content-type": "application/x-www-form-urlencoded" }
});

new Vue({
  vuetify,
  render: h => h(App)
}).$mount("#app");
