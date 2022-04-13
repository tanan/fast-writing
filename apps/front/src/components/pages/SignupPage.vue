<template>
  <div>
    <MainHeader :isLoggedIn="this.isLoggedIn()" />
    <div class="col-10 col-offset-1 mt-6 lg:col-4 lg:col-offset-4">
      <Panel class="col-12 lg:col-9" header="Sign Up">
        <div class="mt-4">
          <i class="pi pi-user"></i>
          <InputText class="col-10 col-offset-1" type="text" v-model="title" @blur="save" placeholder="Email" />
        </div>
        <div class="mt-4">
          <i class="pi pi-lock"></i>
          <InputText class="col-10 col-offset-1" type="password" v-model="title" @blur="save" placeholder="Password" />
        </div>
        <div class="col-12 lg:pr-3">
          <Button class="mt-2 col-4 col-offset-8 lg:mt-4 lg:col-3 lg:col-offset-9" label="Sign Up" @click="signup(this.email, this.password)" />
        </div>
      </Panel>
    </div>
  </div>
</template>


<script>
import { UserServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { User, UserId } from "@/pb/models/user_pb.js"
import app from "@/plugins/firebase.js"
import { getAuth, createUserWithEmailAndPassword } from "firebase/auth"
import MainHeader from '@/components/organisms/MainHeader.vue'
import Panel from 'primevue/panel'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Store from '@/store/index.js'

const client = new UserServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default {
   name: 'SignupPage',
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
    async signup(email, password) {
      const auth = getAuth(app);
      createUserWithEmailAndPassword(auth, email, password)
        .then((userCredential) => {
          const user = userCredential.user;
          this.$store.dispatch("auth", {
            userId: user.uid,
            userToken: user.accessToken
          })
          return user
        })
        .then(async (user) => {
          let req = this.createUser(user.uid, "", this.email)
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
        .then(() => {
          this.$router.push('/')
        })
        .catch((error) => {
          const errorCode = error.code;
          const errorMessage = error.message;
          console.log(errorCode)
          console.log(errorMessage)
        })
    }, createUser (id, name, email) {
      let req = new User()
      let userId = new UserId()
      userId.setId(id)
      req.setId(userId)
      req.setName(name)
      req.setEmail(email)
      return req
    }
  }
};
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