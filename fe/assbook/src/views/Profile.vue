<template>
  <v-dialog v-model="dialog" persistent max-width="600px" min-width="360px">
    <div>
      <v-card class="mx-auto" outlined>
        <v-list-item three-line>
          <v-list-item-content style="'align-items': 'center'">
            <v-list-item-title class="d-flex justify-center mb-6">
              PROFILE
            </v-list-item-title>

            <v-list-item-subtitle class="d-flex justify-center mb-6">
              ID: {{ this.username }}
            </v-list-item-subtitle>

            <v-list-item-subtitle class="d-flex justify-center mb-6">
              <v-avatar class="ma-3" size="125" @click="pickImage">
                <v-img :src="this.avatar"></v-img>
              </v-avatar>
              <input
                type="file"
                style="display: none"
                ref="image"
                accept="image/*"
                @change="onImagePicked"
              />
            </v-list-item-subtitle>

            <v-list-item-subtitle class="d-flex justify-center mb-6">
              <label for="nickname">Nick Name:</label>
              <input
                style="margin-left:10px"
                type="text"
                ref="nickname"
                :value="this.nickname"
                @change="onNicknameChanged"
              />
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>

        <v-card-actions class="d-flex justify-center mb-6">
          <v-btn
            :disabled="!this.modified"
            class="ml-2 mt-5"
            outlined
            rounded
            small
            @click="updateProfile"
          >
            Update
          </v-btn>
        </v-card-actions>
      </v-card>
    </div>
  </v-dialog>
</template>

<script>
import axios from "axios";

export default {
  name: "Profile",

  mounted: function() {
    this.username = this.$route.query.username;
    axios
      .get(`http://localhost:8081/api/v1/profile/${this.username}`, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        withCredentials: true,
      })
      .then((res) => {
        // do something with res
        console.log(res);
        if (res.data.Code != 0) {
          // show error
          alert("Faile to fetch the user profile: " + res.data.Error);
        } else {
          // display the user profile
          this.nickname = res.data.Data.nickname;
          this.avatar = res.data.Data.avatar;
        }
      })
      .catch((err) => {
        // catch error
        alert("Failed to fetch the user profile: " + err);
        console.log(err);
      });
  },

  components: {},

  computed: {},

  methods: {
    validate() {},
    pickImage() {
      this.$refs.image.click();
    },
    onImagePicked(e) {
      const files = e.target.files;
      if (files[0] !== undefined) {
        this.imageName = files[0].name;
        if (this.imageName.lastIndexOf(".") <= 0) {
          return;
        }
        const fr = new FileReader();
        fr.readAsDataURL(files[0]);
        fr.addEventListener("load", () => {
          this.avatar = fr.result;
          this.modified = true;
        });
      } else {
        this.imageName = "";
      }
    },
    onNicknameChanged() {
      this.nickname = this.$refs.nickname.value;
      this.modified = true;
    },
    updateProfile() {
      axios
        .put(
          `http://localhost:8081/api/v1/profile/${this.username}`,
          {
            nickname: this.nickname,
            avatar: this.avatar,
          },
          {
            headers: {},
            withCredentials: true,
          }
        )
        .then((res) => {
          // do something with res
          console.log(res);
          if (res.data.Code != 0) {
            // show error
            alert("Faile to update the user profile: " + res.data.Error);
          } else {
            // reload the page to get the latest profile
            this.$router.go();
          }
        })
        .catch((err) => {
          // catch error
          alert("Failed to fetch the user profile: " + err);
          console.log(err);
        });
    },
  },
  data: () => ({
    dialog: true,
    username: "",
    nickname: "",
    avatar: "",
    modified: false,
  }),
};
</script>
