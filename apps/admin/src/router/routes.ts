import {RouteRecordRaw} from "vue-router";

export const routes:Array<RouteRecordRaw> = [
    {
        path:'/',
        redirect: '/home'
    },
    {
        name: 'login',
        path: '/login',
        component: () => import('../views/login/login.vue'),
        meta: {
            title: '登录页'
        }
    },
    {
        name: 'home',
        path: '/home',
        component: () => import('../views/home/home.vue'),
        meta: {
            title: '后台首页'
        }
    },
    {
        name:'comment',
        path:'/comment',
        component: () => import('../views/comment/comment.vue'),
        meta: {
            title: '评论管理'
        }
    },
    {
        name:'post',
        path:'/post',
        component: () => import('../views/post/post.vue'),
        meta: {
            title: '文章数据'
        }
    }
]