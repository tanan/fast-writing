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
                    <InputText id="nickname" type="text" />
                </div>
                <div class="field mb-4 col-12">
                    <label for="icon" class="font-medium text-900">Icon</label>
                    <div class="flex align-items-center">
                        <!-- <img src="images/blocks/avatars/circle/avatar-f-4.png" class="mr-4" /> -->
                        <Button label="Upload Icon" icon="pi pi-plus" class="w-auto mt-2"></Button>
                        <!-- <FileUpload mode="basic" name="avatar" url="./upload.php" accept="image/*" :maxFileSize="1000000" class="p-button-outlined p-button-plain" chooseLabel="Upload Image"></FileUpload> -->
                    </div>
                </div>
                <div class="field mb-4 col-12 md:col-6">
                    <label for="email" class="font-medium text-900">Email</label>
                    <InputText id="email" type="text" />
                </div>
                <div class="col-12">
                    <Button label="Save Changes" class="w-auto mt-3"></Button>
                </div>
            </div>
        </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'

export default defineComponent({
  name: 'AccountPage',
  components: {
    InputText,
    Button,
  }
})
</script>