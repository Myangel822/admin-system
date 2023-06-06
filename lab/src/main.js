import { createApp } from 'vue'
// import createStore from 'store/index.js'
import App from './App.vue'
import router from './router'
import store from './store'
// import Vue  from 'vue'
import ElementPlus from 'element-plus'
import  'element-plus/theme-chalk/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
// Vue.config.productionTip = false
const app=createApp(App)
for (const[key,component] of Object.entries(ElementPlusIconsVue)){
	app.component(key,component)
}
app.use(store).use(router).use(ElementPlus).mount('#app')
