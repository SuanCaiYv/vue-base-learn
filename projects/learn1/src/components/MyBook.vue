<script setup>
import {ref} from 'vue'

const counter = ref(0)

const increment = function() {
    counter.value++
}

const msg = "aaa"

defineExpose({
    increment,
    msg
})
</script>

<template>
<p>
    {{counter}}
</p>
</template>