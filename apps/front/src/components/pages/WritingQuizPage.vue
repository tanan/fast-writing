<template>
  <div class="writing-quiz-page">
    <div class="quiz-page">
      <QuizHeader class="quiz-header" :contents="contents" :interval="getInterval()" />
      <QuizList class="quiz-list" :indecies="contents.indecies" :questions="contents.questions" :answers="contents.answers" />
    </div>
  </div>
</template>

<script>
import { defineComponent, reactive } from 'vue'
import { useRoute } from 'vue-router'
import Store from '@/store/index.js'
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { ContentsId } from "@/pb/models/contents_pb.js"
import QuizHeader from '@/components/organisms/QuizHeader.vue'
import QuizList from '@/components/organisms/QuizList.vue'

const client = new WritingServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default defineComponent({
  name: 'WritingQuizPage',
  components: {
    QuizList,
    QuizHeader,
  },
  setup () {
    const contents = reactive({
      title: '',
      description: '',
      creator: '',
      lastUpdated: '',
      indecies: [],
      questions: [],
      answers: []
    })

    const route = useRoute()

    const getInterval = () => {
      let item = localStorage.getItem('interval')
      if (!item) {
          return
      }
      let interval = JSON.parse(item)
      return interval.interval
    }

    const createQuizList = async (quizList) => {
      for(var i in quizList) {
        contents.indecies.push(i)
        contents.questions.push(quizList[i].question)
        await new Promise((resolve) => setTimeout(resolve, getInterval()*1000));
        contents.answers.push(quizList[i].answer)
        await new Promise((resolve) => setTimeout(resolve, 2000));
      }
    }

    let req = new ContentsId()
    req.setId(route.params.id)
    const metadata = { 'authorization': 'Bearer ' + Store.state.userToken }
    client.getContents(req, metadata, (err, resp) => {
      if (err) {
        throw new Error(err)
      }
      const obj = resp.toObject()
      contents.title = obj.title
      contents.description = obj.description
      contents.creator = obj.creator
      contents.lastUpdated = obj.lastUpdated
      createQuizList(resp.toObject().quizlistList)
    })

    return {
      contents,
      getInterval,
    }
  }
})
</script>
