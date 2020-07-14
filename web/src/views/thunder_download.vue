<template>
  <div>
    <Card>
    <Row>
      <Col span="6">
      </Col>
      <Col span="6"><h3>当前路径:</h3></Col>
      <Col span="6"></Col>
      <Col span="4"><create select-type="thunder"></create></Col>
    </Row>
    </Card>
    <Card>
      <download vue-msg="thunder"></download>
    </Card>
  </div>
</template>

<script>
import download from '@/components/download'
import create from '@/components/create'
export default {
  name: 'thunderDownload',
  data () {
    return {
      value: ''
    }
  },
  components: {download, create}
}
</script>

<style scoped>
h3 {
  font-weight: normal;
  margin-bottom: 5px;
}
</style>
