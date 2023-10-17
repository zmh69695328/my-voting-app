<script setup>
import { ref } from 'vue';
import { isBetweenZeroAndTwenty } from '../../utils/checkString.js'
const props = defineProps({
    id: { type: Number, default: undefined },
    name: { type: String, default: undefined },
    score: { type: Number, default: undefined },
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
let showFlag=ref(0)
function submit(){
    showFlag.value++
    showFlag.value=1
    setTimeout(() => {
        showFlag.value=2
    }, 500);
}
</script>

<template>
    <div class="basis-full sm:basis-1/2 md:basis-1/3 lg:basis-1/4 p-4">
        <div class="card w-96 bg-base-100 shadow-xl mx-auto">
            <figure><img src="/images/shoes.png" alt="Shoes" /></figure>
            <div class="card-body">
                <h2 class="card-title">{{ id + '. ' + name }}</h2>
                <p>团队成员：</p>
                <div class="card-actions">
                    <div class="flex flex-row w-full">
                        <span class="label-text basis-2/12 text-base">评分：</span>
                        <input v-model="inputValue" @blur="handleBlur" @change="handleChange"
                            class="basis-6/12 input input-bordered input-primary w-8 max-w-xs" type="text"
                            placeholder="请输入0-20分">
                        <div v-if="showFlag>0" class="flex basis-1/12 m-1">
                            <span v-if="showFlag===1" class=" loading loading-spinner text-success"></span>
                            <div v-if="showFlag===2" class="flex justify-center items-center ">
                                <input type="checkbox" @click.stop.prevent="" checked="checked" class="checkbox checkbox-success " />
                            </div>
                        </div>
                        <button class="btn glass basis-3/12  w-6/12 mx-auto text-lg  bg-indigo-900 text-white"
                            @click="submit">
                            提交
                        </button>
                    </div>
                    <span v-show="isShowWarning" class="text-red-600">请输入正确的值</span>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped></style>
