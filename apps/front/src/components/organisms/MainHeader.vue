<template>
  <div>
    <div>
      <Menubar :model="items">
        <template #start>
          <img alt="logo" src="https://www.primefaces.org/wp-content/uploads/2020/05/placeholder.png" height="40" class="mr-2">
        </template>
        <template #end>
          <Button icon="pi pi-user" class="p-button-secondary p-button-rounded" @click="toggleAccountMenu()" />
        </template>
      </Menubar>
    </div>
    <div v-show="isDisplay" class="absolute top-1 right-0">
      <Menu :model="accountItems" />
    </div>
  </div>
</template>

<script>
import Menubar from 'primevue/menubar'
import Menu from 'primevue/menu'
import Button from 'primevue/button'
export default {
  name: 'MainHeader',
  components: {
    Menubar,
    Menu,
    Button,
  },
  props: {
    isLoggedIn: Boolean
  },
  methods: {
    toggleAccountMenu() {
      this.isDisplay = !this.isDisplay
    },
    signout() {
      this.$store.dispatch('signout')
      this.$router.push('/signin')
    }
  },
  data() {
    return {
      isDisplay: false,
      accountItems: [
        {
          label: 'Account',
          icon: 'pi pi-fw pi-user',
          items:[
            {
              label: 'Sign In',
              visible: !this.isLoggedIn,
              to: '/signin'
            },
            {
              label: 'Sign Up',
              visible: !this.isLoggedIn,
              to: '/signup'
            },
            {
              label: 'Sign Out',
              visible: this.isLoggedIn,
              command: () => {
                this.signout()
              }
            }
          ]
        }
      ],
      items: [
        {
          label:'Contents',
          icon:'pi pi-fw pi-file',
        },
      ]
    }
  }
}
</script>

<style lang="scss">

</style>