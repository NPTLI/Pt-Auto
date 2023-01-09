import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import piniaStore from './store';
import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';
import '/@/styles/index.less';
import '/@/styles/reset.less';
import 'uno.css';

// 支持SVG
import 'virtual:svg-icons-register';

const app = createApp(App);
app.use(ArcoVue)
app.use(router);
app.use(piniaStore);
app.mount('#app');
