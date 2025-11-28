<template>
  <div class="banner-container">
    <!-- Branding -->
    <div class="branding">
      <div class="branding-box">
        <svg class="branding-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M12 2C8.13 2 5 5.13 5 9c0 5.25 7 13 7 13s7-7.75 7-13c0-3.87-3.13-7-7-7z"/>
        </svg>
        <span class="branding-text">&lt;くじコレ&gt;</span>
      </div>
    </div>

    <!-- Banner Image and Content -->
    <div class="banner-content" :style="bannerStyle">
      <!-- SAMPLE Watermark -->
      <div class="sample-watermark">SAMPLE</div>

      <div class="banner-overlay">
        <!-- Status Button -->
        <div class="status-button-container">
          <div class="status-button" :class="[`status-${status.type}`]">
            {{ status.text }}
          </div>
        </div>

        <!-- Title -->
        <div class="banner-title">
          {{ category.name }}
        </div>

        <!-- Date Info -->
        <div class="banner-date">
          {{ formatDate(category.salesStartDate) }}～{{ formatDate(category.salesEndDate) }}
        </div>

        <!-- Copyright -->
        <div v-if="category.precautions" class="banner-copyright">
          {{ category.precautions }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  category: {
    type: Object,
    required: true
  }
})

const status = computed(() => {
  const now = new Date()
  const startDate = new Date(props.category.salesStartDate)
  const endDate = new Date(props.category.salesEndDate)

  if (now < startDate) {
    return { type: 'upcoming', text: '販売予定' }
  } else if (now <= endDate) {
    return { type: 'active', text: '販売中' }
  } else {
    return { type: 'ended', text: '販売終了' }
  }
})

const bannerStyle = computed(() => {
  return {
    backgroundImage: `url(${props.category.imgUrl})`,
    backgroundSize: 'cover',
    backgroundPosition: 'center'
  }
})

const formatDate = (dateStr) => {
  const date = new Date(dateStr)
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hour = String(date.getHours()).padStart(2, '0')
  const minute = String(date.getMinutes()).padStart(2, '0')
  
  return `${year}年${month}月${day}日${hour}:${minute}`
}
</script>

<style scoped>
.banner-container {
  position: relative;
  width: 100%;
  margin-bottom: 20px;
}

.branding {
  position: absolute;
  top: 10px;
  left: 10px;
  z-index: 10;
}

.branding-box {
  display: flex;
  align-items: center;
  gap: 8px;
  background-color: #ff9900;
  padding: 8px 12px;
  border-radius: 4px;
}

.branding-icon {
  width: 20px;
  height: 20px;
  color: white;
}

.branding-text {
  font-size: 0.9rem;
  font-weight: bold;
  color: white;
}

.banner-content {
  position: relative;
  width: 100%;
  height: 300px;
  border-radius: 8px;
  overflow: hidden;
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
  font-size: clamp(1.5rem, 8vw, 4rem);
  font-weight: bold;
  color: rgba(0, 0, 0, 0.15);
  transform: rotate(-45deg);
  letter-spacing: clamp(0.1rem, 1vw, 0.5rem);
  white-space: nowrap;
}

.banner-overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 20px;
  background: linear-gradient(to top, rgba(0, 0, 0, 0.8), transparent);
  display: flex;
  flex-direction: column;
  gap: 8px;
  z-index: 6;
}

.status-button-container {
  display: flex;
  justify-content: flex-end;
}

.status-button {
  padding: 6px 12px;
  border-radius: 4px;
  font-size: 0.9rem;
  font-weight: bold;
  color: white;
}

.status-active {
  background-color: #c30000;
}

.status-upcoming {
  background-color: #0066cc;
}

.status-ended {
  background-color: #666;
}

.banner-title {
  font-size: 1.5rem;
  font-weight: bold;
  color: white;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.8);
}

.banner-date {
  font-size: 0.95rem;
  color: white;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.8);
}

.banner-copyright {
  font-size: 0.85rem;
  color: rgba(255, 255, 255, 0.9);
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.8);
}

@media screen and (max-width: 767px) {
  .banner-content {
    height: 200px;
  }

  .banner-overlay {
    padding: 15px;
  }

  .banner-title {
    font-size: 1.2rem;
  }

  .banner-date {
    font-size: 0.85rem;
  }

  .banner-copyright {
    font-size: 0.75rem;
  }
}
</style>

