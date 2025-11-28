<template>
  <div class="privacy-policy-page">
    <div class="privacy-policy-container">
      <!-- ヘッダー -->
      <div class="privacy-policy-header">
        <div class="header-bar"></div>
        <h1 class="privacy-policy-title">個人情報取り扱いについて</h1>
      </div>

      <!-- コンテンツ -->
      <div class="privacy-policy-content">
        <p class="privacy-policy-text">
          株式会社ティームエンタテインメントは、個人情報保護の重要性に鑑み、「個人情報の保護に関する法律」及び本個人情報保護方針を遵守し、お客さまのプライバシー保護に努めます。
        </p>
        <NuxtLink to="/privacy-policy-detail" class="detail-button">
          詳細を見る
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup>
// 個人情報保護方針ページのスクリプト
</script>

<style scoped>
.privacy-policy-page {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  min-height: 60vh;
  padding: 40px 20px;
}

.privacy-policy-container {
  background-color: white;
  border-radius: 8px;
  max-width: 600px;
  width: 100%;
  padding: 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  border: 1px solid #e0e0e0;
}

.privacy-policy-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 20px;
  border-bottom: 3px solid #c30000;
}

.header-bar {
  width: 6px;
  height: 40px;
  background-color: #c30000;
  flex-shrink: 0;
}

.privacy-policy-title {
  font-size: 1.5rem;
  font-weight: bold;
  color: #333;
  margin: 0;
}

.privacy-policy-content {
  padding-top: 20px;
}

.privacy-policy-text {
  font-size: 0.95rem;
  color: #333;
  line-height: 1.8;
  margin: 0 0 20px 0;
}

.detail-button {
  display: inline-block;
  padding: 12px 24px;
  background-color: #c30000;
  color: white;
  text-decoration: none;
  border-radius: 4px;
  font-weight: 500;
  transition: background-color 0.3s;
}

.detail-button:hover {
  background-color: #d31a1a;
}

@media screen and (max-width: 767px) {
  .privacy-policy-page {
    padding: 20px 10px;
  }

  .privacy-policy-container {
    padding: 20px;
  }

  .privacy-policy-title {
    font-size: 1.25rem;
  }

  .privacy-policy-text {
    font-size: 0.9rem;
  }
}
</style>

