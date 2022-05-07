<template>
  <div class="writing-quiz-page">
    <div class="quiz-page">
      <MainHeader />
      <QuizHeader class="quiz-header" :contents="contents" />
      <QuizList class="quiz-list" />
    </div>
  </div>
</template>

<script>
import { defineComponent, reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Store from '@/store/index.js'
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { ContentsId } from "@/pb/models/contents_pb.js"
import MainHeader from '@/components/organisms/MainHeader.vue'
import QuizHeader from '@/components/organisms/QuizHeader.vue'
import QuizList from '@/components/organisms/QuizList.vue'

const client = new WritingServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default defineComponent({
  name: 'WritingQuizPage',
  components: {
    MainHeader,
    QuizList,
    QuizHeader,
  },
  setup () {
    const contents = reactive({})

    const route = useRoute()
    const router = useRouter()

    const createQuizList = (quizList) => {
      let list = []
      for(var i in quizList) {
        let q = {
          index: i,
          question: quizList[i].question,
          answer: quizList[i].answer,
        }
        list.push(q)
      }
      Store.dispatch('updateQuizList', list)
    }

    let req = new ContentsId()
    req.setId(route.params.id)
    const metadata = { 'authorization': 'Bearer ' + Store.state.userToken }
    client.getContents(req, metadata, (err, resp) => {
      if (err) {
        if (err.message === "Unauthenticated") {
          router.push(`/signin?redirect=${route.fullPath}`)
          return
        }
        throw new Error(err)
      }
      const obj = resp.toObject()
      contents.title = obj.title
      contents.description = obj.description
      contents.creator = obj.creator
      contents.lastUpdated = obj.lastUpdated
      contents.tags = obj.tags
      createQuizList(resp.toObject().quizlistList)
    })

    return {
      contents,
    }
  }
})
</script>
