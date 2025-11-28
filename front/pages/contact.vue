<template>
  <div class="contact-page">
    <div class="contact-container">
      <!-- ヘッダー -->
      <div class="contact-header">
        <div class="header-bar"></div>
        <h1 class="contact-title">お問い合わせ</h1>
      </div>

      <!-- FAQ案内 -->
      <div class="faq-notice">
        <p class="faq-notice-text">
          少しでも早い問題解決のため、お客様にはお問い合わせの前に、まずは「よくあるご質問」のご確認をお願いしております。
        </p>
        <NuxtLink to="/faq" class="faq-button">
          よくあるご質問はこちら
        </NuxtLink>
      </div>

      <!-- お問い合わせのご注意 -->
      <section class="notice-section">
        <h2 class="notice-title">お問い合わせのご注意</h2>
        <ul class="notice-list">
          <li>
            お問い合わせの前に、端末キャリア側の迷惑メールフィルターや受信拒否設定などにより、メールが受信できない設定になっていないか必ずご確認ください。メール受信拒否設定を行っている場合は、「<span class="domain-text">@kujico.jp</span>」ドメインからのメールが受信できるように設定してください。
          </li>
          <li>
            ご質問にはメールで返信いたします。ご入力いただくメールアドレスにお間違いがないようご注意ください。お問い合わせ受付後、確認メールが送信されます。端末の受信設定により確認メールが届かない場合、お問い合わせの回答も受信できませんので、ご注意ください。
          </li>
          <li>
            お問い合わせは24時間受け付けておりますが、混雑状況やご質問の内容により、返信までにお時間を頂く場合がございます。あらかじめご了承ください。
          </li>
          <li>
            お電話でのサポートは受け付けておりません。
          </li>
        </ul>
      </section>

      <!-- お問い合わせフォーム -->
      <form class="contact-form" @submit.prevent="handleSubmit">
        <!-- メールアドレス -->
        <div class="form-group">
          <label for="email" class="form-label">メールアドレス</label>
          <input
            id="email"
            v-model="formData.email"
            type="email"
            class="form-input"
            required
            placeholder="メールアドレスを入力してください"
          />
        </div>

        <!-- お問い合わせ内容 -->
        <div class="form-group">
          <label class="form-label">お問い合わせ内容</label>
          <div class="radio-group">
            <label class="radio-label">
              <input
                v-model="formData.inquiryType"
                type="radio"
                value="kujicole"
                class="radio-input"
                required
              />
              <span class="radio-text">くじコレ</span>
            </label>
            <label class="radio-label">
              <input
                v-model="formData.inquiryType"
                type="radio"
                value="print-kujicole"
                class="radio-input"
              />
              <span class="radio-text">プリントくじコレ</span>
            </label>
          </div>
          <textarea
            v-model="formData.content"
            class="form-textarea"
            rows="8"
            required
            placeholder="お問い合わせ内容を記載してください"
          ></textarea>
        </div>

        <!-- 添付ファイル -->
        <div class="form-group">
          <label class="form-label">添付ファイル(オプション)</label>
          <div class="file-upload">
            <input
              id="file-input"
              ref="fileInput"
              type="file"
              multiple
              accept="image/*"
              class="file-input"
              @change="handleFileChange"
            />
            <button
              type="button"
              class="file-button"
              @click="triggerFileInput"
            >
              ファイル選択
            </button>
            <span class="file-status">{{ fileStatus }}</span>
          </div>
          <ul class="file-notice-list">
            <li>
              商品不備でのお問い合わせの場合、不良箇所が分かる写真も一緒にお送りください。
            </li>
            <li>
              送信できるファイルは画像ファイルのみで、5個、20MBまでとなります。
            </li>
          </ul>
        </div>

        <!-- 送信ボタン -->
        <button type="submit" class="submit-button">
          内容を確認する
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const formData = ref({
  email: '',
  inquiryType: 'kujicole',
  content: '',
  files: []
})

const fileInput = ref(null)
const fileStatus = ref('選択されていません')

const triggerFileInput = () => {
  fileInput.value?.click()
}

const handleFileChange = (event) => {
  const files = Array.from(event.target.files || [])
  
  if (files.length > 5) {
    alert('ファイルは5個まで選択できます。')
    event.target.value = ''
    return
  }

  const totalSize = files.reduce((sum, file) => sum + file.size, 0)
  const maxSize = 20 * 1024 * 1024
  
  if (totalSize > maxSize) {
    alert('ファイルの合計サイズは20MBまでです。')
    event.target.value = ''
    return
  }

  const imageFiles = files.filter(file => file.type.startsWith('image/'))
  if (imageFiles.length !== files.length) {
    alert('画像ファイルのみ選択できます。')
    event.target.value = ''
    return
  }

  formData.value.files = files
  
  if (files.length === 0) {
    fileStatus.value = '選択されていません'
  } else if (files.length === 1) {
    fileStatus.value = `1個のファイルが選択されています`
  } else {
    fileStatus.value = `${files.length}個のファイルが選択されています`
  }
}

const handleSubmit = () => {
  console.log('Form submitted:', formData.value)
  alert('内容を確認しました。実際の送信処理は実装が必要です。')
}
</script>

<style scoped>
.contact-page {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  min-height: 60vh;
  padding: 40px 20px;
}

.contact-container {
  background-color: white;
  border-radius: 8px;
  max-width: 600px;
  width: 100%;
  padding: 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.contact-header {
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

.contact-title {
  font-size: 1.5rem;
  font-weight: bold;
  color: #333;
  margin: 0;
}

.faq-notice {
  margin-bottom: 30px;
  padding: 20px;
  background-color: #f8f9fa;
  border-radius: 8px;
}

.faq-notice-text {
  font-size: 0.95rem;
  color: #666;
  line-height: 1.8;
  margin: 0 0 15px 0;
}

.faq-button {
  display: inline-block;
  padding: 12px 24px;
  background-color: #28a745;
  color: white;
  text-decoration: none;
  border-radius: 4px;
  font-weight: 500;
  transition: background-color 0.3s;
}

.faq-button:hover {
  background-color: #218838;
}

.notice-section {
  margin-bottom: 30px;
  padding: 20px;
  background-color: #fff9e6;
  border-radius: 8px;
  border-left: 4px solid #ff9900;
}

.notice-title {
  font-size: 1.2rem;
  font-weight: bold;
  color: #333;
  margin: 0 0 15px 0;
}

.notice-list {
  margin: 0;
  padding-left: 20px;
  list-style-type: disc;
}

.notice-list li {
  margin-bottom: 12px;
  line-height: 1.8;
  color: #333;
  font-size: 0.95rem;
}

.notice-list li:last-child {
  margin-bottom: 0;
}

.domain-text {
  color: #c30000;
  font-weight: 600;
}

.contact-form {
  display: flex;
  flex-direction: column;
  gap: 25px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.form-label {
  font-size: 1rem;
  font-weight: bold;
  color: #333;
}

.form-input {
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  width: 100%;
  box-sizing: border-box;
}

.form-input:focus {
  outline: none;
  border-color: #ff9900;
}

.radio-group {
  display: flex;
  gap: 20px;
  margin-bottom: 10px;
}

.radio-label {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
}

.radio-input {
  width: 18px;
  height: 18px;
  cursor: pointer;
}

.radio-text {
  font-size: 1rem;
  color: #333;
}

.form-textarea {
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  width: 100%;
  box-sizing: border-box;
  resize: vertical;
  font-family: inherit;
}

.form-textarea:focus {
  outline: none;
  border-color: #ff9900;
}

.file-upload {
  display: flex;
  align-items: center;
  gap: 15px;
}

.file-input {
  display: none;
}

.file-button {
  padding: 10px 20px;
  background-color: #f8f9fa;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s;
}

.file-button:hover {
  background-color: #e9ecef;
}

.file-status {
  font-size: 0.95rem;
  color: #666;
}

.file-notice-list {
  margin: 10px 0 0 0;
  padding-left: 20px;
  list-style-type: disc;
}

.file-notice-list li {
  margin-bottom: 8px;
  line-height: 1.6;
  color: #666;
  font-size: 0.9rem;
}

.file-notice-list li:last-child {
  margin-bottom: 0;
}

.submit-button {
  padding: 15px 30px;
  background-color: #0066cc;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1.1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.3s;
  width: 100%;
}

.submit-button:hover {
  background-color: #0052a3;
}

.submit-button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

@media screen and (max-width: 767px) {
  .contact-page {
    padding: 20px 10px;
  }

  .contact-container {
    padding: 20px;
  }

  .contact-title {
    font-size: 1.25rem;
  }

  .radio-group {
    flex-direction: column;
    gap: 10px;
  }
}
</style>

