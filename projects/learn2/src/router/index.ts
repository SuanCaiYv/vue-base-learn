import {createRouter, createWebHistory, Router, RouteRecordRaw} from "vue-router";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/home",
        name: "home",
        component: () => import("../views/Home.vue")
    },
    {
        path: "/about",
        name: "about",
        component: () => import("../views/About.vue")
    },
    {
        path: "/dynamic-route/:param1",
        name: "dynamic-route",
        component: () => import("../views/DynamicRoute.vue")
    }
]

const router: Router = createRouter({
    history: createWebHistory(),
    routes: routes
})

export default router