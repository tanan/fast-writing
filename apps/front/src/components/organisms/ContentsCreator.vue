<template>
  <v-form class="contents-creator-form" ref="form">
    <v-text-field class="title" v-model="title" label="title" required></v-text-field>
    <!-- <ContentsQuizCreator /> -->
    <v-btn @click="save" color="success">SAVE</v-btn>
  </v-form>
</template>

<script>
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { Contents } from "@/pb/models/contents_pb.js"
import { CreateContentsRequest } from "@/pb/fast-writing_pb.js"
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
  }),
  methods: {
    async save () {
      const client = new WritingServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)
      let req = new CreateContentsRequest()
      let contents = new Contents()
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
        })
    }
  }
}
</script>

<style lang="scss" scoped>
  .contents-creator-form {
    margin: 0 8%;
    margin-top: 2%;

    .title {
      width: 50%;
    }
  }
</style>
