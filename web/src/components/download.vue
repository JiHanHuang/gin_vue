<template>
 <div>
    <p>{{list}}</p>
    <p>{{vueMsg}}</p>
    <Row v-for="l in list" :key="l.id" style="height: 35px" type="flex" justify="center" align="middle">
      <Col span="2">
        <Tooltip max-width="400" :content="l.name">
        <p style="width: 100px;overflow: hidden;white-space: nowrap;text-overflow: ellipsis;">
          {{l.name}}
        </p>
        </Tooltip>
      </Col>
      <Col span="18">
        <div v-if="l.status == ''">
        <Progress :percent="l.percent" :stroke-width="15"/>
        </div>
        <div v-else>
        <Progress :percent="l.percent" :stroke-width="15" :status="l.status" />
        </div>
      </Col>
      <Col span="1">
        <Button type="text" icon="md-refresh-circle" @click="info('circle-' + l.id)"></Button>
      </Col>
      <Col span="1">
        <!--<Button type="text" icon="md-download" @click="getFile(l.id)"></Button> -->
        <router-link :to="{path:'/download/local/file',query:{id:l.id}}" target="_blank">
          <Button type="text" icon="md-download"></Button>
        </router-link>
      </Col>
      <Col span="1">
        <Button type="text" icon="md-remove-circle" @click="info('remove')"></Button>
      </Col>
      <Col span="1">
        <Button type="text" icon="md-play" @click="info('play')"></Button>
      </Col>
    </Row>
 </div>
</template>

<script>
import event from '../libs/event.js'
import axios from 'axios'
export default {
  props: ['vueMsg', 'updateFlag'],
  name: 'download',
  data () {
    return {
      timer: '',
      value: '',
      getdata: true,
      list: []
    }
  },
  methods: {
    info (msg) {
      console.log(msg)
      this.$Message.info(msg)
    },
    asyncOK () {
      setTimeout(() => {
        this.modal1 = false
      }, 1000)
      this.$Message.info(this.model2)
    },
    cancel () {
      this.$Message.info('Clicked cancel')
    },
    getDownloadList () {
      if (this.list != null) {
        for (var i = 0; i < this.list.length; i++) {
          if (this.list[i].percent < 100) {
            this.getdata = true
            break
          }
        }
      }
      if (!this.getdata) {
        return
      }
      axios
        .get('/api/v1/download/list')
        .then(response => {
          console.log(response)
          console.log(response.data)
          this.list = response.data.data.list
        }).catch(error => { // 请求失败处理
          this.$Message.error(error.response.data.data)
          console.log(error)
        })
      this.getdata = false
    },
    getFile (id) {
      console.log(id)
      // router.go({name: 'api/v1/getfile', params: {id: id}})
    }
  },
  mounted () {
    this.getDownloadList()
    event.$on('flashList', (val) => {
      this.getdata = true
    })
    this.timer = setInterval(this.getDownloadList, 2000)
  },
  destroyed () {
    event.$off('flashList')
  },
  beforeDestroy () {
    clearInterval(this.timer)
  }
}
</script>

<style scoped>
h3 {
  font-weight: normal;
  margin-bottom: 5px;
}
.box{
  float: left;
}
</style>
