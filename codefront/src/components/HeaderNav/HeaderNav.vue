<script setup lang="ts">
import { useTeamInfoStore } from '@/store/teamInfo';
import { ElMessageBox, ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';

const router = useRouter();

const teamInfo = useTeamInfoStore().getTeamInfo();

const logout = () => {
  ElMessageBox.confirm('确定要退出登录吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    useTeamInfoStore().clearTeamInfo();
    router.push('/login');
    ElMessage({
      type: 'success',
      message: '已退出登录'
    });
  }).catch(() => {
    ElMessage({
      type: 'info',
      message: '已取消'
    });
  });
}

</script>

<template>
  <div class="header-nav">
    <div class="items">
      <div class="item left">
        <div class="title">
          <font-awesome-icon :icon="['fas', 'print']" size="xl" />
          XCPC-打印服务
        </div>
      </div>
      <div class="item right">
        <div class="location">
          <font-awesome-icon :icon="['fas', 'location-crosshairs']" size="lg" />
          位置：{{ teamInfo.location }}
        </div>
        <div class="teamname">
          <font-awesome-icon :icon="['fas', 'people-group']" size="lg" />
          队伍名：{{ teamInfo.teamname }}
        </div>
        <div class="logout" :onclick="logout">
          <font-awesome-icon :icon="['fas', 'right-from-bracket']" size="lg" />
          重新登陆
        </div>
      </div>
    </div>
  </div>

</template>

<style scoped>
.header-nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  padding: 20px;
  background-color: #17a1e0;
}

.items {
  display: flex;
  flex-direction: row;
  align-items: center;
  text-align: center;
  justify-content: space-between;
  width: 100%;
  color: white;
}

.item {
  font-size: 20px;
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  padding: 0 20px;
}

.left {
  text-align: left;
  justify-content: flex-start;
  cursor: default;
  div{
    padding: 0 20px;
    min-width: 200px;
    width: fit-content;
  }
}

.right {
  text-align: right;
  justify-content: flex-end;
  div{
    padding: 0 20px;
  }
}

.title {
  font-weight: bold;
  font-size: 24px;
  line-height: 1;
  color: white;
  cursor: default;
}

.logout {
  font-size: 20px;
  color: white;
  cursor: pointer;
  height: 100%;
  width: fit-content;
  min-width: 100px;
  line-height: 1;
}
.logout:hover {
  color: #9cdb9a;
}
</style>