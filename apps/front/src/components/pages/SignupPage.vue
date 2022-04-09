<template>
  <v-container class="signup" fluid fill-height>
    <v-row justify="center">
      <v-col cols="10" sm="4" md="3">
        <v-card>
          <v-toolbar dark color="secondary">
            <v-toolbar-title>Signup</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-form>
              <div class="text-field">
                <mdicon class="icon" name="account" :width="30" :height="30" />
                <v-text-field name="email" v-model="email" label="Email" type="text"></v-text-field>
              </div>
              <div class="text-field">
                <mdicon class="icon" name="lock" :width="30" :height="30" />
                <v-text-field id="password" name="password" v-model="password" label="Password" type="password"></v-text-field>
              </div>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="secondary" @click="signup(this.email, this.password)">Login</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { UserServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { User, UserId } from "@/pb/models/user_pb.js"
import app from "@/plugins/firebase.js"
import { getAuth, createUserWithEmailAndPassword } from "firebase/auth"

const client = new UserServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default {
   name: 'SignupPage',
   data: () => ({
    email: "",
    password: "",
  }),
  methods: {
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