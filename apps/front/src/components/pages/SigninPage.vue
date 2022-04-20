<template>
  <div>
    <Toast position="bottom-right" :breakpoints="{'920px': {width: '100%', top: '0', right: '0'}}" />
    <div class="col-10 col-offset-1 mt-6 lg:col-4 lg:col-offset-4">
      <Panel class="col-12 lg:col-9" header="Sign in">
        <div class="mt-4">
          <i class="pi pi-user"></i>
          <InputText v-bind:class="{ 'p-invalid': error.email }" class="col-10 col-offset-1" type="text" v-model="email" @blur="validate('email', email)" @keydown.enter="signin($event, email, password)" placeholder="Email" />
        </div>
        <div class="mt-4">
          <i class="pi pi-lock"></i>
          <InputText v-bind:class="{ 'p-invalid': error.password }" class="col-10 col-offset-1" type="password" v-model="password" @blur="validate('password', password)" @keydown.enter="signin($event, email, password)" placeholder="Password" />
        </div>
        <div class="col-12 lg:pr-3">
          <Button class="mt-2 col-6 col-offset-3 lg:mt-4 lg:col-6 lg:col-offset-3" label="ログイン" @click="signin($event, email, password)" />
        </div>
        <hr class="col-offset-1 col-10 mt-4 p-0">
        <div class="mt-2">
          <p class="col-10 col-offset-1" style="text-align: center">他のアカウントでログイン</p>
          <Button class="col-12 lg:col-8 lg:col-offset-2" label="Google アカウントでログイン" @click="signinWithGoogle()" />
        </div>
      </Panel>
    </div>
  </div>
</template>

<script>
import { defineComponent, ref, reactive } from "vue"
import { useRoute, useRouter } from 'vue-router'
import app from "@/plugins/firebase"
import { getAuth, signInWithEmailAndPassword, signInWithPopup, GoogleAuthProvider } from "firebase/auth";
import Panel from 'primevue/panel'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Toast from 'primevue/toast';
import { useToast } from "primevue/usetoast"

import Store from '@/store/index.js'

export default defineComponent({
  name: 'SigninPage',
  components: {
    Panel,
    InputText,
    Button,
    Toast,
  },
  setup () {
    const email = ref('')
    const password = ref('')
    const error = reactive({
      email: false,
      password: false,
    })
    const toast = useToast()
    const route = useRoute()
    const router = useRouter()

    const auth = getAuth(app);
    const provider = new GoogleAuthProvider();
    
    const validate = (k, v) => {
      switch(k) {
        case 'email':
          error.email = v === "" ? true : false
          break
        case 'password':
          error.password = v === "" ? true : false
          break
      }
    }

    const hasInputError = () => {
      error.email = email.value === "" ? true : false
      error.password = password.value === "" ? true : false
      if (error.email || error.password) {
        return true
      }
      return false
    }

    const signinWithGoogle = () => {
      signInWithPopup(auth, provider)
        .then(async (result) => {
          const credential = GoogleAuthProvider.credentialFromResult(result)
          const user = result.user
          Store.dispatch("auth", {
            userId: user.uid,
            userToken: credential.accessToken
          })
          toast.add({severity:'success', summary: 'success', detail: 'login succeeded', life: 5000})
          await new Promise((resolve) => setTimeout(resolve, 1000))
          if (router.query) {
            router.push(route.query.redirect)
          }
          router.push('/')
        }).catch((error) => {
          toast.add({severity:'error', summary: 'failed to login', detail: `${error.code}: ${error.message}`, life: 5000});
        })
    }
    const signin = (event, email, password) => {
      if (event.key == "Enter" && event.keyCode !== 13) return
      if (hasInputError()) {
        return
      }
      signInWithEmailAndPassword(auth, email, password)
        .then((userCredential) => {
          const user = userCredential.user;
          Store.dispatch("auth", {
            userId: user.uid,
            userToken: user.accessToken
          });
          toast.add({severity:'success', summary: 'success', detail: 'login succeeded', life: 5000})
        })
        .then(async () => {
          await new Promise((resolve) => setTimeout(resolve, 1000))
          if (router.query) {
            router.push(route.query.redirect)
          }
          router.push('/')
        })
        .catch((error) => {
          const errorCode = error.code;
          const errorMessage = error.message;
          toast.add({severity:'error', summary: 'failed to login', detail: `${errorCode}: ${errorMessage}`, life: 5000});
        })
     }

    return {
      email,
      password,
      error,
      toast,
      validate,
      signin,
      signinWithGoogle,
    }
  },
  methods: {
   }
})
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