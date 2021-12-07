<template>
    <div>
        <router-view></router-view>
    </div>
</template>

<script lang="ts" setup>
import {ref} from "vue"

const name = ref<String>("Home")
</script>

<style scoped>

</style>