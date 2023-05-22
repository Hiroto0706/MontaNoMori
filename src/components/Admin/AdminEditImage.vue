<template>
  <v-container>
    <!-- 画像のタイトル -->
    <h1 class="image__title mb-5 pb-2 d-flex">
      <div>ID.{{ image.id }} の画像編集</div>
      <v-spacer></v-spacer>
      <div class="d-flex align-center">
        <v-btn
          large
          color="red"
          outlined
          class="font-weight-bold"
          @click="DeleteImage"
        >
          <v-icon class="pr-1">mdi-trash-can-outline</v-icon>削除
        </v-btn>
        <v-btn
          large
          color="green"
          outlined
          class="font-weight-bold ml-2"
          @click="UpdateImage"
        >
          <v-icon class="pr-1">mdi-upload</v-icon>更新
        </v-btn>
      </div>
    </h1>

    <div class="content mb-5">
      <!-- メインコンテンツ -->
      <div class="content__card px-15 pt-15 pb-10 mr-5">
        <div class="content__main">
          <img :src="image.url_path" class="mb-5" />
        </div>
        <v-file-input
          label="画像の編集"
          outlined
          color="black"
          show-size
          counter
          v-model="newImage"
          full-width
        ></v-file-input>
      </div>

      <!-- サブコンテンツ -->
      <div class="content__sub">
        <v-card-text class="pa-0">
          <v-card-title class="text-h6 font-weight-bold pa-0 mb-3 black--text"
            >タイトル</v-card-title
          >
          <v-text-field
            color="black"
            label="画像タイトル"
            outlined
            maxlength="20"
            counter
            class="mb-3"
            v-model="image.title"
          ></v-text-field>
          <v-card-title class="text-h6 font-weight-bold pa-0 mb-3 black--text"
            >説明</v-card-title
          >
          <v-textarea
            outlined
            color="black"
            maxlength="100"
            counter
            class="mb-3"
            label="画像説明"
            v-model="image.description"
          ></v-textarea>
          <v-card-title class="text-h6 font-weight-bold pa-0 mb-3 black--text"
            >タグ</v-card-title
          >
          <v-card flat class="pa-5">
            <!-- aiueo
            <v-btn
              v-bind="attrs"
              v-on="on"
              depressed
              outlined
              block
              x-large
              class="px-3 py-1"
              ><v-icon>mdi-plus</v-icon>
              タグの追加
            </v-btn> -->
            <v-select
              v-model="imageTags"
              :items="tags"
              item-text="name"
              item-value="id"
              label="画像のタグ"
              color="black"
              multiple
              chips
              hint="タグを選択してください"
              persistent-hint
              outlined
            ></v-select>
          </v-card>
        </v-card-text>
        <!-- 
        <div class="hitokoto ma-5">
          <div class="hitokoto__fukidashi">
            <p v-if="image.description" class="ma-0 font-weight-bold text-h7">
              {{ image.description }}
            </p>
            <p v-else class="ma-0 font-weight-bold text-h7" style="color: gray">
              画像の説明
            </p>
          </div>
        </div> -->
      </div>
    </div>
  </v-container>
</template>

<script>
import axios from "axios";

export default {
  name: "DetailImage",
  data() {
    return {
      image: [],
      newImage: [],
      imageTags: [],
      tags: [],
    };
  },
  mounted() {
    this.getImage();
    this.getTags();
    // this.getImages();
  },
  methods: {
    getImage() {
      axios
        .get("http://localhost:8000/get/" + this.$route.params.id)
        .then((response) => {
          this.image = response.data;

          this.image.tags.sort((a, b) => a - b);
          this.image.tags.forEach((tag) => {
            this.imageTags.push(tag.id);
          });
        })
        .catch((error) => {
          console.log(error);
        });
    },
    UpdateImage() {
      console.log(this.imageTags);

      // サーバーに渡すデータ;
      const data = new FormData();
      data.append("title", this.image.title);
      data.append("description", this.image.description);
      data.append("file", this.newImage);

      this.imageTags.sort((a, b) => a - b);
      this.imageTags.forEach((tag) => {
        data.append("tags[]", tag);
      });

      axios
        .put("http://localhost:8000/admin/edit/" + this.image.id, data)
        .then((response) => {
          this.image = response.data;
          this.newImage = [];
        })
        .catch((error) => {
          console.log(error);
        });
    },
    DeleteImage() {
      axios
        .delete("http://localhost:8000/admin/delete/" + this.image.id)
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
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
};
</script>

<style scoped>
.image__title {
  background: linear-gradient(transparent 92%, #000000 92%);
}

.content {
  display: flex;
  min-height: 600px;
}

/* メインコンテンツ */
.content__card {
  width: 55%;
  background-color: white;
  border-radius: 10px;
}
.content__main {
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
  width: calc(100%);
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
  width: calc(100% - 50px);
  height: calc(100% - 250px);
  max-height: 350px;
}

.hitokoto__fukidashi {
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
}

/* 次の画像と前の画像 */
.prev-next__content {
  width: 100%;
  text-align: center;
}
.prev-next__content .prev,
.prev-next__content .next {
  width: calc(50% - 20px);
}
.prev-next__content .prev {
  margin-right: 20px;
}
.prev-next__content .next {
  margin-left: 20px;
}

.prev-next__content__link {
  color: black;
  text-decoration: none;
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
.prev-next__content__image img:hover {
  transform: scale(1.1, 1.1);
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
