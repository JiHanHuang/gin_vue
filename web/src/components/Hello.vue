<template>
 <div>
     <h1>{{ msg }}</h1>
     <button v-on:click="update()">刷新</button>
     <Button v-on:click="update()">刷新</Button>
 </div>
</template>

<script>
export default {
  name: 'Hello',
  data () {
    return {
      msg: 'Hello'
    }
  },
  methods: {
    update () {
      this.msg = 'hello jihan'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
