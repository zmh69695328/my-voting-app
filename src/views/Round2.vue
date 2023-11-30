<script setup>
import List from '@/components/voting/List.vue'
import ListItem from '@/components/voting/ListItem.vue'
import Login from '@/components/login/Login.vue'
import Ranking from '@/components/ranking/Ranking.vue'
import { nextTick, onMounted, reactive, ref, watch, onActivated, computed } from 'vue';
import { useUsernameStore } from '@/stores/username.js'
import Steps from '@/components/Steps.vue'
import { useRoute } from 'vue-router'
import BackToTop from '@/components/BackToTop.vue'
import { isNotEmpty } from '@/utils/utils.js'

const route = useRoute()
// 获取路由名称
console.log(route.name)

// 获取路由参数
console.log(route.params.id)
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
function scrollToElement(id) {
  const targetElement = document.getElementById(id);
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
  const targetElement = document.querySelector('.contexttitle');
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


  let response = await fetch('/api/teams', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json;charset=utf-8'
    }
  });
  try {
    let result = await response?.json();
    console.log(result);
    teamList.length = 0;
    teamList.push(...result.teams.filter(t => t.group !== '最佳落地奖'));
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
    teamVotesList.push(...result.teamvotes.filter(v => v.group !== '最佳落地奖'));
    teamList.forEach((team, index) => {
      if (team.id === result.teamvotes[index].id) {
        team.score = result.teamvotes[index].score
      }
    })
  }
}
async function submit() {
  getTeamVotesList()
}

const unvoted1 = computed(() => {
  if (teamVotesList.length === 0) {
    return ''
  } else {
    return teamList.map((team, index) => {
      if ((team.group === '人工智能' && team.id === teamVotesList[index].id) && (team.score === undefined || team.score === null)) {
        return team.teamname
      } else {
        return ''
      }
    }).filter(item => item !== '').join(',')
  }

})
const unvoted2 = computed(() => {
  if (teamVotesList.length === 0) {
    return ''
  } else {
    return teamList.map((team, index) => {
      if ((team.group === '数据赋能' && team.id === teamVotesList[index].id) && (team.score === undefined || team.score === null)) {
        return team.teamname
      } else {
        return ''
      }
    }).filter(item => item !== '').join(',')
  }

})
const unvoted3 = computed(() => {
  if (teamVotesList.length === 0) {
    return ''
  } else {
    return teamList.map((team, index) => {
      if ((team.group === '融合创新' && team.id === teamVotesList[index].id) && (team.score === undefined || team.score === null)) {
        return team.teamname
      } else {
        return ''
      }
    }).filter(item => item !== '').join(',')
  }

})
// const unvoted2 = computed(() => `${firstName.value} ${lastName.value}`)
// const unvoted3 = computed(() => `${firstName.value} ${lastName.value}`)
async function afterSubmit() {
  // await getTeamVotesList()
  console.log(teamVotesList)
  let flag = true
  for (let i = 0; i < teamVotesList.length; i++) {
    if (!isNotEmpty(teamVotesList[i].score)) {
      console.log(teamVotesList[i])
      flag = false
      scrollToElement(teamVotesList[i].id)
      break
    }
  }
  if (flag) {
    showAlert.value=true
  }
}
let showAlert = ref(false)
function closeTab(){
  window.location.href="about:blank";
}
</script>

<template>
  <Login ref="login" @login="submit"></Login>
  <div class="container mx-auto mt-5">
    <div class="flex flex-col">
      <h1 class="contexttitle ml-2 mb-4 text-3xl font-extrabold text-gray-900 dark:text-white md:text-5xl lg:text-6xl">
        <span class="text-transparent bg-clip-text bg-gradient-to-r to-emerald-600 from-sky-400"></span> 决赛</h1>
      <!-- <CountDown></CountDown> -->
      <h1 class="ml-2 text-lg font-semibold mt-2">投票情况一览（左右滑动）：</h1>
      <Steps :teamVotesList="teamVotesList"></Steps>
      <h1 class="ml-2 text-lg font-semibold mt-2">未投票队伍列表：</h1>
      <p class="ml-2">
        <span class="text-black font-bold">人工智能：</span>
        <span class="text-red-500">{{ unvoted1 }}</span>
      </p>
      <p class="ml-2">
        <span class="text-black font-bold">数据赋能：</span>
        <span class="text-red-500">{{ unvoted2 }}</span>
      </p>
      <p class="ml-2">
        <span class="text-black font-bold">融合创新：</span>
        <span class="text-red-500">{{ unvoted3 }}</span>
      </p>
      <h3 id='voting' class="ml-2 text-lg font-semibold mt-2 pt-6 pb-1">
        提示：点击赛队介绍图片即可放大观看
      </h3>
      <div class="mb-3 flex ">
        <button class="btn ml-2 mr-1" :class="{ 'btn-active': selected === 0 }" @click="selectGroup(0)">所有</button>
        <button class="btn mr-1" :class="{ 'btn-active': selected === 1 }" @click="selectGroup(1)">人工智能</button>
        <button class="btn mr-1" :class="{ 'btn-active': selected === 2 }" @click="selectGroup(2)">数据赋能</button>
        <button class="btn mr-1" :class="{ 'btn-active': selected === 3 }" @click="selectGroup(3)">融合创新</button>
      </div>
      <List>
        <ListItem v-for="team in teamList" :key="team.id" @login="showLogin" v-model:id="team.id"
          v-model:name="team.workname" v-model:score="team.score" v-model:teamName="team.teamname" :leader="team.leader"
          :order="team.order" :group="team.group" :ImageUrl="team.ImageUrl" :selectGroup="selected" @submit="submit">
        </ListItem>
      </List>
      <button class="btn glass text-center w-7/12 mx-auto text-lg m-5 mt-10 bg-orange-500 text-white"
        @click="afterSubmit">
        结束投票
      </button>
    </div>
    <BackToTop @scrollToTop="scrollToSteps"></BackToTop>
  </div>
  <div v-if="showAlert" role="alert" class="alert absolute inset-x-0 top-1/2 w-1/2 center mx-auto">
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-info shrink-0 w-6 h-6">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
        d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
    </svg>
    <span>感谢您的支持与参与！</span>
    <div>
      <button class="btn btn-sm" @click="showAlert = false;">取消</button>
      <button class="btn btn-sm btn-primary" @click="closeTab">离开</button>
    </div>
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
}</style>