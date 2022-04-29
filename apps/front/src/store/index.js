import Vuex from 'vuex'

const init_interval = 6

const store = new Vuex.Store({
  state: {
    userId: "",
    userToken: "",
    interval: init_interval,
    quizList: [],
    isDisplayAnswer: true,
  },
  mutations: {
    updateUser(state, user) {
      state.userId = user.userId
      state.userToken = user.userToken
    },
    updateInterval(state, interval) {
      state.interval = interval
    },
    updateQuizList(state, quizList) {
      state.quizList = quizList
    },
    shiftQuizList(state) {
      state.quizList.shift()
    },
    updateDispalyAnswer(state, isDisplayAnswer) {
      state.isDisplayAnswer = isDisplayAnswer
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
    },
    updateQuizList(context, quizList) {
      context.commit('updateQuizList', quizList)
    },
    shiftQuizList(context) {
      context.commit('shiftQuizList')
    },
    updateDispalyAnswer(context, isDisplayAnswer) {
      context.commit('updateDispalyAnswer', isDisplayAnswer)
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
    },
    getNextQuiz(state) {
      return state.quizList[0]
    },
    hasQuizElement(state) {
      return state.quizList.length > 0 ? true : false
    }
  },
  modules: {},
})

export default store