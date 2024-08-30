import { createRouter, createWebHashHistory, NavigationGuardNext, RouteLocationNormalized } from 'vue-router';
import { routes } from './routes';

const history = createWebHashHistory();
const router = createRouter({
    history,
    routes,
});

// 全局路由守卫
router.beforeEach((to: RouteLocationNormalized, _from: RouteLocationNormalized, next: NavigationGuardNext) => {
    const publicPages = ['/login', '/logout']; // 不需要登录的路由
    const authRequired = !publicPages.includes(to.path); // 需要登录的路由

    // 获取当前登录状态
    const loggedIn = !!localStorage.getItem('user');

    // 如果用户已经登录并尝试访问登录页或其他公开页面
    if (loggedIn && to.path === '/login') {
        next('/'); // 或者你想重定向到的其他页面
    } else if (authRequired && !loggedIn) {
        // 如果需要登录且未登录，跳转到登录页
        next('/login');
    } else {
        // 否则，继续访问
        next();
    }
});

export default router;
