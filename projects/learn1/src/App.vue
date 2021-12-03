<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import MyBook from "./components/MyBook.vue";
import { ref } from "vue";

const child = ref(null);

const childIncrement = function () {
  child.value.increment()
};
</script>

<template>
  <div>
    <MyBook ref="child"></MyBook>
    <button @click="childIncrement">+1</button>
  </div>
</template>

<style>
#app {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
    Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
