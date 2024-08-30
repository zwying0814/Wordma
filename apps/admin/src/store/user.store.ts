import {LoginUser, User, UserLoginResponse} from "@/types/User.ts";
import {alovaBaseUrlInstance} from "@/utils/http.ts";

export const useUserStore = defineStore('user', () => {
    // token信息
    const token = ref<string>();
    // 用户信息
    const userInfo = ref<User>();
// 用户登录
    const login = async (data: LoginUser) => {
        try {
            const res:UserLoginResponse = await alovaBaseUrlInstance.Post('/api/login', data);
            token.value = res.data.token
            userInfo.value = res.data.user
            console.log(userInfo.value)
        } catch (e) {
            console.log(e)
        }
    }
    return {
        token,
        userInfo,
        login
    }
}, {
    persist: true
})