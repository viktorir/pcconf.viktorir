import { createApp } from 'vue'
import { create, NButton, NInput, NH1, NH2, NH3, NText, NIcon} from 'naive-ui';
import AppProviders from './AppProviders.vue'
import router from './router'
import i18n from './i18n'
import store from './store'

const naive = create({
    components: [NButton, NInput, NH1, NH2, NH3, NText, NIcon ], // Зарегистрируйте нужные компоненты
});

createApp(AppProviders).
    use(store).
    use(naive).
    use(router).
    use(i18n).
    mount('body')
