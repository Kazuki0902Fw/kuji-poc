<template>
  <div class="history-page">
    <div class="history-container">
      <!-- ヘッダー -->
      <div class="history-header">
        <div class="header-bar"></div>
        <h1 class="history-title">くじ結果一覧</h1>
      </div>

      <!-- コンテンツエリア -->
      <div class="history-content">
        <p v-if="!hasHistory" class="no-history-message">
          購入履歴はございません。
        </p>
        <!-- 将来的に購入履歴がある場合の表示エリア -->
        <div v-else class="history-list">
          <!-- 購入履歴のリストをここに表示 -->
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { isLoggedIn } from '~/utils/tokenStore';

const loggedIn = computed(() => isLoggedIn());

if (!loggedIn.value) {
  await navigateTo('/login');
}

const hasHistory = ref(false);
</script>

<style scoped>
.history-page {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  min-height: 60vh;
  padding: 40px 20px;
}

.history-container {
  background-color: white;
  border-radius: 8px;
  max-width: 600px;
  width: 100%;
  padding: 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.history-header {
  display: flex;
  align-items: center;
  margin-bottom: 30px;
  border-bottom: 3px solid #c30000;
}

.header-bar {
  width: 6px;
  height: 40px;
  background-color: #c30000;
  margin-right: 10px;
  flex-shrink: 0;
}

.history-title {
  font-size: 1.5rem;
  font-weight: bold;
  color: #333;
  margin: 0;
}

.history-content {
  padding: 0;
}

.no-history-message {
  color: #666;
  font-size: 0.95rem;
  margin: 0;
}

@media screen and (max-width: 767px) {
  .history-page {
    padding: 20px 10px;
  }

  .history-container {
    padding: 20px;
  }

  .history-title {
    font-size: 1.25rem;
  }
}
</style>

