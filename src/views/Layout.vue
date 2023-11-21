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



// 计算子元素的高度 
function getHeights() {
    nextTick(() => {
        const container = document.querySelector('#container');
        const ranking = document.querySelector('#ranking');
        console.log(ranking.offsetHeight)
        container.style.height = ranking.offsetHeight + 100 + 'px';
    })
}

watch(showtabs, async () => {
    // await getTeams()
    if (showtabs.value === 1) {
        nextTick(() => {
            const container = document.querySelector('#container');
            container.style.cssText = ''
        })
    }
})

</script>

<template>
    <Login ref="login" @login="submit"></Login>
    <div id="container" class="relative h-screen bg-indigo-900" :class="{ 'fullscreen': showtabs === 2 }">
        <img v-if="showtabs === 1" src="https://www.qiniu.public.zhuminhao.top/6.svg"
            class="absolute object-cover w-full h-full" />
        <!-- <img v-if="showtabs === 2" src="/images/6.svg" class="absolute object-cover w-full h-[80rem]" /> -->
        <!-- <img v-if="showtabs === 2" src="/images/6.svg" class="absolute object-cover w-full h-screen" /> -->
        <div class="absolute inset-0 bg-black opacity-25"></div>
        <header class="absolute top-0 left-0 right-0 z-20">
            <nav class="container px-6 py-4 mx-auto md:px-12">
                <div class="items-center justify-center md:flex">
                    <div class="items-center md:flex">
                        <div v-if="store.username === '' || store.username === undefined"
                            class="inline-block mx-3 text-lg text-white uppercase cursor-pointer hover:text-gray-300 "
                            @click="showLogin">
                            登录
                        </div>
                        <div v-else
                            class="inline-block mx-3 text-lg text-white uppercase cursor-pointer hover:text-gray-300 "
                            @click="showLogin">
                            {{ store.username }}
                        </div>
                        <a @click="showtabs = 1;$router.push('/home')"
                            class="mx-3 text-lg text-white uppercase cursor-pointer hover:text-gray-300">
                            首页
                        </a>
                        <a v-if="store.username === '黄素梅'" @click="showtabs = 2"
                            class="mx-3 text-lg text-white uppercase cursor-pointer hover:text-gray-300">
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
                    第5届卡猿杯创新大赛公开赛
                </h1>
                <a href="#"
                    class=" rounded-lg block px-4 py-3 mt-10 text-lg font-bold text-white uppercase bg-gray-800 hover:bg-gray-900"
                    @click="scrollToElement">
                    开始投票
                </a>
            </div>
        </div>


    </div>
    <router-view />
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
