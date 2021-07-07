<template>
  <v-dialog v-model="dialog" persistent max-width="600px" min-width="360px">
    <div>
      <v-tabs show-arrows background-color="deep-purple accent-4" icons-and-text dark grow>
        <v-tabs-slider color="purple darken-4"></v-tabs-slider>
        <v-tab>
          <v-icon large>mdi-account</v-icon>
          <div class="caption py-1">LOGIN</div>
        </v-tab>
        <v-tab-item>
          <v-card class="px-4">
            <v-card-text>
              <v-form ref="loginForm" v-model="valid" lazy-validation>
                <v-row>
                  <v-col cols="12">
                    <v-text-field
                      v-model="username"
                      :rules="usernameRules"
                      label="User Name"
                      required
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12">
                    <v-text-field
                      v-model="password"
                      :append-icon="show1 ? 'eye' : 'eye-off'"
                      :rules="[passWordRules.required, passWordRules.min]"
                      :type="show1 ? 'text' : 'password'"
                      name="input-10-1"
                      label="Password"
                      hint="At least 8 characters"
                      counter
                      @click:append="show1 = !show1"
                    ></v-text-field>
                  </v-col>
                  <v-col class="d-flex" cols="12" sm="6" xsm="12"></v-col>
                  <v-spacer></v-spacer>
                  <v-col class="d-flex" cols="12" sm="3" xsm="12" align-end>
                    <v-btn x-large block :disabled="!valid" color="success" @click="validate">Login</v-btn>
                  </v-col>
                </v-row>
              </v-form>
            </v-card-text>
          </v-card>
        </v-tab-item>
      </v-tabs>
    </div>
  </v-dialog>
</template>

<script>
import axios from "axios";
import md5 from "js-md5";

export default {
  name: "Home",

  components: {},
  computed: {
    passwordMatch() {
      return () => this.password === this.verify || "Password must match";
    },
  },
  methods: {
    validate() {
      if (this.$refs.loginForm.validate()) {
        // submit form to server/API here...
        const params = new URLSearchParams();
        params.append("username", this.username);
        params.append("password", md5(this.password));

        axios
          .post("http://localhost:8081/api/v1/login", params, {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
            },
            withCredentials: true
          })
          .then((res) => {
            // do something with res
            console.log(res);
            if (res.data.Code != 0) {
              alert("Login Failed: " + res.data.Error)
            } else {
              this.$router.push({
                path: "/profile",
                query: { username: this.username }
              });
            }
          })
          .catch((err) => {
            // catch error
            alert("Login Failed: " + err)
            console.log(err);
          });
      }
    },
    reset() {
      this.$refs.form.reset();
    },
    resetValidation() {
      this.$refs.form.resetValidation();
    },
  },
  data: () => ({
    dialog: true,
    valid: true,

    username: "",
    usernameRules: [
      (v) => !!v || "Required",
      (v) => /^[0-9]*$/.test(v) || "User Name must be digits",
    ],
    show1: false,

    password: "",
    passWordRules: {
      required: (value) => !!value || "Required.",
      min: (v) => (v && v.length >= 8) || "Min 8 characters",
    },
  }),
};
</script>
