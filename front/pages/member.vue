<template>
  <div class="member-page">
    <div class="member-container">
      <!-- ヘッダー -->
      <div class="member-header">
        <div class="header-bar"></div>
        <h1 class="member-title">会員情報</h1>
      </div>
      <!-- ボタン一覧 -->
      <div class="member-buttons">
        <button class="member-button primary">基本情報の変更</button>
        <button class="member-button secondary">パスワードの変更</button>
        <button class="member-button secondary">配送先住所の変更</button>
        <button class="member-button secondary">メールアドレスの変更</button>
        <button class="member-button secondary">購入履歴</button>
        <button class="member-button danger">退会</button>
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
</script>

<style scoped>
.member-page {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  min-height: 60vh;
  padding: 40px 20px;
}

.member-container {
  background-color: white;
  border-radius: 8px;
  max-width: 600px;
  width: 100%;
  padding: 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.member-header {
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

.member-title {
  font-size: 1.5rem;
  font-weight: bold;
  color: #333;
  margin: 0;
}

.member-buttons {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.member-button {
  padding: 15px 20px;
  border: 2px solid #c30000;
  border-radius: 4px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.3s;
  background-color: white;
  color: #333;
}

.member-button.primary {
  background-color: white;
  color: #333;
}

.member-button.primary:hover {
  background-color: #ffe4e1;
}

.member-button.secondary {
  background-color: white;
  color: #333;
}

.member-button.secondary:hover {
  background-color: #ffe4e1;
}

.member-button.danger {
  background-color: white;
  color: #333;
}

.member-button.danger:hover {
  background-color: #ffe4e1;
}

@media screen and (max-width: 767px) {
  .member-page {
    padding: 20px 10px;
  }

  .member-container {
    padding: 20px;
  }

  .member-title {
    font-size: 1.25rem;
  }
}
</style>


