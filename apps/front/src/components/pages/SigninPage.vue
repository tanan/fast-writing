<template>
  <div>
    <MainHeader :isLoggedIn="this.isLoggedIn()" />
    <div class="col-10 col-offset-1 mt-6 lg:col-4 lg:col-offset-4">
      <Panel class="col-12 lg:col-9" header="Sign In">
        <div class="mt-4">
          <i class="pi pi-user"></i>
          <InputText class="col-10 col-offset-1" type="text" v-model="title" @blur="save" placeholder="Email" />
        </div>
        <div class="mt-4">
          <i class="pi pi-lock"></i>
          <InputText class="col-10 col-offset-1" type="password" v-model="title" @blur="save" placeholder="Password" />
        </div>
        <div class="col-12 lg:pr-3">
          <Button class="mt-2 col-4 col-offset-8 lg:mt-4 lg:col-3 lg:col-offset-9" label="Sign In" @click="signin(this.email, this.password)" />
        </div>
      </Panel>
    </div>
  </div>
</template>

<script>
import app from "@/plugins/firebase"
import { getAuth, signInWithEmailAndPassword } from "firebase/auth";
import MainHeader from '@/components/organisms/MainHeader.vue'
import Panel from 'primevue/panel'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Store from '@/store/index.js'

export default {
  name: 'SigninPage',
  data: () => ({
    email: "",
    password: "",
  }),
  components: {
    MainHeader,
    Panel,
    InputText,
    Button,
  },
  methods: {
    isLoggedIn () {
      if (Store.state.userToken === "") {
        return false
      }
      return true
    },
    signin(email, password) {
      const auth = getAuth(app);
      signInWithEmailAndPassword(auth, email, password)
        .then((userCredential) => {
          const user = userCredential.user;
          this.$store.dispatch("auth", {
            userId: user.uid,
            userToken: user.accessToken
          });
          if (this.$router.query) {
            this.$router.push(this.$route.query.redirect)
          }
          this.$router.push('/')
        })
        .catch((error) => {
          const errorCode = error.code;
          const errorMessage = error.message;
          console.log(errorCode)
          console.log(errorMessage)
        });
     }
   }
};
</script>

<style lang="scss">
  .signin { 
    margin-top: 10%;
    
    .text-field {
      display: flex;
      .icon {
        padding-right: 6px;
        padding-top: 12px;
        &:before {
          content: "";
        }
      }
    }
  }
</style>