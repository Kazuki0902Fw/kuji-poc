<template>
  <div class="category-list-page">
    <!-- カテゴリーグリッド（単一カラム） -->
    <div v-if="data && data.intellectualPropertyCategories.length > 0" class="category-grid">
      <CategoryBanner 
        v-for="category in data.intellectualPropertyCategories" 
        :key="category.id"
        :category="category"
        @click="goToCategoryDetail(category.id)"
      />
    </div>
    
    <div v-else-if="error" class="error-message">
      <p>エラー: {{ error }}</p>
    </div>
    <div v-else class="loading">
      <p>読み込み中...</p>
    </div>

    <!-- お知らせセクション -->
    <section class="announcements-section">
      <div class="announcements-header">
        <h2 class="announcements-title">お知らせ</h2>
        <NuxtLink 
          v-if="hasMoreAnnouncements" 
          to="/announcements" 
          class="more-button"
        >
          もっと見る
        </NuxtLink>
      </div>
      <div class="announcements-list">
        <div v-if="latestAnnouncement" class="announcement-item">
          <div class="announcement-header">
            <span class="announcement-date">{{ formatDate(latestAnnouncement.date) }}</span>
            <span 
              v-if="latestAnnouncement.isImportant" 
              class="announcement-badge"
            >
              重要
            </span>
          </div>
          <h3 class="announcement-title">{{ latestAnnouncement.title }}</h3>
          <p class="announcement-content">{{ latestAnnouncement.content }}</p>
        </div>
        <div v-else class="no-announcements">
          <p>お知らせはありません</p>
        </div>
      </div>
    </section>

    <!-- くじコレとはセクション -->
    <section class="about-section">
      <h2 class="about-title">くじコレとは</h2>
      <div class="about-content">
        <p class="about-description">
          「くじコレ」は、スマホやPCから簡単に購入できる、ハズレなしのオンラインくじ！
          <br>魅力的なキャラクターの描き下ろしイラストやオリジナルグッズを多数ラインナップしています！</br>
          <br><span class="about-note">※商品はくじコレサービス内で再登場する場合がございます。あらかじめご了承ください。</span></br>
        </p>
        <h2 class="about-title">簡単購入ガイド
        </h2>
        <div class="about-features">
          <div class="feature-item">
            <span class="feature-step">ステップ1</span>
            <h3 class="feature-title">遊びたいくじを選ぶ</h3>
            <p class="feature-description">
              販売中のタイトルの中から引きたいくじを選びます。
            </p>
          </div>
          <div class="feature-item">
            <span class="feature-step">ステップ2</span>
            <h3 class="feature-title">ログイン・会員登録</h3>
            <p class="feature-description">
              くじを引くには会員登録してログインが必要です。
            </p>
          </div>
          <div class="feature-item">
            <span class="feature-step">ステップ3</span>
            <h3 class="feature-title">購入個数を決める</h3>
            <p class="feature-description">
              1回の購入で20個までの配送手数料は一律です。
            </p>
          </div>
          <div class="feature-item">
            <span class="feature-step">ステップ4</span>
            <h3 class="feature-title">お支払い(ご購入)</h3>
            <p class="feature-description">
              正常に決済が完了するとくじが引けます。
            </p>
          </div>
          <div class="feature-item">
            <span class="feature-step">ステップ5</span>
            <h3 class="feature-title">くじを引く</h3>
            <p class="feature-description">
              好きなくじを選んで引きます。当選商品がすぐに分かります。
            </p>
          </div>
          <div class="feature-item">
            <span class="feature-step">ステップ6</span>
            <h3 class="feature-title">商品を受け取る</h3>
            <p class="feature-description">
              商品の発送はお届け目安をご確認ください。
            </p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
const router = useRouter();

// キャッシュを無効化して常に最新データを取得
const { data, error } = await useAsyncGql("intellectualPropertyCategories", null, {
  getCachedData: () => null, // キャッシュを無効化
});

// ランクグループの在庫数を計算する関数
const getRankGroupStock = (rankGroup) => {
  if (!rankGroup.properties) return 0;
  return rankGroup.properties.reduce((total, property) => {
    return total + (property.stock || 0);
  }, 0);
};

// カテゴリーの総在庫数を計算する関数
const getTotalStock = (category) => {
  if (!category.rankGroups) return 0;
  return category.rankGroups.reduce((total, rankGroup) => {
    return total + getRankGroupStock(rankGroup);
  }, 0);
};

// カテゴリー詳細ページに遷移
const goToCategoryDetail = (categoryId) => {
  router.push(`/categories/${categoryId}`);
};

// お知らせデータ（仮データ、後でAPIから取得するように変更予定）
const announcements = ref([
  {
    id: 1,
    date: new Date('2024-01-15'),
    title: '新カテゴリー追加のお知らせ',
    content: '新しい知的財産カテゴリーが追加されました。ぜひご覧ください。',
    isImportant: true
  },
  {
    id: 2,
    date: new Date('2024-01-10'),
    title: 'システムメンテナンスのお知らせ',
    content: '2024年1月20日 2:00〜4:00の間、システムメンテナンスを実施いたします。',
    isImportant: false
  }
]);

// 最新のお知らせ（1件のみ）
const latestAnnouncement = computed(() => {
  if (announcements.value.length === 0) return null;
  // 日付でソートして最新のものを取得
  const sorted = [...announcements.value].sort((a, b) => {
    return new Date(b.date) - new Date(a.date);
  });
  return sorted[0];
});

// お知らせが2件以上あるかどうか
const hasMoreAnnouncements = computed(() => {
  return announcements.value.length > 1;
});

// 日付をフォーマットする関数
const formatDate = (date) => {
  const d = new Date(date);
  const year = d.getFullYear();
  const month = String(d.getMonth() + 1).padStart(2, '0');
  const day = String(d.getDate()).padStart(2, '0');
  return `${year}.${month}.${day}`;
};
</script>

<style scoped>
.category-list-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  font-size: 2rem;
  margin-bottom: 30px;
  text-align: center;
}

.category-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
  margin-top: 20px;
}

.category-grid > * {
  cursor: pointer;
  transition: transform 0.3s ease;
}

.category-grid > *:hover {
  transform: translateY(-4px);
}

.error-message,
.loading {
  text-align: center;
  padding: 40px;
  color: #7f8c8d;
}

.announcements-section {
  margin-top: 60px;
  padding-top: 40px;
  border-top: 2px solid #e0e0e0;
}

.announcements-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
}

.announcements-title {
  font-size: 1.8rem;
  font-weight: bold;
  color: #2c3e50;
  margin: 0;
}

.announcements-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.more-button {
  display: inline-block;
  padding: 8px 16px;
  background-color: #c30000;
  color: white;
  text-decoration: none;
  border-radius: 4px;
  font-weight: 600;
  font-size: 0.9rem;
  transition: background-color 0.3s;
  white-space: nowrap;
}

.announcement-item {
  background: white;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: box-shadow 0.3s;
}

.announcement-item:hover {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
}

.announcement-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.announcement-date {
  font-size: 0.9rem;
  color: #7f8c8d;
}

.announcement-badge {
  background-color: #e74c3c;
  color: white;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 0.75rem;
  font-weight: 600;
}

.announcement-title {
  font-size: 1.2rem;
  font-weight: 600;
  color: #2c3e50;
  margin: 0 0 10px 0;
}

.announcement-content {
  font-size: 0.95rem;
  color: #34495e;
  line-height: 1.6;
  margin: 0;
}

.no-announcements {
  text-align: center;
  padding: 40px;
  color: #7f8c8d;
}

.about-section {
  margin-top: 60px;
  padding-top: 40px;
  border-top: 2px solid #e0e0e0;
}

.about-title {
  font-size: 1.8rem;
  font-weight: bold;
  color: #2c3e50;
  margin-bottom: 30px;
  text-align: center;
}

.about-content {
  max-width: 800px;
  margin: 0 auto;
}

.about-description {
  font-size: 1.1rem;
  line-height: 1.8;
  color: #34495e;
  text-align: center;
  margin-bottom: 40px;
}

.about-note {
  font-size: 0.85rem;
}

.about-features {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 30px;
  margin-top: 40px;
}

.feature-item {
  position: relative;
  background: white;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 25px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s, box-shadow 0.3s;
}

.feature-step {
  position: absolute;
  top: -12px;
  left: 12px;
  background-color: #c30000;
  color: white;
  padding: 4px 12px;
  border-radius: 4px;
  font-size: 0.85rem;
  font-weight: 600;
  z-index: 1;
}

.feature-item:hover {
  transform: translateY(-4px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.feature-title {
  font-size: 1.3rem;
  font-weight: 600;
  color: #2c3e50;
  margin: 0 0 15px 0;
}

.feature-description {
  font-size: 0.95rem;
  line-height: 1.7;
  color: #34495e;
  margin: 0;
}

/* レスポンシブ対応 */
@media screen and (max-width: 767px) {
  .announcements-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 15px;
  }

  .announcements-title {
    font-size: 1.5rem;
  }

  .more-button {
    align-self: flex-end;
  }

  .about-features {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .about-description {
    font-size: 1rem;
    padding: 0 10px;
  }

  .category-grid {
    grid-template-columns: 1fr;
  }
}
</style>

