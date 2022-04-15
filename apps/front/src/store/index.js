import Vuex from 'vuex'

const init_interval = 6

const store = new Vuex.Store({
  state: {
    userId: "",
    userToken: "",
    interval: init_interval
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
        let value = JSON.stringify({interval: init_interval})
        localStorage.setItem('interval', value)
        return
      }
      let interval = JSON.parse(item)
      context.commit('updateInterval', interval)
    }
  },
  getters: {
    interval (state) {
      return state.interval
    },
    isLoggedIn (state) {
      if (state.userToken === "") {
        return false
      }
        return true
    }
  },
  modules: {},
})

export default store