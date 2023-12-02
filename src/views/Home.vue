<script setup>
import { useRoute } from 'vue-router'
const route = useRoute()
import { onMounted, onUpdated,watch } from 'vue'
import { ref } from 'vue'
console.log(route.query.flag,111)
const flag = ref(false)
// const {query:{flag}}=route

watch(
      () => route.query.flag,
      () => {
        console.log('侦听到 ID 变化')
        flag.value=true
      }
    )
</script>
<template>
    <div v-show="flag===true" role="alert" class="alert absolute inset-x-0 top-3 w-1/2 center mx-auto">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-info shrink-0 w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        <span>感谢您的支持与参与！</span>
        <!-- <div>
      <button class="btn btn-sm" @click="showAlert = false;">取消</button>
      <button class="btn btn-sm btn-primary" @click="afterSubmit">离开</button>
    </div> -->
    </div>
    <div class="flex w-full justify-center mt-40 mb-8">

        <!-- Your HTML template code here -->
        <div class="w-1/2 flexjustify-center">
            <div class="flex flex-col justify-center w-full">
                <!-- <div class="grid h-20 card bg-base-300 rounded-box place-items-center">第一轮</div> -->
                <button class="btn" @click="$router.push('/round1')">上届落地项目</button>

                <div class="divider"></div>
                <!-- <div class="grid h-20 card bg-base-300 rounded-box place-items-center">第二轮</div> -->
                <button class="btn" @click="$router.push('/round2')">决赛项目</button>

            </div>

        </div>
    </div>
    <div class="h-3"></div>
</template>
