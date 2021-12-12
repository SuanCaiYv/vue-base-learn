<template>
    <div>
        <router-view></router-view>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {useRoute} from "vue-router";

const name = ref<String>("Home")

const route = useRoute()
document.title = String(route.meta.title)
</script>

<style scoped>

</style>