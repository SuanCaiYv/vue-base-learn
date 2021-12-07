<template>
    <div>
        <WithVModelEmitsS v-model:age="myAge" v-model:msg="name" v-model:time-str="timeString"></WithVModelEmitsS>
        <br>
        {{ timeString }}
    </div>
</template>

<script lang="ts" setup>
import {ref} from "vue"
import WithVModelEmitsS from "./WithVModelEmitsS.vue";

const name = ref<String>("WithVModelEmitsF")
const myAge = ref<Number>(21)

// 此变量和子组件的timeStr属性完成绑定，所以会被同步更新
const timeString = ref<String>("")
</script>

<style scoped>

</style>