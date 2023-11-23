<script setup>
import Ranking from '@/components/ranking/Ranking.vue'
const groups = ['最佳落地奖','人工智能','数据赋能','融合创新']
</script>
<template>
  <Ranking v-for="group in groups" :group="group"/>
</template>
