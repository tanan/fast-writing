<template>
  <MainHeader />
  <div class="surface-section px-4 py-8 md:px-6 lg:px-8">
    <Toast position="bottom-right" :breakpoints="{'920px': {width: '100%', top: '0', right: '0'}}" />
    <div class="grid">
        <div class="col-12 lg:col-2">
            <div class="text-900 font-medium text-xl mb-3">Profile</div>
            <p class="m-0 p-0 text-600 line-height-3 mr-3">Please complete your profile.</p>
        </div>
        <div class="col-12 lg:col-10">
            <div class="grid formgrid p-fluid">
                <div class="field mb-4 col-12 md:col-6">
                    <label for="nickname" class="font-medium text-900">Display name</label>
                    <InputText id="nickname" type="text" v-model="profile.name" />
                </div>
                <div class="field mb-4 col-12">
                    <label for="icon" class="font-medium text-900">Icon</label>
                    <div class="flex align-items-center mt-2">
                        <div class="profile-icon mr-4">
                          <img v-show="profile.hasIcon" :src="profile.image" class="pl-4" style="max-width: 64px; max-height: 64px" />
                          <Button v-show="!profile.hasIcon" icon="pi pi-user" class="p-button-rounded p-button-secondary" style="width: 48px; height: 48px;" />
                        </div>
                        <FileUpload mode="basic" name="Upload Icon" :customUpload="true" @uploader="uploadIcon">
                          <template #empty>
                            <p>Drag and drop files to here to upload.</p>
                          </template>
                        </FileUpload>
                    </div>
                </div>
                <div class="field mb-4 col-12 md:col-6">
                  <label for="email" class="font-medium text-900">Email</label>
                  <InputText id="email" type="text" readonly v-model="profile.email" />
                </div>
                <div class="col-12">
                  <Button label="Save Changes" class="w-auto mt-3" @click="save()"></Button>
                </div>
            </div>
        </div>
    </div>
  </div>
  <MainFooter />
</template>

<script>
import { defineComponent, reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Store from '@/store/index.js'
import app from "@/plugins/firebase"
import { UserServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { getStorage, ref, uploadBytes, getDownloadURL } from "firebase/storage"
import { User, UserId } from "@/pb/models/user_pb.js"
import MainHeader from '@/components/organisms/MainHeader.vue'
import MainFooter from '@/components/organisms/MainFooter.vue'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import FileUpload from 'primevue/fileupload'
import Toast from 'primevue/toast'
import { useToast } from "primevue/usetoast"

const client = new UserServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default defineComponent({
  name: 'AccountPage',
  components: {
    MainHeader,
    MainFooter,
    InputText,
    Button,
    FileUpload,
    Toast,
  },
  setup () {
    const route = useRoute()
    const router = useRouter()
    const storage = getStorage(app)
    const storageRef = ref(storage, `users/${Store.state.userId}/profilePicture.png`)

    const toast = useToast()

    const profile = reactive({
      name: '',
      email: '',
      image: '',
      hasIcon: false,
    })

    const setProfile = (user) => {
      profile.name = user.name
      profile.email = user.email
      profile.image = downloadIcon()
    }

    const uploadIcon = (e) => {
      console.log(e.files[0])
      uploadBytes(storageRef, e.files[0], {contentType: 'image/png'}).then((snapshot) => {
        console.log(snapshot)
        console.log('Uploaded a blob or file!')
      })
    }

    const downloadIcon = () => {
      getDownloadURL(storageRef)
        .then((url) => {
          profile.image = url
          profile.hasIcon = true
        })
        .catch((error) => {
          switch (error.code) {
            case 'storage/object-not-found':
              break;
          }
        })
    }

    // init data
    const metadata = { 'authorization': 'Bearer ' + Store.state.userToken }
    const req = new UserId()
    req.setId(Store.state.userId)
    client.getUser(req, metadata, (err, resp) => {
      if (err) {
        if (err.message === "Unauthenticated") {
          router.push(`/signin?redirect=${route.fullPath}`)
          return
        }
        console.log("cannot find user: ", err)
        toast.add({severity:'error', summary: 'user not found', detail: `再ログインしてください。${err.message}`, life: 5000})
        return
      }
      setProfile(resp.toObject())
    })

    const save = () => {
      const metadata = { 'authorization': 'Bearer ' + Store.state.userToken }
      let req = new User()
      let userId = new UserId()
      userId.setId(Store.state.userId)
      req.setId(userId)
      req.setName(profile.name)
      req.setEmail(profile.email)
      client.updateUser(req, metadata, async (err) => {
      if (err) {
        console.log("failed to update user: ", err)
      }
      toast.add({severity:'success', summary: 'updated', detail: 'success to update user', life: 3000})
      await new Promise((resolve) => setTimeout(resolve, 1000))
      router.push('/')
    })
    }

    return {
      profile,
      uploadIcon,
      save,
    }
  }
})
</script>
