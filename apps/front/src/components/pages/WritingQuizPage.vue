<template>
  <div class="writing-quiz-page">
    <MainHeader :isLoggedIn="isLoggedIn" />
    <div class="quiz-page">
      <QuizHeader class="quiz-header" :contents="contents" :interval="getInterval()" />
      <QuizList class="quiz-list" :indecies="indecies" :questions="questions" :answers="answers" />
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import Store from '@/store/index.js'
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { ContentsId } from "@/pb/models/contents_pb.js"
import QuizHeader from '@/components/organisms/QuizHeader.vue'
import QuizList from '@/components/organisms/QuizList.vue'
import MainHeader from '@/components/organisms/MainHeader.vue'

const client = new WritingServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default defineComponent({
  name: 'WritingQuizPage',
  components: {
    MainHeader,
    QuizList,
    QuizHeader,
  },
  data: () => ({
    contents: {},
    indecies: [],
    questions: [],
    answers: [],
    interval: Store.state.interval,
  }),
  async created() {
   this.getContentsById(this.$route.params.id)
  },
  setup () {
    const isLoggedIn = Store.getters.isLoggedIn
    return {
      isLoggedIn
    }
  },
  methods: {
    getInterval() {
      let item = localStorage.getItem('interval')
      if (!item) {
          return
      }
      let interval = JSON.parse(item)
      return interval.interval
    },
    async getContentsById (id) {
      let req = new ContentsId()
      req.setId(id)
      const metadata = { 'authorization': 'Bearer ' + Store.state.userToken }
      await client.getContents(req, metadata, (err, resp) => {
        if (err) {
          throw new Error(err)
        }
        this.contents = resp.toObject()
        this.createQuizList(resp.toObject().quizlistList)
      })
    },
    async createQuizList(quizList) {
      for(var i in quizList) {
        this.indecies.push(quizList[i].id)
        this.questions.push(quizList[i].question)
        await new Promise((resolve) => setTimeout(resolve, this.getInterval()*1000));
        this.answers.push(quizList[i].answer)
        await new Promise((resolve) => setTimeout(resolve, 2000));
      }
    }
  }
})
</script>

<style lang="scss">
</style>