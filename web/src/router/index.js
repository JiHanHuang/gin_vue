import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import home from '@/views/home'
import videoDownload from '@/views/video_download'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/hello',
      component: Hello
    },
    {
      path: '/home',
      component: home
    },
    {
      path: '/video_download',
      component: videoDownload
    }
  ]
})
