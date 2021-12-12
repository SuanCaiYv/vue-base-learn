import {createStore, Store} from "vuex";

const store: Store<any> = createStore({
    state() {
        return {
            counter: Number
        }
    },
    mutations: {},
    getters: {}
})

export default store