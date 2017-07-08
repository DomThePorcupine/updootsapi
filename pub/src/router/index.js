import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import VueMaterial from 'vue-material'
import VueResource from 'vue-resource'
import 'vue-material/dist/vue-material.css'

Vue.use(Router)
Vue.use(VueMaterial)
Vue.use(VueResource)

Vue.material.registerTheme('default', {
  primary: 'blue'
})

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    }
  ]
})
