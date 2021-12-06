<template>
    <div>
        <button @click="router.push('cp-two')">跳往组件二</button>
        <br>
        {{msg}}
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {useRoute, useRouter} from "vue-router";

const name = ref<String>("CpOne")
const msg = ref<String>("")
const route = useRoute()
const router = useRouter()

const title = String(route.meta.title)
const requireAuthed = Boolean(route.meta.requireAuthed)
document.title = title
if (requireAuthed) {
    msg.value = "此页面需要认证"
} else {
    msg.value = "此页面允许游客访问"
}
</script>

<style scoped>
button {
    background-color: white;
    border: 1px solid black;
    border-radius: 6px;
    width: 120px;
}
</style>