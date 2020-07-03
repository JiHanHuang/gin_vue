<template>
  <div>
    <Card>
      <show name="image" image-path="../assets/logo.png"></show>
    </Card>
  </div>
</template>

<script>
import show from '@/components/show'
export default {
  name: 'imageShow',
  data () {
    return {
      value: ''
    }
  },
  components: {show}
}
</script>

<style scoped>
h3 {
  font-weight: normal;
  margin-bottom: 5px;
}
</style>
