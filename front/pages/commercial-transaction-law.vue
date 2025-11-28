<template>
  <div class="commercial-transaction-law-page">
    <div class="commercial-transaction-law-container">
      <!-- ヘッダー -->
      <div class="commercial-transaction-law-header">
        <div class="header-bar"></div>
        <h1 class="commercial-transaction-law-title">特定商取引法に基づく表示</h1>
      </div>

      <!-- 情報テーブル -->
      <div class="info-table">
        <!-- 販売事業者 -->
        <div class="info-row">
          <div class="info-label">販売事業者</div>
          <div class="info-value">株式会社ティームエンタテインメント</div>
        </div>

        <!-- 運営責任者 -->
        <div class="info-row">
          <div class="info-label">運営責任者</div>
          <div class="info-value">代表取締役社長 賀島義成</div>
        </div>

        <!-- 住所 -->
        <div class="info-row">
          <div class="info-label">住所</div>
          <div class="info-value">〒101-0003 東京都千代田区一ツ橋2-4-3 光文恒産ビル7階</div>
        </div>

        <!-- 電話番号 -->
        <div class="info-row">
          <div class="info-label">電話番号</div>
          <div class="info-value">03-3222-6511</div>
        </div>

        <!-- メールアドレス -->
        <div class="info-row">
          <div class="info-label">メールアドレス</div>
          <div class="info-value">
            support@kujico.jp<br>
            <span class="info-note">※配信専用メールアドレスとなります。お問い合わせは、専用ページからお願いいたします。</span>
          </div>
        </div>

        <!-- URL -->
        <div class="info-row">
          <div class="info-label">URL</div>
          <div class="info-value">https://kujico.jp/</div>
        </div>

        <!-- 販売価格 -->
        <div class="info-row">
          <div class="info-label">販売価格</div>
          <div class="info-value">各商品の購入ページ参照</div>
        </div>

        <!-- 商品代金以外の必要代金 -->
        <div class="info-row">
          <div class="info-label">商品代金以外の必要代金</div>
          <div class="info-value">
            配送手数料・消費税<br>
            各商品の購入ページ参照<br>
            <span class="info-note">※配送手数料の詳細は、<NuxtLink to="/terms" class="content-link">注意事項</NuxtLink>の配送手数料についてを参照</span>
          </div>
        </div>

        <!-- 注文方法 -->
        <div class="info-row">
          <div class="info-label">注文方法</div>
          <div class="info-value">
            WEBサイトから<br>
            <span class="info-note">※商品購入に期間がある場合、各商品の購入ページ参照</span>
          </div>
        </div>

        <!-- お支払い方法 -->
        <div class="info-row">
          <div class="info-label">お支払い方法</div>
          <div class="info-value">
            クレジットカード(VISA、MASTER、JCB) ※1回払いのみ<br>
            Vプリカ(VISAプリペイドカード)、Vプリカギフト<br>
            ドコモ払い、ソフトバンクまとめて支払い、auかんたん決済<br>
            PayPay、メルペイ、Amazon Pay<br>
            <span class="info-note">※決済は決済代行業者「GMOペイメントゲートウェイ株式会社」を通じて行われます。</span><br>
            <span class="info-note">※ご利用するクレジットカード等は、お客様の本人名義である必要があります。</span><br>
            <span class="info-note">※ご請求は「くじコレ」という表示でさせていただきます。</span>
          </div>
        </div>

        <!-- 支払期限 -->
        <div class="info-row">
          <div class="info-label">支払期限</div>
          <div class="info-value">
            ご請求はご購入完了後(くじの抽選結果の表示後)、即時となります。<br>
            お客様の口座からの引き落とし時期につきましては、ご利用のクレジットカード会社・キャリアに準じます。
          </div>
        </div>

        <!-- 引渡し時期 -->
        <div class="info-row">
          <div class="info-label">引渡し時期</div>
          <div class="info-value">
            商品お届け時期は、<NuxtLink to="/terms" class="content-link">注意事項</NuxtLink>参照
          </div>
        </div>

        <!-- 返品・交換・キャンセルについて -->
        <div class="info-row">
          <div class="info-label">返品・交換・キャンセルについて</div>
          <div class="info-value">
            商品購入後の返品・返却・キャンセルはお受けいたしかねます。<br>
            商品が不良の場合にのみ交換いたします。
          </div>
        </div>

        <!-- 商品交換の配送手数料 -->
        <div class="info-row">
          <div class="info-label">商品交換の配送手数料</div>
          <div class="info-value">不良品の場合は弊社が負担いたします。</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// 特定商取引法に基づく表示ページのスクリプト
</script>

<style scoped>
.commercial-transaction-law-page {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  min-height: 60vh;
  padding: 40px 20px;
}

.commercial-transaction-law-container {
  background-color: white;
  border-radius: 8px;
  max-width: 600px;
  width: 100%;
  padding: 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.commercial-transaction-law-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 30px;
  border-bottom: 3px solid #c30000;
}

.header-bar {
  width: 6px;
  height: 40px;
  background-color: #c30000;
  flex-shrink: 0;
}

.commercial-transaction-law-title {
  font-size: 1.5rem;
  font-weight: bold;
  color: #333;
  margin: 0;
}

.info-table {
  display: flex;
  flex-direction: column;
}

.info-row {
  display: flex;
  border-bottom: 1px solid #e0e0e0;
  padding: 15px 0;
}

.info-row:first-child {
  border-top: none;
}

.info-row:last-child {
  border-bottom: none;
}

.info-label {
  width: 200px;
  font-weight: bold;
  color: #333;
  font-size: 0.95rem;
  flex-shrink: 0;
  padding-right: 20px;
}

.info-value {
  flex: 1;
  color: #333;
  font-size: 0.95rem;
  line-height: 1.8;
}

.info-note {
  color: #666;
  font-size: 0.9rem;
}

.content-link {
  color: #0066cc;
  text-decoration: underline;
  font-weight: 500;
}

.content-link:hover {
  color: #0052a3;
}

@media screen and (max-width: 767px) {
  .commercial-transaction-law-page {
    padding: 20px 10px;
  }

  .commercial-transaction-law-container {
    padding: 20px;
  }

  .commercial-transaction-law-title {
    font-size: 1.25rem;
  }

  .info-row {
    flex-direction: column;
    gap: 10px;
  }

  .info-label {
    width: 100%;
    padding-right: 0;
  }

  .info-value {
    font-size: 0.9rem;
  }
}
</style>


