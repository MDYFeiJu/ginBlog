<template>
  <a-button type="primary" danger class="exitBtn" @click="logOut">
    退出
  </a-button>
</template>

<script>
import { useRouter } from 'vue-router'

export default {
  name: 'myHeader',
  setup() {
    const router = useRouter()
    const logOut = () => {
      window.sessionStorage.clear()
      router.push('login')
    }
    return {
      logOut,
    }
  },
}
</script>

<style></style>
