<template>
  <div>
    <div class="grid">
      <div class="mt-6 col-offset-1 col-10 md:col-10 lg:col-6 lg:col-offset-3">
        <h2 class="pl-2 pb-2">レッスン名</h2>
        <span class="col-12">
          <InputText class="col-12 p-inputtext-lg" type="text" v-model="contents.title" v-on:keyup="save()" placeholder="Title" />
        </span>
        <h2 class="pl-2 pt-4">説明</h2>
        <span class="col-12">
          <InputText class="col-12 p-inputtext-lg" type="text" v-model="contents.description" v-on:keyup="save()" placeholder="Description" />
        </span>
        <h2 class="pl-2 pt-4">公開設定</h2>
        <span class="col-2">
          <SelectButton v-model="contents.scope" :options="scopes" @click="save()" />
        </span>
      </div>
      <div class="col-offset-1 col-10 md:col-10 lg:col-6 lg:col-offset-3 p-0">
        <h2 class="pl-2 pt-4">クイズ</h2>
        <div class="quiz mt-4" v-for="(quiz, i) in contents.quizzes" :key="i">
          <div class="flex question">
            <h3 class="col-1 p-2 mr-4 lg:mr-0">Q{{ i+1 }}.</h3>
            <InputText
              class="col-10 lg:col-11 align-items-center p-inputtext"
              type="text"
              v-model="contents.quizzes[i].question"
              v-on:keyup="save"
              placeholder="Question"
            />
            <i class="icon pi pi-trash" @click="removeQuiz(i)"></i>
          </div>
          <div class="flex quiz-answer mt-2 p-0">
            <h3 class="col-1 p-2 mr-4 lg:mr-0">A{{ i+1 }}.</h3>
            <InputText
              class="col-10 lg:col-11 align-items-center p-inputtext"
              type="text"
              v-model="contents.quizzes[i].answer"
              v-on:keyup="save"
              placeholder="Answer"
            />
          </div>
        </div>
        <div class="pr-2">
          <Button class="mt-2 mb-4 lg:mt-4 p-button-rounded" icon="pi pi-plus" @click="addQuiz()" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, reactive } from "vue"
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { Contents, ContentsId, Quiz } from "@/pb/models/contents_pb.js"
import { CreateContentsRequest } from "@/pb/fast-writing_pb.js"
import { UserId } from "@/pb/models/user_pb.js"
import Store from '@/store/index.js'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import SelectButton from 'primevue/selectbutton'
import { debounce } from 'lodash'

const client = new WritingServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default defineComponent ({
  name: 'ContentsCreator',
  components: {
    InputText,
    Button,
    SelectButton,
  },
  props: {
    id: Number
  },
  setup (props) {
    const scopes = ['private', 'public']

    const contents = reactive({
      id: props.id,
      title: '',
      description: '',
      creator: '',
      scope: 'private',
      quizzes: [{"question": "", "answer": ""}]
    })

    if (contents.id) {
      let req = new ContentsId()
      req.setId(props.id)
      const metadata = { 'authorization': 'Bearer ' + Store.state.userToken }
      client.getContents(req, metadata, (err, resp) => {
        if (err != null) {
          throw new Error("Could not receive the data from API!")
        }
        let c = resp.toObject()
        contents.title = c.title
        contents.description = c.description
        contents.creator = c.creator
        let list = []
        for(var v of c.quizlistList) {
          list.push({
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

    const save = debounce(async () => {
      let req = createContentsRequest(contents)
      console.log(req)
      const metadata = { 'authorization': 'Bearer ' + Store.state.userToken }
      let response = await new Promise((resolve, reject) => {
        client.createUserContents( req, metadata, ( err, resp ) => {
          if ( err ) {
            reject(err)
          }
          resolve(resp)
        })
      })
      contents.id = await response.toObject().contents.id.id
    }, 2000)

    const createContentsRequest = (contents) => {
      let req = new CreateContentsRequest()
      let cid = new ContentsId()
      let c = new Contents()
      let userId = new UserId()
      for (const [index, quiz] of contents.quizzes.entries()) {
        c.addQuizlist(createQuiz(index, quiz))
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

    const createQuiz = (i, q) => {
      let quiz = new Quiz()
      quiz.setQuestion(q.question)
      quiz.setAnswer(q.answer)
      quiz.setOrder(i+1)
      return quiz
    }

    const removeQuiz = (i) => {
      contents.quizzes.splice(i, 1)
      save()
    }
    
    return {
      scopes,
      contents,
      save,
      addQuiz,
      createQuiz,
      createContentsRequest,
      removeQuiz,
    }
  }
})
</script>

<style lang="scss" scoped>
  .icon {
    font-size: 1.4rem;
    color: gray;
    padding-top: 16px;
    padding-left: 8px;
  }
</style>
