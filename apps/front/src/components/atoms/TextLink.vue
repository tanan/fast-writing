<template>
    <li><a :href="url" target="_blank" rel="noopener">{{ text }}</a></li>
</template>

<script>
export default {
  name: 'TextLink',
  props: {
    url: String,
    text: String
  }
}
</script>


<style scoped lang="scss">
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
