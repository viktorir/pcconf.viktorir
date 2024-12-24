<template>
  <n-h2 style="margin: 1em auto;">Ready-Made configuration</n-h2>
  <n-spin :show="loading" size="large" description="Loading...">
    <template #default>
      <n-alert v-if="error" type="error" style="margin-bottom: 20px;">
        {{ error }}
      </n-alert>

      <div class="cards-container-wrapper" style="display: flex; flex-direction: column; align-items: center;">
        <div v-if="paginatedConfigs.length" class="cards-container">
          <n-card
            v-for="config in paginatedConfigs"
            :key="config.id_config"
            style="margin: 10px; width: 100%;" 
          >
            <img src="@/assets/image.png" alt="pc" style="width: 100%; height: auto; border-radius: 8px; margin-bottom: 10px;" />

            <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 10px;">
              <n-h3>{{ config.name }}</n-h3>
              <n-button size="small" type="default">Settings</n-button>
            </div>

            <n-descriptions size="small" :column="1">
              <n-descriptions-item v-if="config.description" label="Description">{{ config.description }}</n-descriptions-item>
              <n-descriptions-item v-if="config.cpu" label="CPU">{{ config.cpu.name }}</n-descriptions-item>
              <n-descriptions-item v-if="config.card" label="GPU">{{ config.card.name }}</n-descriptions-item>
              <n-descriptions-item v-if="config.ram" label="RAM">{{ config.ram.name }}</n-descriptions-item>
              <n-descriptions-item v-if="config.storage" label="Storage">{{ config.storage.name }}</n-descriptions-item>
            </n-descriptions>
          </n-card>
        </div>
        <div v-else style="display: flex; flex-direction: column; align-items: center; padding: 8em 0;">
          <n-alert type="info" style="margin-bottom: 20px;">
            No more configurations available. But you can create one!
          </n-alert>
          <n-button type="primary" size="large" @click="goToConfigurator">
            Go to Configurator
          </n-button>
        </div>

        <n-pagination
          :page="currentPage"
          :page-count="totalPages"
          @update:page="updatePage"
          @update:page-size="updatePageSize"
          style="margin: 1em 0;"
        />
      </div>
    </template>
  </n-spin>
</template>

<script>
import {
  NCard,
  NSpin,
  NAlert,
  NButton,
  NH3,
  NH2,
  NDescriptions,
  NDescriptionsItem,
  NPagination,
} from 'naive-ui';
import { mapGetters, mapActions } from 'vuex';

export default {
  name: 'ConfigCards',
  components: {
    NCard,
    NSpin,
    NAlert,
    NButton,
    NH3,
    NH2,
    NDescriptions,
    NDescriptionsItem,
    NPagination,
  },
  computed: {
    ...mapGetters('readyMade', ['paginatedConfigs', 'loading', 'error', 'currentPage', 'totalPages']),
  },
  methods: {
    ...mapActions('readyMade', ['fetchConfigs', 'setCurrentPage']),
    updatePage(page) {
      this.setCurrentPage(page);
    },
    updatePageSize(size) {
      this.$store.commit('readyMade/SET_CONFIGS_PER_PAGE', size); // Предположим, есть mutation для configsPerPage
      this.setCurrentPage(1);
      this.fetchConfigs(1);
    },
  },
  mounted() {
    this.fetchConfigs(this.currentPage);
  },
};
</script>

<style scoped>
.cards-container-wrapper {
  display: flex;
  justify-content: center;
  padding: 0 16em; /* Отступ слева */
}

.cards-container {
  display: grid;
  grid-template-columns: repeat(3, 1fr); /* 3 столбца */
  gap: 16px;
  justify-items: center;
}

@media (max-width: 1024px) {
  .cards-container-wrapper {
    padding: 0 8em; /* Отступ слева */
  }
}

@media (max-width: 768px) {
  .cards-container-wrapper {
    padding: 0 4em; /* Отступ слева */
  }

  .cards-container {
    grid-template-columns: repeat(2, 1fr); /* 2 столбца на устройствах с шириной экрана до 768px */
  }
}

@media (max-width: 480px) {
  .cards-container-wrapper {
    padding: 0 1em; /* Отступ слева */
  }

  .cards-container {
    grid-template-columns: 1fr; /* 1 столбец на мобильных устройствах */
    padding-left: 1em; /* Для мобильных устройств можно уменьшить отступ */
    padding-right: 1em; /* Для мобильных устройств можно уменьшить отступ */
  }
}
</style>
