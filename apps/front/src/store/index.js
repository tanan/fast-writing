import Vuex from 'vuex'

const store = new Vuex.Store({
    state: {
        userId: "",
        userToken: "",
        interval: 6
    },
    mutations: {
        updateUser(state, user) {
            state.userId = user.userId
            state.userToken = user.userToken
        },
        updateInterval(state, interval) {
            state.interval = interval
        }
    },
    actions: {
        auth(context, user) {
            context.commit('updateUser', user)
            let value = JSON.stringify(user)
            localStorage.setItem('user', value)
        },
        signout(context) {
            context.commit('updateUser', {
                userId: "",
                userToken: "",
            })
            localStorage.removeItem('user')
        },
        loadState(context) {
            let item = localStorage.getItem('user')
            if (!item) {
                return
            }
            let user = JSON.parse(item)
            context.commit('updateUser', user)
        },
        interval(context, interval) {
            context.commit('updateInterval', interval)
            let value = JSON.stringify(interval)
            localStorage.setItem('interval', value)
        },
        loadInterval(context) {
            let item = localStorage.getItem('interval')
            if (!item) {
                return
            }
            let interval = JSON.parse(item)
            context.commit('updateInterval', interval)
        }
    },
    getters: {
        interval (state) {
            return state.interval
        }
    },
    modules: {},
})

export default store