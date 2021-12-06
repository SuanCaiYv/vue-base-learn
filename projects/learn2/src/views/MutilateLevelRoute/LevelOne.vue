<template>
    <div>
        我是一级路由组件
    </div>
    <router-view></router-view>
</template>

<script lang="ts" setup>
import {ref} from "vue";

const name = ref<String>("LevelOne")
</script>

<style scoped>

</style>