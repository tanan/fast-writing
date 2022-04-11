<template>
  <v-container>
    <div class="title"><h2>作成したコンテンツ</h2></div>
    <v-row no-gutters>
      <v-col v-for="contents in contentsList" :key="contents.id" cols="12" sm="4">
        <ContentsCard :contents=contents />
      </v-col>
    </v-row>
    <div class="title mt-6"><h2>公開コンテンツ</h2></div>
    <v-row no-gutters>
      <v-col v-for="contents in contentsList" :key="contents.id" cols="12" sm="4">
        <ContentsCard :contents=contents />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { QueryParams, ContentsQueryParams, UserContentsQueryParams } from "@/pb/models/query_pb.js"
import { UserId } from "@/pb/models/user_pb.js"
import ContentsCard from "@/components/molecules/ContentsCard.vue"
import Store from '@/store/index.js'

const client = new WritingServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default {
  name: 'ContentsList',
  components: {
    ContentsCard
  },
  data: () => ({
    contentsList: [],
    userContentsList: []
  }),
  async created() {
   this.getContentsList()
   this.getUserContentsList()

  },
  methods: {
    async getContentsList () {
      let req = new ContentsQueryParams()
      let queryParams = new QueryParams()
      req.setParams(queryParams)
      const token = Store.state.userToken
      const metadata = { 'authorization': 'Bearer ' + token }
      let response = await new Promise((resolve, reject) => {
        client.getContentsList(req, metadata, (err, resp) => {
          if (err) {
            reject(err)
          }
          resolve(resp)
        })
      })
      this.toContentsList(response.toObject().contentslistList)
    },
    async getUserContentsList () {
      let req = new UserContentsQueryParams()
      let queryParams = new QueryParams()
      let userId = new UserId()
      req.setParams(queryParams)
      userId.setId(Store.state.userId)
      req.setId(userId)
      const token = Store.state.userToken
      const metadata = { 'authorization': 'Bearer ' + token }
      let response = await new Promise((resolve) => {
        client.getUserContentsList(req, metadata, (err, resp) => {
          if (err) {
            // reject(err)
          }
          console.log(resp.toObject().contentslistList)
          resolve(resp.toObject().contentslistList)
        })
      })
      this.toContentsList(response)
    },
    toContentsList(list) {
      for(var contents of list) {
        let c = {}
        c.id = contents.id.id
        c.title = contents.title
        c.lastUpdated = contents.lastUpdated
        c.creator = contents.creator
        this.contentsList.push(c)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
  .title {
    margin: 8px 16px;
  }
</style>