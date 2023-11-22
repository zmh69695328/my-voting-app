<script setup>
import List from '@/components/voting/List.vue'
import ListItem from '@/components/voting/ListItem.vue'
import Login from '@/components/login/Login.vue'
import Ranking from '@/components/ranking/Ranking.vue'
import { nextTick, onMounted, reactive, ref, watch, onActivated, computed } from 'vue';
import { useUsernameStore } from '@/stores/username.js'
import Steps from '@/components/Steps.vue'
import { useRoute } from 'vue-router'
const route = useRoute()
// 获取路由名称
console.log(route.name)
const showtabs = ref(1) //1是投票页，2是排行榜页
const login = ref()
const selected = ref(0)
function selectGroup(index) {
  selected.value = index
}
function showLogin() {
  login.value.showModal = true
}
const store = useUsernameStore()
function scrollToElement() {

  if (store.username === '' || store.username === undefined) {
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
}
function scrollToSteps() {
  const targetElement = document.querySelector('.steps');
  if (targetElement) {
    setTimeout(() => {
      targetElement.scrollIntoView({
        behavior: 'smooth', // 使用平滑滚动效果
        block: 'start',     // 滚动到元素的顶部
      })
    }, 10);

  }
}

function checkScroll() {
  // 当滚动时检查滚动位置和元素是否在视窗中
  const scrollY = window.scrollY;
  const windowHeight = window.innerHeight;
  const element = document.querySelector('.steps'); // 替换为要检查的元素的实际ID

  if (scrollY > 0 || (element && element.getBoundingClientRect().top < windowHeight)) {
    this.showButton = true;
  } else {
    this.showButton = false;
  }
}

let teamList = reactive([])
async function getTeams() {

  let req={group:'最佳落地奖'}
  let response = await fetch('/api/teams', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8'
    },
    body:JSON.stringify(req)
  });
  try {
    let result = await response?.json();
    console.log(result);
    teamList.length = 0;
    teamList.push(...result.teams);
  } catch (error) {
    console.error('Error:', error)
  }

}

// 计算子元素的高度 
function getHeights() {
  nextTick(() => {
    const container = document.querySelector('#container');
    const ranking = document.querySelector('#ranking');
    console.log(ranking.offsetHeight)
    container.style.height = ranking.offsetHeight + 100 + 'px';
  })
}
onMounted(async () => {
  await getTeams()
  if (showtabs.value === 2) {
    getHeights()
  }
  if (store.username !== '' && store.username !== undefined) {
    getTeamVotesList()
  }
})
watch(showtabs, async () => {
  // await getTeams()
  if (showtabs.value === 1) {
    nextTick(() => {
      const container = document.querySelector('#container');
      container.style.cssText = ''
    })
  }
})


const teamVotesList = reactive([])
async function getTeamVotesList() {

  let req = {
    username: store.username
  }

  let response = await fetch('/api/teamvotes', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8'
    },
    body: JSON.stringify(req)
  }).catch(error => {
    console.error('Error:', error)
  })
  if (response.status === 200) {
    teamVotesList.length = 0;
    let result = await response.json();
    teamVotesList.push(...result.teamvotes);
    teamList.forEach((team) => {
      // debugger
      result.teamvotes.forEach(v=>{
        if (team.id === v.id) {
        team.score = v.score
      }
      })
      
    })
  }
}
async function submit() {
  getTeamVotesList()
}

</script>

<template>
  <div class="container mx-auto mt-5">
    <div class="flex flex-col">

      <h1 class="ml-2 mb-4 text-3xl font-extrabold text-gray-900 dark:text-white md:text-5xl lg:text-6xl"><span
          class="text-transparent bg-clip-text bg-gradient-to-r to-emerald-600 from-sky-400">最佳</span> 落地奖</h1>
      <!-- <p class="text-lg font-normal text-gray-500 lg:text-xl dark:text-gray-400">Here at Flowbite we focus on markets where technology, innovation, and capital can unlock long-term value and drive economic growth.</p> -->

      <!-- <CountDown></CountDown> -->
      <h1 class="ml-2 text-lg font-semibold mt-2">投票情况一览（左右滑动）：</h1>
      <Steps :teamVotesList="teamVotesList"></Steps>

      <h3 id='voting' class="ml-2 text-lg font-semibold mt-2 pt-6 pb-1">
        提示：点击赛队介绍图片即可放大观看
      </h3>
      <List>
        <ListItem v-for="team in teamList" :key="team.id" @login="showLogin" v-model:id="team.id"
          v-model:name="team.workname" v-model:score="team.score" v-model:teamName="team.teamname" :leader="team.leader"
          :group="team.group" :ImageUrl="team.ImageUrl" :selectGroup="selected" @submit="submit"></ListItem>
      </List>
      <!-- <button class="btn glass text-center w-6/12 mx-auto text-lg m-6 bg-indigo-900 text-white" @click="submit">
        投票
      </button> -->
    </div>
  </div>
  <div id="topButton" class="rightside fixed bottom-20 h-10 right-2 w-10 rounded-full">
    <button class="button" @click="scrollToSteps">
      <svg t="1698144093962" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
        p-id="1446" width="32" height="32">
        <path
          d="M848.69632 642.33472L537.344 330.9824a35.83488 35.83488 0 0 0-50.68288 0L175.3088 642.33472a35.84 35.84 0 0 0 50.68288 50.688L512 407.00928l286.00832 286.00832a35.83488 35.83488 0 0 0 50.68288 0 35.82976 35.82976 0 0 0 0.00512-50.68288z"
          fill="#666666" p-id="1447"></path>
      </svg> </button>
  </div>
</template>
<style scoped>
.fullscreen {
  /* background-image: url("/images/6.svg"); */
  background-image: url('http://s32m14foq.bkt.clouddn.com/6.svg');
  background-size: cover;
  background-position: center center;
  background-attachment: fixed;
  height: 100%;
  width: 100%;
  position: relative;
  /* overflow: hidden; */
}
</style>