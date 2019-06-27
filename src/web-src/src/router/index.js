import Vue from 'vue'
import Router from 'vue-router'

import Index from '@/components/Index'
import ChatRoom from '@/components/ChatRoom'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index
    }, {
      path: '/ws',
      name: 'ChatRoom',
      component: ChatRoom
    }
  ]
})
