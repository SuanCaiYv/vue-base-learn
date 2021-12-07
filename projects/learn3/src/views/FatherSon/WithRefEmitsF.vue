<template>
    <div>
        <!--设置子组件的引用-->
        <WithRefEmitsS ref="child" @showTime="updateTimeStr"></WithRefEmitsS>
        <br>
        {{ timeStr }}
    </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue"
import WithRefEmitsS from "./WithRefEmitsS.vue";

const name = ref<String>("WithRefEmitsF")
const timeStr = ref<String>("")
// 设置一个变量，保存子组件的引用
const child = ref<HTMLElement>(null)

onMounted(() => {
    console.log(child)
    child.value.msg = name.value
    child.value.age = 21
    child.value.f()
})

const updateTimeStr = function (val: String): void {
    timeStr.value = val
}
</script>

<style scoped>

</style>