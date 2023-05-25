<template>
  <v-footer padless class="footer-content">
    <v-card flat class="pa-15">
      <v-spacer></v-spacer>

      <!-- このサイトについてのフッター -->
      <v-card-text class="px-10 black--text d-flex">
        <router-link
          to="/about"
          class="mr-8 font-weight-bold a_tag black--text"
          style="text-decoration: none"
          >もんたのフリイラについて</router-link
        >
        <router-link
          to="/about"
          class="mr-8 font-weight-bold a_tag black--text"
          style="text-decoration: none"
          >お問い合わせ</router-link
        >

        <v-spacer></v-spacer>

        <div class="gray--text">&copy; 2023 もんたのフリイラ</div>
      </v-card-text>
    </v-card>
  </v-footer>
</template>

<script>
export default {};
</script>

<style scoped>
.footer-content {
  width: 100%;
  display: block;
}
</style>
