<template>
    <div>
        <router-view></router-view>
    </div>
</template>

<script lang="ts" setup>
import {ref} from "vue"

const name = ref<String>("Home")

// ref(1) === reactive({value: 1})
</script>

<style scoped>

</style>