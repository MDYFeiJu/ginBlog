<template>
  <a-layout-sider
    style="background-color: #fff"
    breakpoint="lg"
    v-model:collapsed="collapsed"
  >
    <div class="log">
      <span>{{ collapsed ? 'Blog' : 'My Blog' }}</span>
    </div>
    <a-menu
      v-model:openKeys="openKeys"
      v-model:selectedKeys="selectedKeys"
      style="width: 100%"
      mode="inline"
      @click="goToPage"
    >
      <a-menu-item key="index">
        <template #icon>
          <pie-chart-outlined />
        </template>
        <span>仪表盘</span>
      </a-menu-item>
      <a-sub-menu key="sub1">
        <template #icon>
          <MailOutlined />
        </template>
        <template #title>文章管理</template>
        <a-menu-item key="addart">
          <edit-outlined />
          <span>写文章</span>
        </a-menu-item>
        <a-menu-item key="artlist">
          <unordered-list-outlined />
          <span>文章列表</span>
        </a-menu-item>
      </a-sub-menu>
      <a-menu-item key="catelist">
        <unordered-list-outlined />
        <span>分类列表</span>
      </a-menu-item>
      <a-menu-item key="userlist">
        <user-outlined />
        <span>用户列表</span>
      </a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>

<script>
import { reactive, ref, toRefs } from 'vue'
import {
  MailOutlined,
  QqOutlined,
  AppstoreOutlined,
  SettingOutlined,
  PieChartOutlined,
  EditOutlined,
  UnorderedListOutlined,
  UserOutlined,
} from '@ant-design/icons-vue'
import router from '@/router'

export default {
  name: 'Nav',
  components: {
    MailOutlined,
    QqOutlined,
    AppstoreOutlined,
    SettingOutlined,
    PieChartOutlined,
    EditOutlined,
    UnorderedListOutlined,
    UserOutlined,
  },
  setup() {
    const state = reactive({
      selectedKeys: ['1'],
      openKeys: [],
    })
    const collapsed = ref(true)
    const goToPage = (item) => {
      router.push('/admin/' + item.key)
    }
    return {
      ...toRefs(state),
      collapsed,
      goToPage,
    }
  },
}
</script>

<style>
.nav {
  background: #fff;
}
</style>
