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
import { initializeApp } from "firebase/app";
import { getAuth, signInWithEmailAndPassword } from "firebase/auth";

const firebaseConfig = {
  apiKey: "AIzaSyCgSiRqYgLaqNpg_Agn1u5kocxIAVB4OU8",
  authDomain: "anan-project.firebaseapp.com",
  databaseURL: "https://anan-project.firebaseio.com",
  projectId: "anan-project",
  storageBucket: "anan-project.appspot.com",
  messagingSenderId: "474794270878",
  appId: "1:474794270878:web:0a2bcc65d7f97da62e5b8f"
};

const app = initializeApp(firebaseConfig);

export default {
   name: 'SigninPage',
   data: () => ({
    email: "",
    password: "",
  }),
  methods: {
    signin(email, password) {
      console.log(email)
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