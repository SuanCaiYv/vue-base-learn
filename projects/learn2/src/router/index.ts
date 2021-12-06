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
    },
    // 使用router进行跳转和回退，以及在跳转时包含参数信息
    {
        path: "/jump",
        name: "jump",
        component: () => import("../views/Home.vue"),
        children: [
            {
                path: "view-one",
                name: "view-one",
                component: () => import("../views/Jump/ViewOne.vue")
            },
            {
                path: "view-two",
                name: "view-two",
                component: () => import("../views/Jump/ViewTwo.vue")
            },
            {
                path: "view-three",
                name: "view-three",
                component: () => import("../views/Jump/ViewThree.vue")
            }
        ]
    },
    {
        path: "/meta",
        name: "meta",
        component: () => import("../views/Home.vue"),
        children: [
            {
                path: "cp-one",
                name: "cp-one",
                component: () => import("../views/CustomMeta/CpOne.vue"),
                // 绑定到这个路由组件，可以在跳转完成后，组件内访问这些值，所以我们可以在此设置标题，是否允许游客访问等属性
                meta: {
                    title: "组件一",
                    requireAuthed: false,
                }
            },
            {
                path: "cp-two",
                name: "cp-two",
                component: () => import("../views/CustomMeta/CpTwo.vue"),
                meta: {
                    title: "组件二",
                    requireAuthed: true
                }
            }
        ]
    },
    {
        path: "/",
        // 配置重定向，可携带参数，也可以设置跳转逻辑，即redirect为一个函数，根据逻辑返回一个路径字符串，作为重定向目的地
        redirect: "/home",
        // 组件独享的钩子函数
        beforeEnter: (to, from) => {
            document.title = "Home"
        }
    },
    {
        path: "/profile",
        // 配置别名，访问/me效果和访问/profile是一样的，关于它和重定向的区别，在于它不会修改显示的URL，而重定向会把URL改成实际访问的路径。
        alias: "/me",
        component: () => import("../views/About.vue")
    }
]

const router: Router = createRouter({
    history: createWebHistory(),
    routes: routes
})

// 路由跳转前触发
router.beforeEach((to, from) => {
    console.log("来自" + String(from.name))
    console.log("去往" + String(to.name))
    if (to.meta.requireAuthed !== undefined) {
        const requireAuthed = Boolean(to.meta.requireAuthed)
        if (requireAuthed) {
            console.log("此页面需要认证")
        } else {
            console.log("此页面允许游客访问")
        }
    }
})

// 组件解析之后，正式触发导航跳转之前触发，一个场景是申请摄像头和麦克风权限
router.beforeResolve((to, from) => {
    // 不太常用
})

// 组件解析完成，但是又是每个组件都需要的操作，比如点击次数统计，流量分析
router.afterEach((to, from) => {
    console.log("已跳转至" + String(to.name))
})

export default router