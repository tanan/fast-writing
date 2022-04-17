<template>
  <div>
    <div class="grid" v-show="hasUserContents()">
      <div class="title col-10 col-offset-1 lg:col-8 lg:col-offset-2">
        <h2 class="mt-4 ml-2 lg:ml-4">作成したコンテンツ</h2>
      </div>
    </div>
    <div class="grid">
      <div class="flex align-items-stretch flex-wrap col-10 col-offset-1 lg:col-8 lg:col-offset-2">
        <div class="flex col-12 lg:col-4 lg:pl-4 lg:pr-4" v-for="c in contents.userContentsList" :key="c.id">
          <ContentsCard :contents="c" />
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
        <div class="flex col-12 lg:col-4 lg:pl-4 lg:pr-4" v-for="c in contents.contentsList" :key="c.id">
          <ContentsCard :contents="c" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { QueryParams, ContentsQueryParams, UserContentsQueryParams } from "@/pb/models/query_pb.js"
import { UserId } from "@/pb/models/user_pb.js"
import ContentsCard from "@/components/molecules/ContentsCard.vue"
import Store from '@/store/index.js'

const client = new WritingServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default defineComponent({
  name: 'ContentsList',
  components: {
    ContentsCard
  },
  setup () {
    const route = useRoute()
    const router = useRouter()

    const contents = reactive({
      contentsList: [],
      userContentsList: []
    })

    const hasUserContents = () => {
      if (contents.userContentsList.length > 0) {
        return true
      }
      return false
    }

    const toContentsList = (type, list) => {
      for(var v of list) {
        let c = {
          id: v.id.id,
          title: v.title,
          lastUpdated: v.lastUpdated,
          description: v.description,
          creator: v.creator,
        }
        if (type === "user") {
          contents.userContentsList.push(c)  
        }
        contents.contentsList.push(c)
      }
    }

    const getContentsList = async () => {
      let req = new ContentsQueryParams()
      let queryParams = new QueryParams()
      req.setParams(queryParams)
      const metadata = { 'authorization': 'Bearer ' + Store.state.userToken }
      let response = await new Promise((resolve, reject) => {
        client.getContentsList(req, metadata, (err, resp) => {
          if (err) {
            console.log(err)
            reject(err)
          }
          resolve(resp.toObject().contentslistList)
        })
      })
      toContentsList('public', response)
    }

    const getUserContentsList = async () => {
      let req = new UserContentsQueryParams()
      let queryParams = new QueryParams()
      let userId = new UserId()
      req.setParams(queryParams)
      userId.setId(Store.state.userId)
      req.setId(userId)
      const metadata = { 'authorization': 'Bearer ' + Store.state.userToken }
      let response = await new Promise((resolve) => {
        client.getUserContentsList(req, metadata, (err, resp) => {
          if (err) {
            if (err.message.indexOf("ID token has expired at") > -1) {
              console.log("test")
              Store.dispatch('signout')
              router.push(`/signin?redirect=${route.fullPath}`)
            }
            console.log(err)
          }
          resolve(resp.toObject().contentslistList)
        })
      })
      toContentsList('user', response)
    }

    getUserContentsList()
    getContentsList()

    return {
      contents,
      hasUserContents,
    }
  },
  async created() {
  },
  methods: {
    
  }
})
</script>

<style lang="scss" scoped>
  .title {
    margin: 8px 16px;
  }
</style>