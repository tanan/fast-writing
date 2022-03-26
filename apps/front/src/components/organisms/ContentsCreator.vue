<template>
  <v-form class="contents-creator-form" ref="form">
    <v-container>
      <v-row>
        <v-col>
          <v-text-field dense clearable class="title my-0 py-0" v-model="title" label="title" required></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <span>Quiz</span>
          <div class="quiz" v-for="(quiz, i) in quizzes" :key="i">
            <div class="mb-4 question">
              <span class="index">{{ i+1 }}.</span>
              <v-text-field dense hide-details v-model="quizzes[i].question" label="question"></v-text-field>
            </div>
            <div class="quiz-answer">
              <span class="index"></span>
              <v-text-field dense hide-details class="answer" v-model="quizzes[i].answer" label="answer"></v-text-field>
            </div>
          </div>
        </v-col>
      </v-row>
      <v-btn @click="addQuiz" class="mx-2 append-btn" color="indigo">
        <v-icon dark>mdi-plus</v-icon>
      </v-btn>
      <v-btn @click="save" color="success">SAVE</v-btn>
    </v-container>
  </v-form>
</template>

<script>
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { Contents, ContentsId, Quiz } from "@/pb/models/contents_pb.js"
import { CreateContentsRequest, CreateQuizRequest } from "@/pb/fast-writing_pb.js"
import { UserId } from "@/pb/models/user_pb.js"

export default {
  name: 'ContentsCreator',
  props: {
    indecies: Array,
    questions: Array,
    answers: Array,
  },
  data: () => ({
    title: "",
    quizzes: [
      {"question": "question1", "answer": "answer1"}
    ]
  }),
  methods: {
    addQuiz () {
      let quiz = {"question": "", "answer": ""}
      this.quizzes.push(quiz)
    },
    async save () {
      const client = new WritingServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)
      let req = new CreateContentsRequest()
      let contents = new Contents()
      let contentsId = new ContentsId()
      let userId = new UserId()
      userId.setId("11eae085-55e6-e2ca-a15d-0242ac110002")
      contents.setTitle(this.title)
      req.setContents(contents)
      req.setUserid(userId)
      await client.createUserContents(req, {}, (err, resp) => {
        if (err != null) {
          throw new Error("Could not receive the data from API!")
        }
        console.log(resp.toObject())
        contentsId.setId(resp.toObject().contentsId)
      })
      let req2 = CreateQuizRequest()
      let quiz = new Quiz()
      quiz.setQuestion(this.quizzes[0].question)
      quiz.setAnswer(this.quizzes[0].answer)
      await client.createUserQuiz(req2, {}, (err, resp) => {
        if (err != null) {
          throw new Error("Could not receive the data from API!")
        }
        console.log(resp.toObject())
      })
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
