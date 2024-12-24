<template>
  <n-flex vertical align="center" style="margin: 0 1em;">
      <n-h2 style="margin-top: 1em;">Ready-made community configurations</n-h2>

      <n-spin :show="loading" size="large" description="Loading..." style="margin-bottom: 20px;">
        <template #default>
          <n-alert v-if="error" type="error" style="margin-bottom: 20px;">
            {{ error }}
          </n-alert>

          <div v-if="configs.length" class="carousel">
            <n-button @click="scrollLeft" circle size="small" style="margin-right: 10px;">←</n-button>
            <div class="card-container" ref="cardContainer">
              <n-card
                v-for="(config, i) in displayedConfigs"
                :key="config.id_config"
                style="margin: 10px; width: 300px;"
              >
                <img src="@/assets/image.png" alt="pc" style="width: 100%; height: auto; border-radius: 8px; margin-bottom: 10px;" />

                <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 10px;">
                  <n-h4>{{ config.name }}</n-h4>
                  <n-button size="small" type="default">Settings</n-button>
                </div>

                <n-descriptions size="small" column="1">
                  <n-descriptions-item label="Description">{{ config.description }}</n-descriptions-item>
                  <n-descriptions-item label="CPU">{{ config.cpu.name }}</n-descriptions-item>
                  <n-descriptions-item label="GPU">{{ config.card.name }}</n-descriptions-item>
                  <n-descriptions-item label="RAM">{{ config.ram.name }}</n-descriptions-item>
                  <n-descriptions-item label="Storage">{{ config.storage.name }}</n-descriptions-item>
                </n-descriptions>
              </n-card>
            </div>
            <n-button @click="scrollRight" circle size="small" style="margin-left: 10px;">→</n-button>
          </div>

          <n-alert v-else type="info">No configurations available.</n-alert>
        </template>
      </n-spin>

      <n-h2>Create your own configuration</n-h2>
      <n-button type="primary" @click="navigateToConfigurator">Configurate</n-button>
  </n-flex>
</template>

<script>
import {
  NCard,
  NSpin,
  NAlert,
  NButton,
  NH2,
  NH4,
  NText,
  NFlex,
  NDescriptions,
  NDescriptionsItem
} from 'naive-ui';
import { mapGetters, mapActions } from 'vuex';

export default {
  components: {
    NCard,
    NSpin,
    NAlert,
    NButton,
    NH2,
    NH4,
    NText,
    NFlex,
    NDescriptions,
    NDescriptionsItem
  },
  data() {
    return {
      currentIndex: 0, 
    };
  },
  computed: {
    ...mapGetters(['configs', 'loading', 'error']),
    displayedConfigs() {
      return this.configs.slice(this.currentIndex, this.currentIndex + 3);
    },
  },
  methods: {
    ...mapActions(['fetchConfigs']),
    scrollLeft() {
      if (this.currentIndex > 0) this.currentIndex--;
    },
    scrollRight() {
      if (this.currentIndex < this.configs.length - 3) this.currentIndex++;
    },
    navigateToConfigurator() {
      this.$router.push('/configurator');
    },
  },
  mounted() {
    // Загружаем конфигурации при монтировании
    this.fetchConfigs();
  },
};
</script>

<style scoped>
.carousel {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}
.card-container {
  display: flex;
  overflow-x: auto;
  flex-wrap: nowrap;
}
</style>