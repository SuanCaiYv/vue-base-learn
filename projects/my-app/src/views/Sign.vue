<template>
    <div class="sign">
        <div class="s-o-one"></div>
        <div class="main">
            <div class="input-box">
                <div class="l1">
                    <div class="msg1 center">
                        <span>用户名</span>
                    </div>
                    <input class="input1" type="text" v-model="username">
                </div>
                <div class="l2">
                    <div class="msg2 center">
                        <span>密&nbsp;&nbsp;&nbsp;&nbsp;码</span>
                    </div>
                    <input class="input2" type="password" v-model="credential">
                </div>
                <div class="l3">
                    <div class="msg3 center" @click="sendCode">
                        <span>验证码</span>
                    </div>
                    <input class="input3" type="text" v-model="verCode">
                </div>
                <div class="l4">
                    <button class="login center" type="submit" @click="login">
                        <span>登录</span>
                    </button>
                    <button class="signup center" type="submit" @click="signup">
                        <span>注册</span>
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import HttpClient from "../net/index"
import {AxiosResponse} from "axios";

const name = ref<String>("Sign")

let props = defineProps({
    username: String,
    credential: String,
    verCode: String
})

const login = function () {
    const resp = ref<Object>({})
    HttpClient.put("/user", {
        username: props.username,
        credential: props.credential
    }, function (resp: AxiosResponse) {
        if (resp.status !== 200) {
            alert("?")
        } else {
            if (resp.data.code === 200) {
                alert("OK")
            } else {
                alert(resp.data.msg)
            }
        }
    })
}

const sendCode = function () {
    console.log("已发送")
}

const signup = function () {
}
</script>

<style scoped>
.sign {
    height: 100%;
    display: grid;
    grid-template-areas: "s-o-one main";
    grid-template-columns: 1fr 1fr;
    grid-template-rows: 1fr;
    background-color: cornflowerblue;
}

.s-o-one {
    grid-area: s-o-one;
    background-color: #42b983;
}

.main {
    grid-area: main;
    background-color: dodgerblue;
}

.input-box {
    width: 400px;
    height: 300px;
    margin: 200px auto auto;
    border: 1px solid black;
    border-radius: 8px;
    display: grid;
    grid-template-areas:
        "l1"
        "l2"
        "l3"
        "l4";
    grid-template-rows: 1fr 1fr 1fr 1fr;
    background-color: cadetblue;
}

.l1 {
    border: 1px solid black;
    border-radius: 8px;
    grid-area: l1;
    margin: 10px;
    display: grid;
    grid-template-areas: "msg1 input1";
    grid-template-columns: 1fr 3fr;
    background-color: goldenrod;
}

.msg1 {
    border: 1px solid black;
    border-radius: 8px;
    grid-area: msg1;
    margin: 10px;
    background-color: bisque;
}

.input1 {
    border: 1px solid black;
    border-radius: 8px;
    grid-area: input1;
    margin: 10px;
    font-size: 1rem;
    background-color: aquamarine;
}

.l2 {
    border: 1px solid black;
    border-radius: 8px;
    grid-area: l2;
    margin: 10px;
    display: grid;
    grid-template-areas: "msg2 input2";
    grid-template-columns: 1fr 3fr;
    background-color: gold;
}

.msg2 {
    border: 1px solid black;
    border-radius: 8px;
    grid-area: msg2;
    margin: 10px;
    background-color: darkorange;
}

.input2 {
    border: 1px solid black;
    border-radius: 8px;
    grid-area: input2;
    margin: 10px;
    font-size: 1rem;
    background-color: darkorchid;
}

.l3 {
    border: 1px solid black;
    border-radius: 8px;
    grid-area: l3;
    margin: 10px;
    display: grid;
    grid-template-areas: "msg3 input3";
    grid-template-columns: 1fr 3fr;
    background-color: cornsilk;
}

.msg3 {
    border: 1px solid black;
    border-radius: 8px;
    grid-area: msg3;
    margin: 10px;
    background-color: deepskyblue;
}

.input3 {
    border: 1px solid black;
    border-radius: 8px;
    grid-area: input3;
    margin: 10px;
    font-size: 1rem;
    background-color: khaki;
}

.l4 {
    border: 1px solid black;
    border-radius: 8px;
    grid-area: l4;
    margin: 10px;
    display: grid;
    grid-template-areas: "login signup";
    background-color: azure;
}

.login {
    border: 1px solid black;
    border-radius: 8px;
    grid-area: login;
    margin: 10px 40px;
    font-size: 1.2rem;
    background-color: hotpink;
}

.login:hover {
    background-color: lightpink;
}

.login:active {
    background-color: pink;
}

.signup {
    border: 1px solid black;
    border-radius: 8px;
    grid-area: signup;
    margin: 10px 40px;
    font-size: 1.2rem;
    background-color: lightpink;
}

.center {
    display: flex;
    justify-content: center;
    align-items: center;
}

span {
    display: inline-block;
    margin-top: auto;
    margin-bottom: auto;
    font-stretch: ultra-expanded;
}
</style>