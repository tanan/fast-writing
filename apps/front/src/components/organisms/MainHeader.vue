<template>
  <div>
    <div>
      <Menubar class="menu" :model="items">
        <template #start>
          <img alt="logo" src="https://www.primefaces.org/wp-content/uploads/2020/05/placeholder.png" height="40" class="mr-2">
        </template>
        <template #end>
          <Button icon="pi pi-user" class="p-button-rounded p-button-secondary mr-2" @click="toggleAccountMenu()" />
        </template>
      </Menubar>
    </div>
    <div v-show="isDisplay" class="absolute top-1 right-0">
      <Menu :model="accountItems" @click="disableAccountMenu()" />
    </div>
  </div>
</template>

<script>
import { defineComponent, ref, watch } from "vue"
import { useRouter } from 'vue-router'
import Store from '@/store/index.js'
import Menubar from 'primevue/menubar'
import Menu from 'primevue/menu'
import Button from 'primevue/button'

export default defineComponent ({
  name: 'MainHeader',
  components: {
    Menubar,
    Menu,
    Button,
  },
  setup () {
    const router = useRouter()

    const isLoggedIn = ref(Store.getters.isLoggedIn)
    watch([isLoggedIn], () => {
      console.log(isLoggedIn.value)
    })
    const isDisplay = ref(false)
    const items = ref([
      {
        label:'Contents',
        icon:'pi pi-fw pi-file',
        items: [
          {
            label:'List',
            icon:'pi pi-fw pi-align-justify',
            to: '/',
          },
          {
            label:'Create',
            icon:'pi pi-fw pi-plus',
            to: '/contents/create',
          }
        ]
      },
    ])
    const accountItems = [
      {
        label: 'Account',
        icon: 'pi pi-fw pi-user',
        items:[
          {
            label: 'Sign In',
            visible: !isLoggedIn.value,
            to: '/signin'
          },
          {
            label: 'Sign Up',
            visible: !isLoggedIn.value,
            to: '/signup'
          },
          {
            label: 'Profile',
            visible: isLoggedIn.value,
            to: '/account'
          },
          {
            label: 'Sign Out',
            visible: isLoggedIn.value,
            command: () => {
              signout()
            }
          }
        ]
      }
    ]

    const disableAccountMenu = () => {
      isDisplay.value = false
    }

    const toggleAccountMenu = () => {
      isDisplay.value = !isDisplay.value
    }

    const signout = () => {
      Store.dispatch('signout')
      router.push('/signin')
    }

    return {
      isDisplay,
      items,
      accountItems,
      disableAccountMenu,
      toggleAccountMenu,
      signout,
    }
  }
})
</script>

<style lang="scss">
  .menu {
    outline: none;
  }
</style>