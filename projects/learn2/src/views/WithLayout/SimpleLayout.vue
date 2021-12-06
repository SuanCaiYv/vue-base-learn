<template>
    <div class="app">
        <MyHeader>我是页头</MyHeader>
        <main class="main">
            <router-view></router-view>
        </main>
        <MyFooter>我是页尾</MyFooter>
    </div>
</template>

<script lang="ts" setup>
import MyHeader from "../MyHeader.vue";
import MyFooter from "../MyFooter.vue";</script>

<style scoped>
.main {
    margin-top: 25px;
}

.app {
    width: 100%;
}
</style>