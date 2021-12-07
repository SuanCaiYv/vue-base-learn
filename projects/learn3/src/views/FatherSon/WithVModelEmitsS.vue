<template>
    <div></div>
</template>

<script lang="ts" setup>
import {ref} from "vue"

const name = ref<String>("WithVModelEmitsS")

const props = defineProps({
    msg: String,
    age: Number,
    timeStr: String
})

console.log(props.msg + " + " + props.age)

const emits = defineEmits([
    // v-model版本的父子通信只能触发update类型的事件
    "update:msg",
    "update:age",
    "update:timeStr"
])

setInterval(() => {
    // 只能触发已定义的属性的update事件，且如果想通知父组件，必须事先把父组件变量和子组件属性完成绑定
    emits("update:timeStr", new Date().toTimeString())
}, 1000)
</script>

<style scoped>

</style>