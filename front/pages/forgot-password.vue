<template>
  <div class="forgot-password-page">
    <div class="forgot-password-container">
      <!-- パスワードを忘れた場合 -->
      <section class="forgot-section">
        <h2 class="section-title">
          <span class="section-bar"></span>
          パスワードを忘れた場合
        </h2>
        <div class="section-content">
          <p class="instruction-text">
            新しいパスワードを設定するURLを登録メールアドレスに送信致します。<br>
            くじコレIDとして登録されているメールアドレスを入力してください。
          </p>
          <div class="form-group">
            <label for="email" class="form-label">
              <svg class="form-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/>
                <polyline points="22,6 12,13 2,6"/>
              </svg>
              メールアドレス
            </label>
            <input
              id="email"
              v-model="email"
              type="email"
              class="form-input"
              placeholder="メールアドレス"
              required
            />
          </div>
          <ul class="notice-list">
            <li>※パスワードを再設定すると現在のパスワードは破棄されます。</li>
            <li>※「<span class="domain-text">@kujico.jp</span>」ドメインからのメールを受信できるように設定してください。</li>
          </ul>
          <div class="button-group">
            <button type="button" class="back-button" @click="handleBack">
              戻る
            </button>
            <button type="button" class="send-button" @click="handleSendEmail" :disabled="!email || isSending">
              {{ isSending ? '送信中...' : 'メール送信' }}
            </button>
          </div>
        </div>
      </section>

      <!-- IDを忘れた場合 -->
      <section class="forgot-section">
        <h2 class="section-title">
          <span class="section-bar"></span>
          IDを忘れた場合
        </h2>
        <div class="section-content">
          <p class="instruction-text">
            IDはメールアドレスとなっております。<br>
            恐れ入りますがお手持ちのメールアドレスをご確認ください。
          </p>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const email = ref('')
const isSending = ref(false)

const handleBack = () => {
  navigateTo('/login')
}

const handleSendEmail = async () => {
  if (!email.value) {
    alert('メールアドレスを入力してください。')
    return
  }

  isSending.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    alert('パスワード再設定用のメールを送信しました。')
    email.value = ''
  } catch (error) {
    console.error('メール送信エラー:', error)
    alert('メールの送信に失敗しました。もう一度お試しください。')
  } finally {
    isSending.value = false
  }
}
</script>

<style scoped>
.forgot-password-page {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  min-height: 60vh;
  padding: 40px 20px;
}

.forgot-password-container {
  background-color: white;
  border-radius: 8px;
  max-width: 600px;
  width: 100%;
  padding: 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.forgot-section {
  margin-bottom: 40px;
}

.forgot-section:last-child {
  margin-bottom: 0;
}

.section-title {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 1.3rem;
  font-weight: bold;
  color: #333;
  margin: 0 0 15px 0;
}

.section-bar {
  width: 6px;
  height: 30px;
  background-color: #c30000;
  flex-shrink: 0;
}

.section-content {
  color: #333;
  line-height: 1.8;
  font-size: 0.95rem;
}

.instruction-text {
  font-size: 0.95rem;
  color: #333;
  line-height: 1.8;
  margin: 0 0 20px 0;
}

.form-group {
  margin-bottom: 20px;
}

.form-label {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.9rem;
  font-weight: 500;
  color: #333;
  margin-bottom: 8px;
}

.form-icon {
  width: 18px;
  height: 18px;
  color: #666;
}

.form-input {
  width: 100%;
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  box-sizing: border-box;
}

.form-input:focus {
  outline: none;
  border-color: #ff9900;
}

.notice-list {
  margin: 20px 0;
  padding-left: 20px;
  list-style-type: disc;
}

.notice-list li {
  margin-bottom: 8px;
  line-height: 1.6;
  color: #333;
  font-size: 0.9rem;
}

.notice-list li:last-child {
  margin-bottom: 0;
}

.domain-text {
  color: #c30000;
  font-weight: 600;
}

.button-group {
  display: flex;
  gap: 15px;
  margin-top: 30px;
}

.back-button {
  flex: 1;
  padding: 12px 24px;
  background-color: #e0e0e0;
  color: #333;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.3s;
}

.back-button:hover {
  background-color: #d0d0d0;
}

.send-button {
  flex: 1;
  padding: 12px 24px;
  background-color: #e91e63;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.3s;
}

.send-button:hover:not(:disabled) {
  background-color: #c2185b;
}

.send-button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

@media screen and (max-width: 767px) {
  .forgot-password-page {
    padding: 20px 10px;
  }

  .forgot-password-container {
    padding: 20px;
  }

  .section-title {
    font-size: 1.1rem;
  }

  .button-group {
    flex-direction: column;
  }
}
</style>

