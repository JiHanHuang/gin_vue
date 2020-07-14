<template>
 <div>
    <p>{{list}}</p>
    <p>{{vueMsg}}</p>
    <Row v-for="l in list" :key="l.id" style="height: 35px">
      <Col span="2">
        <Tooltip max-width="400" :content="l.name">
        <p style="width: 100px;overflow: hidden;white-space: nowrap;text-overflow: ellipsis;">
          {{l.name}}
        </p>
        </Tooltip>
      </Col>
      <Col span="18">
        <div v-if="l.status === ''">
        <Progress :percent="l.percent" :stroke-width="15"/>
        </div>
        <div v-else>
        <Progress :percent="l.percent" :stroke-width="15" :status="l.status" />
        </div>
      </Col>
      <Col span="1"><Icon type="md-refresh-circle"  size="24" /></Col>
      <Col span="1"><Icon type="md-download" size="24" /></Col>
      <Col span="1"><Icon type="md-remove-circle" size="24" /></Col>
      <Col span="1"><Icon type="md-play" size="24" /></Col>
    </Row>
 </div>
</template>

<script>
import axios from 'axios'
export default {
  props: ['vueMsg', 'updateFlag'],
  name: 'download',
  data () {
    return {
      timer: '',
      value: '',
      list: []
    }
  },
  methods: {
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
      axios
        .get('/api/v1/download/list')
        .then(response => {
          console.log(response)
          console.log(response.data)
          this.list = response.data.data.list
        }).catch(function (error) { // 请求失败处理
          console.log(error)
        })
    }
  },
  mounted () {
    this.timer = setInterval(this.getDownloadList, 2000)
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
