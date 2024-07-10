import { createRouter } from "vue-router";
import { createWebHistory } from "vue-router";


const routes = [
  {
    path: "/",
    alias: "/login",
    component: () => import("@/views/LoginPage/LoginPage.vue"),
  },
  {
    path: "/login",
    name: "LoginPage",
    component: () => import("@/views/LoginPage/LoginPage.vue"),
  }, {
    path: "/print",
    name: "PrintPage",
    component: () => import("@/views/PrintPage/PrintPage.vue"),
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});


// 路由守卫
router.beforeEach((to, from, next) => {
  // console.log(from);
  // console.log(to.path);
  const teamInfo = JSON.parse(window.localStorage.getItem('teamInfo') || '{}');
  if (from.path === '/' && to.path === '/') {
    // 初始进入页面时，判断是否有token, 有则直接跳转到打印页面
    console.log(teamInfo);
    if (teamInfo.token) {
      next({ name: 'PrintPage'});
    } else {
      next({ name: 'LoginPage'});
    }
  }
  else if (to.path === '/print') {
    // 前往打印页面之前，判断是否有token
    if (teamInfo.token) {
      next();
    } else {
      next({ name: 'LoginPage'});
    }
  }
  else if (from.path === '/print' &&(to.path === '/login' || to.path === '/')) {
    // 从打印页面返回时，清空缓存
    window.localStorage.clear();
    next();
  }
  else {
    next();
  }
});


export default router;