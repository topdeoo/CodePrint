<script lang="ts" setup>
import { ref, reactive } from 'vue';
import type { FormInstance, FormRules } from 'element-plus';
import { ElMessage } from 'element-plus';
import { useTeamInfoStore } from '@/store/teamInfo';
import { useRouter } from 'vue-router';
import { teamLogin } from '@/apis/apis';

const router = useRouter();

interface Form {
  teamname: string
  password: string
  location: string
};

const ruleFormRef = ref<FormInstance>()

const form = ref<Form>({
  teamname: '',
  password: '',
  location: ''
});

const rule = reactive<FormRules<Form>>({
  teamname: [
    { required: true, message: '请输入队伍名称', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' }
  ],
  location: [
    { required: true, message: '请输入座位号', trigger: 'blur' }
  ]
})

const submitLogin = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate((valid, fields) => {
    if (valid) {
      console.log('submit');
      // 登陆请求
      console.log(form.value);
      teamLogin(form.value)
        .then((res: any) => {
          console.log(res);
          if (res.data.message !== 'Success') {
            ElMessage({
              type: 'error',
              message: res.data.message
            });
          }
          else {
            ElMessage({
              type: 'success',
              message: '登录成功'
            });
            // 存token
            const token = res.data.token;
            useTeamInfoStore().setTeamInfo(form.value.teamname, form.value.location, token)
            // 跳转
            router.push({ path: '/print' });
          }
        }).catch((err: any) => {
          console.log(err);
          ElMessage({
            type: 'error',
            message: '登录失败'
          });
        });
    } else {
      console.log('error submit!!', fields)
      ElMessage({
        type: 'error',
        message: '请检查输入'
      });
    }
  });
};

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
}

</script>

<template>
  <div class="shell">
    <el-form ref="ruleFormRef" :model="form" label-position="top" :rules="rule" class="form-style">
      <h3 class="form-title">队伍登录</h3>
      <el-form-item prop="teamname" label="队伍名称">
        <el-input v-model="form.teamname" placeholder="Teamname"></el-input>
      </el-form-item>
      <el-form-item prop="password" label="密码">
        <el-input v-model="form.password" placeholder="Password"></el-input>
      </el-form-item>
      <el-form-item prop="location" label="位置号">
        <el-input v-model="form.location" placeholder="Location"></el-input>
      </el-form-item>
      <el-form-item>
        <div class="btn-shell">
          <el-button @click="submitLogin(ruleFormRef)" type="primary" large round>登录</el-button>
          <el-button @click="resetForm(ruleFormRef)" type="default" large round>重置</el-button>
        </div>
      </el-form-item>
    </el-form>
  </div>
</template>
<style scoped>
.shell {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

.form-style {
  width: 400px;
  font-size: 20px;
  box-shadow: rgba(0, 0, 0, 0.12) 2px 2px 14px 0px;
  height: fit-content;
  padding: 20px;
}

.form-title {
  text-align: center;
  font-size: 24px;
  margin-bottom: 16px;
}

.btn-shell {
  display: flex;
  width: 100%;
  justify-content: center;
}

.btn-shell button {
  padding: 0px 10px;
  flex: 1;
}
</style>