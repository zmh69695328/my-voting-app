<script setup>
import { ref, onMounted, computed } from 'vue';
import { isBetweenZeroAndTwenty } from '../../utils/checkString.js'
import VueEasyLightbox from 'vue-easy-lightbox'
onMounted(() => {

});
const props = defineProps({
    id: { type: Number, default: undefined },
    name: { type: String, default: undefined },
    score: { type: Number, default: undefined },
    teamName: { type: String, default: undefined },
    leader: { type: String, default: undefined },
});
let inputValue = ref(props.score)
const isShowWarning = ref(false)
const emit = defineEmits(['update:score'])
function handleChange(val) {
    console.log(val)
    emit('update:score', inputValue.value)
}
function handleBlur() {
    isShowWarning.value = !isBetweenZeroAndTwenty(inputValue.value)
}
let showFlag = ref(0)
function submit() {
    showFlag.value++
    showFlag.value = 1
    setTimeout(() => {
        showFlag.value = 2
    }, 500);
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
    imgsRef.value = 'teams/team1.png'
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
    <div class="basis-full sm:basis-1/2 md:basis-1/3 lg:basis-1/4 p-4">
        <div class="card w-96 bg-base-100 shadow-xl mx-auto">
            <figure>

                <img src="/teams/team1.png" @click="showSingle" alt="Shoes" />
                <!-- <div :class="classes">
                    <img src="/teams/team1.png" alt="" title="" />
                    <a href="teams/team1.png"><img src="teams/team1.png" alt="" title="Beautiful Image"/></a>
                </div> -->
            </figure>
            <div class="card-body">
                <h2 class="card-title">{{ id + '. ' + name }}</h2>
                <p>赛队名：{{ teamName }}</p>
                <!-- <p>作品名：</p> -->
                <p>队长：{{ leader }}</p>
                <div class="card-actions">
                    <div class="flex flex-row w-full">
                        <span class="label-text basis-2/12 text-base">评分：</span>
                        <input v-model="inputValue" @blur="handleBlur" @change="handleChange"
                            class="basis-6/12 input input-bordered input-primary w-8 max-w-xs" type="text"
                            placeholder="请输入0-20分">
                        <div v-if="showFlag > 0" class="flex basis-1/12 m-1">
                            <span v-if="showFlag === 1" class=" loading loading-spinner text-success"></span>
                            <div v-if="showFlag === 2" class="flex justify-center items-center ">
                                <input type="checkbox" @click.stop.prevent="" checked="checked"
                                    class="checkbox checkbox-success " />
                            </div>
                        </div>
                        <button class="btn glass basis-3/12  w-6/12 mx-auto text-lg  bg-indigo-900 text-white"
                            @click="submit">
                            投票
                        </button>
                    </div>
                    <span v-show="isShowWarning" class="text-red-600">请输入正确的值</span>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped></style>
