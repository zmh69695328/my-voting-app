<script setup>
import { ref, onMounted, computed, nextTick, watch } from 'vue';
import { isBetweenZeroAndTwenty } from '../../utils/checkString.js'
import VueEasyLightbox from 'vue-easy-lightbox'
import { useUsernameStore } from '@/stores/username.js'
onMounted(() => {

});
const props = defineProps({
    id: { type: Number, default: undefined },
    name: { type: String, default: undefined },
    score: { type: Number, default: undefined },
    teamName: { type: String, default: undefined },
    leader: { type: String, default: undefined },
    group: { type: String, default: '' },
    ImageUrl: { type: String, default: '/teams/team1.png' },
    selectGroup: { type: Number, default: 0 },
    order: { type: String, default: '' },
});
let inputValue = ref(props.score)
const selectGroupMap={
    1:'人工智能',
    2:'数据赋能',
    3:'融合创新',
}
watch(() => props.score, (newVal) => {
    const store = useUsernameStore()
    if (store.username&&props.score !== undefined && props.score !== null&&isBetweenZeroAndTwenty(props.score)&&props.score!=='') {
        showFlag.value = 2
    }
    inputValue.value = newVal
})
const mode = computed(() => {
    return import.meta.env.MODE
})
let showMessage = ref(false)
const isShowWarning = ref(false)
const isShowWarningUnderline = ref(false)
const emit = defineEmits(['update:score', 'login', 'submit'])
function handleChange(val) {
    console.log(val)
    emit('update:score', inputValue.value)
}
function handleBlur() {
    isShowWarning.value = !isBetweenZeroAndTwenty(inputValue.value)
}
let showFlag = ref(0)
function debounce(func, timeout = 300){
  let timer;
  return (...args) => {
    clearTimeout(timer);
    timer = setTimeout(() => { func.apply(this, args); }, timeout);
  };
}

const submitDebounce = debounce(() => submit());
async function submit() {
    const store = useUsernameStore()
    // check if login
    if (store.username === '' || store.username === undefined) {
        showMessage.value = true
        setTimeout(() => {
            showMessage.value = false
        }, 1000)
        emit('login')
        return;
    }
    // 对输入值进行校验，0-20 
    // validate input
    if (inputValue.value === '' || !isBetweenZeroAndTwenty(inputValue.value)) {
        isShowWarning.value = true
        isShowWarningUnderline.value = true
        setTimeout(() => {
            isShowWarningUnderline.value = false
        }, 300)
        return
    }
    showFlag.value = 1
    // setTimeout(() => {
    //     showFlag.value = 2
    // }, 500);

    const username = store.username
    let req = {
        teamid: props.id,
        score: parseInt(inputValue.value),
        username
    }

    let response = await fetch('/api/vote', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json;charset=utf-8'
        },
        body: JSON.stringify(req)
    }).catch(error => {
        showFlag.value = 3
        console.error('Error:', error)
    });
    if (response.status !== 200) {
        showFlag.value = 3
        console.log(response)
    } else {
        let result = await response.json();
        console.log(result)
        setTimeout(() => {
            showFlag.value = 2
            emit('submit')
        }, 200);
        // showFlag.value = 2
    }
}
const visibleRef = ref(false)
const indexRef = ref(0)
const imgsRef = ref([])
// Img Url , string or Array of string
// ImgObj { src: '', title: '', alt: '' }
// 'src' 是必须值
// 允许混合
const onShow = () => {
    visibleRef.value = true
}
const showSingle = () => {
    imgsRef.value = props.ImageUrl
    // or
    // imgsRef.value  = {
    //   title: 'this is a placeholder',
    //   src: 'http://via.placeholder.com/350x150'
    // }
    onShow()
}
const onHide = () => (visibleRef.value = false)
</script>

<template>
    <vue-easy-lightbox :visible="visibleRef" :imgs="imgsRef" :index="indexRef" @hide="onHide">
    </vue-easy-lightbox>
    <div class="toast z-50" v-show="showMessage">
        <div class="alert alert-info">
            <span>请先登录！</span>
        </div>
    </div>
    <div :id="props.id" v-show="selectGroup===0||selectGroupMap[selectGroup]===group" class="basis-full sm:basis-1/2 md:basis-1/3 lg:basis-1/4 p-4">
        <div class="card w-96 bg-base-100 shadow-xl mx-auto">
            <figure>
                <img v-if="mode==='prod'" :src="ImageUrl" @click="showSingle" />
            </figure>
            <div class="card-body">
                <h2 class="card-title">{{ order + '. ' + name }}</h2>
                <p>赛队名：{{ teamName }}</p>
                <p>赛道：{{ group }}</p>
                <p>队长：{{ leader }}</p>
                <div class="card-actions">
                    <div class="flex flex-row w-full">
                        <span class="label-text basis-2/12 text-base">评分：</span>
                        <input v-model="inputValue" type="number" pattern="[0-9]*" @blur="handleBlur" @change="handleChange"
                            class="basis-6/12 input input-bordered input-primary w-8 max-w-xs"
                            placeholder="请输入0-100分">
                        <div v-if="showFlag > 0" class="flex basis-1/12 m-1">
                            <span v-if="showFlag === 1" class=" loading loading-spinner text-success"></span>
                            <div v-if="showFlag === 2" class="flex justify-center items-center ">
                                <input type="checkbox" @click.stop.prevent="" checked="checked"
                                    class="checkbox checkbox-success " />
                            </div>
                            <div v-if="showFlag === 3" class="flex justify-center items-center alert-error">
                                <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-8 w-8" fill="none"
                                    viewBox="0 0 24 24">
                                    <path stroke="white" fill="red" stroke-linecap="round" stroke-linejoin="round"
                                        stroke-width="2"
                                        d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
                                </svg>
                            </div>
                        </div>
                        <button class="btn glass basis-3/12  w-6/12 mx-auto text-lg  bg-indigo-900 text-white"
                            @click="submitDebounce">
                            投票
                        </button>
                    </div>
                    <span v-show="isShowWarning" class="text-red-600"
                        :class="{ underline: isShowWarningUnderline }">请输入正确的值</span>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped></style>
