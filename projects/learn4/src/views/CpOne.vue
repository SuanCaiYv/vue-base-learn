<template>
    <div>
        <input v-model="count" type="number">
        <br>
        <button @click="incByCount">计数器+N</button>
        <br>
        <input v-model="item" type="text">
        <br>
        <button @click="pushItem">Push元素至列表</button>
        <br>
        {{counter}}
        <br>
        {{msg}}
        <br>
        {{lastItem()}}
    </div>
</template>

<script lang="ts" setup>
import {computed, ref} from "vue"
import {useStore} from "vuex";

const name = ref<String>("CpOne")

const store = useStore()

const counter = computed(() => store.state.counter)
const msg = computed(() => store.state.msg)
const findIdx = store.getters.findIdx
const listLen = computed(() => store.getters.listLen)

const props = defineProps({
    count: Number,
    item: String
})

const incByCount = function () {
    store.commit("incByNum", props.count)
}

const pushItem = function () {
    store.commit("pushItem", props.item)
}

const lastItem = function () {
    return findIdx(listLen.value-1)
}
</script>

<style scoped>
button {
    border: 1px solid black;
    border-radius: 6px;
    background-color: white;
    width: 120px;
}
</style>