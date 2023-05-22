<template>
  <v-container>
    <v-card flat class="pa-10">
      <v-card-title class="px-0 mb-5 text-h5 font-weight-bold"
        >管理者ログインフォーム</v-card-title
      >
      <v-text-field
        v-model="email"
        color="black"
        label="メールアドレス"
        outlined
        ref="email"
        type="email"
        class="mb-5"
        @keydown.enter="$refs.password.focus()"
      ></v-text-field>
      <v-text-field
        v-model="password"
        :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
        :type="show1 ? 'text' : 'password'"
        label="パスワード"
        outlined
        color="black"
        ref="password"
        @click:append="show1 = !show1"
        @keydown.enter="login"
      ></v-text-field>
      <p class="mb-6 red--text font-weight-bold">{{ errorMessage }}</p>
      <v-btn
        block
        x-large
        depressed
        color="primary"
        dark
        class="font-weight-bold"
        @click="login"
        >ログインする</v-btn
      >
    </v-card>
  </v-container>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      email: "",
      password: "",
      show1: false,
      errorMessage: "",
    };
  },
  mounted() {
    this.$refs.email.focus();
  },
  methods: {
    login() {
      const data = new FormData();
      data.append("email", this.email);
      data.append("password", this.password);

      axios
        .post("http://localhost:8000/login", data)
        .then((response) => {
          localStorage.setItem("token", response.data.uuid);
          this.$router.push("/admin");
        })
        .catch((error) => {
          this.errorMessage = "入力内容に誤りがあります。";
          console.log(error);
        });
    },
  },
};
</script>
