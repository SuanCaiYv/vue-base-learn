import {createRouter, createWebHistory, RouteRecordRaw} from "vue-router";

const routes: Array<RouteRecordRaw> = [
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
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes: routes
})

export default router