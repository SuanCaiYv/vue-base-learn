<template>
    <div class="msg">
        {{ msg }}
    </div>
</template>

<script lang="ts" setup>
import {ref, watchEffect} from "vue";
import {useRoute} from "vue-router";

const msg = ref<String>("")

const route = useRoute()

const f = (): void => {
    msg.value = String(route.params.msg)
}
watchEffect(f)
</script>

<style scoped>
.msg {
    background-color: white;
    font-size: medium;
}
</style>