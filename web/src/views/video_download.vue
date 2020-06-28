<template>
  <div>
    <Card>
    <h3>当前路径:</h3>
    <Input v-model="value" placeholder="Enter something..." style="width: 300px" />
    <Button>修改路径</Button>
    </Card>
    <Card>
      <download vue-msg="video"></download>
    </Card>
  </div>
</template>

<script>
import download from '@/components/download'
export default {
  name: 'videoDownload',
  data () {
    return {
      value: ''
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
</style>
