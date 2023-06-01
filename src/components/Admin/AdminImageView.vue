<template>
  <v-container>
    <h1 class="my-10 pl-3 pb-3 text-h4 font-weight-bold page__title">
      {{ $store.state.title }}
    </h1>
    <v-row>
      <v-col cols="3" v-for="image in $store.state.images" :key="image.id">
        <router-link
          :to="'/admin/' + image.id"
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
              >{{ image.id }}. {{ image.title }}</router-link
            >
          </p>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
// import axios from "axios";

export default {
  name: "HelloWorld",
  data() {
    return {
      images: [],
    };
  },
  // mounted() {
  //   axios
  //     .get("http://localhost:8000/")
  //     .then((response) => {
  //       this.images = response.data;
  //     })
  //     .catch((error) => {
  //       console.log(error);
  //     });
  // },
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
