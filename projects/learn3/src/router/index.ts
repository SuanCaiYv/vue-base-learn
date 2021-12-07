import {createRouter, createWebHistory, RouteRecordRaw} from "vue-router";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/home",
        name: "home",
        component: () => import("../views/Home.vue")
    },
    {
        path: "/",
        redirect: "/home",
        component: () => import("../views/Home.vue")
    },
    // 父子组件通信的几种方式
    {
        path: "/f-s",
        name: "f-s",
        component: () => import("../views/Home.vue"),
        children: [
            // props-emits形式
            {
                path: "p-e",
                name: "p-e",
                component: () => import("../views/FatherSon/WithPropsEmitsF.vue")
            },
            // v-model-emits形式
            {
                path: "v-e",
                name: "v-e",
                component: () => import("../views/FatherSon/WithVModelEmitsF.vue")
            },
            // ref-emits形式，不好用
            {
                path: "r-e",
                name: "r-e",
                component: () => import("../views/FatherSon/WithRefEmitsF.vue")
            }
        ]
    },
    {
        path: "/g-f-s",
        name: "g-f-s",
        component: () => import("../views/Home.vue"),
        children: []
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes: routes
})

export default router