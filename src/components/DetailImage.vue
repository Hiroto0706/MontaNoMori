<template>
  <v-container>
    <!-- 画像のタイトル -->
    <h1 v-if="image.title" class="image__title mb-5 pb-2 d-flex">
      <span>{{ image.title }}</span>
      <v-spacer></v-spacer>
      <v-btn icon x-large>
        <v-icon
          @click="favorite(image.id)"
          :color="isFavorite(image.id) ? 'red lighten-1' : ''"
          >mdi-heart</v-icon
        >
      </v-btn>
    </h1>
    <h1 v-else class="image__title mb-5 pb-2" style="color: gray">
      画像タイトル
    </h1>

    <div class="content mb-5">
      <!-- メインコンテンツ -->
      <div class="content__main pa-15 mr-5">
        <img :src="image.url_path" />
      </div>

      <!-- サブコンテンツ -->
      <div class="content__sub">
        <!-- ハッシュタグ一覧 -->
        <v-card-text class="pa-0 mb-10">
          <v-card-title class="text-h5 font-weight-bold pa-0 mb-3 black--text"
            >タグたち</v-card-title
          >
          <v-card class="d-flex flex-wrap pa-2" outlined>
            <v-btn
              v-for="imageTag in imageTags"
              :key="imageTag.id"
              class="mx-1 my-1 px-3 py-1"
              outlined
              rounded
              @click="getImagesByTag(imageTag.name)"
            >
              # {{ imageTag.name }}
            </v-btn>
          </v-card>
        </v-card-text>

        <!-- ひとこと -->
        <v-card-title class="text-h5 font-weight-bold pa-0 mb-3 black--text"
          >ひとこと</v-card-title
        >
        <v-card class="px-10 pt-10 pb-15 mb-10 hitokoto" outlined>
          <p class="font-weight-bold text-h7 pb-15 mb-15">
            {{ image.description }}
          </p>
          <img src="@/assets/description.png" />
        </v-card>

        <!-- 画像のダウンロード -->
        <div class="content__sub__btn-content d-flex">
          <v-btn
            class="content__sub__btn-content__btn content__sub__btn-content__btn-png py-7"
            depressed
            x-large
            block
            color="green"
            dark
            @click="downloadImage(image)"
            id="download"
          >
            <v-icon left large class="px-5">mdi-tray-arrow-down</v-icon
            >画像をダウンロード</v-btn
          >
        </div>
      </div>
    </div>

    <!-- 前の画像と後ろの画像 -->
    <div class="d-flex prev-next__content mb-15 p b-10">
      <!-- 前の画像 -->
      <v-card
        class="prev-none-link pa-5 ml-0 mr-5"
        flat
        v-if="prev_image.id === 0"
      >
        <div
          style="
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100%;
          "
        >
          <span class="font-weight-bold">これより前はないよ!</span>
        </div>
      </v-card>
      <router-link
        :to="'/' + parseInt(prev_image.id)"
        class="prev-next__content__link ml-0 mr-5"
        v-else
      >
        <v-card class="prev pa-5" flat>
          <span class="font-weight-bold">一つ前の画像も見てみる</span>
          <div class="prev-next__content__image">
            <img class="pa-5" :src="prev_image.url_path" />
            <p class="ma-0 font-weight-bold text-h7">
              {{ prev_image.title }}
            </p>
          </div>
        </v-card>
      </router-link>

      <!-- 次の画像 -->
      <v-card
        class="prev-none-link pa-5 mr-0 ml-5"
        flat
        v-if="next_image.id === 0"
      >
        <div
          style="
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100%;
          "
        >
          <span class="font-weight-bold">これより次はないよ!</span>
        </div>
      </v-card>
      <router-link
        :to="'/' + parseInt(next_image.id)"
        class="prev-next__content__link mr-0 ml-5"
        v-else
      >
        <v-card class="prev pa-5" flat>
          <span class="font-weight-bold">一つ次の画像も見てみる</span>
          <div class="prev-next__content__image">
            <img class="pa-5" :src="next_image.url_path" />
            <p class="ma-0 font-weight-bold text-h7">{{ next_image.title }}</p>
          </div>
        </v-card>
      </router-link>
    </div>

    <!-- おすすめ画像 -->
    <h1 class="others__title mb-5 pb-2">ほかにもこんな仲間たちがいます</h1>
    <v-row class="featured-image__content">
      <v-col cols="3" v-for="image in images.slice(0, 8)" :key="image.id">
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
          <v-btn icon class="ml-auto">
            <v-icon
              @click="favorite(image.id)"
              :color="isFavorite(image.id) ? 'red lighten-1' : ''"
              >mdi-heart</v-icon
            >
          </v-btn>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";
import { saveAs } from "file-saver";

export default {
  name: "DetailImage",
  data() {
    return {
      image: [],
      images: [],
      prev_image: [],
      next_image: [],
      imageTags: [],

      storedIDs: [],
    };
  },
  mounted() {
    this.getImage();
    this.getImages();

    this.storedIDs = JSON.parse(localStorage.getItem("favoriteIDs")) || [];
  },
  watch: {
    $route() {
      location.reload();
    },
  },
  methods: {
    getImage() {
      axios
        .get("http://localhost:8000/get/" + this.$route.params.id)
        .then((response) => {
          this.image = response.data;
          this.getPrevImage(this.$route.params.id);
          this.getNextImage(this.$route.params.id);
          this.imageTags = this.image.tags;
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getImages() {
      axios
        .get("http://localhost:8000/recommended")
        .then((response) => {
          this.images = response.data;

          for (var i = 0; i < this.images.length; i++) {
            if (this.images[i].id == this.$route.params.id) {
              this.images.splice(i, 1);
            }
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getPrevImage(id) {
      axios
        .get("http://localhost:8000/get/prev/" + parseInt(id))
        .then((response) => {
          this.prev_image = response.data;
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getNextImage(id) {
      axios
        .get("http://localhost:8000/get/next/" + parseInt(id))
        .then((response) => {
          this.next_image = response.data;
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getImagesByTag(tagName) {
      console.log(tagName);
      this.$store.dispatch("getImagesByTag", tagName).then(() => {
        this.$router.push({
          path: "/tag/" + tagName,
        });
      });
    },
    async downloadImage(image) {
      try {
        const response = await axios.get(image.url_path, {
          responseType: "blob",
        });
        const fileName = image.url_path.substring(
          image.url_path.lastIndexOf("/") + 1
        );
        saveAs(response.data, fileName);
      } catch (error) {
        console.error("画像のダウンロードに失敗しました。", error);
      }
    },

    // お気に入り機能
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
.image__title {
  background: linear-gradient(transparent 92%, #000000 92%);
}
.others__title {
  background: linear-gradient(transparent 92%, #000000 92%);
}

.content {
  display: flex;
  min-height: 600px;
}

/* メインコンテンツ */
.content__main {
  width: 55%;
  background-color: white;
  border-radius: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
}

img {
  width: 100%;
  height: auto;
  transition: 0.3s ease;
}

/* サブコンテンツ */
.content__sub {
  width: 45%;
}

.content__sub__btn-content__btn {
  /* width: calc(100%); */
  font-size: 20px;
  font-weight: bold;
}

/* .content__sub__btn-content__btn-png {
  margin-right: 5px;
} */

/* .content__sub__btn-content__btn-jpg {
  margin-left: 5px;
} */

.hitokoto {
  /* width: calc(100% - 50px);
  height: calc(100% - 250px); */
  /* max-height: 350px; */
  position: relative;
}

.hitokoto img {
  position: absolute;
  width: 30%;
  min-width: 100px;
  bottom: 0;
  left: 0px;
}

/* .hitokoto__fukidashi {
  padding: 40px;
  height: 100%;
  border: 3px solid black;
  border-radius: 20px;
  background-color: white;
  position: relative;
}
.hitokoto__fukidashi img {
  position: absolute;
  bottom: 0px;
  left: 0px;
  width: 43%;
  height: auto;
} */

/* 次の画像と前の画像 */
.prev-next__content {
  width: 100%;
  text-align: center;
}

.prev-next__content .prev-none-link,
.prev-next__content .next-none-link {
  width: calc(50% - 20px);
  display: block;
}
/* .prev-next__content .prev {
  margin-right: 20px;
}
.prev-next__content .next {
  margin-left: 20px;
} */

.prev-next__content__link {
  color: black;
  text-decoration: none;
  display: block;
  width: calc(50% - 20px);
  transition: ease 0.1s;
}

.prev-next__content__link:hover {
  cursor: pointer;
}

.prev-next__content .prev-next__content__image {
  width: 100%;
  height: auto;
}

.prev-next__content__image img {
  width: 160px;
  height: auto;
}
.prev-next__content__link:hover {
  transform: scale(1.08, 1.08);
}

/* おすすめ画像表示 */
.featured-image__content img {
  width: 100%;
  height: auto;
  transition: 0.3s ease;
}

.featured-image__content img:hover {
  transform: scale(1.1, 1.1);
}
</style>
