<template>
  <div>
    <div class="grid">
      <div class="title col-10 col-offset-1 lg:col-8 lg:col-offset-2">
        <h2 class="mt-4 ml-2 lg:ml-4">作成したコンテンツ</h2>
      </div>
    </div>
    <div class="grid">
      <div class="flex align-items-stretch flex-wrap col-10 col-offset-1 lg:col-8 lg:col-offset-2">
        <div class="flex col-12 lg:col-4 lg:pl-4 lg:pr-4" v-for="contents in userContentsList" :key="contents.id">
          <ContentsCard :contents=contents />
        </div>
      </div>
    </div>
    <div class="grid">
      <div class="title col-10 col-offset-1 lg:col-8 lg:col-offset-2">
        <h2 class="mt-4 ml-2 lg:ml-4">公開コンテンツ</h2>
      </div>
    </div>
    <div class="grid">
      <div class="flex align-items-stretch flex-wrap col-10 col-offset-1 lg:col-8 lg:col-offset-2">
        <div class="flex col-12 lg:col-4 lg:pl-4 lg:pr-4" v-for="contents in contentsList" :key="contents.id">
          <ContentsCard :contents=contents />
        </div>
      </div>
    </div>
  </div>
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
      this.toContentsList('public', response.toObject().contentslistList)
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
      this.toContentsList('user', response)
    },
    toContentsList(type, list) {
      for(var contents of list) {
        let c = {}
        c.id = contents.id.id
        c.title = contents.title
        c.lastUpdated = contents.lastUpdated
        c.description = contents.description
        c.creator = contents.creator
        if (type === 'user') {
          this.userContentsList.push(c)
        } else {
          this.contentsList.push(c)
        }
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