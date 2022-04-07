<template>
  <v-container>
    <v-row no-gutters>
      <v-col v-for="contents in contentsList" :key="contents.id" cols="12" sm="4">
        <ContentsCard :contents=contents />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { QueryParams, ContentsQueryParams } from "@/pb/models/query_pb.js"
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
  }),
  async created() {
   this.shows = await this.getShows();
  },
  methods: {
    async getShows () {
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
