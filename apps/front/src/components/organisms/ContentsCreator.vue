<template>
  <v-form class="contents-creator-form" ref="form">
    <v-container>
      <v-row>
        <v-col>
          <h2>Title</h2>
          <v-text-field dense clearable class="title my-0 py-0" v-model="title" @blur="save" @keydown.enter="save" label="title" required></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <h2>Description</h2>
          <v-text-field dense clearable class="title my-0 py-0" v-model="description" @blur="save" @keydown.enter="save" label="description" required></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <h2>Quiz</h2>
          <div class="quiz" v-for="(quiz, i) in quizzes" :key="i">
            <div class="mb-4 question">
              <span class="index">{{ i+1 }}.</span>
              <v-text-field dense hide-details v-model="quizzes[i].question" @blur="saveQuiz" @keydown.enter="saveQuiz" label="question"></v-text-field>
            </div>
            <div class="quiz-answer">
              <span class="index"></span>
              <v-text-field dense hide-details class="answer" v-model="quizzes[i].answer" @blur="saveQuiz" @keydown.enter="saveQuiz" label="answer"></v-text-field>
            </div>
          </div>
        </v-col>
      </v-row>
      <v-btn @click="addQuiz" class="mx-2 append-btn" color="indigo">
        追加
      </v-btn>
    </v-container>
  </v-form>
</template>

<script>
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { Contents, ContentsId, Quiz, QuizId } from "@/pb/models/contents_pb.js"
import { CreateContentsRequest, CreateQuizRequest } from "@/pb/fast-writing_pb.js"
import { UserId } from "@/pb/models/user_pb.js"
import Store from '@/store/index.js'

const client = new WritingServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default {
  name: 'ContentsCreator',
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
