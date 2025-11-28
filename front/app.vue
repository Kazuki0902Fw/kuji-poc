<template>
  <div class="app-container">
    <nav class="navbar">
      <div class="navbar-container">
        <div class="navbar-header">
          <div class="navbar-brand">
            <NuxtLink to="/" class="brand-link">くじコレ</NuxtLink>
          </div>
          <button 
            class="hamburger-button"
            :class="{ active: isMenuOpen }"
            @click="toggleMenu"
            aria-label="メニューを開く"
          >
            <span class="hamburger-line"></span>
            <span class="hamburger-line"></span>
            <span class="hamburger-line"></span>
          </button>
        </div>
        <ul class="navbar-menu">
          <li class="navbar-item">
            <NuxtLink to="/history" class="navbar-link">
              <svg class="navbar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M9 5H7a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2M9 5a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2M9 5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2"/>
              </svg>
              <span>くじ結果一覧</span>
            </NuxtLink>
          </li>
          <li class="navbar-item">
            <NuxtLink :to="isLoggedIn ? '/member' : '/login'" class="navbar-link">
              <svg class="navbar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                <circle cx="12" cy="7" r="4"/>
              </svg>
              <span>会員情報</span>
            </NuxtLink>
          </li>
          <li class="navbar-item">
            <NuxtLink to="/faq" class="navbar-link">
              <svg class="navbar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <circle cx="12" cy="12" r="10"/>
                <path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/>
                <line x1="12" y1="17" x2="12.01" y2="17"/>
              </svg>
              <span>よくある質問</span>
            </NuxtLink>
          </li>
          <li class="navbar-item">
            <NuxtLink to="/terms" class="navbar-link">
              <svg class="navbar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                <polyline points="14 2 14 8 20 8"/>
                <line x1="16" y1="13" x2="8" y2="13"/>
                <line x1="16" y1="17" x2="8" y2="17"/>
                <polyline points="10 9 9 9 8 9"/>
              </svg>
              <span>注意事項</span>
            </NuxtLink>
          </li>
          <li class="navbar-item">
            <NuxtLink to="/contact" class="navbar-link">
              <svg class="navbar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/>
                <polyline points="22,6 12,13 2,6"/>
              </svg>
              <span>お問い合わせ</span>
            </NuxtLink>
          </li>
          <li class="navbar-item" v-if="!isLoggedIn">
            <NuxtLink to="/login" class="navbar-link">
              <svg class="navbar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4"/>
                <polyline points="10 17 15 12 10 7"/>
                <line x1="15" y1="12" x2="3" y2="12"/>
              </svg>
              <span>ログイン</span>
            </NuxtLink>
          </li>
          <li class="navbar-item" v-if="isLoggedIn">
            <button @click="handleLogout" class="navbar-link logout-button">
              <svg class="navbar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>
                <polyline points="16 17 21 12 16 7"/>
                <line x1="21" y1="12" x2="9" y2="12"/>
              </svg>
              <span>ログアウト</span>
            </button>
          </li>
        </ul>
      </div>
    </nav>
    
    <!-- ドロワーオーバーレイ -->
    <div 
      v-if="isMenuOpen"
      class="drawer-overlay"
      @click="closeMenu"
    ></div>
    
    <!-- ドロワーメニュー -->
    <div 
      class="drawer-menu"
      :class="{ active: isMenuOpen }"
    >
      <ul class="drawer-menu-list">
        <li class="drawer-menu-item">
          <NuxtLink to="/history" class="drawer-menu-link" @click="closeMenu">
            <svg class="drawer-menu-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M9 5H7a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2M9 5a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2M9 5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2"/>
            </svg>
            <span>くじ結果一覧</span>
          </NuxtLink>
        </li>
        <li class="drawer-menu-item">
          <NuxtLink :to="isLoggedIn ? '/member' : '/login'" class="drawer-menu-link" @click="closeMenu">
            <svg class="drawer-menu-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
              <circle cx="12" cy="7" r="4"/>
            </svg>
            <span>会員情報</span>
          </NuxtLink>
        </li>
        <li class="drawer-menu-item">
          <NuxtLink to="/faq" class="drawer-menu-link" @click="closeMenu">
            <svg class="drawer-menu-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="12" cy="12" r="10"/>
              <path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/>
              <line x1="12" y1="17" x2="12.01" y2="17"/>
            </svg>
            <span>よくある質問</span>
          </NuxtLink>
        </li>
        <li class="drawer-menu-item">
          <NuxtLink to="/terms" class="drawer-menu-link" @click="closeMenu">
            <svg class="drawer-menu-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
              <polyline points="14 2 14 8 20 8"/>
              <line x1="16" y1="13" x2="8" y2="13"/>
              <line x1="16" y1="17" x2="8" y2="17"/>
              <polyline points="10 9 9 9 8 9"/>
            </svg>
            <span>注意事項</span>
          </NuxtLink>
        </li>
        <li class="drawer-menu-item">
          <NuxtLink to="/contact" class="drawer-menu-link" @click="closeMenu">
            <svg class="drawer-menu-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/>
              <polyline points="22,6 12,13 2,6"/>
            </svg>
            <span>お問い合わせ</span>
          </NuxtLink>
        </li>
        <li class="drawer-menu-item" v-if="!isLoggedIn">
          <NuxtLink to="/login" class="drawer-menu-link" @click="closeMenu">
            <svg class="drawer-menu-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4"/>
              <polyline points="10 17 15 12 10 7"/>
              <line x1="15" y1="12" x2="3" y2="12"/>
            </svg>
            <span>ログイン</span>
          </NuxtLink>
        </li>
        <li class="drawer-menu-item" v-if="isLoggedIn">
          <button @click="handleLogout" class="drawer-menu-link logout-button">
            <svg class="drawer-menu-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>
              <polyline points="16 17 21 12 16 7"/>
              <line x1="21" y1="12" x2="9" y2="12"/>
            </svg>
            <span>ログアウト</span>
          </button>
        </li>
      </ul>
    </div>
    
    <main class="main-content">
      <NuxtPage />
    </main>

    <!-- フッター -->
    <footer class="footer">
      <div class="footer-container">
        <NuxtLink to="/" class="footer-link">TOP</NuxtLink>
        <NuxtLink to="/terms-of-service" class="footer-link">利用規約</NuxtLink>
        <NuxtLink to="/commercial-transaction-law" class="footer-link">特定商取引法表示</NuxtLink>
        <NuxtLink to="/privacy-policy" class="footer-link">個人情報保護方針</NuxtLink>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { isLoggedIn as checkIsLoggedIn } from '~/utils/tokenStore';

const isMenuOpen = ref(false);

const { data: meData, error: meError } = await useAsyncGql("me", null, {
  getCachedData: () => undefined,
  clientId: 'default'
});

const isLoggedIn = computed(() => {
  return checkIsLoggedIn();
});

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value;
};

const closeMenu = () => {
  isMenuOpen.value = false;
};

const handleLogout = async () => {
  try {
    await useAsyncGql("logout", null, {
      getCachedData: () => undefined,
      clientId: 'default'
    });
    
    const { clearTokenFromLocalStorage } = await import('~/utils/tokenStore');
    clearTokenFromLocalStorage();
    
    useGqlToken(null);
    
    await navigateTo('/');
    window.location.reload();
  } catch (error) {
    console.error('ログアウトに失敗しました:', error);
    const { clearTokenFromLocalStorage } = await import('~/utils/tokenStore');
    clearTokenFromLocalStorage();
    useGqlToken(null);
    await navigateTo('/');
    window.location.reload();
  }
};
</script>

<style scoped>
.app-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.navbar {
  background-color: #c30000;
  color: white;
  padding: 0;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  position: sticky;
  top: 0;
  z-index: 1001;
}

.navbar-container {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 50px;
}

.navbar-header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}

.navbar-brand {
  min-width: 96px;
  padding-left: 20px;
  font-size: 1.5rem;
  font-weight: bold;
}

.brand-link {
  color: white;
  text-decoration: none;
  transition: opacity 0.3s;
}

.brand-link:hover {
  opacity: 0.8;
}

.hamburger-button {
  display: none;
  flex-direction: column;
  justify-content: space-around;
  width: 30px;
  height: 30px;
  background: transparent;
  border: none;
  cursor: pointer;
  padding: 0;
  z-index: 10;
}

.hamburger-line {
  width: 100%;
  height: 3px;
  background-color: white;
  border-radius: 3px;
  transition: all 0.3s ease;
}

.hamburger-button.active .hamburger-line:nth-child(1) {
  transform: rotate(45deg) translate(8px, 8px);
}

.hamburger-button.active .hamburger-line:nth-child(2) {
  opacity: 0;
}

.hamburger-button.active .hamburger-line:nth-child(3) {
  transform: rotate(-45deg) translate(7px, -7px);
}

.navbar-menu {
  display: flex;
  list-style: none;
  margin: 0;
  padding: 0;
  gap: 20px;
}

.navbar-item {
  margin: 0;
}

.navbar-link {
  color: white;
  text-decoration: none;
  padding: 8px 16px;
  border-radius: 4px;
  transition: background-color 0.3s;
  display: flex;
  align-items: center;
  gap: 8px;
}

.navbar-icon {
  width: 20px;
  height: 20px;
  flex-shrink: 0;
}

.logout-button {
  background: none;
  border: none;
  cursor: pointer;
  width: 100%;
  text-align: left;
}

.navbar-link:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.navbar-link.router-link-active {
  background-color: rgba(255, 255, 255, 0.2);
  font-weight: 600;
}

.main-content {
  flex: 1;
  max-width: 1200px;
  width: 100%;
  margin: 0 auto;
  padding: 20px;
}

/* ドロワーオーバーレイ */
.drawer-overlay {
  display: none;
}

/* ドロワーメニュー */
.drawer-menu {
  display: none;
}

/* レスポンシブ対応 - ハンバーガーメニュー（ドロワー形式） */
@media screen and (max-width: 920px) {
  .navbar-container {
    position: relative;
    flex-direction: row;
    justify-content: space-between;
  }

  .navbar-header {
    width: 100%;
  }

  .navbar-brand {
    width: 100%;
  }

  .hamburger-button {
    display: flex;
    position: relative;
    z-index: 1002;
    margin-right: 20px;
    margin-top: 3px;
  }

  .navbar-menu {
    display: none;
  }

  /* ドロワーオーバーレイ */
  .drawer-overlay {
    display: block;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 999;
    animation: fadeIn 0.3s ease;
  }

  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }

  /* ドロワーメニュー */
  .drawer-menu {
    display: block;
    position: fixed;
    top: 0;
    right: -300px;
    width: 300px;
    max-width: 80%;
    height: 100%;
    background-color: #c30000;
    box-shadow: -2px 0 8px rgba(0, 0, 0, 0.2);
    z-index: 1000;
    transition: right 0.3s ease;
    overflow-y: auto;
  }

  .drawer-menu.active {
    right: 0;
  }

  .drawer-menu-list {
    list-style: none;
    margin: 0;
    padding: 50px 0 20px 0;
  }

  .drawer-menu-item {
    border-bottom: 1px solid white;
  }

  .drawer-menu-item:first-child {
    border-top: 1px solid white;
  }

  .drawer-menu-link {
    display: flex;
    align-items: center;
    gap: 12px;
    color: white;
    text-decoration: none;
    padding: 20px 30px;
    transition: background-color 0.3s;
    font-size: 1rem;
    width: 100%;
    text-align: left;
    background: none;
    border: none;
    cursor: pointer;
  }

  .drawer-menu-icon {
    width: 24px;
    height: 24px;
    flex-shrink: 0;
  }

  .drawer-menu-link:hover {
    background-color: rgba(255, 255, 255, 0.1);
  }

  .drawer-menu-link.router-link-active {
    background-color: rgba(255, 255, 255, 0.2);
    font-weight: 600;
  }
}

/* フッター */
.footer {
  background-color: #333;
  color: white;
  padding: 30px 20px;
  margin-top: auto;
}

.footer-container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 30px;
  flex-wrap: wrap;
}

.footer-link {
  color: white;
  text-decoration: none;
  font-size: 0.95rem;
  transition: opacity 0.3s;
}

.footer-link:hover {
  opacity: 0.7;
}

@media screen and (max-width: 767px) {
  .footer-container {
    flex-direction: column;
    gap: 15px;
  }
}
</style>
