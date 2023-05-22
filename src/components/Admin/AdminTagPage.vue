<template>
  <v-container>
    <h1 class="my-10 pl-3 pb-3 text-h4 font-weight-bold page__title">
      タグ一覧
    </h1>

    <!-- タグ一覧 -->
    <div class="d-flex flex-wrap mb-10">
      <div v-for="tag in tags" :key="tag.id">
        <v-dialog v-model="dialog[tag.id]" max-width="600px">
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              v-bind="attrs"
              v-on="on"
              depressed
              outlined
              rounded
              style="text-transform: none"
              class="mx-1 my-1 px-3 py-1"
            >
              # {{ tag.name }}
            </v-btn>
          </template>
          <v-card>
            <v-card-title>
              <span class="text-h5">タグの編集</span>
              <v-spacer></v-spacer>
              <v-btn color="grey darken-1" icon @click="dialog[tag.id] = false">
                <v-icon>mdi-close</v-icon>
              </v-btn>
            </v-card-title>
            <v-card-text class="pb-0">
              <v-text-field
                outlined
                label="タグの名前"
                class="mt-5"
                color="black"
                v-model="tag.name"
              ></v-text-field>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                color="red darken-1"
                text
                @click="
                  dialog[tag.id] = false;
                  tagDelete(tag.id);
                "
              >
                <v-icon small class="pr-1">mdi-trash-can-outline</v-icon>
                Delete
              </v-btn>
              <v-btn color="blue darken-1" text @click="tagEdit(tag)">
                <v-icon small class="pr-1">mdi-file-edit-outline</v-icon>
                Edit
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </div>
    </div>

    <!-- 追加のダイアログ -->
    <v-dialog v-model="dialogN" max-width="600px">
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          v-bind="attrs"
          v-on="on"
          depressed
          outlined
          block
          x-large
          class="mx-1 my-1 px-3 py-1"
          ><v-icon>mdi-plus</v-icon>
          タグの追加
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="text-h5">タグの追加</span>
          <v-spacer></v-spacer>
          <v-btn color="grey darken-1" icon @click="dialogN = false">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-card-title>
        <v-card-text class="pb-0">
          <v-text-field
            outlined
            label="追加するタグの名前"
            class="mt-5"
            color="black"
            v-model="newTagName"
          ></v-text-field>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="blue darken-1"
            text
            :rules="[rules.required, rules.counter]"
            @click="
              dialogN = false;
              tagCreate();
            "
          >
            <v-icon small>mdi-plus</v-icon>
            Create
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      dialog: {},
      dialogN: false,
      newTagName: "",
      tags: [],
      rules: {
        required: (value) => !!value || "必須項目だよ。",
        counter: (value) => value.length <= 20 || "最大20文字までだよ。",
      },
    };
  },
  mounted() {
    this.getTags();
  },
  methods: {
    tagCreate() {
      console.log("create tag");
      console.log(this.newTagName);

      // サーバーに渡すデータ;
      const data = new FormData();
      data.append("name", this.newTagName);

      axios
        .post("http://localhost:8000/admin/tag/new", data)
        .then((response) => {
          console.log(response.data);
          this.newTagName = "";
          this.getTags();
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
          console.log(this.tags);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    tagEdit(tag) {
      axios
        .put("http://localhost:8000/admin/tag/edit", tag)
        .then((response) => {
          console.log(response.data);
          this.getTags();
        })
        .catch((error) => {
          console.log(error);
        });
    },
    tagDelete(id) {
      const data = new FormData();
      data.append("id", id);

      axios
        .delete("http://localhost:8000/admin/tag/delete/" + parseInt(id))
        .then((response) => {
          console.log(response.data);
          this.getTags();
        })
        .catch((error) => {
          console.log(error);
        });
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
