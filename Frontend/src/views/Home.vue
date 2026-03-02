<template>
  <div class="drawer-container">
    <div class="header">
      <h1>Drawer 工具箱</h1>
      <p>你的个人专属数字化抽屉</p>
    </div>

    <div class="tools-grid">
      <el-card class="tool-card" shadow="hover" @click="goTo('/ledger')">
        <div class="card-content">
          <div class="icon">💰</div>
          <h3>记账本</h3>
          <p>个人财务管理与收支统计</p>
        </div>
      </el-card>

      <el-card class="tool-card is-disabled" shadow="never">
        <div class="card-content">
          <div class="icon">🛠️</div>
          <h3>敬请期待</h3>
          <p>更多工具正在开发中...</p>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goTo = (path) => {
  router.push(path)
}
</script>

<style scoped>
.drawer-container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 40px 20px;
}

.header {
  text-align: center;
  margin-bottom: 50px;
}

.header h1 {
  font-size: 2.5rem;
  color: #303133;
  margin-bottom: 10px;
}

.header p {
  color: #909399;
  font-size: 1.1rem;
}

.tools-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
}

.tool-card {
  cursor: pointer;
  border-radius: 12px;
  transition: all 0.3s;
}

.tool-card:hover {
  transform: translateY(-5px);
}

.tool-card.is-disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.tool-card.is-disabled:hover {
  transform: none;
}

.card-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  padding: 20px 0;
}

.icon {
  font-size: 40px;
  margin-bottom: 15px;
}

h3 {
  margin: 0 0 10px 0;
  color: #303133;
}

p {
  margin: 0;
  color: #606266;
  font-size: 0.9rem;
}
</style>