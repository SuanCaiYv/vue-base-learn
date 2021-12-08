<template>
    <div>
        {{ msg }}
        <br>
        {{ obj.name }}
        <br>
        {{ obj.age }}
    </div>
</template>

<script lang="ts" setup>
import {inject, ref} from "vue"

const name = ref<String>("WithProvideInjectReactiveS")

// 注入的响应式对象可以直接获取，然后直接作为响应式对象使用
const msg = inject("msg")
const obj = inject("obj")
</script>

<style scoped>

</style>