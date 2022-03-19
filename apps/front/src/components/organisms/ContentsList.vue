<template>
  <div class="contents-section">
    <h2>レッスン一覧</h2>
    <p>{{ contents.title }}</p>
    <p>{{ contents.creator }}</p>
    <div class="card-list-container">
    </div>
  </div>    
</template>

<script>
import { WritingServiceClient } from "@/pb/fast-writing_grpc_web_pb.js";
import { ContentsId } from "@/pb/models/contents_pb.js";
export default {
  name: 'ContentsList',
  data () {
    return {
      contents: {}
    };
  },
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

<style lang="scss">
@import "@/styles/_variables.scss";
  .contents-section {
    width: 58%;
    margin-right: auto;
    margin-left: auto;

    h2 {
      font-size: 1.8em;
      text-align: center;
    }
  }
  .card-list-container {
    display: flex;
    flex-flow: row wrap;
    justify-content: space-between;
  }

  .card-list-container::before{
    content:"";
    display: block;
    width:23%;
    order:1;
  }

  .card-list-container::after {
    display: block;
    content:"";
    width: 23%;
  }
</style>