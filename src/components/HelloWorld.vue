<template>
  <v-container>
    <h1 class="my-10 pl-3 pb-3 text-h4 font-weight-bold page__title">
      {{ $store.state.title }}
    </h1>
    <v-row v-if="$store.state.images === null"
      ><v-col class="mt-5 mb-8 font-weight-bold"
        >画像が見つかりませんでした</v-col
      ></v-row
    >
    <v-row v-else>
      <v-col cols="3" v-for="image in $store.state.images" :key="image.id">
        <router-link
          :to="'/' + image.id"
          class="router-link-style"
          style="color: black"
        >
          <v-card class="pa-10" flat>
            <img :src="image.url_path" />
          </v-card>
        </router-link>
        <div class="my-5 d-flex">
          <p class="font-weight-bold my-auto">
            <router-link
              :to="'/' + image.id"
              class="router-link-style"
              style="color: black"
              >{{ image.title }}</router-link
            >
          </p>
          <v-btn icon class="ml-auto" @click="favorite(image.id)">
            <v-icon :color="isFavorite(image.id) ? 'red lighten-1' : ''"
              >mdi-heart</v-icon
            >
          </v-btn>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "HelloWorld",
  data() {
    return {
      storedIDs: [],
    };
  },
  mounted() {
    this.storedIDs = JSON.parse(localStorage.getItem("favoriteIDs")) || [];
  },
  computed: {
    tags() {
      return this.$store.state.tags;
    },
  },
  methods: {
    favorite(id) {
      if (!this.storedIDs.includes(id)) {
        this.storedIDs.push(id);
      } else {
        this.storedIDs = this.storedIDs.filter((item) => item !== id);
      }
      localStorage.setItem("favoriteIDs", JSON.stringify(this.storedIDs));
    },
    isFavorite(id) {
      return this.storedIDs.includes(id);
    },
  },
};
</script>

<style scoped>
.page__title {
  background: linear-gradient(transparent 92%, #000000 92%);
}

img {
  width: 100%;
  height: auto;
  transition: 0.3s ease;
}

img:hover {
  transform: scale(1.1, 1.1);
}
</style>
