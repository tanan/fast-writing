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

import app from "@/plugins/firebase.js"
import { getAuth, createUserWithEmailAndPassword } from "firebase/auth";

export default {
   name: 'SignupPage',
   data: () => ({
    email: "",
    password: "",
  }),
  methods: {
    signup(email, password) {
      console.log(email)
      const auth = getAuth(app);
      createUserWithEmailAndPassword(auth, email, password)
        .then((userCredential) => {
          const user = userCredential.user;
          this.$store.dispatch("auth", {
            userId: user.uid,
            userToken: user.accessToken
          })
          this.$router.push('/')
        })
        .catch((error) => {
          const errorCode = error.code;
          const errorMessage = error.message;
          console.log(errorCode)
          console.log(errorMessage)
        })
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