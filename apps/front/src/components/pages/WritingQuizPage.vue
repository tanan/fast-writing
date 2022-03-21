<template>
  <div>
    <MainHeader />
    <div class="quiz-page">
      <!-- <QuizHeader /> -->
      <QuizList :indecies="indecies" :questions="questions" :answers="answers" />
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { ContentsId } from "@/pb/models/contents_pb.js"
import QuizList from '@/components/organisms/QuizList.vue'
import MainHeader from '@/components/organisms/MainHeader.vue'

export default defineComponent({
  name: 'WritingQuizPage',
  components: {
    MainHeader,
    QuizList
  },
  data: () => ({
    contents: {},
    indecies: [],
    questions: [],
    answers: [],
  }),
  async created() {
   this.getContentsById(this.$route.params.id)
  },
  methods: {
    async getContentsById (id) {
      const client = new WritingServiceClient(`${process.env.VUE_APP_WRITING_API_HOST}`, null, null)
      let req = new ContentsId()
      req.setId(id)
        await client.getContents(req, {}, (err, resp) => {
          if (err != null) {
            throw new Error("Could not receive the data from API!")
          }
          this.contents = resp.toObject()
          this.createQuizList(resp.toObject().quizlistList)
        })
    },
    async createQuizList(quizList) {
      for(var i in quizList) {
        this.indecies.push(quizList[i].id)
        this.questions.push(quizList[i].question)
        await new Promise((resolve) => setTimeout(resolve, 7000));
        this.answers.push(quizList[i].answer)
        await new Promise((resolve) => setTimeout(resolve, 3000));
      }
    }
  }
})
</script>
