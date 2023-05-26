<template>
  <div class="content pa-10 my-10">
    <h1 class="content__title">お問い合わせ</h1>
    <div class="d-flex flex-column justify-center align-center my-5">
      <v-img
        src="@/assets/contact-me.png"
        lazy-src="@/assets/contact-me.png"
        max-height="350"
        max-width="350"
        aspect-ratio="1"
      />
    </div>
    <div class="py-5">
      もんたの森に関するお問い合わせはこちらからお願いします！
      <br />
      <br />
      <ul class="ma-5">
        <li>もんたの森に関する質問</li>
        <li>もんたの森の感想</li>
        <li>リクエスト</li>
        <li>お仕事の依頼（ないか〜）</li>
      </ul>

      <br />
      などなど！なんでもお問合せください！<br />
      <br />
    </div>
    <v-form @submit="submitForm">
      <v-text-field
        v-model="name"
        label="名前 ※必須だよ"
        outlined
        color="black"
      ></v-text-field>
      <v-text-field
        v-model="email"
        label="メールアドレス ※必須だよ"
        outlined
        color="black"
      ></v-text-field>
      <v-text-field
        v-model="title"
        label="題名 ※必須だよ"
        outlined
        color="black"
      ></v-text-field>
      <v-textarea
        v-model="message"
        label="本文 ※必須だよ"
        outlined
        color="black"
      ></v-textarea>
      <v-btn type="submit" color="primary" depressed x-large>送信する</v-btn>
    </v-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      name: "",
      email: "",
      title: "",
      message: "",
    };
  },
  methods: {
    submitForm() {
      // フォームの送信処理を行う
      // 例えば、APIリクエストを送信したり、メールを送信したりする処理をここに記述します
      console.log("Form submitted");
    },
  },
};
</script>

<style scoped>
.content {
  width: 100%;
  background-color: white;
  border-radius: 10px;
}

.content__title {
  display: block;
  width: 100%;
  background: linear-gradient(transparent 92%, #000000 92%);
}
</style>
