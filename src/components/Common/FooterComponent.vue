<template>
  <v-footer padless class="footer-content">
    <v-card flat class="pa-15">
      <!-- ハッシュタグ一覧 -->
      <v-card-text class="pa-10">
        <v-card-title class="text-h5 font-weight-bold pl-2 mb-3 black--text"
          >ハッシュタグたち</v-card-title
        >
        <div class="d-flex flex-wrap">
          <v-btn
            v-for="tag in tags.slice(0, 30)"
            :key="tag.id"
            class="mx-1 my-1 px-3 py-1"
            outlined
            rounded
            @click="getImagesByTag(tag.name)"
          >
            # {{ tag.name }}
          </v-btn>
        </div>
      </v-card-text>

      <v-spacer></v-spacer>

      <!-- このサイトについてのフッター -->
      <v-card-text class="px-10 black--text d-flex">
        <div class="mr-8 font-weight-bold a_tag">もんたのフリイラについて</div>
        <div class="mr-8 font-weight-bold a_tag">よくある質問</div>
        <div class="mr-8 font-weight-bold a_tag">お問い合わせ</div>

        <v-spacer></v-spacer>

        <div class="gray--text">&copy; 2023 もんたの森</div>
      </v-card-text>
    </v-card>
  </v-footer>
</template>

<script>
export default {
  computed: {
    tags() {
      return this.$store.state.tags;
    },
  },
  created() {
    this.$store.dispatch("fetchTags");
  },
  methods: {
    getImagesByTag(tagName) {
      console.log(tagName);
      this.$store.dispatch("getImagesByTag", tagName).then(() => {
        this.$router.push({
          path: "/tag/" + tagName,
        });
      });
    },
  },
};
</script>

<style scoped>
.footer-content {
  width: 100%;
  display: block;
}
</style>
