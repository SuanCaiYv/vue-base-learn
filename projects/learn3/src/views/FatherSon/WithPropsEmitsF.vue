<template>
    <div>
        <!--@子组件事件名="<父组件函数>的形式可以接收子组件发出的事件并处理"-->
        <!--v-bind把父组件的变量的值绑定到子组件属性上-->
        <WithPropsEmitsS :age="21" :msg="name" @showTime="timeStr"></WithPropsEmitsS>
        <br>
        {{ time }}
    </div>
</template>

<script lang="ts" setup>
import {ref} from "vue"
import WithPropsEmitsS from "./WithPropsEmitsS.vue"

const name = ref<String>("WithPropsEmitsF")
const time = ref<String>("")

const timeStr = function (val: String): void {
    time.value = val
}
</script>

<style scoped>

</style>