<template>
    <div>
        {{ msg }}
        <br>
        {{ age }}
    </div>
</template>

<script lang="ts" setup>
import {ref} from "vue"

const name = ref<String>("WithPropsEmitsS")

// 定义属性，用来接收来自父组件的值
const props = defineProps({
    msg: String,
    age: Number
})

console.log(props.msg + " + " + props.age)

// 子组件定义要发出的事件
const emits = defineEmits([
    "showTime"
])

setInterval(() => {
    // 发出事件，并设置事件参数，由父组件接收
    emits("showTime", new Date().toTimeString())
}, 1000)
</script>

<style scoped>

</style>