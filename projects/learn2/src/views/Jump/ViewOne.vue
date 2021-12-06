<template>
    <div>
        我是视图一
        <br>
        {{msg}}
        <br>
        <button @click="jump">跳转至视图二</button>
        <br>
        <button @click="back">回退</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {useRoute, useRouter} from "vue-router";

const name = ref<String>("ViewOne")
// 获取路由器
const router = useRouter()
// 获取当前路由
const route = useRoute()
// 通过route访问路由信息
const path = route.matched
const msg = ref<String>(String(route.params.msg))
for (let i = 0; i < path.length; ++ i) {
    // 在这里可以访问到从父级到当前路由的路径信息
    console.log(path[i].path)
}

// 通过router进行跳转，并携带参数
const jump = (): void => {
    router.push({
        name: "view-two",
        params: {
            msg: "我来自试图一"
        }
    })
}
// 通过router进行回退
const back = ():void => {
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