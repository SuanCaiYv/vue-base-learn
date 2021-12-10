<template>
    <div>
        现在我们来学学VueX的简单用法。
        <br>
        本来learn3应该还有一个是兄弟组件通信/全局组件通信的，但是那个方法处理起来比较麻烦。
        <br>
        所以对于全局组件/兄弟组件通信我们借由VueX来完成，所以在存储这一章中顺手给介绍了。
        <br>
        <router-view></router-view>
    </div>
</template>

<script lang="ts" setup>
import {ref} from "vue"

const name = ref<String>("Home")
</script>

<style scoped>

</style>