<template>
    <div>
        <el-row>
            <el-col :span="24"><div class="grid-content bg-purple-dark">aaa</div></el-col>
        </el-row>
        <el-row>
            <el-col :span="12"><div class="grid-content bg-purple"></div>ccc</el-col>
            <el-col :span="12"><div class="grid-content bg-purple-light"></div></el-col>
        </el-row>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Test")
</script>

<style scoped>

</style>