<template>
  <v-container>
    <v-row no-gutters>
      <v-col cols="12" sm="4">
        <ContentsCard :contents=contents />
      </v-col>
      <v-col cols="12" sm="4">
        <ContentsCard :contents=contents />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { ContentsId } from "@/pb/models/contents_pb.js"
import ContentsCard from "@/components/molecules/ContentsCard.vue"

export default {
  name: 'ContentsList',
  components: {
    ContentsCard
  },
  data: () => ({
    contents: {},
  }),
  async created() {
   this.shows = await this.getShows();
  },
  methods: {
    getShows () {
      const client = new WritingServiceClient("http://localhost:10000", null, null)
      let req = new ContentsId()
      req.setId(1)
      try {
        client.getContents(req, {}, (err, resp) => {
          this.contents = resp.toObject()
        })
        console.log(this.contents)
      } catch (error) {
        throw new Error("Could not receive the data from API!")
      }
      
    }
  }
}
</script>
