<script lang="ts" setup>
import { reactive, ref } from 'vue';
import { ElMessage } from 'element-plus';
import HeaderNav from '@/components/HeaderNav/HeaderNav.vue';
import { codePrint } from '@/apis/apis';

const printContent = ref<string>('');
const inputRef = ref<null | HTMLInputElement>(null);


const submitPrint = () => {
  // 发送打印
  console.log(printContent.value);
  if (!printContent.value) {
    ElMessage({
      type: 'error',
      message: '请输入打印内容!'
    });
    return;
  }

  const data = {
    code: printContent.value
  };
  codePrint(data).then((res: any) => {
    console.log(res);
    if (res.message !== 'Success') {
      ElMessage({
        type: 'error',
        message: '发送打印失败'
      });
    }
    else {
      ElMessage({
        type: 'success',
        message: '发送打印成功'
      });
    };
  });
};

const resetPrint = () => {
  printContent.value = '';
  // 反馈
  ElMessage({
    type: 'success',
    message: '已重置!'
  })
}

const font = reactive({
  color: 'rgba(0, 0, 0, .15)',
  fontFamily: 'Arial',
  textAlign: 'center',
});


</script>

<template>
  <el-watermark content="打印服务" :font="font">
    <HeaderNav />
    <div class="shell">
      <div class="input-shell">
        <el-input id="code" ref="inputRef" type='textarea' v-model="printContent" placeholder="请输入打印内容"
          :autosize="{ minRows: 25, maxRows: 25 }" class="input-area" />

        <div class="btn-shell">
          <el-button @click="submitPrint" large round type="primary">打印</el-button>
          <el-button @click="resetPrint" type="default" large round>清空</el-button>
        </div>
      </div>
    </div>
  </el-watermark>
</template>

<style scoped>
.shell {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  height: fit-content;
  padding-top: 60px;
  background-color: #FBF9FC;
}

.input-shell {
  width: 70%;
  z-index: 999;
}

.input-area {
  margin: 20px 0;
  font-size: 18px;
}

.btn-shell {
  display: flex;
  width: 100%;
  justify-content: space-around;
  margin-bottom: 20px;
}

.btn-shell button {
  padding: 10px 20px;
  font-size: large;
  flex: 0.2;
}
</style>