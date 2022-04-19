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
          scope: v.scope,
        }
        if (type === "user") {
          contents.userContentsList.push(c)
        } else {
          contents.contentsList.push(c)
        }
      }
    }

    const getContentsList = async () => {
      let req = new ContentsQueryParams()
      let queryParams = new QueryParams()
      req.setParams(queryParams)
      client.getContentsList(req, {}, (err, resp) => {
        if (err) {
          console.log(err)
          throw new Error(err.message)
        }
        if (resp) {
          toContentsList('public', resp.toObject().contentslistList)
        }
      })
    }

    const getUserContentsList = async () => {
      let req = new UserContentsQueryParams()
      let queryParams = new QueryParams()
      let userId = new UserId()
      req.setParams(queryParams)
      userId.setId(Store.state.userId)
      req.setId(userId)
      const metadata = { 'authorization': 'Bearer ' + Store.state.userToken }
      
      client.getUserContentsList(req, metadata, (err, resp) => {
        if (err) {
          console.log(err)
          if (err.code === 2) {
            Store.dispatch('signout')
            router.push(`/signin?redirect=${route.fullPath}`)
          }
          if (err.code === 16) {
            Store.dispatch('signout')
            router.push(`/signin?redirect=${route.fullPath}`)
          }
        }
        if (resp) {
          toContentsList('user', resp.toObject().contentslistList)
        }
      })
    }

    getUserContentsList()
    getContentsList()

    return {
      contents,
      hasUserContents,
    }
  }
})
</script>

<style lang="scss" scoped>
  .title {
    margin: 8px 16px;
  }
</style>