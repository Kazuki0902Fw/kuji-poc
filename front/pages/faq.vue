<template>
  <div class="faq-page">
    <div class="faq-container">
      <div class="faq-header">
        <div class="header-bar"></div>
        <h1 class="faq-title">よくある質問</h1>
      </div>

      <!-- ログインについて -->
      <section class="faq-section">
        <div class="faq-section-header">
          <svg class="faq-section-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/>
            <polyline points="3.27 6.96 12 12.01 20.73 6.96"/>
            <line x1="12" y1="22.08" x2="12" y2="12"/>
          </svg>
          <h2 class="faq-section-title">ログインについて</h2>
        </div>
        <div class="faq-items">
          <div class="faq-item">
            <p class="faq-question">■会員登録しないと買えないの?</p>
            <p class="faq-answer">このサービスは会員登録後に商品を購入することができます。</p>
          </div>
          <div class="faq-item">
            <p class="faq-question">■IDを忘れた</p>
            <p class="faq-answer">IDはメールアドレスとなります。そのため、個人情報に該当するためお答えできません。メールアドレスをご確認ください。</p>
          </div>
          <div class="faq-item">
            <p class="faq-question">■パスワードを忘れた</p>
            <p class="faq-answer">パスワードの再設定が可能です。「<NuxtLink to="/forgot-password" class="faq-link">パスワードを忘れた</NuxtLink>」ページより、登録されているメールアドレス宛にパスワード再設定メールを送信いたしますので、案内に従って新しいパスワードを設定してください。</p>
          </div>
          <div class="faq-item">
            <p class="faq-question">■ログインできない</p>
            <p class="faq-answer">PCやスマートフォンの機種変更を行った場合、機種や動作環境によってはログインできない場合があります。大変お手数ですが、「<NuxtLink to="/contact" class="faq-link">お問い合わせ</NuxtLink>」よりご連絡ください。</p>
          </div>
          <div class="faq-item">
            <p class="faq-question">■退会したい</p>
            <p class="faq-answer">退会をご希望の場合は、ログイン後「<NuxtLink to="/member" class="faq-link">こちら</NuxtLink>」より退会が可能です。なお、退会されますとご利用中のサービスに関する情報はすべて破棄されますのでご注意ください。</p>
          </div>
          <div class="faq-item">
            <p class="faq-question">■個人情報の取り扱いについて知りたい</p>
            <p class="faq-answer">「<NuxtLink to="/privacy-policy" class="faq-link">個人情報保護方針</NuxtLink>」をご確認ください。</p>
          </div>
        </div>
      </section>

      <!-- お支払いについて -->
      <section class="faq-section">
        <div class="faq-section-header">
          <svg class="faq-section-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/>
            <polyline points="3.27 6.96 12 12.01 20.73 6.96"/>
            <line x1="12" y1="22.08" x2="12" y2="12"/>
          </svg>
          <h2 class="faq-section-title">お支払いについて</h2>
        </div>
        <div class="faq-items">
          <div class="faq-item">
            <p class="faq-question">■どんな支払い方法があるの?</p>
            <div class="faq-answer">
              <p>以下の方法でお支払いが可能です。</p>
              <ul class="faq-list">
                <li>①クレジットカード（VISA、MASTER、JCB）※一括払いのみ</li>
                <li>②Vプリカ/Vプリカギフト</li>
                <li>③ドコモ払い</li>
                <li>④ソフトバンクまとめて支払い</li>
                <li>⑤auかんたん決済</li>
                <li>⑥PayPay</li>
                <li>⑦メルペイ</li>
                <li>⑧Amazon Pay</li>
              </ul>
              <p>※Vプリカ（Visaプリペイドカード）の詳細は「<NuxtLink to="/vpreca" class="faq-link">こちら</NuxtLink>」、Vプリカギフトの詳細は「<NuxtLink to="/vpreca-gift" class="faq-link">こちら</NuxtLink>」をご確認ください。</p>
            </div>
          </div>
          <div class="faq-item">
            <p class="faq-question">■領収書は発行してもらえるの?</p>
            <p class="faq-answer">領収書の発行は行っておりません。</p>
          </div>
          <div class="faq-item">
            <p class="faq-question">■購入完了と購入内容は確認できるの?</p>
            <p class="faq-answer">お支払いが正常に完了した場合、登録されているメールアドレス宛に「購入完了通知メール」を送信いたします。<br>※購入完了通知メールの送信まで時間がかかる場合がありますのでご了承ください。</p>
          </div>
          <div class="faq-item">
            <p class="faq-question">■何が当たったのか確認したい</p>
            <p class="faq-answer">ログイン後、くじ結果一覧より獲得した商品の詳細を確認することができます。<br>※獲得した商品が反映されるまで時間がかかる場合がありますのでご了承ください。</p>
          </div>
          <div class="faq-item">
            <p class="faq-question">■返品・交換・キャンセルはできるの?</p>
            <div class="faq-answer">
              <p>くじ結果が発生するサービスの性質上、お客様都合による返品・交換・キャンセルはできません。ご了承ください。</p>
              <p>商品に不備がある場合は、大変お手数ですが「<NuxtLink to="/terms" class="faq-link">注意事項</NuxtLink>」内の「返品について」をご確認の上、商品到着後7日以内に「<NuxtLink to="/contact" class="faq-link">お問い合わせ</NuxtLink>」よりご連絡ください。</p>
            </div>
          </div>
        </div>
      </section>

      <!-- 配送・配送手数料について -->
      <section class="faq-section">
        <div class="faq-section-header">
          <svg class="faq-section-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M1 3h15v13H1zM16 8h4l3 3v5h-7V8z"/>
            <path d="M1 9l9-7 9 7"/>
          </svg>
          <h2 class="faq-section-title">配送・配送手数料について</h2>
        </div>
        <div class="faq-items">
          <div class="faq-item">
            <p class="faq-question">■配送手数料はどのくらいかかるの?</p>
            <div class="faq-answer">
              <p>配送手数料は購入ごとに発生します。配送手数料は全国一律となります。配送手数料は各商品の価格とあわせて記載されています。</p>
              <p>1回の購入で購入できる商品の上限は20個です。複数の商品を購入される場合は、1回の取引で「まとめ買い」として購入していただくと配送手数料がお得になります。</p>
              <p>※購入商品の合計が20個以下でも、購入回数が複数回の場合は購入回数分の配送手数料が発生しますのでご注意ください。</p>
            </div>
          </div>
          <div class="faq-item">
            <p class="faq-question">■商品は購入ごとに送られて来るの?</p>
            <p class="faq-answer">当選した商品によっては、購入回数より配送回数が多くなる場合や、購入回数より少ない配送回数にまとめられる場合があります。また、郵便や宅配便など複数の配送方法が混在する場合があります。</p>
          </div>
          <div class="faq-item">
            <p class="faq-question">■配送はいつになるの?</p>
            <div class="faq-answer">
              <p>商品のお届け目安は「くじを引く」ボタンの上に記載されています。配送が完了した際は、登録されているメールアドレス宛に「商品発送通知メール」を送信いたします。</p>
              <p>複数の商品が配送される場合、配送日が同日にならない場合や、商品によっては若干遅れる場合があります。ご了承ください。</p>
            </div>
          </div>
          <div class="faq-item">
            <p class="faq-question">■購入完了後に届け先住所を変更したい</p>
            <p class="faq-answer">購入した商品の販売期間中は届け先住所の変更が可能です。期間中は、会員情報変更ページで設定した最新の届け先住所に商品が配送されます。</p>
          </div>
          <div class="faq-item">
            <p class="faq-question">■運送会社の保管期間が過ぎてしまった商品の再配送をしてほしい</p>
            <p class="faq-answer">「<NuxtLink to="/terms" class="faq-link">注意事項</NuxtLink>」内の「商品の再配送について」「商品の保管期間について」をご確認ください。</p>
          </div>
        </div>
      </section>

      <!-- くじコレについて -->
      <section class="faq-section">
        <div class="faq-section-header">
          <svg class="faq-section-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/>
            <polyline points="3.27 6.96 12 12.01 20.73 6.96"/>
            <line x1="12" y1="22.08" x2="12" y2="12"/>
          </svg>
          <h2 class="faq-section-title">くじコレについて</h2>
        </div>
        <div class="faq-items">
          <div class="faq-item">
            <p class="faq-question">■パソコン、スマートフォン、フィーチャーフォンで利用できるの?</p>
            <p class="faq-answer">PCまたはスマートフォンからご利用いただけます。フィーチャーフォンには対応しておりません。<br>※一部のPCやスマートフォンの機種には対応していない場合がありますのでご了承ください。</p>
          </div>
          <div class="faq-item">
            <p class="faq-question">■利用可能な年齢は?</p>
            <p class="faq-answer">未成年者は、法定代理人（保護者等、以下同じ）の同意なく本サービスを利用することはできません。ただし、未成年者が商品を購入した場合、当該未成年者が本サービスの利用規約に同意したものとみなし、この場合、法定代理人の同意があったものとみなします。未成年者が購入した場合は、法定代理人の許可を得たものとみなし、購入完了後のキャンセルや返品はお受けできません。ご了承ください。</p>
          </div>
          <div class="faq-item">
            <p class="faq-question">■商品・サービスに関する問い合わせ方法は?</p>
            <p class="faq-answer">商品および本サービスに関するお問い合わせは「<NuxtLink to="/contact" class="faq-link">こちら</NuxtLink>」よりお願いいたします。</p>
          </div>
          <div class="faq-item">
            <p class="faq-question">■海外対応はしているの?</p>
            <p class="faq-answer">本サービスは日本国内のみの対応となります。ご了承ください。</p>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
// FAQページのスクリプト
</script>

<style scoped>
.faq-page {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  min-height: 60vh;
  padding: 40px 20px;
}

.faq-container {
  background-color: white;
  border-radius: 8px;
  max-width: 600px;
  width: 100%;
  padding: 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.faq-header {
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
  margin-right: 10px;
  flex-shrink: 0;
}

.faq-title {
  font-size: 1.5rem;
  font-weight: bold;
  color: #333;
  margin: 0;
}

.faq-section {
  margin-bottom: 50px;
}

.faq-section-header {
  display: flex;
  align-items: center;
  gap: 15px;
  margin-bottom: 25px;
  padding-bottom: 15px;
  border-bottom: 2px solid #e0e0e0;
}

.faq-section-icon {
  width: 32px;
  height: 32px;
  color: #c30000;
  flex-shrink: 0;
}

.faq-section-title {
  font-size: 1.5rem;
  font-weight: bold;
  color: #333;
  margin: 0;
}

.faq-items {
  display: flex;
  flex-direction: column;
  gap: 25px;
}

.faq-item {
  padding: 20px;
  background-color: #f8f9fa;
  border-radius: 8px;
  border-left: 4px solid #c30000;
}

.faq-question {
  font-size: 1.1rem;
  font-weight: 600;
  color: #333;
  margin: 0 0 12px 0;
}

.faq-answer {
  font-size: 0.95rem;
  color: #666;
  line-height: 1.8;
  margin: 0;
}

.faq-answer p {
  margin: 0 0 10px 0;
}

.faq-answer p:last-child {
  margin-bottom: 0;
}

.faq-list {
  margin: 10px 0;
  padding-left: 25px;
  list-style-type: decimal;
}

.faq-list li {
  margin-bottom: 8px;
  line-height: 1.6;
}

.faq-link {
  color: #0066cc;
  text-decoration: underline;
  font-weight: 500;
}

.faq-link:hover {
  color: #0052a3;
}

@media screen and (max-width: 767px) {
  .faq-page {
    padding: 20px 10px;
  }

  .faq-container {
    padding: 20px;
  }

  .faq-title {
    font-size: 1.25rem;
    margin-bottom: 30px;
  }

  .faq-section-title {
    font-size: 1.25rem;
  }

  .faq-item {
    padding: 15px;
  }

  .faq-question {
    font-size: 1rem;
  }

  .faq-answer {
    font-size: 0.9rem;
  }
}
</style>

