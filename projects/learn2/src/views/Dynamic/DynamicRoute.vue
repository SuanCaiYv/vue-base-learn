<template>
    <div>
        {{ id }}
        <br>
        {{ age }}
    </div>
</template>

<script lang="ts" setup>
import {ref, watchEffect} from "vue";
import {useRoute} from "vue-router";

const name = ref<String>("dynamic-route")
const route = useRoute()
const id = ref<String>("")
const age = ref<Number>(0)

const f = (): void => {
    id.value = String(route.params.id)
    age.value = Number(route.params.age)
}
// 我他妈直接无语，watch不生效，改成watchEffect就行了
watchEffect(f)
</script>

<style scoped>

</style>