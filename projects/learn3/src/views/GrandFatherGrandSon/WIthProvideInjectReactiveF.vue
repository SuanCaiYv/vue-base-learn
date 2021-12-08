<template>
    <div>
        <!--假设这里嵌套了很多组件-->
        <WithProvideInjectReactiveS></WithProvideInjectReactiveS>
    </div>
</template>

<script lang="ts" setup>
import {provide, reactive, ref} from "vue"
import WithProvideInjectReactiveS from "./WithProvideInjectReactiveS.vue";

const name = ref<String>("WIthProvideInjectReactiveF")

const msg = ref<String>("")
const obj = reactive({
    name: "",
    age: 21
})

// 在这里我们直接注入响应式对象，子孙组件可以直接获取作为响应式对象使用
provide("msg", msg)
provide("obj", obj)

setInterval(() => {
    msg.value = new Date().toTimeString()
}, 1000)

setInterval(() => {
    obj.name = new Date().toTimeString()
    obj.age = 21
}, 1000)
</script>

<style scoped>

</style>