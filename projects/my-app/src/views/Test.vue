<template>
    <div>
        <button @click="login">测试</button>
        <br>
        {{resp}}
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import HttpClient from "../net/index"
import {AxiosResponse} from "axios";

const name = ref<String>("Test")
const resp = ref<Object>({})

const login = function () {
    HttpClient.put("/user", {
        username: "2508826394@qq.com",
        credential: "123456"
    }, function (resp: AxiosResponse) {
        console.log(resp)
    })
}
</script>

<style scoped>

</style>