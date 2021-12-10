<template>
    <div>
        我是视图二
        <br>
        {{ msg }}
        <br>
        <button @click="jump">跳转至视图三</button>
        <br>
        <button @click="back">回退</button>
    </div>
</template>

<script lang="ts" setup>
import {ref} from "vue"
import {useRoute, useRouter} from "vue-router";

const name = ref<String>("ViewTwo")
const router = useRouter()
const route = useRoute()
// 通过route访问路由信息
const path = route.matched
const msg = ref<String>(String(route.params.msg))
for (let i = 0; i < path.length; ++i) {
    // 在这里可以访问到从父级到当前路由的路径信息
    console.log(path[i].path)
}

// 通过router进行跳转，并携带参数
const jump = (): void => {
    router.push({
        name: "views-three",
        params: {
            msg: "我来自试图二"
        }
    })
}
// 通过router进行回退
const back = (): void => {
    router.back()
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