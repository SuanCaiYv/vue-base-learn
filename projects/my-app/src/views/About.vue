<template>
    <div></div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {useRoute} from "vue-router";

const name = ref<String>("About")
const route = useRoute()
document.title = String(route.meta.title)
</script>

<style scoped>

</style>