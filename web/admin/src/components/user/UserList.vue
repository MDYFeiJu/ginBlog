<template>
  <h3>用户列表界面</h3>
  <a-card>
    <a-row :gutter="20">
      <a-col :span="6">
        <a-input-search placeholder="请输入用户名查找" enter-button />
      </a-col>
      <a-col :span="1">
        <a-button type="primary">新增</a-button>
      </a-col>
    </a-row>
    <a-table
      :columns="columns"
      :pagination="paginationOption"
      :data-source="userlist"
      @change="handleTableChange"
      bordered
    >
      <template #role="role"></template>
      <template slot="action"></template>
    </a-table>
  </a-card>
</template>

<script>
import { getCurrentInstance, onMounted, reactive, ref } from 'vue'
import { message } from 'ant-design-vue'

export default {
  name: 'UserList',
  setup() {
    const app = getCurrentInstance().appContext.config.globalProperties
    let userlist = ref([])
    const columns = [
      {
        title: 'ID',
        dataIndex: 'ID',
        key: 'id',
        width: '20%',
      },
      {
        title: '用户名',
        dataIndex: 'username',
        key: 'username',
        width: '20%',
      },
      {
        title: '角色',
        dataIndex: 'role',
        key: 'role',
        width: '20%',
        scopedSlots: { customRender: 'role' },
        customRender: (text) => {
          return text.text === 1 ? '管理员' : '订阅者'
        },
      },
      {
        title: '操作',
        key: 'action',
        width: '30%',
        scopedSlots: { customRender: 'action' },
        customRender: () => {
          return (
            <div className="actionSlot">
              <a-button
                type="primary"
                style="margin-right:15px;border-radius:15px"
              >
                编辑
              </a-button>
              <a-button type="danger" style="border-radius:15px">
                删除
              </a-button>
            </div>
          )
        },
      },
    ]
    const paginationOption = reactive({
      pageSizeOptions: ['2', '5', '10'],
      defaultPageSize: 2,
      current: 1,
      total: 0,
      showSizeChanger: true,
      showTotal: (total) => `共${total}条`,
      onChange: (page, pagesize) => {
        paginationOption.defaultPageSize = pagesize
        paginationOption.current = page
        getUserList()
      },
      onShowSizeChange: (current, size) => {
        paginationOption.defaultPageSize = size
        paginationOption.current = current
        getUserList()
      },
    })
    const getUserList = async () => {
      const { data: res } = await app.$http.get('users', {
        params: {
          pagesize: paginationOption.defaultPageSize,
          pagenum: paginationOption.current,
        },
      })
      if (res.status !== 200) return message.error(res.message)
      paginationOption.total = res.total
      userlist.value = [...res.data]
    }
    getUserList()

    const handleTableChange = () => {}
    return {
      userlist,
      // getUserList,
      columns,
      paginationOption,
    }
  },
}
</script>

<style>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>
