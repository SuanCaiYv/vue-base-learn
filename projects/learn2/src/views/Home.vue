<template>
    <div class="app">
        我是主页
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue";

const name = ref<String>("home")
</script>

<style scoped>
.app {
    color: red;
}
</style>