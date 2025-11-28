<template>
  <div class="category-detail-page">
    <div v-if="data" class="category-detail-container">
      <!-- 戻るリンク -->
      <NuxtLink to="/" class="back-link">← 一覧に戻る</NuxtLink>

      <!-- カテゴリヘッダー -->
      <div class="category-header">
        <!-- カテゴリ画像 -->
        <div class="category-image-container">
          <img 
            :src="data.intellectualPropertyCategory.imgUrl" 
            :alt="data.intellectualPropertyCategory.name"
            class="category-image"
          />
          <div class="sample-watermark">SAMPLE</div>
        </div>

        <!-- 商品情報セクション -->
        <div class="product-info">
          <div class="info-row">
            <span class="info-label">【販売価格】</span>
            <span class="info-value">1個{{ Math.round(data.intellectualPropertyCategory.price).toLocaleString() }}円(税込)</span>
          </div>
          <div class="info-row">
            <span class="info-label">【販売期間】</span>
            <span class="info-value">{{ formatSalesPeriod(data.intellectualPropertyCategory.salesStartDate, data.intellectualPropertyCategory.salesEndDate) }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">【配送手数料】</span>
            <span class="info-value">{{ Math.round(data.intellectualPropertyCategory.fee).toLocaleString() }}円(税込)</span>
          </div>
          <div class="info-row">
            <span class="info-label">【お届け目安】</span>
            <span class="info-value">ご購入後、約1～2週間でお届け予定</span>
          </div>
          <div class="info-row">
            <span class="info-label">【注意事項】</span>
            <span class="info-value">{{ data.intellectualPropertyCategory.precautions || '特になし' }}</span>
          </div>
        </div>

        <!-- まとめ買い情報 -->
        <div class="bulk-purchase">
          <svg class="bulk-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <rect x="1" y="3" width="15" height="13"></rect>
            <polygon points="16 8 20 8 23 11 23 16 16 16 16 8"></polygon>
            <circle cx="5.5" cy="18.5" r="2.5"></circle>
            <circle cx="18.5" cy="18.5" r="2.5"></circle>
          </svg>
          <span class="bulk-text">
            1回の購入で20個までの配送手数料は一律です。<br>
            複数購入される場合はまとめてのご購入がお得です！
          </span>
        </div>

        <!-- SNSボタン -->
        <div class="social-buttons">
          <button class="social-button twitter">
            <span class="social-icon">𝕏</span>
            <span>ポストする</span>
          </button>
          <button class="social-button facebook">
            <span class="social-icon">f</span>
            <span>シェアする</span>
          </button>
          <button class="social-button line">
            <span class="social-icon">LINE</span>
            <span>送る</span>
          </button>
        </div>

        <!-- FAQボタン -->
        <NuxtLink to="/faq" class="faq-button">
          よくあるご質問はこちら
        </NuxtLink>
      </div>

      <!-- 商品ラインナップセクション -->
      <div class="rank-groups-section">
        <h2>商品ラインナップ</h2>
        <div class="rank-groups-list">
          <div 
            v-for="rankGroup in data.intellectualPropertyCategory.rankGroups" 
            :key="rankGroup.id"
            class="rank-group-card"
          >
            <div class="rank-group-header">
              <span :class="['rank-badge', `rank-badge-${rankGroup.rank.toLowerCase()}`]">{{ rankGroup.rank }}賞</span>
              <h3 class="rank-group-name">{{ rankGroup.name }}</h3>
              <span class="emission-rate">当選確率: {{ rankGroup.emissionRate }} / 100</span>
            </div>

            <!-- Properties Grid -->
            <div class="properties-grid">
              <div 
                v-for="(property, index) in rankGroup.properties" 
                :key="property.id"
                class="property-item"
              >
                <div class="property-image-container">
                  <img 
                    v-if="property.imgUrl" 
                    :src="property.imgUrl" 
                    :alt="property.name" 
                    class="property-image"
                  />
                  <div class="sample-watermark">SAMPLE</div>
                </div>
                <p class="property-name">{{ property.name }}</p>
                <p class="property-copyright" v-if="data.intellectualPropertyCategory.precautions">
                  {{ data.intellectualPropertyCategory.precautions }}
                </p>
              </div>
            </div>

            <p v-if="rankGroup.explanation" class="explanation">{{ rankGroup.explanation }}</p>
          </div>
        </div>
      </div>

      <!-- 固定された抽選セクション -->
      <div class="draw-section" ref="drawSectionRef" :style="{ bottom: drawSectionBottom }">
        <!-- カウントダウンタイマー -->
        <div class="countdown-timer">
          購入終了まで: {{ countdownText }}
        </div>

        <div class="draw-controls">
          <!-- 個数選択 -->
          <div class="quantity-selector">
            <label for="drawCount">個数:</label>
            <select 
              id="drawCount"
              v-model="drawCount"
              class="draw-select"
            >
              <option :value="0">選択してください</option>
              <option 
                v-for="i in Math.min(getTotalStock(data.intellectualPropertyCategory), 100)" 
                :key="i" 
                :value="i"
              >
                {{ i }}枚
              </option>
            </select>
          </div>

          <!-- くじを引くボタン -->
          <button 
            @click="drawIntellectualProperty"
            :disabled="!drawCount || drawCount <= 0 || isDrawing"
            class="draw-button"
          >
            <span>{{ isDrawing ? '抽選中...' : 'くじを引く' }}</span>
            <svg class="draw-button-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polyline points="13 17 18 12 13 7"></polyline>
              <polyline points="6 17 11 12 6 7"></polyline>
            </svg>
          </button>
        </div>
      </div>

      <!-- 抽選結果モーダル -->
      <div v-if="showResult" class="result-modal-overlay" @click="closeResult">
        <div class="result-modal" @click.stop>
          <h2 class="result-title">抽選結果</h2>
          <div class="result-items">
            <div v-for="(item, index) in drawResult" :key="index" class="result-item">
              <div :class="['result-rank-badge', `rank-badge-${item.rank.toLowerCase()}`]">
                {{ item.rank }}賞
              </div>
              <img v-if="item.imgUrl" :src="item.imgUrl" :alt="item.name" class="result-image" />
              <p class="result-name">{{ item.name }}</p>
            </div>
          </div>
          <button @click="closeResult" class="close-button">閉じる</button>
        </div>
      </div>
    </div>
    <div v-else-if="error" class="error-message">
      <p>エラー: {{ error }}</p>
    </div>
    <div v-else class="loading">
      <p>読み込み中...</p>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { isLoggedIn } from '~/utils/tokenStore'

const route = useRoute()
const router = useRouter()
const categoryId = route.params.id

const drawCount = ref(0)
const isDrawing = ref(false)
const drawResult = ref([])
const showResult = ref(false)
const drawSectionRef = ref(null)
const drawSectionBottom = ref('0px')

const { data, error } = await useAsyncGql("intellectualPropertyCategory", {
  ipCategoryId: categoryId
}, {
  getCachedData: () => undefined,
  clientId: 'default'
})

// カウントダウンタイマー
const countdownText = ref('')
let countdownInterval = null

const updateCountdown = () => {
  if (!data.value) return
  
  const endDate = new Date(data.value.intellectualPropertyCategory.salesEndDate)
  const now = new Date()
  const diff = endDate - now
  
  if (diff <= 0) {
    countdownText.value = '販売終了'
    if (countdownInterval) {
      clearInterval(countdownInterval)
    }
    return
  }
  
  const days = Math.floor(diff / (1000 * 60 * 60 * 24))
  const hours = Math.floor((diff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
  const minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60))
  const seconds = Math.floor((diff % (1000 * 60)) / 1000)
  
  countdownText.value = `${days}日 ${hours}時間 ${minutes}分 ${seconds}秒`
}

onMounted(() => {
  updateCountdown()
  countdownInterval = setInterval(updateCountdown, 1000)
  
  // スクロールイベントリスナーを追加
  window.addEventListener('scroll', handleScroll)
  handleScroll()
})

onUnmounted(() => {
  if (countdownInterval) {
    clearInterval(countdownInterval)
  }
  window.removeEventListener('scroll', handleScroll)
})

const handleScroll = () => {
  const footer = document.querySelector('footer')
  if (!footer || !drawSectionRef.value) return
  
  const footerRect = footer.getBoundingClientRect()
  const windowHeight = window.innerHeight
  const drawSectionHeight = drawSectionRef.value.offsetHeight
  
  if (footerRect.top < windowHeight) {
    const overlap = windowHeight - footerRect.top
    drawSectionBottom.value = `${overlap}px`
  } else {
    drawSectionBottom.value = '0px'
  }
}

const getTotalStock = (category) => {
  if (!category.rankGroups) return 0
  return category.rankGroups.reduce((total, rankGroup) => {
    return total + (rankGroup.properties || []).reduce((sum, property) => {
      return sum + (property.stock || 0)
    }, 0)
  }, 0)
}

const drawIntellectualProperty = async () => {
  if (!drawCount.value || drawCount.value <= 0) {
    alert('抽選枚数を選択してください')
    return
  }

  // ログインしてるかチェックして、ログインしてなかったらログインページ飛ばす
  if (!isLoggedIn()) {
    router.push('/login')
    return
  }

  isDrawing.value = true
  drawResult.value = []

  try {
    const { data: drawData, error: drawError } = await useAsyncGql("drawIntellectualProperty", {
      input: {
        ipCategoryId: categoryId,
        drawCount: parseInt(drawCount.value)
      }
    }, {
      getCachedData: () => undefined,
      clientId: 'default'
    })

    if (drawError.value) {
      throw new Error(drawError.value.message || '抽選に失敗しました')
    }

    if (drawData.value && drawData.value.drawIntellectualProperty) {
      drawResult.value = drawData.value.drawIntellectualProperty.map(item => ({
        rank: item.rank,
        name: item.name,
        imgUrl: item.imgUrl
      }))
      showResult.value = true
    }
  } catch (error) {
    console.error('抽選エラー:', error)
    alert('抽選に失敗しました。もう一度お試しください。')
  } finally {
    isDrawing.value = false
  }
}

const closeResult = () => {
  showResult.value = false
  drawCount.value = 0
  window.location.reload()
}

const formatSalesPeriod = (startDate, endDate) => {
  const formatDate = (dateStr) => {
    const date = new Date(dateStr)
    const year = date.getFullYear()
    const month = date.getMonth() + 1
    const day = date.getDate()
    const hour = String(date.getHours()).padStart(2, '0')
    const minute = String(date.getMinutes()).padStart(2, '0')
    
    return `${year}年${month}月${day}日${hour}:${minute}`
  }
  
  return `${formatDate(startDate)}~${formatDate(endDate)}まで`
}
</script>

<style scoped>
.category-detail-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  padding-bottom: 150px;
}

.back-link {
  display: inline-block;
  margin-bottom: 20px;
  color: #0066cc;
  text-decoration: none;
  font-weight: 500;
}

.back-link:hover {
  text-decoration: underline;
}

.category-header {
  background-color: white;
  border-radius: 8px;
  padding: 30px;
  margin-bottom: 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.category-image-container {
  position: relative;
  width: 100%;
  margin-bottom: 20px;
  border-radius: 8px;
  overflow: hidden;
}

.category-image {
  width: 100%;
  height: auto;
  display: block;
}

.sample-watermark {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  z-index: 5;
  display: flex;
  align-items: center;
  justify-content: center;
  background-image: repeating-linear-gradient(
    45deg,
    transparent,
    transparent 20px,
    rgba(0, 0, 0, 0.02) 20px,
    rgba(0, 0, 0, 0.02) 40px
  );
}

.sample-watermark::after {
  content: "SAMPLE";
  position: absolute;
  font-size: clamp(2rem, 10vw, 6rem);
  font-weight: bold;
  color: rgba(0, 0, 0, 0.15);
  transform: rotate(-45deg);
  letter-spacing: clamp(0.2rem, 1.5vw, 0.8rem);
  white-space: nowrap;
  /* アスペクト比に応じてスケール */
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.product-info {
  margin-bottom: 20px;
}

.info-row {
  display: flex;
  padding: 10px 0;
  border-bottom: 1px solid #f0f0f0;
  font-size: 0.95rem;
}

.info-row:last-child {
  border-bottom: none;
}

.info-label {
  font-weight: bold;
  color: #333;
  margin-right: 10px;
  flex-shrink: 0;
}

.info-value {
  color: #666;
  line-height: 1.6;
}

.bulk-purchase {
  display: flex;
  align-items: flex-start;
  gap: 15px;
  padding: 15px;
  background-color: #fff9e6;
  border-radius: 8px;
  margin-bottom: 20px;
}

.bulk-icon {
  width: 32px;
  height: 32px;
  color: #ff9900;
  flex-shrink: 0;
}

.bulk-text {
  font-size: 0.95rem;
  color: #333;
  line-height: 1.6;
}

.social-buttons {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

.social-button {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 10px;
  border: none;
  border-radius: 4px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: opacity 0.3s;
}

.social-button:hover {
  opacity: 0.8;
}

.social-button.twitter {
  background-color: #000;
  color: white;
}

.social-button.facebook {
  background-color: #1877f2;
  color: white;
}

.social-button.line {
  background-color: #00c300;
  color: white;
}

.social-icon {
  font-weight: bold;
}

.faq-button {
  display: block;
  width: 100%;
  padding: 15px;
  background-color: #28a745;
  color: white;
  text-align: center;
  text-decoration: none;
  border-radius: 4px;
  font-weight: 500;
  transition: background-color 0.3s;
}

.faq-button:hover {
  background-color: #218838;
}

.rank-groups-section {
  background-color: white;
  border-radius: 8px;
  padding: 30px;
  margin-bottom: 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.rank-groups-section h2 {
  font-size: 1.8rem;
  font-weight: bold;
  color: #333;
  margin-bottom: 20px;
}

.rank-groups-list {
  display: flex;
  flex-direction: column;
  gap: 30px;
}

.rank-group-card {
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 20px;
  background-color: #fafafa;
}

.rank-group-header {
  display: flex;
  align-items: center;
  gap: 15px;
  margin-bottom: 20px;
}

.rank-badge {
  padding: 8px 16px;
  border-radius: 4px;
  font-size: 1.1rem;
  font-weight: bold;
  color: white;
  flex-shrink: 0;
}

.rank-badge-s {
  background-color: #ffd700;
  color: #333;
  border-color: #ffc0cb;
}

.rank-badge-a {
  background-color: #ff6b6b;
  border-color: #ffc0cb;
}

.rank-badge-b {
  background-color: #4ecdc4;
  border-color: #ffc0cb;
}

.rank-badge-c {
  background-color: #95e1d3;
  color: #333;
  border-color: #ffc0cb;
}

.rank-badge-d {
  background-color: #28a745;
  color: white;
  border-color: #ffc0cb;
}

.property-rank-badge.rank-badge-d {
  background-color: #28a745;
  color: white;
  border-color: #ffc0cb;
}

.rank-badge-e {
  background-color: #dcedc1;
  color: #333;
  border-color: #ffc0cb;
}

.rank-group-name {
  font-size: 1.3rem;
  font-weight: bold;
  color: #333;
  margin: 0;
  flex-grow: 1;
}

.emission-rate {
  font-size: 0.95rem;
  color: #666;
  white-space: nowrap;
}

.properties-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
  margin-bottom: 20px;
}

.property-item {
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
  background-color: white;
}

.property-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.property-name {
  font-size: 0.95rem;
  font-weight: 500;
  color: #333;
  margin: 0;
}

.property-copyright {
  font-size: 0.8rem;
  color: #999;
  margin: 0;
}

.explanation {
  padding: 15px;
  background-color: #f8f9fa;
  border-radius: 4px;
  font-size: 0.95rem;
  color: #666;
  line-height: 1.6;
  margin: 0;
}

.draw-section {
  position: fixed;
  left: 0;
  right: 0;
  background-color: white;
  border-top: 2px solid #e0e0e0;
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.1);
  padding: 15px 20px;
  z-index: 100;
  transition: bottom 0.3s ease;
}

.countdown-timer {
  text-align: center;
  font-size: 0.95rem;
  font-weight: bold;
  color: #c30000;
  margin-bottom: 10px;
}

.draw-controls {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.quantity-selector {
  display: flex;
  align-items: center;
  gap: 10px;
}

.quantity-selector label {
  font-weight: bold;
  color: #333;
}

.draw-select {
  padding: 10px 15px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
}

.draw-button {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 30px;
  background: linear-gradient(135deg, #ff9900 0%, #ff6600 100%);
  background-image: repeating-linear-gradient(
    45deg,
    transparent,
    transparent 10px,
    rgba(255, 255, 255, 0.1) 10px,
    rgba(255, 255, 255, 0.1) 20px
  );
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1.1rem;
  font-weight: bold;
  cursor: pointer;
  transition: transform 0.3s;
}

.draw-button:hover:not(:disabled) {
  transform: scale(1.05);
}

.draw-button:disabled {
  background: #ccc;
  background-image: none;
  cursor: not-allowed;
}

.draw-button-icon {
  width: 24px;
  height: 24px;
}

.result-modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.7);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.result-modal {
  background-color: white;
  border-radius: 8px;
  padding: 30px;
  max-width: 800px;
  max-height: 80vh;
  overflow-y: auto;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
}

.result-title {
  font-size: 1.8rem;
  font-weight: bold;
  color: #333;
  text-align: center;
  margin-bottom: 20px;
}

.result-items {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 20px;
  margin-bottom: 20px;
}

.result-item {
  text-align: center;
}

.result-rank-badge {
  display: inline-block;
  padding: 8px 16px;
  border-radius: 4px;
  font-size: 1.1rem;
  font-weight: bold;
  color: white;
  margin-bottom: 10px;
}

.result-image {
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 8px;
  margin-bottom: 10px;
}

.result-name {
  font-size: 1rem;
  font-weight: 500;
  color: #333;
}

.close-button {
  display: block;
  width: 100%;
  padding: 15px;
  background-color: #0066cc;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1.1rem;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.3s;
}

.close-button:hover {
  background-color: #0052a3;
}

.error-message,
.loading {
  text-align: center;
  padding: 40px;
  font-size: 1.2rem;
  color: #666;
}

@media screen and (max-width: 767px) {
  .category-detail-page {
    padding: 10px;
    padding-bottom: 150px;
  }

  .category-header {
    padding: 15px;
  }

  .rank-groups-section {
    padding: 15px;
  }

  .rank-group-header {
    flex-wrap: wrap;
  }

  .properties-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 15px;
  }

  .draw-controls {
    flex-direction: column;
    gap: 10px;
  }

  .quantity-selector {
    width: 100%;
    justify-content: center;
  }

  .draw-button {
    width: 100%;
  }

  .social-buttons {
    flex-direction: column;
  }

  .result-items {
    grid-template-columns: 1fr;
  }
}

@media screen and (max-width: 479px) {
  .properties-grid {
    grid-template-columns: 1fr;
  }
}
</style>

