import {createAlova, Method} from 'alova';
import VueHook from 'alova/vue';
import adapterFetch from "alova/fetch";
import {message} from "ant-design-vue";
import router from "@/router";

/**
 * 创建一个Alova实例，带token等信息
 */
export const alovaBaseUrlInstance = createAlova({
    baseURL: import.meta.env.VITE_BASE_API,
    beforeRequest(method: Method) {
        const user = JSON.parse(localStorage.getItem('user') || '{}');
        if (user.token) {
            method.config.headers.Authorization = `Bearer ${user.token}`;
        }
    },
    cacheFor: null, // 关闭Alova的请求缓存
    responded: {
        onSuccess: async (response: Response) => {
            if (response.status===401){
                message.error('登录已过期，请重新登录')
                localStorage.removeItem('user')
                return await router.replace('/login')
            }
            if (response.status > 200) {
                message.error(response.statusText)
                return
            }
            const json = await response.json();
            if (response.status != 200) {
                message.error(json.message || response.statusText)
                console.log(response.statusText)
                return
            }
            if (json.code !== 200) {
                message.error(json.message)
                return
            }
            return json;
        },
        onError: async (err) => {
            console.log(err)
            message.error(err.message)
        },
    },
    statesHook: VueHook,
    requestAdapter: adapterFetch()
})