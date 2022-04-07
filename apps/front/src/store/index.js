import Vuex from 'vuex'

const store = new Vuex.Store({
    state: {
        userId: "",
        userToken: ""
    },
    mutations: {
        updateUser(state, user) {
            state.userId = user.userId;
            state.userToken = user.userToken;
        }
    },
    actions: {
        auth(context, user) {
            context.commit('updateUser', user);
            let value = JSON.stringify(user);
            localStorage.setItem('user', value);
        },
        loadState(context) {
            let item = localStorage.getItem('user')
            if (!item) {
                return
            }
            let user = JSON.parse(item)
            context.commit('updateUser', user)
        }
    },
    modules: {},
})

export default store