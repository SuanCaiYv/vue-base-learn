<template>
    <div class="sign">
        <div class="s-o-one"></div>
        <div class="main">
            <div class="input-box">
                <div class="l1">
                    <div class="msg1 center">
                        <span>用户名</span>
                    </div>
                    <div class="input1"></div>
                </div>
                <div class="l2">
                    <div class="msg2 center">
                        <span>密码</span>
                    </div>
                    <div class="input2"></div>
                </div>
                <div class="l3"></div>
                <div class="l4"></div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Sign")
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
    grid-area: l1;
    margin: 10px;
    display: grid;
    grid-template-areas: "msg1 input1";
    grid-template-columns: 1fr 3fr;
    background-color: goldenrod;
}

.msg1 {
    grid-area: msg1;
    margin: 10px;
    background-color: bisque;
}

.input1 {
    grid-area: input1;
    margin: 10px;
    background-color: aquamarine;
}

.l2 {
    grid-area: l2;
    margin: 10px;
    display: grid;
    grid-template-areas: "msg2 input2";
    grid-template-columns: 1fr 3fr;
    background-color: gold;
}

.msg2 {
    grid-area: msg2;
    margin: 10px;
    background-color: darkorange;
}

.input2 {
    grid-area: input2;
    margin: 10px;
    background-color: darkorchid;
}

.l3 {
    grid-area: l3;
    margin: 10px;
    background-color: cornsilk;
}

.l4 {
    grid-area: l4;
    margin: 10px;
    background-color: azure;
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