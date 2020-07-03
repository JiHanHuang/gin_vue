<template>
 <div>
    <Card style="width:350px">
      <p slot="title">
          <Icon type="ios-film-outline"></Icon>
          {{name}}
      </p>
      <p> {{imagePath}} </p>
      <img :src="imgpath" width="100%" />
      <img :src="image" width="100%" :alt="imagePath"/>
    </Card>
 </div>
</template>

<script>
export default {
  props: ['name', 'imagePath'],
  name: 'show',
  data () {
    return {
      // imgpath: require(this.imagePath), ???? fault
      image: require('../assets/logo.png')
    }
  }
}
</script>
