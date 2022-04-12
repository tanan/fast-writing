<template>
  <div>
    <div class="grid">
      <div class="mt-6 col-offset-1 col-10 md:col-10 lg:col-6 lg:col-offset-3">
        <h2 class="pl-2 pb-2">レッスン名</h2>
        <span class="col-12">
          <InputText class="col-12 p-inputtext-lg" type="text" v-model="title" @blur="save" placeholder="Title" />
        </span>
        <h2 class="pl-2 pt-4">説明</h2>
        <span class="col-12">
          <InputText class="col-12 p-inputtext-lg" type="text" v-model="description" @blur="save" placeholder="Description" />
        </span>
      </div>
      <div class="col-offset-1 col-10 md:col-10 lg:col-6 lg:col-offset-3 p-0">
        <h2 class="pl-2 pt-4">クイズ</h2>
        <div class="quiz mt-4" v-for="(quiz, i) in quizzes" :key="i">
          <div class="flex question">
            <h3 class="col-1 p-2 mr-4 lg:mr-0">Q{{ i+1 }}.</h3>
            <InputText
              class="col-10 lg:col-11 align-items-center p-inputtext-lg"
              type="text"
              v-model="quizzes[i].question"
              placeholder="Question"
            />
          </div>
          <div class="flex quiz-answer mt-2 p-0">
            <h3 class="col-1 p-2 mr-4 lg:mr-0">A{{ i+1 }}.</h3>
            <InputText
              class="col-10 lg:col-11 align-items-center p-inputtext-lg"
              type="text"
              v-model="quizzes[i].answer"
              @blur="saveQuiz"
              placeholder="Answer"
            />
          </div>
        </div>
        <div class="col-offset-2 lg:col-offset-1">
          <Button class="mt-2 lg:mt-4" label="追加" icon="pi pi-plus" @click="addQuiz" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { Contents, ContentsId, Quiz, QuizId } from "@/pb/models/contents_pb.js"
import { CreateContentsRequest, CreateQuizRequest } from "@/pb/fast-writing_pb.js"
import { UserId } from "@/pb/models/user_pb.js"
import Store from '@/store/index.js'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'

const client = new WritingServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default {
  name: 'ContentsCreator',
  components: {
    InputText,
    Button,
  },
  data: () => ({
    title: "",
    description: "",
    contentsId: null,
    scope: "public",
    quizzes: [
      {"id": null, "question": "question1", "answer": "answer1"}
    ]
  }),
  methods: {
    addQuiz () {
      let quiz = {"question": "", "answer": ""}
      this.quizzes.push(quiz)
    },
    async save () {
      let req = this.createContentsRequest()
      const token = Store.state.userToken
      const metadata = { 'authorization': 'Bearer ' + token }
      let response = await new Promise((resolve, reject) => {
        client.createUserContents( req, metadata, ( err, resp ) => {
          if ( err ) {
            reject(err)
          }
          console.log("contents was created")
          resolve(resp)
        })
      })
      this.contentsId = await response.toObject().contents.id.id
      this.setQuizList(await response.toObject().contents.quizlistList)
    },
    async saveQuiz () {
      const token = Store.state.userToken
      const metadata = { 'authorization': 'Bearer ' + token }
      for (const [index, quiz] of this.quizzes.entries()) {
        let quizReq = await this.createQuizRequest(quiz, this.contentsId)
        await new Promise((resolve, reject) => {
          client.createUserQuiz(quizReq, metadata, (err, resp) => {
            if (err != null) {
              reject(err)
            }
            this.quizzes[index].id = resp.toObject().quizid.id
            console.log("quiz was created")
            resolve(resp)
          })
        })
      }
    },
    setQuizList (quizList) {
      let quizzes = []
      for (const q of quizList) {
        let quiz = {
          id: q.id.id,
          answer: q.answer,
          question: q.question,
          order: q.order
        }
        quizzes.push(quiz)
      }
      this.quizzes = quizzes
    },
    createContentsRequest () {
      let req = new CreateContentsRequest()
      let contentsId = new ContentsId()
      let contents = new Contents()
      let userId = new UserId()
      for (const quiz of this.quizzes) {
        contents.addQuizlist(this.createQuiz(quiz))
      }
      userId.setId(Store.state.userId)
      contents.setTitle(this.title)
      contents.setDescription(this.description)
      if (this.contentsId) {
        contentsId.setId(this.contentsId)
      }
      contents.setId(contentsId)
      contents.setScope(this.scope)
      req.setContents(contents)
      req.setUserid(userId)
      return req
    },
    createQuiz (q) {
      let quiz = new Quiz()
      let quizId = new QuizId()
      quizId.setId(q.id)
      quiz.setId(quizId)
      quiz.setQuestion(q.question)
      quiz.setAnswer(q.answer)
      return quiz
    },
    createQuizRequest (q, cid) {
      let req = new CreateQuizRequest()
      let contentsId = new ContentsId()
      contentsId.setId(cid)
      let quiz = new Quiz()
      let quizId = new QuizId()
      quizId.setId(q.id)
      quiz.setId(quizId)
      quiz.setQuestion(q.question)
      quiz.setAnswer(q.answer)
      req.setContentsid(contentsId)
      req.setQuiz(quiz)
      return req
    }
  }
}
</script>

<style lang="scss" scoped>
  .contents-creator-form {
    margin: 0 8%;
    margin-top: 2%;

    .quiz {
      .index {
        font-size: 1.2em;
        vertical-align: bottom;
        min-width: 4%;
        margin-right: 2%;
      }

      .question {
        display: flex;
        align-items: center;
      }

      .quiz-answer {
        display: flex;
        align-items: center;
      }
    }
  }
</style>
