<template>
    <div>
        {{ msg0 }}
        <br>
        {{ obj0.name }}
        <br>
        {{ obj0.age }}
        <br>
        {{ counter0 }}
    </div>
</template>

<script lang="ts" setup>
import {inject, reactive, ref} from "vue"

const name = ref<String>("WithProvideInjectS")

const msg = inject("msg")
const obj = inject("obj")
const counter = inject("counter")
// 用于模板渲染的响应式变量
const msg0 = ref<String>("")
const obj0 = reactive({
    name: "",
    age: 21,
})
const counter0 = ref<Number>(0)

// 这里说一下为什么不用模板渲染，是因为模板渲染只适用于响应式数据，当数据不是响应式时，模板渲染结果不会因为数据更新而更新
// 但是我们可以用间接传值实现在模板中看到最新的值，即定时器+新的响应式变量
setInterval(() => {
    msg0.value = String(msg)
}, 2000)

// 通过引用类型可以传递被更新的值，但是无法做到和模板响应式渲染结合
setInterval(() => {
    obj0.name = obj.name
    obj0.age = obj.age
}, 2000)

// 传个函数过来，可以获得最新的值，但是每次都要调用，很麻烦
setInterval(() => {
    counter0.value = counter()
}, 1000)
</script>

<style scoped>

</style>