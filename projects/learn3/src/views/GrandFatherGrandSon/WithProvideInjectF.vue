<template>
    <div>
        <!--假设这里有很多嵌套组件-->
        <WithProvideInjectS></WithProvideInjectS>
    </div>
</template>

<script lang="ts" setup>
import {provide, ref} from "vue"
import WithProvideInjectS from "./WithProvideInjectS.vue";

const name = ref<String>("WithProvideInjectF")

// 我们来测试普通数据类型和对象的传递，这里可以剧透一下，因为对象传的是引用，所以更新对象字段，子孙组件可以收到字段的更新
let msg = name.value
// 这里传递一个对象
const obj = {
    name: "msl",
    age: 21
}
let val = 0
// 函数传参，也能起到更新效果，但是需要在接收端多次调用传递过来的函数，才能得到新的值，不好用。
const counter = function (): Number {
    return val
}

provide("msg", msg)
provide("obj", obj)
provide("counter", (): Number => {
    return val
})

// 尝试不停更新msg的值，看子孙组件是否会更新
setInterval(() => {
    msg = new Date().toTimeString()
}, 1000)

setInterval(() => {
    obj.name = new Date().toTimeString()
    obj.age = 21
}, 1000)

setInterval(() => {
    val += 1
}, 1000)
</script>

<style scoped>

</style>