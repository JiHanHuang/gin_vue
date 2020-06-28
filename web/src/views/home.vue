<template>
 <div>
     <h1>{{ msg }}</h1>
     <button v-on:click="update()">刷新</button>
     <Button v-on:click="update()">刷新</Button>
 </div>
</template>

<script>
export default {
  name: 'home',
  data () {
    return {
      msg: 'home'
    }
  },
  methods: {
    update () {
      this.msg = 'JiHan'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
</style>
