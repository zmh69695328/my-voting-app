<script setup>
import {onMounted} from 'vue'
import Ranking from '@/components/ranking/Ranking.vue'
const groups = ['最佳落地奖','人工智能','数据赋能','融合创新']
onMounted(async () => {
  // debugger
  // scrolltoElement()
  const targetElement = document.querySelector('.contexttitle');
    if (targetElement) {
        setTimeout(() => {
            targetElement.scrollIntoView({
                behavior: 'smooth', // 使用平滑滚动效果
                block: 'start',     // 滚动到元素的顶部
            })
        }, 150);

    }
})
</script>
<template>
  <div class="contexttitle"></div>
  <Ranking v-for="group in groups" :key="group" :group="group"/>
</template>
