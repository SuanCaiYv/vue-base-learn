<template>
    <div class="sign">
        <div class="s-o-one"></div>
        <div class="main">
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
    width: 200px;
    height: 300px;
    border: 1px solid black;
    background-color: cadetblue;
}
</style>