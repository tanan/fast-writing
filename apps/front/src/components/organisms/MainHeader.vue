<template>
  <div class="header">
    <v-container class="ma-0 pa-2">
      <v-row no-gutters>
        <v-col cols="12" sm="3"><MainLogo /></v-col>
        <v-col cols="10" offset="1" sm="4" class="d-flex align-center"><HeaderForm /></v-col>
        <v-spacer></v-spacer>
        <v-col>
          <div class="hamburger-menu">
            <input type="checkbox" id="menu-btn-check">
            <label for="menu-btn-check" class="menu-btn"><span></span></label>
            <div class="menu-content">
              <ul>
                  <li>
                      <a href="#">メニューリンク1</a>
                  </li>
                  <li>
                      <a href="#">メニューリンク2</a>
                  </li>
                  <li>
                      <a href="#">メニューリンク3</a>
                  </li>
              </ul>
            </div>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </div>    
</template>

<script>
import MainLogo from '@/components/atoms/MainLogo.vue'
import HeaderForm from '@/components/molecules/HeaderForm.vue'

export default {
  name: 'MainHeader',
  components: {
    MainLogo,
    HeaderForm,
  },
  props: {
    page: String
  }
}
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
.header {
  background-color: #55c500;
}

.menu-btn {
  position: fixed;
  top: 0px;
  right: 0px;
  display: flex;
  height: 64px;
  width: 64px;
  justify-content: center;
  align-items: center;
  z-index: 90;
  
  span:before {
    bottom: 8px;
  }

  span:after {
    top: 8px;
  }
  
  span, span:before, span:after {
    content: '';
    display: block;
    height: 3px;
    width: 25px;
    border-radius: 3px;
    background-color: #ffffff;
    position: absolute;
  }
}

#menu-btn-check:checked ~ .menu-btn span {
    background-color: rgba(255, 255, 255, 0);/*メニューオープン時は真ん中の線を透明にする*/
}
#menu-btn-check:checked ~ .menu-btn span::before {
    bottom: 0;
    transform: rotate(45deg);
}
#menu-btn-check:checked ~ .menu-btn span::after {
    top: 0;
    transform: rotate(-45deg);
}

#menu-btn-check {
    display: none;
}

.menu-content {
  width: 100%;
  height: 100%;
  position: fixed;
  top: 0;
  left: 100%;
  z-index: 80;
  background-color: #3584bb;
  transition: all 0.5s;

  ul {
    padding: 70px 10px 0;
  }
  ul li {
    border-bottom: solid 1px #ffffff;
    list-style: none;
  }
  ul li a {
    display: block;
    width: 100%;
    font-size: 15px;
    box-sizing: border-box;
    color:#ffffff;
    text-decoration: none;
    padding: 9px 15px 10px 0;
    position: relative;

    &:before {
      content: "";
      width: 7px;
      height: 7px;
      border-top: solid 2px #ffffff;
      border-right: solid 2px #ffffff;
      transform: rotate(45deg);
      position: absolute;
      right: 11px;
      top: 16px;
    }
  }
}

#menu-btn-check:checked ~ .menu-content {
    left: 0;/*メニューを画面内へ*/
}
</style>