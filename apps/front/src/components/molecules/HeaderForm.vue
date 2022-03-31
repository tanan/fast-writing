<template>
  <div class="header-form">
    <form action="/search">
      <div class="search-box">
        <input name="q" type="text" autocomplete="off" placeholder="カテゴリを検索">
        <mdicon class="icon" name="magnify" :width="30" :height="30" />
      </div>
    </form>
  </div>    
</template>

<script>
export default {
  name: 'HeaderForm',
}
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
.header-form {
  width: 100%;
  .search-box {
    position: relative;
    input {
      background-color: white;
      border-radius: 4px;
      width: 100%;
      height: 36px;
      padding-left: 36px;
    }
    .icon {
      position: absolute;
      left: 4px;
      top: 4px;
      &:before {
        content: "";
      }
    }
  }
}
</style>