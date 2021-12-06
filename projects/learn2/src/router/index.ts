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
        // 动态参数配置
        path: "/dynamic-route/:id/:age",
        name: "dynamic-route",
        component: () => import("../views/Dynamic/DynamicRoute.vue")
    },
    // 这是一个嵌套路由的例子，访问路径可以通过/l1/l2/l3进行访问到各个组件
    {
        path: "/l1",
        name: "l1",
        component: () => import("../views/MutilateLevelRoute/LevelOne.vue"),
        children: [
            {
                path: "l2",
                name: "l2",
                component: () => import("../views/MutilateLevelRoute/LevelTwo.vue"),
                children: [
                    {
                        path: "l3",
                        name: "l3",
                        component: () => import("../views/MutilateLevelRoute/LevelThree.vue")
                    }
                ]
            }
        ]
    },
    // 带有模板组件的例子
    {
        path: "/layout",
        name: "layout",
        component: () => import("../views/WithLayout/SimpleLayout.vue"),
        children: [
            {
                path: "msg/:msg",
                component: () => import("../views/MsgShow.vue")
            }
        ]
    }
]

const router: Router = createRouter({
    history: createWebHistory(),
    routes: routes
})

export default router