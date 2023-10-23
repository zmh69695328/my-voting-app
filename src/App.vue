<script setup>
import List from './components/voting/List.vue'
import ListItem from './components/voting/ListItem.vue'
import Login from './components/login/Login.vue'
import Ranking from './components/ranking/Ranking.vue'
import { nextTick, onMounted, reactive, ref } from 'vue';
import { useUsernameStore } from '@/stores/username.js'
const showtabs=ref(2) //1是投票页，2是排行榜页
const login = ref()
function showLogin() {
  login.value.showModal = true
}
const store = useUsernameStore()
function scrollToElement() {
  
  if(store.username === ''|| store.username === undefined){
        showLogin()
    }
  const targetElement = document.getElementById('voting');
  if (targetElement) {
    setTimeout(() => {
      targetElement.scrollIntoView({
        behavior: 'smooth', // 使用平滑滚动效果
        block: 'start',     // 滚动到元素的顶部
      })
    }, 10);

  }
} let teamList = reactive([])
async function getTeams() {


  let response = await fetch('/api/teams', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json;charset=utf-8'
    }
  });

  let result = await response.json();
  console.log(result);
  teamList.length = 0;
  teamList.push(...result.teams);
}
onMounted(() => {
  getTeams()
})


function submit() {
  console.log(teamList)
}
</script>

<template>
  <Login ref="login"></Login>
  <div class="relative h-screen bg-indigo-900" :class="{ 'fullscreen': showtabs === 2 }">
    <img v-if="showtabs === 1" src="/images/6.svg" class="absolute object-cover w-full h-full" />
    <!-- <img v-if="showtabs === 2" src="/images/6.svg" class="absolute object-cover w-full h-[80rem]" /> -->
    <!-- <img v-if="showtabs === 2" src="/images/6.svg" class="absolute object-cover w-full h-screen" /> -->
    <div class="absolute inset-0 bg-black opacity-25"></div>
    <header class="absolute top-0 left-0 right-0 z-20">
      <nav class="container px-6 py-4 mx-auto md:px-12">
        <div class="items-center justify-center md:flex">
          <div class="items-center md:flex">
            <a v-if="store.username === ''|| store.username === undefined" class="mx-3 text-lg text-white uppercase cursor-pointer hover:text-gray-300 " @click="showLogin">
              登录
            </a>
            <a v-else class="mx-3 text-lg text-white uppercase cursor-pointer hover:text-gray-300 " @click="showLogin">
              {{ store.username }}
            </a>
            <a @click="showtabs = 1" class="mx-3 text-lg text-white uppercase cursor-pointer hover:text-gray-300">
              首页
            </a>
            <a @click="showtabs = 2" class="mx-3 text-lg text-white uppercase cursor-pointer hover:text-gray-300">
              排行榜
            </a>
            <a class="mx-3 text-lg text-white uppercase cursor-pointer hover:text-gray-300">
              关于
            </a>

          </div>
        </div>
      </nav>
    </header>
    <div v-if="showtabs === 1" class="container relative z-10 flex items-center px-6 py-32 mx-auto md:px-12 xl:py-40">
      <div class="relative z-10 flex flex-col items-center w-full">
        <h1 class="mt-4 font-extrabold leading-tight text-center text-white text-5xl sm:text-5xl">
          第3届卡猿杯创新大赛公开赛
        </h1>
        <a href="#"
          class=" rounded-lg block px-4 py-3 mt-10 text-lg font-bold text-white uppercase bg-gray-800 hover:bg-gray-900"
          @click="scrollToElement">
          开始投票
        </a>
      </div>
    </div>
    <div v-if="showtabs === 2" class=" absolute top-16 w-full">
      <Ranking group="人工智能"></Ranking>
      <Ranking group="数据赋能"></Ranking>
      <Ranking group="融合创新"></Ranking>
    </div>
  </div>

  <div v-if="showtabs === 1" class="container mx-auto">
    <div class="flex flex-col">
      <h3 id='voting' class="text-lg font-semibold mt-2 pt-6 pb-1">
        提示：点击图片即可放大观看
      </h3>
      <!-- <CountDown></CountDown> -->
      <h1 class="text-lg font-semibold">投票情况：</h1>
      <List>
        <ListItem v-for="team in teamList" :key="team.id" @login="showLogin" v-model:id="team.id" v-model:name="team.workname"
          v-model:score="team.score" v-model:teamName="team.teamname" :leader="team.leader" :group="team.group" :ImageUrl="team.ImageUrl"></ListItem>
      </List>
      <!-- <button class="btn glass text-center w-6/12 mx-auto text-lg m-6 bg-indigo-900 text-white" @click="submit">
        投票
      </button> -->
    </div>
  </div>
</template>

<style scoped>
.fullscreen {
  background-image: url("/images/6.svg");
  height: 200vh;
}
</style>
