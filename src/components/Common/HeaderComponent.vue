<template>
  <v-app-bar app color="white" height="100" flat>
    <!-- タイトル -->
    <div
      class="align-center title_style mx-12"
      style="color: black"
      @click="getImages()"
    >
      <span class="text-h4 font-weight-bold">もんたの森</span>
      <v-spacer></v-spacer>
      <span class="text-subtitle-2 text-center">ゆる〜い無料イラストたち.</span>
    </div>

    <v-spacer></v-spacer>

    <!-- 検索欄 -->
    <v-text-field
      v-model.lazy="search"
      append-icon="mdi-magnify"
      outlined
      label="キーワード検索"
      hide-details
      color="black"
      class="cursor-pointer"
      @click:append="searchItems()"
      @keydown.enter="searchItems()"
    ></v-text-field>

    <v-spacer></v-spacer>

    <!-- タブ -->
    <v-btn
      class="mx-1 px-0 mx-md-2 px-md-3 font-weight-bold"
      depressed
      color="white"
      rounded
      x-large
      v-on:click.native="$router.push('/')"
    >
      <v-icon class="mr-2">mdi-image-multiple</v-icon>全てのイラスト
    </v-btn>
    <v-btn
      class="mx-1 px-0 mx-md-2 px-md-3 font-weight-bold"
      depressed
      color="white"
      rounded
      x-large
      @click.native="$router.push('/favorite')"
    >
      <v-icon class="mr-2">mdi-heart</v-icon>お気に入り
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
    async searchItems() {
      if (this.search == "") {
        return;
      }

      await this.$store.dispatch("searchItems", this.search).then(() => {
        this.$router.push({
          name: "search",
          query: { q: this.search },
        });
      });
    },
    async getImages() {
      this.search = "";
      await this.$store.dispatch("getImages").then(() => {
        this.$router.push({
          name: "home",
        });
      });
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
