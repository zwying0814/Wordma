import {defineConfig, loadEnv} from 'vite'
import type {UserConfig, ConfigEnv} from 'vite';
import {resolve} from 'node:path';
import vue from '@vitejs/plugin-vue'
import UnoCSS from 'unocss/vite'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {AntDesignVueResolver} from "unplugin-vue-components/resolvers";

// https://vitejs.dev/config/
export default defineConfig(({mode}: ConfigEnv): UserConfig => {
    // 获取当前环境的配置
    const config = loadEnv(mode, './')
    return {
        plugins: [
            vue(),
            UnoCSS(),
            AutoImport({
                imports: ["vue", "vue-router", "pinia"],
                dts: 'src/auto-imports.d.ts',
                resolvers: [AntDesignVueResolver()],
            }),
            Components({
                resolvers: [AntDesignVueResolver(
                    {
                        importStyle: false, // css in js
                        resolveIcons: true,
                    }
                )],
                dts: 'src/components.d.ts'
            }),
        ],
        resolve: {
            alias: {
                "@": resolve(__dirname, "./src")
            }
        },
        server: {
            proxy: {
                '/api': {
                    target: config.VITE_BASE_API, // 目标地址 --> 服务器地址
                    changeOrigin: true, // 允许跨域
                    // 重写路径 --> 作用与vue配置pathRewrite作用相同
                    rewrite: (path) => path.replace(/^\/api/, '/api/v1'),
                    bypass: (req, res, options: any) => {
                        const proxy = options.target + options.rewrite(req.url);
                        console.log(`请求代理: ${req.url} -> ${proxy}`);
                        req.headers['x-req-proxy'] = proxy;
                        res.setHeader('x-req-proxy', proxy)
                    }
                },
            },
        }
    }
})
