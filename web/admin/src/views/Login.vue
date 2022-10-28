<template>
  <div class="container">
    <div class="loginBox">
      <h2 class="textBox">Login</h2>
      <div class="loginForm">
        <a-form
          ref="loginFormRef"
          :model="formState"
          :rules="rules"
          @finish="finish"
          @finishFailed="finishFailed"
        >
          <a-form-item name="username">
            <a-input v-model:value="formState.username">
              <template #prefix>
                <UserOutlined class="site-form-item-icon" />
              </template>
            </a-input>
          </a-form-item>
          <a-form-item name="password">
            <a-input v-model:value="formState.password">
              <template #prefix>
                <LockOutlined class="site-form-item-icon" />
              </template>
            </a-input>
          </a-form-item>
          <a-form-item class="loginBtn">
            <a-button type="primary" style="margin: 10px" html-type="submit">
              登录
            </a-button>
            <a-button type="info" @click="resetForm">取消</a-button>
          </a-form-item>
        </a-form>
      </div>
    </div>
  </div>
</template>

<script>
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue'
import { reactive, ref, getCurrentInstance } from 'vue'
import { message } from 'ant-design-vue'
import { useRouter } from 'vue-router'
export default {
  name: 'login-page',
  components: {
    UserOutlined,
    LockOutlined,
  },
  setup() {
    const instance = getCurrentInstance()

    const router = useRouter()
    const props = instance.appContext.config.globalProperties
    const formState = reactive({
      username: '',
      password: '',
    })
    const rules = {
      username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
        {
          min: 4,
          max: 12,
          message: '用户名必须在4-12字符之间',
          trigger: 'blur',
        },
      ],
      password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        {
          min: 6,
          max: 20,
          message: '密码必须在6-20字符之间',
          trigger: 'blur',
        },
      ],
    }
    const loginFormRef = ref()
    const resetForm = () => {
      loginFormRef.value.resetFields()
    }
    const finish = async () => {
      const { data: res } = await props.$http.post('login', formState)
      console.log(res)
      if (res.status !== 200) {
        return message.error(res.message)
      }
      window.sessionStorage.setItem('token', res.token)
      message.success('登录成功')
      await router.push('admin')
    }
    const finishFailed = () => {
      message.error('输入格式不正确')
    }
    return {
      resetForm,
      formState,
      loginFormRef,
      rules,
      finish,
      finishFailed,
    }
  },
}
</script>

<style>
.container {
  height: 100%;
  background: #282c34;
}
.loginBox {
  width: 450px;
  height: 300px;
  background-color: #fff;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}
.loginForm {
  width: 100%;
  position: absolute;
  bottom: 8%;
  padding: 5px 20px;
  box-sizing: border-box;
}
.loginBtn {
  display: flex;
  text-align: right;
}
.textBox {
  display: flex;
  justify-content: center;
  padding-top: 20px;
}
</style>
