<template>
  <div>
    <Toast position="bottom-right" :breakpoints="{'920px': {width: '100%', top: '0', right: '0'}}" />
    <MainHeader :isLoggedIn="isLoggedIn" />
    <div class="col-10 col-offset-1 mt-6 lg:col-4 lg:col-offset-4">
      <Panel class="col-12 lg:col-9" header="Sign in">
        <div class="mt-4">
          <i class="pi pi-user"></i>
          <InputText v-bind:class="{ 'p-invalid': error.email }" class="col-10 col-offset-1" type="text" v-model="email" @blur="validate('email', email)" placeholder="Email" />
        </div>
        <div class="mt-4">
          <i class="pi pi-lock"></i>
          <InputText v-bind:class="{ 'p-invalid': error.password }" class="col-10 col-offset-1" type="password" v-model="password" @blur="validate('password', password)" placeholder="Password" />
        </div>
        <div class="col-12 lg:pr-3">
          <Button class="mt-2 col-4 col-offset-8 lg:mt-4 lg:col-3 lg:col-offset-9" label="Sign In" @click="signin(email, password)" />
        </div>
      </Panel>
    </div>
  </div>
</template>

<script>
import { defineComponent, ref, reactive } from "vue"
import { useRoute, useRouter } from 'vue-router'
import app from "@/plugins/firebase"
import { getAuth, signInWithEmailAndPassword } from "firebase/auth";
import MainHeader from '@/components/organisms/MainHeader.vue'
import Panel from 'primevue/panel'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Toast from 'primevue/toast';
import { useToast } from "primevue/usetoast"

import Store from '@/store/index.js'

export default defineComponent({
  name: 'SigninPage',
  components: {
    MainHeader,
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
    
    const isLoggedIn = Store.getters.isLoggedIn

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

    const signin = (email, password) => {
      const auth = getAuth(app);
      signInWithEmailAndPassword(auth, email, password)
        .then(async (userCredential) => {
          const user = userCredential.user;
          Store.dispatch("auth", {
            userId: user.uid,
            userToken: user.accessToken
          });
          toast.add({severity:'success', summary: 'created', detail: 'success create user', life: 5000})
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
      isLoggedIn,
      validate,
      signin,
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