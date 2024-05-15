import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

// https://github.com/vuetifyjs/vuetify-loader/tree/next/packages/vite-plugin
import vuetify from 'vite-plugin-vuetify'


// https://vitejs.dev/config/
export default defineConfig({
	// server: {
	// 	headers: {
	// 		'Cross-Origin-Opener-Policy': 'same-origin',
	// 		'Cross-Origin-Embedder-Policy': 'require-corp'
	// 	}
	// },
	plugins: [
		vue(),
		vuetify({ autoImport: true }),
	],
	resolve: {
		alias: [ // 配置 @ 指代 src
		{
			find: "@",
			replacement: resolve(__dirname, "./src"),
		},
		],
	},
})
