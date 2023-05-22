import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    // 状態を定義
    title: "イラストたち",
    images: [],
    tags: [],
    tag: "",
    mode: "new",
  },
  mutations: {
    // 状態を変更する関数を定義
    SET_IMAGES(state, images) {
      state.images = images;
    },
    SET_TITLE(state, title) {
      state.title = title;
    },
    SET_TAGS(state, tags) {
      state.tags = tags;
    },
    SET_TAG(state, tag) {
      state.tag = tag;
    },
    SET_MODE(state, mode) {
      state.mode = mode;
    },
  },
  actions: {
    // 非同期処理を含む関数を定義
    async getImages({ commit }) {
      axios
        .get("http://localhost:8000/")
        .then((response) => {
          commit("SET_TITLE", "イラストたち");
          commit("SET_MODE", "new");
          commit("SET_IMAGES", response.data);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    // 検索機能を実施する処理
    async searchItems({ commit }, target) {
      const data = new FormData();
      data.append("target", target);
      console.log(target);

      await axios
        .get("http://localhost:8000/search", { params: { q: target } })
        .then((response) => {
          commit("SET_TITLE", "「" + target + "」" + "で検索");
          commit("SET_MODE", "search");
          commit("SET_IMAGES", response.data);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    async getImagesByTag({ commit }, tagName) {
      await axios
        .get("http://localhost:8000/tag/" + tagName)
        .then((response) => {
          console.log(response.data);
          commit("SET_TITLE", "「#" + tagName + "」" + "でタグ検索");
          commit("SET_MODE", "tag");
          commit("SET_IMAGES", response.data);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    // タグをとってくる機能
    async fetchTags({ commit }) {
      await axios
        .get("http://localhost:8000/admin/tag/get")
        .then((response) => {
          const tags = response.data;
          commit("SET_TAGS", tags);
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
  getters: {
    // 状態を加工して返す関数を定義
  },
});
