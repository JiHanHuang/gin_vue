<template>
  <div>
    <create select-type="thunder"></create>
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
