<template>
    <div>
        {{ param1 }}}
    </div>
</template>

<script setup lang="ts">
import {ref, watch} from "vue";
import {useRouter} from "vue-router";

const name = ref<String>("dynamic-route")
const param1 = ref<String>("")
const param2 = ref<String>("")
const router = useRouter()

watch(router, (to, from) => {
    console.log(to)
})
</script>

<style scoped>

</style>