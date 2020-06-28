<template>
 <div>
    <p>{{vueMsg}}</p>
    <Row v-for="l in list" :key="l.id" style="height: 35px">
      <Col span="20">
        <Progress :percent="25" :stroke-width="15" />
      </Col>
      <Col span="1"><Icon type="md-refresh-circle"  size="24" /></Col>
      <Col span="1"><Icon type="md-download" size="24" /></Col>
      <Col span="1"><Icon type="md-remove-circle" size="24" /></Col>
      <Col span="1"><Icon type="md-play" size="24" />{{l.percent}}</Col>
    </Row>
 </div>
</template>

<script>
export default {
  props: ['vueMsg'],
  name: 'download',
  data () {
    return {
      list: [
        {id: 1, percent: '40', status: 'active'},
        {id: 2, percent: '50', status: 'fault'}
      ]
    }
  }
}
</script>
