<template>
    <div>
        <button @click="counterInc">计数器+1</button>
        <br>
        <input v-model="msgVal" type="text">
        <br>
        <button @click="updateMsg">更新消息</button>
        <br>
        {{ counter }}
        <br>
        {{ msg }}
        <br>
        {{ listLen }}
    </div>
</template>

<script lang="ts" setup>
import {computed, ref} from "vue"
import {useStore} from "vuex";

const name = ref<String>("CpTwo")

// 获取store对象
const store = useStore()

// 通过计算属性获取响应式状态
const counter = computed(() => store.state.counter)
const msg = computed(() => store.state.msg)
const listLen = computed(() => store.getters.listLen)

const counterInc = function () {
    // 只能通过commit提交状态更新事件
    store.commit("inc")
}

const props = defineProps({
    msgVal: String
})

const updateMsg = function () {
    store.commit("updateMsg", {
        value: props.msgVal
    })
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