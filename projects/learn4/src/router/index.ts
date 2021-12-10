import {createRouter, createWebHistory, Router, RouteRecordRaw} from "vue-router";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/home",
        alias: "/",
        name: "home",
        component: () => import("../views/Home.vue")
    },
    {
        path: "/store1",
        name: "store1",
        component: () => import("../views/CpOne.vue")
    },
    {
        path: "/store2",
        name: "store2",
        component: () => import("../views/CpTwo.vue")
    }
]

const router: Router = createRouter({
    history: createWebHistory(),
    routes: routes
})

export default router