<template>
  <div>
    <div class="grid">
      <div class="mt-6 col-offset-1 col-10 md:col-10 lg:col-6 lg:col-offset-3">
        <h2 class="pl-2 pb-2">レッスン名</h2>
        <span class="col-12">
          <InputText class="col-12 p-inputtext-lg" type="text" v-model="contents.title" @blur="save" placeholder="Title" />
        </span>
        <h2 class="pl-2 pt-4">説明</h2>
        <span class="col-12">
          <InputText class="col-12 p-inputtext-lg" type="text" v-model="contents.description" @blur="save" placeholder="Description" />
        </span>
      </div>
      <div class="col-offset-1 col-10 md:col-10 lg:col-6 lg:col-offset-3 p-0">
        <h2 class="pl-2 pt-4">クイズ</h2>
        <div class="quiz mt-4" v-for="(quiz, i) in contents.quizzes" :key="i">
          <div class="flex question">
            <h3 class="col-1 p-2 mr-4 lg:mr-0">Q{{ i+1 }}.</h3>
            <InputText
              class="col-10 lg:col-11 align-items-center p-inputtext-lg"
              type="text"
              v-model="contents.quizzes[i].question"
              placeholder="Question"
            />
          </div>
          <div class="flex quiz-answer mt-2 p-0">
            <h3 class="col-1 p-2 mr-4 lg:mr-0">A{{ i+1 }}.</h3>
            <InputText
              class="col-10 lg:col-11 align-items-center p-inputtext-lg"
              type="text"
              v-model="contents.quizzes[i].answer"
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
import { defineComponent, reactive } from "vue"
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { Contents, ContentsId, Quiz, QuizId } from "@/pb/models/contents_pb.js"
import { CreateContentsRequest, CreateQuizRequest } from "@/pb/fast-writing_pb.js"
import { UserId } from "@/pb/models/user_pb.js"
import Store from '@/store/index.js'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'

const client = new WritingServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)
const token = Store.state.userToken
const metadata = { 'authorization': 'Bearer ' + token }

export default defineComponent ({
  name: 'ContentsCreator',
  components: {
    InputText,
    Button,
  },
  props: {
    id: Number
  },
  setup (props) {
    const contents = reactive({
      id: props.id,
      title: '',
      description: '',
      creator: '',
      scope: 'public',
      quizzes: [{"id": null, "question": "question1", "answer": "answer1"}]
    })

    if (contents.id) {
      let req = new ContentsId()
      req.setId(props.id)
      client.getContents(req, metadata, (err, resp) => {
        if (err != null) {
          throw new Error("Could not receive the data from API!")
        }
        let c = resp.toObject()
        contents.title = c.title
        contents.description = c.description
        contents.creator = c.creator
        // contents.scope = c.scope
        let list = []
        for(var v of c.quizlistList) {
          list.push({
            id: v.id.id,
            question: v.question,
            answer: v.answer,
            order: v.order,
          })
        }
        contents.quizzes = list
      })
    }

    const addQuiz = () => {
      let quiz = {"question": "", "answer": ""}
      contents.quizzes.push(quiz)
    }

    const save = async () => {
      let req = createContentsRequest(contents)
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
      contents.id = await response.toObject().contents.id.id
      setQuizList(contents, await response.toObject().contents.quizlistList)
    }

    const saveQuiz = async () => {
      const token = Store.state.userToken
      const metadata = { 'authorization': 'Bearer ' + token }
      for (const [index, quiz] of contents.quizzes.entries()) {
        let quizReq = await createQuizRequest(quiz, contents.id)
        await new Promise((resolve, reject) => {
          client.createUserQuiz(quizReq, metadata, (err, resp) => {
            if (err != null) {
              reject(err)
            }
            contents.quizzes[index].id = resp.toObject().quizid.id
            console.log("quiz was created")
            resolve(resp)
          })
        })
      }
    }

    const setQuizList = (contents, quizList) => {
      let list = []
      for (const q of quizList) {
        let quiz = {
          id: q.id.id,
          answer: q.answer,
          question: q.question,
          order: q.order
        }
        list.push(quiz)
      }
      contents.quizzes = list
    }
    const createContentsRequest = (contents) => {
      console.log(contents)
      let req = new CreateContentsRequest()
      let cid = new ContentsId()
      let c = new Contents()
      let userId = new UserId()
      for (const quiz of contents.quizzes) {
        c.addQuizlist(createQuiz(quiz))
      }
      userId.setId(Store.state.userId)
      if (contents.id) {
        cid.setId(contents.id)
      }
      c.setId(cid)
      c.setTitle(contents.title)
      c.setDescription(contents.description)
      c.setScope(contents.scope)
      req.setContents(c)
      req.setUserid(userId)
      return req
    }
    const createQuiz = (q) => {
      let quiz = new Quiz()
      let quizId = new QuizId()
      quizId.setId(q.id)
      quiz.setId(quizId)
      quiz.setQuestion(q.question)
      quiz.setAnswer(q.answer)
      return quiz
    }
    const createQuizRequest = (q, cid) => {
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
    
    return {
      contents,
      save,
      saveQuiz,
      addQuiz,
      createQuizRequest,
      createQuiz,
      createContentsRequest,
    }
  }
})
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
