<template>
  <div class="box">
    <Button type="text" size="large" @click="modal1 = true"><Icon type="md-add" size="30" /></Button>
    <Modal
    v-model="modal1"
    title="创建新的下载"
    icon="md-cloud-download"
    :loading="loading"
    @on-ok="create"
    @on-cancel="cancel">
    <Select v-model="model2" style="width:200px">
        <Option v-for="item in typeList" :value="item.value" :key="item.value">{{ item.label }}</Option>
    </Select>
    <h3>名称:</h3>
    <Input v-model="name" placeholder="name" style="width: 300px" />
    <h3>下载链接:</h3>
    <Input v-model="addr" placeholder="address" style="width: 300px" />
    </Modal>
  </div>
</template>

<script>
import download from '@/components/download'
import axios from 'axios'
export default {
  props: ['selectType'],
  name: 'create',
  data () {
    return {
      value: '',
      modal1: false,
      loading: true,
      model2: this.selectType,
      id: 0,
      name: 'Test',
      addr: '',
      typeList: [
        {
          value: 'thunder',
          label: '迅雷'
        },
        {
          value: 'file',
          label: '文件'
        },
        {
          value: 'other',
          label: '其他'
        }],
      testGet: 'xxx'
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
    create () {
      axios
        .post('/api/v1/download', {
          addr: this.addr,
          downloadPath: './runtime/',
          id: (new Date()).valueOf(),
          name: this.name
        })
        .then(response => {
          console.log(response)
          console.log(response.data)
          this.testGet = response.data.msg
        }).catch(function (error) { // 请求失败处理
          console.log(error)
        })
      this.modal1 = false
    }
  },
  components: {download}
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
