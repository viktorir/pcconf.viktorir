<template>
  <h2>Ready-made community configurations</h2>
  <div v-if="loading">Загрузка...</div>
  <div v-if="error">{{ error }}</div>
  <div v-if="configs.length" class="carousel">
    <button @click="scrollLeft" class="scroll-button">←</button>
    <div class="card-container" ref="cardContainer">
      <div class="card" v-for="(config, i) in displayedConfigs" :key="config.id_config">
        <img src="@/assets/image.png" alt="pc">
        <div id="name_row">
          <h3>{{ config.name }}</h3>
          <button>stngs</button>
        </div>
        <p>{{ config.description }}</p>
        <p>{{ config.cpu.name }}</p>
        <p>{{ config.card.name }}</p>
        <p>{{ config.ram.name }}</p>
        <p>{{ config.storage.name }}</p>
      </div>
    </div>
    <button @click="scrollRight" class="scroll-button">→</button>
  </div>
  <h2>Create your own configuration</h2>
  <button>Configurate</button>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

export default {
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
      if (this.currentIndex > 0) {
        this.currentIndex--;
      }
    },
    scrollRight() {
      if (this.currentIndex < this.configs.length - 3) {
        this.currentIndex++;
      }
    },
  },
  mounted() {
    this.fetchConfigs();
  }
};
</script>

<style>
h2 {
  color: #FFFFFF;
}

.carousel {
  padding: 0 10%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.card-container { 
  display: flex;
  overflow: hidden;
}

.card {
  width: 286px;
  min-width: 80px;

  padding: 8px;
  margin: 5px;

  border: 1px solid #ccc;
  border-radius: 4px;
  
  img {
    max-width: 100%;

    border-bottom: #404040 solid 2px;
  }

  h3 {
    color: #E31C25;
  }
  p {
    margin-bottom: 0;
    padding-bottom: 8px;
    color: #FFFFFF;
    border-bottom: #404040 solid 2px;

    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  #name_row {
    padding-bottom: 4px;
    display: flex;
    justify-content: space-between;

    border-bottom: #404040 solid 2px;
  }
}

.scroll-button {
  cursor: pointer;
  padding: 10px;
}
</style>