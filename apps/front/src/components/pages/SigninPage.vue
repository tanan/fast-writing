<template>
  <v-container class="signin" fluid fill-height>
    <v-row justify="center">
      <v-col cols="10" sm="4" md="3">
        <v-card>
          <v-toolbar dark color="primary">
            <v-toolbar-title>Login form</v-toolbar-title>
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
            <v-btn color="primary" @click="signin(this.email, this.password)">Login</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import app from "@/plugins/firebase"
import { getAuth, signInWithEmailAndPassword } from "firebase/auth";

export default {
   name: 'SigninPage',
   data: () => ({
    email: "",
    password: "",
  }),
  methods: {
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