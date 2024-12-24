<template>
  <n-layout>
    <n-layout-header>
      <nav style="display: flex; justify-content: space-between; align-items: center; padding: 8px 16px;">
        <div>
          <router-link to="/" id="main_page" style="margin-right: 16px;">pc-conf</router-link>
          <router-link to="/readymade" id="ready_made" style="margin-right: 16px;">Ready-made</router-link>
          <router-link to="/configurator" id="configurator">Configurator</router-link>
        </div>


        <div style="display: flex; align-items: center; gap: 16px;">
          <a href="https://sso.viktorir.ru/signin?redirect_url=https://pcconf.viktorir.ru" style="text-decoration: none;">
            <n-button>Sign In</n-button>
          </a>

          <n-button-group>
            <n-button @click="setLocale('ru')" :type="locale === 'ru' ? 'primary' : 'default'">
              Ru
            </n-button>
            <n-button @click="setLocale('en')" :type="locale === 'en' ? 'primary' : 'default'">
              En
            </n-button>
          </n-button-group>

          <n-button-group>
            <n-button @click="setTheme('dark')" :type="theme === 'dark' ? 'primary' : 'default'">
              <template #icon>
                <n-icon><moon-outline /></n-icon>
              </template>
            </n-button>
            <n-button @click="setTheme('light')" :type="theme === 'light' ? 'primary' : 'default'">
              <template #icon>
                <n-icon><sunny-outline /></n-icon>
              </template>
            </n-button>
          </n-button-group>
        </div>
      </nav>
    </n-layout-header>
  </n-layout>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';
import { NLayout, NLayoutHeader, NButtonGroup, NButton, NIcon } from 'naive-ui';
import { MoonOutline, SunnyOutline } from '@vicons/ionicons5';

export default {
  components: {
    NLayout,
    NLayoutHeader,
    NButtonGroup,
    NButton,
    NIcon,
    MoonOutline,
    SunnyOutline,
  },
  computed: {
    ...mapGetters('settings', ['theme', 'locale']),
  },
  methods: {
    ...mapActions('settings', ['setTheme', 'setLocale']),
  },
  watch: {
    locale(newLocale) {
      this.$i18n.locale = newLocale;
    },
  },
};
</script>