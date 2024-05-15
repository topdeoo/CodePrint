<script setup lang="ts">
import { useTeamInfoStore } from '@/store/teamInfo';

const teamInfo = useTeamInfoStore().getTeamInfo();

</script>

<template>
  <div class="header-nav">
    <div class="items">
      <div class="item">位置：{{ teamInfo.location }}</div>
      <div class="item title">打印服务</div>
      <div class="item">队伍名：{{ teamInfo.teamname }}</div>
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
  background-color: #cccccc;
}
.items{
  display: flex;
  flex-direction: row;
  align-items: center;
  text-align: center;
  justify-content: space-around;
}

.item {
  font-size: 20px;
  flex: 1;
}
.title {
  font-size: 24px;
  font-weight: bold;
}
</style>