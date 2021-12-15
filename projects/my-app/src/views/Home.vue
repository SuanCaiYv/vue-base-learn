<template>
    <div class="home">
        <div class="header">
            <div class="h-offset-one"></div>
            <div class="logo">logo</div>
            <div class="home-page">home</div>
            <div class="h-offset-two"></div>
            <div class="search">search</div>
            <div class="h-offset-three"></div>
        </div>
        <div class="main">
            <div class="user">
                <div class="u-o-one"></div>
                <div class="info">
                    <div class="avatar">
                        <img src="public/favicon.ico">
                    </div>
                    <div class="nickname">
                        小白白白
                    </div>
                    <div class="email">
                        2508826394@qq.com
                    </div>
                    <div class="brief">
                        小白，拜拜
                    </div>
                </div>
            </div>
            <div class="articles">
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
                <br>
            </div>
            <div class="topics">
                <div class="t-o-one"></div>
                <div class="list">list</div>
            </div>
        </div>
        <div class="footer">ccc</div>
    </div>
</template>

<script lang="ts" setup>
import {ref} from "vue"
import {useRoute} from "vue-router";

const name = ref<String>("Home")

const route = useRoute()
document.title = String(route.meta.title)
</script>

<style scoped>
.home {
    height: 100%;
    display: grid;
    grid-template-areas:
        "header"
        "main"
        "footer";
    grid-template-columns: 1fr;
    grid-template-rows: 30px 1fr 30px;
}

.header {
    grid-area: header;
    display: grid;
    grid-template-areas:
        "h-offset-one logo home-page h-offset-two search h-offset-three";
    grid-template-columns: 60px 60px 60px 1fr 240px 120px;
    grid-column-gap: 10px;
    background-color: lightskyblue;
}

.h-offset-one {
    grid-area: h-offset-one;
}

.logo {
    grid-area: logo;
    border: 0 solid black;
    border-radius: 2px;
}

.home-page {
    grid-area: home-page;
    border: 0 solid black;
    border-radius: 2px;
}

.h-offset-two {
    grid-area: h-offset-two;
}

.search {
    grid-area: search;
}

.h-offset-three {
    grid-area: h-offset-three;
}

.main {
    grid-area: main;
    display: grid;
    grid-template-areas:
        "user articles topics";
    grid-template-columns: 200px 1fr 200px;
    background-color: #42b983;
}

.user {
    grid-area: user;
    display: grid;
    grid-template-areas:
        "u-o-one"
        "info"
        "u-o-two";
    grid-template-rows: 120px 1fr;
    background-color: antiquewhite;
}

.u-o-one {
    grid-area: u-o-one;
    background-color: blueviolet;
}

.info {
    grid-area: info;
    display: grid;
    grid-template-areas:
        "avatar"
        "nickname"
        "email"
        "brief";
    grid-template-rows: 200px 40px 40px 1fr;
    background-color: gold;
}

.avatar {
    grid-area: avatar;
    margin: 20px;
    height: 160px;
    width: 160px;
}

img {
    width: 100%;
    height: 100%;
    object-fit: contain;
}

.nickname {
    grid-area: nickname;
    width: 180px;
    margin: 10px;
    font-size: 0.8rem;
    background-color: cadetblue;
}

.email {
    grid-area: email;
    width: 180px;
    margin: 10px;
    font-size: 0.8rem;
    background-color: dodgerblue;
}

.brief {
    grid-area: brief;
    width: 180px;
    margin: 10px;
    font-size: 0.8rem;
    background-color: cornflowerblue;
}

.articles {
    grid-area: articles;
    background-color: azure;
    margin-top: 120px;
    margin-left: 30px;
    margin-right: 30px;
}

.topics {
    grid-area: topics;
    display: grid;
    grid-template-areas:
        "t-o-one"
        "list";
    grid-template-rows: 120px 1fr;
    background-color: beige;
}

.t-o-one {
    background-color: cornsilk;
}

.list {
    background-color: goldenrod;
}

.footer {
    grid-area: footer;
}
</style>