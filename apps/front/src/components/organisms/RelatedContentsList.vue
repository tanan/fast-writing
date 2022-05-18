<template>
  <div>
    <div class="col-offset-1 p-0 col-10 md:col-10 lg:col-6 lg:col-offset-3">
      <h3 class="text-700" style="font-size: 1.4rem">関連クイズ一覧</h3>
    </div>
    <div class="grid">
      <div class="flex align-items-stretch flex-wrap mt-4 col-offset-1 col-10 md:col-10 lg:col-7 lg:col-offset-3">
        <div class="flex col-12 lg:col-4 lg:pl-0 lg:pr-4" v-for="c in contents.contentsList" :key="c.id">
          <ContentsCard :contents="c" :editable="false" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Store from '@/store/index.js'
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { QueryParams, TagQueryParams } from "@/pb/models/query_pb.js"
import ContentsCard from "@/components/molecules/ContentsCard.vue"

const client = new WritingServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default defineComponent({
  name: 'ContentsList',
  components: {
    ContentsCard,
  },
  props: {
    tags: String,
  },
  setup (props) {
    const route = useRoute()
    const router = useRouter()

    const contents = reactive({
      contentsList: [],
    })

    const getTag = (tags) => {
      if (tags === undefined) {
        return "Unrecognized"
      }
      console.log(tags)
      return tags.split(',')[0]
    }

    const toContentsList = (list) => {
      for(var v of list) {
        let c = {
          id: v.id.id,
          title: v.title,
          lastUpdated: v.lastUpdated,
          description: v.description,
          creator: v.creator,
          scope: v.scope,
          tags: v.tags,
        }
        contents.contentsList.push(c)
      }
    }

    const getContentsListByTags = async () => {
      let req = new TagQueryParams()
      let queryParams = new QueryParams()
      req.setParams(queryParams)
      console.log(props.tags)
      req.setTags(getTag(props.tags))
      const metadata = { 'authorization': 'Bearer ' + Store.state.userToken }
      
      client.getContentsListByTags(req, metadata, (err, resp) => {
        if (err) {
          console.log(err)
          if (err.code === 2) {
            Store.dispatch('signout')
            router.push(`/signin?redirect=${route.fullPath}`)
            return
          }
          if (err.code === 16) {
            Store.dispatch('signout')
            router.push(`/signin?redirect=${route.fullPath}`)
            return
          }
        }
        if (resp) {
          toContentsList(resp.toObject().contentslistList)
        }
      })
    }

    getContentsListByTags()

    return {
      contents,
    }
  }
})
</script>

<style lang="scss" scoped>
  .title {
    margin: 8px 16px;
  }
</style>