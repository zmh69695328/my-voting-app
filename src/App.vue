<script setup>
import List from './components/voting/List.vue'
import ListItem from './components/voting/ListItem.vue'
import Login from './components/login/Login.vue'
import { nextTick, onMounted, reactive, ref } from 'vue';
const login = ref()
function showLogin() {
  login.value.showModal = true
}
function scrollToElement() {
  const targetElement = document.getElementById('voting');
  if (targetElement) {
    setTimeout(() => {
      targetElement.scrollIntoView({
        behavior: 'smooth', // 使用平滑滚动效果
        block: 'start',     // 滚动到元素的顶部
      })
    }, 50);

  }
} let teamList = reactive([
  {
    id: 1,
    name: 'AI防火墙',
    teamName: '智慧队',
    score: 12
  },
  {
    id: 2,
    name: 'AIGC图片生成工具',
    teamName: '天马行空共赢队',
    score: 12
  },
  {
    id: 3,
    name: 'Andy',
    score: 12
  },
  {
    id: 4,
    name: 'Andy',
    score: 12
  },
  {
    id: 5,
    name: 'Andy',
    score: 12
  },
  {
    id: 6,
    name: 'Andy',
    score: 12
  },
])
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
  <div class="relative h-screen overflow-hidden bg-indigo-900">
    <img src="/images/6.svg" class="absolute object-cover w-full h-full" />
    <div class="absolute inset-0 bg-black opacity-25"></div>
    <header class="absolute top-0 left-0 right-0 z-20">
      <nav class="container px-6 py-4 mx-auto md:px-12">
        <div class="items-center justify-center md:flex">
          <div class="items-center md:flex">
            <a class="mx-3 text-lg text-white uppercase cursor-pointer hover:text-gray-300 " @click="showLogin">
              登录

            </a>
            <a class="mx-3 text-lg text-white uppercase cursor-pointer hover:text-gray-300">
              排行榜
            </a>
            <a class="mx-3 text-lg text-white uppercase cursor-pointer hover:text-gray-300">
              关于
            </a>

          </div>
        </div>
      </nav>
    </header>
    <div class="container relative z-10 flex items-center px-6 py-32 mx-auto md:px-12 xl:py-40">
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
  </div>

  <div class="container mx-auto">
    <div class="flex flex-col">
      <h3 id='voting' class="text-lg font-semibold mt-2 pt-6 pb-1">
        提示：点击图片即可放大观看
      </h3>
      <!-- <CountDown></CountDown> -->
      <h1 class="text-lg font-semibold">投票情况：</h1>
      <List>
        <ListItem v-for="team in teamList" :key="team.id" v-model:id="team.id" v-model:name="team.workname"
          v-model:score="team.score" v-model:teamName="team.teamname" :leader="team.leader" :group="team.group" :ImageUrl="team.ImageUrl"></ListItem>
      </List>
      <!-- <button class="btn glass text-center w-6/12 mx-auto text-lg m-6 bg-indigo-900 text-white" @click="submit">
        投票
      </button> -->
    </div>
  </div>
</template>

<style scoped></style>
