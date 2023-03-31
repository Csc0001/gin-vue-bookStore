import Vue from 'vue';
import VueRouter from 'vue-router';
import store from '@/store';
import HomeView from '../views/HomeView.vue';
import userRoutes from './module/user';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue'),
  },
  ...userRoutes,
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.meta.auth) { // 是否需要判断登录
    // 判断用户是否登录

    if (store.state.userModule.token) {
      // 判断token有效性
      next();
    } else {
      router.push({ name: 'login' });
    }
  } else {
    next();
  }
});
export default router;
