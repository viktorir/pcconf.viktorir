<template>
    <header>
        <nav>
            <router-link to="/" id="main_page">pc-conf</router-link>
            <router-link to="/readymade" id="ready_made">Ready-made</router-link>
            <router-link to="/configurator" id="configurator">Configurator</router-link>
        </nav>
        <div>
            <button>Search</button>
            <a href="http://localhost:8081/signin?redirect_url=http://localhost:8080">Sign In</a>
        </div>
    </header>
</template>

<script>
export default {
  data() {
    return {
      showModal: false,
      email: "",
      password: "",
    };
  },
  methods: {
    submitForm() {
      console.log("Email:", this.email);
      console.log("Password:", this.password);
      this.showModal = false; // Закрыть модальное окно после отправки
    },
  },
}
</script>

<style scoped>
header {
    height: 80px;
    padding: 8px 10%;
    background-color: #404040;

    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
}

nav {
    
}

#main_page {
    padding: 8px;
    background-color: #E31C25;
    color: #FFFFFF;
    font-size: 32px;
    text-decoration: none;
    border-radius: 8px;
}

#ready_made {
    padding: 8px;
    color: #FFFFFF;
    font-size: 16px;
    font-weight: bold;
    text-decoration: none;
}

#configurator {
    padding: 8px;
    color: #FFFFFF;
    font-size: 16px;
    font-weight: bold;
    text-decoration: none;
}

button {
    padding: 8px;
    border: none;
    font-size: 16px;
    font-weight: bold;
    color: #FFFFFF;
    background-color: transparent;
    cursor: pointer;
}

a{
    all: unset;

    padding: 8px;
    border: none;
    font-size: 16px;
    font-weight: bold;
    color: #FFFFFF;
    background-color: transparent;
    cursor: pointer;
}
</style>