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
        <!-- 購入履歴のリスト -->
        <div v-else class="history-list">
          <div 
            v-for="transaction in purchaseTransactions" 
            :key="transaction.id"
            class="transaction-card"
          >
            <div class="transaction-header">
              <div class="transaction-info">
                <div class="info-row">
                  <span class="info-label">購入個数:</span>
                  <span class="info-value">{{ transaction.purchaseQuantity }}個</span>
                </div>
                <div class="info-row">
                  <span class="info-label">購入金額:</span>
                  <span class="info-value">¥{{ formatPrice(transaction.purchasePrice) }}</span>
                </div>
                <div class="info-row">
                  <span class="info-label">支払日時:</span>
                  <span class="info-value">{{ formatDateTime(transaction.paidAt) }}</span>
                </div>
                <div class="info-row">
                  <span class="info-label">支払い方法:</span>
                  <span class="info-value">{{ formatPaymentMethod(transaction.paymentMethod) }}</span>
                </div>
                <div class="info-row">
                  <span class="info-label">ステータス:</span>
                  <span class="info-value" :class="`status-${transaction.status}`">
                    {{ formatStatus(transaction.status) }}
                  </span>
                </div>
              </div>
              <button 
                class="detail-button"
                @click="openDetail(transaction)"
              >
                詳細を見る
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 詳細モーダル -->
    <div v-if="selectedTransaction" class="modal-overlay" @click="closeDetail">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h2 class="modal-title">当選商品詳細</h2>
          <button class="modal-close" @click="closeDetail">×</button>
        </div>
        <div class="modal-body">
          <div class="properties-list">
            <div 
              v-for="history in selectedTransaction.purchaseHistories" 
              :key="history.id"
              class="property-card"
            >
              <div class="property-image-container">
                <img 
                  v-if="history.intellectualProperty.imgUrl" 
                  :src="history.intellectualProperty.imgUrl" 
                  :alt="history.intellectualProperty.name" 
                  class="property-image"
                />
                <div class="sample-watermark">SAMPLE</div>
              </div>
              <div class="property-info">
                <div class="property-rank-badge" :class="`rank-badge-${history.intellectualProperty.rankGroup?.rank?.toLowerCase() || ''}`">
                  {{ history.intellectualProperty.rankGroup?.rank || '' }}賞
                </div>
                <h3 class="property-name">{{ history.intellectualProperty.name }}</h3>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { isLoggedIn } from '~/utils/tokenStore';

const loggedIn = computed(() => isLoggedIn());

if (!loggedIn.value) {
  await navigateTo('/login');
}

// 購入履歴を取得
const { data, error } = await useAsyncGql('purchaseTransactions', null, {
  getCachedData: () => null,
});

const hasHistory = computed(() => {
  return data.value && data.value.purchaseTransactions && data.value.purchaseTransactions.length > 0;
});

const purchaseTransactions = computed(() => {
  return data.value?.purchaseTransactions || [];
});

const selectedTransaction = ref(null);

const openDetail = (transaction) => {
  selectedTransaction.value = transaction;
};

const closeDetail = () => {
  selectedTransaction.value = null;
};

const formatPrice = (price) => {
  return new Intl.NumberFormat('ja-JP').format(price);
};

const formatDateTime = (dateTime) => {
  if (!dateTime) return '-';
  const date = new Date(dateTime);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hour = String(date.getHours()).padStart(2, '0');
  const minute = String(date.getMinutes()).padStart(2, '0');
  return `${year}年${month}月${day}日 ${hour}:${minute}`;
};

const formatPaymentMethod = (method) => {
  const methods = {
    credit_card: 'クレジットカード',
    bank_transfer: '銀行振込',
    other: 'その他'
  };
  return methods[method] || method;
};

const formatStatus = (status) => {
  const statuses = {
    payment_pending: '支払い待ち',
    payment_success: '支払い完了',
    payment_failed: '支払い失敗',
    draw_success: '抽選成功',
    draw_failed: '抽選失敗'
  };
  return statuses[status] || status;
};
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

.history-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.transaction-card {
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 20px;
  background-color: #fafafa;
}

.transaction-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 20px;
}

.transaction-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.info-row {
  display: flex;
  align-items: center;
  gap: 10px;
}

.info-label {
  font-weight: bold;
  color: #666;
  min-width: 100px;
}

.info-value {
  color: #333;
}

.status-draw_success {
  color: #28a745;
  font-weight: bold;
}

.status-payment_success {
  color: #28a745;
  font-weight: bold;
}

.status-payment_pending {
  color: #ffc107;
  font-weight: bold;
}

.status-payment_failed,
.status-draw_failed {
  color: #dc3545;
  font-weight: bold;
}

.detail-button {
  padding: 10px 20px;
  background-color: #c30000;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 0.95rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.3s;
  flex-shrink: 0;
}

.detail-button:hover {
  background-color: #d31a1a;
}

/* モーダル */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  padding: 20px;
}

.modal-content {
  background-color: white;
  border-radius: 8px;
  max-width: 800px;
  width: 100%;
  max-height: 90vh;
  overflow-y: auto;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 30px;
  border-bottom: 1px solid #e0e0e0;
}

.modal-title {
  font-size: 1.5rem;
  font-weight: bold;
  color: #333;
  margin: 0;
}

.modal-close {
  background: none;
  border: none;
  font-size: 2rem;
  color: #666;
  cursor: pointer;
  padding: 0;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  line-height: 1;
}

.modal-close:hover {
  color: #333;
}

.modal-body {
  padding: 30px;
}

.properties-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 20px;
}

.property-card {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.property-image-container {
  position: relative;
  width: 100%;
  aspect-ratio: 1;
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid #e0e0e0;
}

.property-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.sample-watermark {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%) rotate(-45deg);
  font-size: 2rem;
  font-weight: bold;
  color: rgba(255, 0, 0, 0.3);
  pointer-events: none;
  white-space: nowrap;
}

.property-info {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.property-rank-badge {
  display: inline-block;
  padding: 4px 12px;
  border-radius: 4px;
  font-size: 0.85rem;
  font-weight: bold;
  color: white;
  width: fit-content;
}

.rank-badge-s {
  background-color: #ffd700;
  color: #333;
}

.rank-badge-a {
  background-color: #ff6b6b;
}

.rank-badge-b {
  background-color: #4ecdc4;
}

.rank-badge-c {
  background-color: #95e1d3;
  color: #333;
}

.rank-badge-d {
  background-color: #28a745;
}

.rank-badge-e {
  background-color: #6c757d;
}

.property-name {
  font-size: 1rem;
  font-weight: 500;
  color: #333;
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

  .transaction-header {
    flex-direction: column;
  }

  .detail-button {
    width: 100%;
  }

  .properties-list {
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    gap: 15px;
  }

  .modal-content {
    max-width: 100%;
  }

  .modal-header {
    padding: 15px 20px;
  }

  .modal-body {
    padding: 20px;
  }
}
</style>


