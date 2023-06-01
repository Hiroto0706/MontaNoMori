<template>
  <div>
    <!-- <h1 class="my-5">For Admin Page</h1> -->
    <v-card flat class="pa-10 my-15">
      <h2 class="mb-5">画像の追加</h2>
      <v-card-text class="pa-0">
        <v-text-field
          color="black"
          label="タイトル"
          outlined
          maxlength="20"
          counter
          ref="title"
          v-model="title"
          :rules="[rules.required, rules.counter_for_title]"
          class="mb-5"
          @keydown.enter="$refs.description.focus()"
        ></v-text-field>
        <v-textarea
          color="black"
          outlined
          label="説明"
          value=""
          maxlength="100"
          counter
          ref="description"
          v-model="description"
          :rules="[rules.required, rules.counter_for_description]"
          class="mb-5"
          @keydown.enter="$refs.image.focus()"
        ></v-textarea>
        <v-file-input
          label="画像"
          outlined
          color="black"
          show-size
          counter
          ref="image"
          v-model="image"
          :rules="[rules.required]"
          class="mb-5"
          @keydown.enter="$refs.tag.focus()"
        ></v-file-input>
        <v-select
          color="black"
          label="タグ"
          chips
          multiple
          v-model="imageTags"
          :items="tags"
          item-text="name"
          item-value="id"
          class="mb-10"
          outlined
          ref="tag"
          persistent-hint
          :rules="[rules.required]"
          @keydown.enter="PostImage(title, description, image, imageTags)"
          hint="タグを選択してください"
        ></v-select>
        <v-btn
          x-large
          color="primary"
          block
          class="font-weight-bold"
          @click="PostImage(title, description, image, imageTags)"
        >
          画像をアップロードする
        </v-btn>
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      title: "",
      description: "",
      image: [],
      tags: [],
      imageTags: [],
      rules: {
        required: (value) => !!value || "必須項目だよ。",
        counter_for_title: (value) =>
          value.length <= 20 || "最大20文字までだよ。",
        counter_for_description: (value) =>
          value.length <= 100 || "最大100文字までだよ。",
      },
    };
  },
  mounted() {
    this.getTags();
    this.$refs.title.focus();
  },
  methods: {
    PostImage(title, description, image, tags) {
      // サーバーに渡すデータ;
      const data = new FormData();
      data.append("title", title);
      data.append("description", description);
      data.append("file", image, "image.png");
      tags.sort((a, b) => a - b);
      tags.forEach((tag) => {
        data.append("tags[]", tag);
      });

      // POSTリクエストを送信
      axios
        .post("http://localhost:8000/admin/upload", data)
        .then(() => {
          this.$router.push("/admin/images");
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getTags() {
      axios
        .get("http://localhost:8000/admin/tag/get")
        .then((response) => {
          this.tags = response.data;
          console.log(this.tags);
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
};
</script>
