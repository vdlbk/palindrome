<script lang="ts">

import { defineComponent } from 'vue'

interface dataProperties {
  setImage: string | ArrayBuffer | null
}

export default defineComponent({
  name: 'ImagePalindrome',
  data(): dataProperties {
    return {
      setImage: null
    }
  },
  methods: {
    handleChange(e: any) {
      console.log(e)
      const reader = new FileReader();

      reader.readAsDataURL(e[0]);

      reader.onload = (onloadEvent) => {
        if (onloadEvent && onloadEvent.target) {
          console.log(onloadEvent.target.result)
          this.setImage = onloadEvent.target.result
        }
      };
    },
    run() {
      console.log("dodo")
    }
  }
})
</script>

<template>
  <div>
    <div>
      <span v-if="setImage == null">No image uploaded</span>
      <div v-else>
        <NuxtImg
            :src="setImage"
        />
      </div>
    </div>

    <div class="mt-4 space-y-2">
      <UInput
          type="file"
          size="md"
          accept="image/*"
          name="fileInput"
          @change="handleChange"

      />

      <UButton color="primary" variant="solid" @click="run">Generate Palindromic Image</UButton>
    </div>

  </div>
</template>

<style scoped>

</style>