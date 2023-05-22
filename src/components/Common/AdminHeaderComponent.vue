<template>
  <v-app-bar app color="white" height="100" flat>
    <!-- タイトル -->
    <router-link
      class="align-center title_style ml-12 mr-5"
      to="/admin"
      style="color: black"
    >
      <span class="text-h4 font-weight-bold">もんたの森</span>
      <v-spacer></v-spacer>
      <span class="text-subtitle-2 text-center">ゆる〜い無料イラストたち.</span>
    </router-link>
    <span class="text-subtitle1 grey--text text-right">for 管理者</span>

    <v-spacer></v-spacer>

    <!-- 検索欄 -->
    <v-text-field
      v-model="search"
      append-icon="mdi-magnify"
      outlined
      label="キーワード検索"
      hide-details
      color="black"
      class="cursor-pointer"
      @click:append="searchItems"
    ></v-text-field>

    <v-spacer></v-spacer>

    <!-- タブ -->
    <v-btn
      class="mx-1 px-0 mx-md-2 px-md-3 font-weight-bold"
      depressed
      color="white"
      rounded
      x-large
      v-on:click.native="$router.push('/admin/post')"
    >
      <v-icon class="mr-2">mdi-image-plus</v-icon>画像追加
    </v-btn>
    <v-btn
      class="mx-1 px-0 mx-md-2 px-md-3 font-weight-bold"
      depressed
      color="white"
      rounded
      x-large
      v-on:click.native="$router.push('/admin/images')"
    >
      <v-icon class="mr-2">mdi-image-multiple</v-icon>画像一覧
    </v-btn>
    <v-btn
      class="mx-1 px-0 mx-md-2 px-md-3 font-weight-bold"
      depressed
      color="white"
      rounded
      x-large
      v-on:click.native="$router.push('/admin/tag')"
    >
      <v-icon class="mr-2">mdi-tag-multiple</v-icon>タグ一覧
    </v-btn>
  </v-app-bar>
</template>

<script>
export default {
  data() {
    return {
      search: "",
    };
  },
  methods: {
    searchItems() {
      console.log("search!!");
    },
  },
};
</script>

<style scoped>
.a_tag:hover {
  cursor: pointer;
}

.title_style,
.router-link-style {
  color: inherit;
  text-decoration: none;
  cursor: pointer;
}
</style>
