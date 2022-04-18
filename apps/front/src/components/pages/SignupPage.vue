<template>
  <div>
    <Toast position="bottom-right" :breakpoints="{'920px': {width: '100%', top: '0', right: '0'}}" />
    <MainHeader :isLoggedIn="isLoggedIn" />
    <div class="col-10 col-offset-1 mt-6 lg:col-4 lg:col-offset-4">
      <Panel class="col-12 lg:col-9" header="Sign Up">
        <div class="mt-4">
          <i class="pi pi-user"></i>
          <InputText v-bind:class="{ 'p-invalid': error.email }" class="col-10 col-offset-1" type="text" v-model="email" @blur="validate('email', email)" @keydown.enter="signup($event, email, password)" placeholder="Email" />
        </div>
        <div class="mt-4">
          <i class="pi pi-lock"></i>
          <InputText v-bind:class="{ 'p-invalid': error.password }" class="col-10 col-offset-1" type="password" v-model="password" @blur="validate('password', password)" @keydown.enter="signup($event, email, password)" placeholder="Password" />
        </div>
        <div class="col-12 lg:pr-3">
          <Button class="mt-2 col-4 col-offset-8 lg:mt-4 lg:col-3 lg:col-offset-9" label="Sign Up" @click="signup($event, email, password)" />
        </div>
      </Panel>
    </div>
  </div>
</template>


<script>
import { defineComponent, ref, reactive } from "vue"
import { useRouter } from 'vue-router'
import { UserServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { User, UserId } from "@/pb/models/user_pb.js"
import app from "@/plugins/firebase.js"
import { getAuth, createUserWithEmailAndPassword } from "firebase/auth"
import MainHeader from '@/components/organisms/MainHeader.vue'
import Panel from 'primevue/panel'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Toast from 'primevue/toast';
import { useToast } from "primevue/usetoast"
import Store from '@/store/index.js'

const client = new UserServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default defineComponent({
  name: 'SignupPage',
  components: {
    MainHeader,
    Panel,
    InputText,
    Button,
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
          router.push('/')
        })
        .catch((error) => {
          const errorCode = error.code;
          const errorMessage = error.message;
          toast.add({severity:'error', summary: 'failed to create user', detail: `${errorCode}: ${errorMessage}`, life: 5000});
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
      isLoggedIn,
      validate,
      signup,
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