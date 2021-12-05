<template>
    <div class="app">
        我是作者信息
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue";

const name = ref<String>("about")
</script>

<style scoped>
.app {
    color: green;
}
</style>