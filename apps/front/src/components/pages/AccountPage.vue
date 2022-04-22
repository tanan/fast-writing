<template>
  <div class="surface-section px-4 py-8 md:px-6 lg:px-8">
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
                    <div class="flex align-items-center">
                        <!-- <img src="images/blocks/avatars/circle/avatar-f-4.png" class="mr-4" /> -->
                        <FileUpload name="Upload Icon" :customUpload="true" @uploader="uploadIcon">
                          <template #empty>
                            <p>Drag and drop files to here to upload.</p>
                          </template>
                        </FileUpload>
                        <!-- <Button label="Upload Icon" icon="pi pi-plus" class="w-auto mt-2" @click="uploadIcon"></Button> -->
                        <!-- <FileUpload mode="basic" name="avatar" url="./upload.php" accept="image/*" :maxFileSize="1000000" class="p-button-outlined p-button-plain" chooseLabel="Upload Image"></FileUpload> -->
                    </div>
                </div>
                <div class="field mb-4 col-12 md:col-6">
                    <label for="email" class="font-medium text-900">Email</label>
                    <InputText id="email" type="text" v-model="profile.email" />
                </div>
                <div class="col-12">
                    <Button label="Save Changes" class="w-auto mt-3" @click="save()"></Button>
                </div>
            </div>
        </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, reactive } from 'vue'
import Store from '@/store/index.js'
import app from "@/plugins/firebase"
import { UserServiceClient } from "@/pb/fast-writing_grpc_web_pb.js"
import { getStorage, ref, uploadBytes } from "firebase/storage"
import { UserId } from "@/pb/models/user_pb.js"
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import FileUpload from 'primevue/fileupload'

const client = new UserServiceClient(`${process.env.VUE_APP_WRITING_API_ENDPOINT}`, null, null)

export default defineComponent({
  name: 'AccountPage',
  components: {
    InputText,
    Button,
    FileUpload,
  },
  setup () {
    const storage = getStorage(app);
    const storageRef = ref(storage, `users/${Store.state.userId}/profilePicture.png`);

    const profile = reactive({
      name: '',
      email: ''
    })

    const setProfile = (user) => {
      profile.name = user.name
      profile.email = user.email
    }

    const uploadIcon = (e) => {
      uploadBytes(storageRef, e.file).then((snapshot) => {
        console.log(snapshot)
        console.log('Uploaded a blob or file!')
      })
    }

    // init data
    const metadata = { 'authorization': 'Bearer ' + Store.state.userToken }
    const req = new UserId()
    req.setId(Store.state.userId)
    client.getUser(req, metadata, (err, resp) => {
      if (err) {
        console.log("cannot find user: ", err)
      }
      setProfile(resp.toObject())
    })

    const save = () => {

    }

    return {
      profile,
      uploadIcon,
      save,
    }
  }
})
</script>