<template>
    <div></div>
</template>

<script lang="ts" setup>
import {ref} from "vue"

const name = ref<String>("WithRefEmitsS")

const msg = ref<String>("")
const age = ref<Number>(0)

const f = function () {
    console.log(msg.value + " + " + age.value)
}

const emits = defineEmits([
    "showTime"
])

setInterval(() => {
    emits("showTime", new Date().toTimeString())
}, 1000)
</script>

<style scoped>

</style>