import {createStore, Store} from "vuex";

const store: Store<any> = createStore({
    // state创建的变量也是响应式的
    state() {
        return {
            counter: 0,
            msg: "",
            list: []
        }
    },
    // 更改事件，state的更改不能手动设置，只能通过触发更改事件来完成
    mutations: {
        inc(state) {
            state.counter++
        },
        incByNum(state, num) {
            state.counter += num
        },
        pushItem(state, item) {
            state.list.push(item)
        },
        // 通过有效载荷更新法
        updateMsg(state, payload) {
            console.log("run1")
            state.msg = payload.value
        }
    },
    // 可以把getters理解成计算属性，也会缓存属性值，适用于需要全局使用的方法
    getters: {
        listLen: (state) => {
            return state.list.length
        },
        // 返回方法的计算属性，不会缓存值，每次调用返回的方法，都会触发新的计算。
        findIdx: (state) => (idx: any) => {
            return state.list[idx]
        }
    }
})

export default store