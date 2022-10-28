<template>
  <a-layout class="container">
    <Nav></Nav>
    <a-layout>
      <a-layout-header class="header"><my-header></my-header></a-layout-header>
      <a-layout-content><router-view></router-view></a-layout-content>
      <a-layout-footer><myFooter></myFooter></a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import Nav from '../components/admin/Nav'
import myFooter from '../components/admin/myFooter'
import myHeader from '@/components/admin/myHeader'

export default {
  name: 'admin-page',
  components: {
    Nav,
    myFooter,
    myHeader,
  },
  data() {
    return {}
  },
  methods: {},
}
</script>

<style>
.container {
  height: 100%;
}
.header {
  text-align: right;
}
.log {
  height: 32px;
  margin: 20px 10px 10px;
  display: flex;
  font-size: 20px;
  background-color: #fff;
  align-items: center;
  justify-content: center;
}
</style>
