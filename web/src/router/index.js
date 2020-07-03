import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import home from '@/views/home'
import videoDownload from '@/views/video_download'
import thunderDownload from '@/views/thunder_download'
import imageShow from '@/views/image_show'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/hello',
      component: Hello,
      meta: [
        {name: 'Hello'}
      ]
    },
    {
      path: '/home',
      component: home,
      meta: [
        {name: 'Home', icon: 'ios-home'}
      ]
    },
    {
      path: '/show/image_show',
      name: 'imageView',
      component: imageShow,
      meta: [
        {name: 'Show', icon: 'md-download', url: '/home'},
        {name: 'Image'}
      ]
    },
    {
      path: '/download/video_download',
      component: videoDownload,
      meta: [
        {name: 'Download', icon: 'md-download', url: '/home'},
        {name: 'Video'}
      ]
    },
    {
      path: '/download/thunder_download',
      component: thunderDownload,
      meta: [
        {name: 'Download', icon: 'md-download', url: '/home'},
        {name: 'Thunder'}
      ]
    }
  ],
  mode: 'history'
})
