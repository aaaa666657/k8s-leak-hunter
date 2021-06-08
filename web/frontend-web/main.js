import index from './index.js'
import host from './host.js'
import service from './service.js'
const router = new VueRouter({
    routes: [
      {
        path: '/', 
        component: index
      },
      {
        path: '/service',
        component: service
      },
      {
        path: '/host',
        component: host
      }
    ]
  })
  
  new Vue({
    router: router
  }).$mount('#app')