<template>
  <div class="banner-container">
    <!-- Banner Image and Content -->
    <div class="banner-content">
      <!-- Background Image -->
      <div class="banner-background" :style="bannerStyle"></div>
      
      <!-- SAMPLE Watermark -->
      <div class="sample-watermark">SAMPLE</div>

      <div class="banner-overlay">
        <!-- Status Button -->
        <div class="status-button-container">
          <div class="status-button" :class="[`status-${status.type}`]">
            {{ status.text }}
          </div>
        </div>

        <!-- Title and Date Container -->
        <div class="banner-content-right">
          <!-- Title -->
          <div class="banner-title">
            {{ category.name }}
          </div>

          <!-- Date Info -->
          <div class="banner-date">
            {{ formatDate(category.salesEndDate) }}まで
          </div>
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
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.banner-content {
  position: relative;
  width: 100%;
  height: 300px;
  border-radius: 8px;
  overflow: hidden;
}

.banner-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-size: cover;
  background-position: center;
  border-radius: 8px;
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
  padding: 12px 20px;
  background-color: black;
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 15px;
  z-index: 6;
}

.status-button-container {
  display: flex;
  flex-shrink: 0;
}

.banner-content-right {
  display: flex;
  flex-direction: column;
  gap: 2px;
  flex: 1;
  justify-content: center;
}

.status-button {
  padding: 10px 12px;
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
  font-size: 1.1rem;
  font-weight: bold;
  color: white;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.8);
  line-height: 1.2;
}

.banner-date {
  font-size: 0.85rem;
  color: white;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.8);
  line-height: 1.2;
}

@media screen and (max-width: 767px) {
  .banner-content {
    height: 200px;
  }

  .banner-overlay {
    padding: 10px 15px;
  }

  .banner-title {
    font-size: 1rem;
  }

  .banner-date {
    font-size: 0.75rem;
  }
}
</style>


