// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Element from 'element-ui'
import 'element-ui/lib/theme-default/index.css'
import axios from 'axios'
import VueAxios from 'vue-axios'
import '@progress/kendo-theme-default/dist/all.css'
import '@progress/kendo-ui'
import { Chat, ChatInstaller } from '@progress/kendo-chat-vue-wrapper'
Vue.use(ChatInstaller)
Vue.use(VueAxios, axios)
Vue.use(Element)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: {
    App,
    Chat},
  template: '<App/>'
})
