<template>
  <div class="content-list">
    <a class="home" v-bind:class="{ active: isHome }" href="/">ホーム</a>
    <a class="about" v-bind:class="{ active: isAbout }" href="/about/4MpKUQwtphCr4Qvil1ub75">about</a>
    <a class="profile" v-bind:class="{ active: isProfile }" href="/profile">profile</a>
    <a class="health-care" v-bind:class="{ active: isHealthCare }" href="/health-care">health care</a>
  </div>    
</template>

<script>
export default {
  name: 'HeaderContent',
  props: {
    page: String
  },
  data () {
    return {
      isHome: false,
      isAbout: false,
      isProfile: false,
      isHealthCare: false,
    }
  },
  created() {
    switch(this.page) {
      case 'home':
        this.isHome = true
        break
      case 'about':
        this.isAbout = true
        break
      case 'profile':
        this.isProfile = true
        break
      case 'health-care':
        this.isHealthCare = true
        break
      default:
    }
  },
}
</script>

<style scoped lang="scss">
  .content-list {
    height: 48px;
    padding-left: 24px;
    display: flex;
    flex-grow: 0;
    align-items: stretch;
    
    a {
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 0 20px;
      font-size: 16px;
      text-decoration: none;
      font-weight: 600;
      color: white;
      opacity: 0.3;

      &:hover {
        opacity: 1;
      }
    }
    .active {
      opacity: 1;
      border-bottom: solid 3px #408100;
    }
  }
</style>