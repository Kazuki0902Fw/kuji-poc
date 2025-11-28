<template>
  <div class="login-page">
    <!-- エラーバナー -->
    <div v-if="errorMessage" class="error-banner">
      {{ errorMessage }}
    </div>
    <div class="login-card">
      <h1 class="login-title">ログイン</h1>
      <form @submit.prevent="handleLogin" class="login-form">
        <div class="form-group">
          <label for="mailAddress" class="form-label">メールアドレス</label>
          <input
            id="mailAddress"
            v-model="mailAddress"
            type="email"
            class="form-input"
            required
            placeholder="メールアドレスを入力"
          />
        </div>
        <div class="form-group">
          <label for="password" class="form-label">パスワード</label>
          <input
            id="password"
            v-model="password"
            type="password"
            class="form-input"
            required
            placeholder="パスワードを入力"
          />
        </div>
        <button type="submit" class="login-button" :disabled="isLoading">
          {{ isLoading ? 'ログイン中...' : 'ログイン' }}
        </button>
      </form>
      <div class="forgot-link-container">
        <NuxtLink to="/forgot-password" class="forgot-link">
          ID・パスワードを忘れた場合 >>
        </NuxtLink>
      </div>
      <div class="register-section">
        <h2 class="register-title">新規会員登録</h2>
        <p class="register-text">初めてご利用の方はこちらから会員登録をお願いいたします。</p>
        <NuxtLink to="/register" class="register-button">新規会員登録</NuxtLink>
      </div>
      <div class="amazon-pay-section">
        <h2 class="amazon-pay-title">Amazon Payでログイン</h2>
        <button class="amazon-pay-button">Amazon Payでログイン</button>
      </div>
      <div class="disclaimer">
        <p>※ログインすることで、<NuxtLink to="/terms-of-service" class="disclaimer-link">利用規約</NuxtLink>および<NuxtLink to="/privacy-policy" class="disclaimer-link">個人情報保護方針</NuxtLink>に同意したものとみなされます。</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { setTokenToLocalStorage } from '~/utils/tokenStore'

const mailAddress = ref('')
const password = ref('')
const isLoading = ref(false)
const errorMessage = ref('')

const handleLogin = async () => {
  if (!mailAddress.value || !password.value) {
    errorMessage.value = 'メールアドレスとパスワードを入力してください。'
    return
  }

  isLoading.value = true
  errorMessage.value = ''

  try {
    const { data, error } = await useAsyncGql('login', {
      input: {
        mailAddress: mailAddress.value,
        password: password.value
      }
    }, {
      getCachedData: () => undefined,
      clientId: 'default'
    })

    if (error.value?.cause?.gqlErrors?.[0]?.message) {
      errorMessage.value = "ログインに失敗しました。"
    }

    if (data.value && data.value.login) {
      setTokenToLocalStorage(data.value.login)
      useGqlToken(data.value.login.accessToken)
      await navigateTo('/')
      window.location.reload()
    }
  } catch (error) {
    console.error('ログインエラー:', error)
    errorMessage.value = 'ログインに失敗しました。もう一度お試しください。'
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.login-page {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  min-height: 60vh;
  padding: 40px 20px;
  background-color: #f5f5f5;
}

.error-banner {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background-color: #f44336;
  color: white;
  padding: 15px;
  text-align: center;
  z-index: 1000;
  font-weight: 500;
}

.login-card {
  background-color: white;
  border-radius: 8px;
  max-width: 500px;
  width: 100%;
  padding: 40px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-top: 60px;
}

.login-title {
  font-size: 2rem;
  font-weight: bold;
  color: #333;
  margin-bottom: 30px;
  text-align: center;
}

.login-form {
  margin-bottom: 30px;
}

.form-group {
  margin-bottom: 20px;
}

.form-label {
  display: block;
  font-size: 0.95rem;
  font-weight: 500;
  color: #333;
  margin-bottom: 8px;
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

.login-button {
  width: 100%;
  padding: 15px;
  background-color: #ff9900;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1.1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.3s;
}

.login-button:hover:not(:disabled) {
  background-color: #e68900;
}

.login-button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.forgot-link-container {
  text-align: center;
  margin-bottom: 30px;
}

.forgot-link {
  color: #0066cc;
  text-decoration: underline;
  font-size: 0.9rem;
}

.register-section {
  border-top: 1px solid #e0e0e0;
  padding-top: 30px;
  margin-bottom: 30px;
}

.register-title {
  font-size: 1.3rem;
  font-weight: bold;
  color: #333;
  margin-bottom: 10px;
}

.register-text {
  font-size: 0.95rem;
  color: #666;
  margin-bottom: 20px;
}

.register-button {
  display: block;
  width: 100%;
  padding: 15px;
  background-color: #0066cc;
  color: white;
  text-decoration: none;
  text-align: center;
  border-radius: 4px;
  font-size: 1.1rem;
  font-weight: 500;
  transition: background-color 0.3s;
}

.register-button:hover {
  background-color: #0052a3;
}

.amazon-pay-section {
  border-top: 1px solid #e0e0e0;
  padding-top: 30px;
  margin-bottom: 30px;
}

.amazon-pay-title {
  font-size: 1.3rem;
  font-weight: bold;
  color: #333;
  margin-bottom: 20px;
}

.amazon-pay-button {
  width: 100%;
  padding: 15px;
  background-color: #ff9900;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1.1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.3s;
}

.amazon-pay-button:hover {
  background-color: #e68900;
}

.disclaimer {
  border-top: 1px solid #e0e0e0;
  padding-top: 20px;
  text-align: center;
}

.disclaimer p {
  font-size: 0.85rem;
  color: #666;
  line-height: 1.6;
  margin: 0;
}

.disclaimer-link {
  color: #0066cc;
  text-decoration: underline;
}

@media screen and (max-width: 767px) {
  .login-page {
    padding: 20px 10px;
  }

  .login-card {
    padding: 20px;
  }

  .login-title {
    font-size: 1.5rem;
  }
}
</style>

