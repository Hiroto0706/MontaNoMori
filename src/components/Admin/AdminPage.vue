<template>
  <v-container>
    <v-card flat class="pa-10 my-15"
      ><h1 class="page__title mb-5 pb-2">管理者画面</h1>

      <div>
        <router-link to="/admin/post" class="link-style">
          <v-btn
            block
            depressed
            x-large
            outlined
            color="primary"
            class="mb-5 py-8 font-weight-bold"
          >
            <v-icon class="px-2">mdi-image-plus-outline</v-icon>
            画像追加
          </v-btn>
        </router-link>
        <router-link to="/admin/images" class="link-style">
          <v-btn
            block
            depressed
            x-large
            outlined
            color="blue darken-3
"
            class="mb-5 py-8 font-weight-bold"
          >
            <v-icon class="px-2">mdi-image-multiple-outline</v-icon>
            画像一覧
          </v-btn>
        </router-link>
        <router-link to="/admin/tag" class="link-style">
          <v-btn
            block
            depressed
            x-large
            outlined
            color="primary"
            class="mb-5 py-8 font-weight-bold"
          >
            <v-icon class="px-2">mdi-tag-multiple-outline</v-icon>
            タグ一覧
          </v-btn>
        </router-link>
        <v-btn
          block
          depressed
          x-large
          outlined
          color="error"
          class="mb-5 py-8 font-weight-bold"
          @click="logout"
        >
          <v-icon class="px-1">mdi-logout</v-icon>ログアウト</v-btn
        >
      </div>
    </v-card>
  </v-container>
</template>

<script>
export default {
  data() {
    return {};
  },
  methods: {
    logout() {
      localStorage.removeItem("token");
      this.$router.push("login");
    },
  },
};
</script>

<style scoped>
.page__title {
  background: linear-gradient(transparent 92%, #000000 92%);
}

a {
  text-decoration: none;
}
</style>
