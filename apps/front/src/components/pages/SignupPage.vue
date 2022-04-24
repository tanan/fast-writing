<template>
  <div style="background: url('images/signup/signup-bg.jpg') no-repeat; background-size: cover" class="px-4 py-8 md:px-6 lg:px-8">
    <Toast position="bottom-right" :breakpoints="{'920px': {width: '100%', top: '0', right: '0'}}" />
    <div class="flex flex-wrap">
      <div class="w-full lg:w-6 p-4 lg:p-7" style="background-color: rgba(255,255,255,.7)">
        <img src="images/logos/bastion-purple.svg" alt="Image" height="50" class="mb-6">
        <div class="text-xl text-900 font-500 mb-3">Welcome to SokuDoku</div>
        <p class="text-600 line-height-3 mt-0 mb-6">クイズに答えて英語力をアップしましょう！！</p>
        <ul class="list-none p-0 m-0">
          <li class="flex align-items-start mb-4">
            <div>
              <span class="flex align-items-center justify-content-center bg-purple-400" style="width:38px;height:38px;border-radius:10px">
                <i class="text-xl text-white pi pi-inbox"></i>
              </span>
            </div>
            <div class="ml-3">
              <span class="font-medium text-900">単語力UP</span>
              <p class="mt-2 mb-0 text-600 line-height-3">英作文をすることで、使える英単語を覚えます</p>
            </div>
          </li>
          <li class="flex align-items-start mb-4">
            <div>
              <span class="flex align-items-center justify-content-center bg-purple-400" style="width:38px;height:38px;border-radius:10px">
                <i class="text-xl text-white pi pi-shield"></i>
              </span>
            </div>
            <div class="ml-3">
              <span class="font-medium text-900">会話スピードUP</span>
              <p class="mt-2 mb-0 text-600 line-height-3">トレーニングを積むことで話すスピードが向上します</p>
            </div>
          </li>
          <li class="flex align-items-start">
            <div>
              <span class="flex align-items-center justify-content-center bg-purple-400" style="width:38px;height:38px;border-radius:10px">
                <i class="text-xl text-white pi pi-globe"></i>
              </span>
            </div>
          <div class="ml-3">
            <span class="font-medium text-900">自信UP</span>
            <p class="mt-2 mb-0 text-600 line-height-3">会話ができるようになると自分に自信が持てます</p>
          </div>
        </li>
      </ul>
    </div>
      <div class="w-full lg:w-6 p-4 lg:p-7 surface-card">
        <div class="text-900 text-2xl font-medium mb-6">ユーザー登録</div>
        <label for="email3" class="block text-900 font-medium mb-2">Email</label>
        <InputText id="email3" type="text" class="w-full mb-4" v-bind:class="{ 'p-invalid': error.email }" v-model="email" @blur="validate('email', email)" @keydown.enter="signup($event, email, password)" />

        <label for="password3" class="block text-900 font-medium mb-2">Password</label>
        <InputText id="password3" type="password" class="w-full mb-4" v-bind:class="{ 'p-invalid': error.password }" v-model="password" @blur="validate('password', password)" @keydown.enter="signup($event, email, password)" />
    
        <div class="flex align-items-center justify-content-between mb-6">
            <!-- <div class="flex align-items-center">
                <Checkbox id="rememberme3" :binary="true" v-model="checked3" class="mr-2"></Checkbox>
                <label for="rememberme3">Remember me</label>
            </div> -->
          <a href="/signin" class="font-medium no-underline ml-2 text-blue-500 text-right cursor-pointer">アカウントをお持ちの方</a>
        </div>
    
        <Button label="Sign Up" icon="pi pi-user" class="w-full p-button-help" @click="signup($event, email, password)"></Button>

        <Divider align="center" class="my-6">
          <span class="text-600 font-normal text-sm">OR</span>
        </Divider>

        <Button label="Sign Up with Google" icon="pi pi-google" class="w-full p-button-danger"></Button>

        <!-- <div class="mt-6 text-center text-600">
            Don't have an account? <a tabindex="0" class="font-medium text-blue-500">Sign up</a>
        </div> -->
      </div>
    </div>
  </div>
</template>


<script>
import { defineComponent, ref, reactive } from "vue"
import { useRoute, useRouter } from 'vue-router'
import { UserServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { User, UserId } from "@/pb/models/user_pb.js"
import app from "@/plugins/firebase.js"
import { getAuth, createUserWithEmailAndPassword, signInWithPopup, GoogleAuthProvider } from "firebase/auth"
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Divider from 'primevue/divider'
import Toast from 'primevue/toast'
import { useToast } from "primevue/usetoast"
import Store from '@/store/index.js'

const client = new UserServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default defineComponent({
  name: 'SignupPage',
  components: {
    InputText,
    Button,
    Divider,
    Toast
  },
  setup() {
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

    const signup = (event, email, password) => {
      if (event.key == "Enter" && event.keyCode !== 13) return
      if (hasInputError()) {
        return
      }
      const auth = getAuth(app);
      createUserWithEmailAndPassword(auth, email, password)
        .then((userCredential) => {
          const user = userCredential.user;
          Store.dispatch("auth", {
            userId: user.uid,
            userToken: user.accessToken
          })
          return user
        })
        .then(async (user) => {
          let req = createUser(user.uid, "", email)
          const token = user.accessToken
          const metadata = { 'authorization': 'Bearer ' + token }
          await new Promise((resolve, reject) => {
            client.createUser(req, metadata, (err, resp) => {
              if (err) {
                reject(err)
              }
              resolve(resp)
            })
          })
        })
        .then(async () => {
          toast.add({severity:'success', summary: 'created', detail: 'success create user', life: 5000})
          await new Promise((resolve) => setTimeout(resolve, 1000))
          router.push('/profile')
        })
        .catch((error) => {
          const errorCode = error.code;
          const errorMessage = error.message;
          toast.add({severity:'error', summary: 'failed to create user', detail: `${errorCode}: ${errorMessage}`, life: 5000});
        })
    }

    const signupWithGoogle = () => {
      signInWithPopup(auth, provider)
        .then((result) => {
          // const credential = GoogleAuthProvider.credentialFromResult(result)
          return result.user
        })
        .then((user) => {
          console.log(user)
          let req = createUser(user.uid, "", user.email)
          const metadata = { 'authorization': 'Bearer ' + user.accessToken }
          client.createUser(req, metadata, (err) => {
            if (err) {
              throw new Error(err)
            }
          })
          return user
        })
        .then((user) => {
          Store.dispatch("auth", {
            userId: user.uid,
            userToken: user.accessToken
          })
        })
        .then(async () => {
          toast.add({severity:'success', summary: 'success', detail: 'login succeeded', life: 5000})
          await new Promise((resolve) => setTimeout(resolve, 1000))
          if (route.query.redirect) {
            router.push(route.query.redirect)
            return
          }
          router.push('/')
        })
        .catch((error) => {
          toast.add({severity:'error', summary: 'failed to login', detail: `${error.code}: ${error.message}`, life: 5000});
        })
    }

    const createUser = (id, name, email) => {
      let req = new User()
      let userId = new UserId()
      userId.setId(id)
      req.setId(userId)
      req.setName(name)
      req.setEmail(email)
      return req
    }

    return {
      email,
      password,
      error,
      toast,
      validate,
      signup,
      signupWithGoogle,
      createUser,
    }
  }
})
</script>

<style lang="scss">
  .signup { 
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