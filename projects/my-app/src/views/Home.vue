<template>
    <header>
        <el-affix :offset="0" class="max-width">
            <div class="top max-width">
                <el-row class="max-width">
                    <el-col :span="1" :offset="1">
                        <div>aaa</div>
                    </el-col>
                    <el-col :span="2" :offset="16">
                        <div>bbb</div>
                    </el-col>
                </el-row>
            </div>
        </el-affix>
    </header>
    <main>
        <div class="body">
            aaa
        </div>
    </main>
    <footer>
        <div>
            ccc
        </div>
    </footer>
</template>

<script lang="ts" setup>
import {ref} from "vue"
import {useRoute} from "vue-router";

const name = ref<String>("Home")

const route = useRoute()
document.title = String(route.meta.title)
</script>

<style scoped>

.header {
    background-color: pink;
    position: relative;
    width: 100%;
    height: 30px;
    display: flex;
    justify-content: center;
    align-items: center;
}

.header-font {
    font-size: 1.5rem;
}

.main {
    background-color: lightskyblue;
    position: relative;
    min-height: 99%;
    width: 100%;
}

.footer {
    position: relative;
    top: 0;
    width: 100%;
    min-height: 70px;
    background-color: #42b983;
    display: flex;
    justify-content: center;
    align-items: center;
}

.top {
    background-color: antiquewhite;
    height: 60px;
    display: flex;
    justify-content: center;
    align-items: center;
}

.body {
    background-color: lightskyblue;
}

.max-width {
    width: 100%;
}
</style>