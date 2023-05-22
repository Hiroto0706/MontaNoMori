<template>
  <v-app style="background-color: #f5f5f7">
    <!-- ヘッダー -->
    <router-view name="header"></router-view>

    <v-main>
      <v-container>
        <router-view></router-view>
      </v-container>
    </v-main>

    <!-- フッター -->
    <router-view name="footer"></router-view>
  </v-app>
</template>

<script>
export default {
  data() {
    return {};
  },
};
</script>

<style>
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
