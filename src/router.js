import Vue from "vue";
import VueRouter from "vue-router";
import store from "./store";

import HelloWorld from "./components/HelloWorld";
import DetailImage from "./components/DetailImage";
import LoginPage from "./components/LoginPage";

import AdminPage from "./components/Admin/AdminPage";
import NewImageComponent from "./components/Admin/NewImageComponent";
import AdminImageView from "./components/Admin/AdminImageView";
import AdminEditImage from "./components/Admin/AdminEditImage";
import AdminTagPage from "./components/Admin/AdminTagPage";
import HeaderComponent from "./components/Common/HeaderComponent";
import AdminHeaderComponent from "./components/Common/AdminHeaderComponent";
import FooterComponent from "./components/Common/FooterComponent";
import AdminFooterComponent from "./components/Common/AdminFooterComponent";
import axios from "axios";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    components: {
      default: HelloWorld,
      header: HeaderComponent,
      footer: FooterComponent,
    },
    beforeEnter: (to, from, next) => {
      // ルート遷移前にアクションをディスパッチする
      store
        .dispatch("getImages")
        .then(() => {
          // アクションの処理が完了したらルート遷移を続ける
          next();
        })
        .catch((error) => {
          // エラーが発生した場合は、ルート遷移を中止する
          console.error(error);
          next(false);
        });
    },
  },
  {
    path: "/search",
    name: "search",
    components: {
      default: HelloWorld,
      header: HeaderComponent,
      footer: FooterComponent,
    },
    beforeEnter: (to, from, next) => {
      console.log(to.query.q);
      store
        .dispatch("searchItems", to.query.q)
        .then(() => {
          // アクションの処理が完了したらルート遷移を続ける
          next();
        })
        .catch((error) => {
          // エラーが発生した場合は、ルート遷移を中止する
          console.error(error);
          next(false);
        });
    },
  },
  {
    path: "/tag/:name([^/]+)",
    name: "tag",
    components: {
      default: HelloWorld,
      header: HeaderComponent,
      footer: FooterComponent,
    },
    beforeEnter: (to, from, next) => {
      console.log(to.params.name);
      store
        .dispatch("getImagesByTag", to.params.name)
        .then(() => {
          // アクションの処理が完了したらルート遷移を続ける
          next();
        })
        .catch((error) => {
          // エラーが発生した場合は、ルート遷移を中止する
          console.error(error);
          next(false);
        });
    },
  },
  {
    path: "/login",
    name: "login",
    components: {
      default: LoginPage,
      header: HeaderComponent,
      footer: FooterComponent,
    },
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem("token"); // TokenをlocalStorageから取得する
      const data = new FormData();
      data.append("token", token);
      axios
        .post("http://localhost:8000/login/check", data)
        .then(() => {
          next("/admin");
        })
        .catch(() => {
          next();
        });
    },
  },
  {
    path: "/admin",
    name: "admin",
    components: {
      default: AdminPage,
      header: AdminHeaderComponent,
      footer: AdminFooterComponent,
    },
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem("token"); // TokenをlocalStorageから取得する
      const data = new FormData();
      data.append("token", token);
      axios
        .post("http://localhost:8000/login/check", data)
        .then(() => {
          next();
        })
        .catch(() => {
          next("/login");
        });
    },
  },
  {
    path: "/admin/post",
    components: {
      default: NewImageComponent,
      header: AdminHeaderComponent,
      footer: AdminFooterComponent,
    },
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem("token"); // TokenをlocalStorageから取得する
      const data = new FormData();
      data.append("token", token);
      axios
        .post("http://localhost:8000/login/check", data)
        .then(() => {
          next();
        })
        .catch(() => {
          next("/login");
        });
    },
  },
  {
    path: "/admin/images",
    components: {
      default: AdminImageView,
      header: AdminHeaderComponent,
      footer: AdminFooterComponent,
    },
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem("token"); // TokenをlocalStorageから取得する
      const data = new FormData();
      data.append("token", token);
      axios
        .post("http://localhost:8000/login/check", data)
        .then(() => {
          next();
        })
        .catch(() => {
          next("/login");
        });
    },
  },
  {
    path: "/admin/tag",
    components: {
      default: AdminTagPage,
      header: AdminHeaderComponent,
      footer: AdminFooterComponent,
    },
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem("token"); // TokenをlocalStorageから取得する
      const data = new FormData();
      data.append("token", token);
      axios
        .post("http://localhost:8000/login/check", data)
        .then(() => {
          next();
        })
        .catch(() => {
          next("/login");
        });
    },
  },
  {
    path: "/admin/:id",
    components: {
      default: AdminEditImage,
      header: AdminHeaderComponent,
      footer: AdminFooterComponent,
    },
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem("token"); // TokenをlocalStorageから取得する
      const data = new FormData();
      data.append("token", token);
      axios
        .post("http://localhost:8000/login/check", data)
        .then(() => {
          next();
        })
        .catch(() => {
          next("/login");
        });
    },
  },
  {
    path: "/:id(\\d+)",
    name: "detail",
    components: {
      default: DetailImage,
      header: HeaderComponent,
      footer: FooterComponent,
    },
  },
  {
    path: "*",
    components: {
      default: HelloWorld,
      header: HeaderComponent,
      footer: FooterComponent,
    },
  },
];

const scrollBehaviorFunc = function (to, from, savedPosition) {
  // 特定のページのnameをチェックする
  if (to.name === "detail") {
    return { x: 0, y: 0 };
  } else {
    // それ以外の場合は通常のscrollBehaviorを実行する
    return savedPosition || { x: 0, y: 0 };
  }
};

const router = new VueRouter({
  mode: "history",
  scrollBehavior: scrollBehaviorFunc, // 変数として渡す
  routes,
});

export default router;
