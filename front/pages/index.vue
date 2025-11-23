<template>
  <div>
    <h1>トップページ</h1>
    <!-- Serverからガチャの個数を取得してSelectBox表示する -->
    <!-- <select v-model="gachaCount">
      <option v-for="gacha in gachas" :key="gacha.id" :value="gacha.id">{{ gacha.name }}</option>
    </select> -->
    <!-- ガチャを引くボタンを用意する -->
    <!-- <button @click="drawGacha">ガチャを引く</button> -->
    <!-- ガチャの結果を表示する -->
    <!-- <div v-if="gachaResult">
      <p>ガチャの結果: {{ gachaResult }}</p>
    </div> -->
    
    <div v-if="data">
      <h2>知的財産カテゴリー一覧</h2>
      <ul>
        <li v-for="category in data.intellectualPropertyCategories" :key="category.id">
          <div>
            <h3>{{ category.name }} - ¥{{ category.price }}</h3>
            <p>総在庫: {{ getTotalStock(category) }}枚</p>
            <!-- 各ランクの在庫数を表示する -->
            <ul>
              <li v-for="rankGroup in category.rankGroups" :key="rankGroup.id">
                <p>{{ rankGroup.name }} - {{ getRankGroupStock(rankGroup) }}枚</p>
              </li>
            </ul>
            <!-- 在庫分のselectboxを用意する -->
            <select v-model="drawCounts[category.id]">
              <option :value="0">選択してください</option>
              <option 
                v-for="i in Math.min(getTotalStock(category), 100)" 
                :key="i" 
                :value="i"
              >
                {{ i }}枚
              </option>
            </select>
            <button 
              @click="drawIntellectualProperty(category.id, drawCounts[category.id])"
              :disabled="!drawCounts[category.id] || drawCounts[category.id] <= 0"
            >
              くじを引く
            </button>
          </div>
        </li>
      </ul>
    </div>
    <div v-else-if="error">
      <p>エラー: {{ error }}</p>
    </div>
    <div v-else>
      <p>読み込み中...</p>
    </div>
  </div>
</template>

<script setup>
// キャッシュを無効化して常に最新データを取得
// keyにタイムスタンプを含めることで、毎回新しいリクエストとして扱われる
const { data, error, refresh } = await useAsyncGql("intellectualPropertyCategories", null, {
  getCachedData: () => null, // キャッシュを無効化
});

// 各カテゴリーごとの抽選枚数を管理する
const drawCounts = ref({});

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

// くじを引く関数
const drawIntellectualProperty = async (categoryId, drawCount) => {
  // ログインしてるかチェックして、ログインしてなかったらログイン画面にリダイレクト
  // if (!isLoggedIn.value) {
  //   router.push('/login');
  //   return;
  // }
  if (!drawCount || drawCount <= 0) {
    alert('抽選枚数を選択してください');
    return;
  }

  // TODO: 決済方法を選択して、決済方法を選択する画面を用意しておく予定

  try {
    const { data: result } = await useAsyncGql('drawIntellectualProperty', {
      input: {
        ipCategoryId: categoryId,
        drawCount: parseInt(drawCount)
      }
    });
    
    if (result.value) {
      console.log('抽選結果:', result.value.drawIntellectualProperty.results);
      // ここで結果を表示する処理を追加
      refresh();
    }
  } catch (err) {
    console.error('抽選エラー:', err);
    alert('抽選に失敗しました');
  }
};

if (data.value) {
  console.log(data.value.intellectualPropertyCategories);
} else if (error.value) {
  console.error("Error loading intellectual property categories:", error.value);
}
</script>

