<script setup>
import { ref, onMounted, computed } from 'vue';
onMounted(async () => {
    search()
})

const teamName = ref('')
const userName = ref('')
const VoteTeam = ref([])
async function search() {

    let response = await fetch('/api/votesbyteamnameandusername', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json;charset=utf-8'
        },
        body: JSON.stringify({
            teamName: teamName.value.trim(),
            userName: userName.value.trim(),
            isDuplicate:isDuplicate.value
        })
    }).catch(error => {
        console.error('Error:', error)
    });
    if (response.status !== 200) {
        console.log(response)
    } else {
        let result = await response.json();
        VoteTeam.value = result.VoteTeam
    }
}
const isDuplicate = ref(true)
function showIsDuplicate(e) {
    console.log('showIsDuplicate', e)
    isDuplicate.value = !isDuplicate.value
}
</script>
<template>
    <div class="flex justify-between m-5">

        <h1 class="contexttitle ml-2 mb-4 text-3xl font-extrabold text-gray-900 dark:text-white md:text-5xl lg:text-6xl">
            <span class="text-transparent bg-clip-text bg-gradient-to-r to-emerald-600 from-sky-400">所有</span> 明细
        </h1>
        <button class="btn w-30" @click="search">
            查询
            <svg class="inline-block w-[20px] h-[20px]" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path
                    d="M5.46257 4.43262C7.21556 2.91688 9.5007 2 12 2C17.5228 2 22 6.47715 22 12C22 14.1361 21.3302 16.1158 20.1892 17.7406L17 12H20C20 7.58172 16.4183 4 12 4C9.84982 4 7.89777 4.84827 6.46023 6.22842L5.46257 4.43262ZM18.5374 19.5674C16.7844 21.0831 14.4993 22 12 22C6.47715 22 2 17.5228 2 12C2 9.86386 2.66979 7.88416 3.8108 6.25944L7 12H4C4 16.4183 7.58172 20 12 20C14.1502 20 16.1022 19.1517 17.5398 17.7716L18.5374 19.5674Z">
                </path>
            </svg>
        </button>
    </div>
    <div class="flex flex-col m-5">
        <h1>筛选条件：</h1>
        <div class="m-2">
            队名：
            <input v-model="teamName" type="text" placeholder="" class="input input-bordered input-sm w-full max-w-xs" />
        </div>
        <div class="m-2">
            评委名称：
            <input v-model="userName" type="text" placeholder=" " class="input input-bordered input-sm w-full max-w-xs" />
        </div>
        <div class="m-2 " >
            <div class="form-control inline-block" @click="showIsDuplicate">
                <label class="label cursor-pointer justify-start" >
                    <span class="label-text w-full">是否去重：</span>
                    <input type="checkbox" class="toggle toggle-sm" checked/>
                </label>
            </div>
        </div>
    </div>
    <div>
        <div class="overflow-x-auto">
            <table class="table table-xs table-pin-rows table-pin-cols table-zebra">
                <thead>
                    <tr>
                        <th></th>
                        <td>队伍名</td>
                        <td>作品名</td>
                        <td>组别</td>
                        <td>队长</td>
                        <td>评委名称</td>
                        <td>分数</td>
                        <td>时间</td>
                        <td>ID</td>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(item, index) in VoteTeam" :key="item.id">
                        <th>{{ index + 1 }}</th>
                        <td>{{ item.teamname }}</td>
                        <td>{{ item.workname }}</td>
                        <td>{{ item.groupname }}</td>
                        <td>{{ item.leader }}</td>
                        <td>{{ item.username }}</td>
                        <td>{{ item.score }}</td>
                        <th>{{ item.date }}</th>
                        <th>{{ item.id }}</th>
                    </tr>
                </tbody>
                <tfoot>
                    <tr>
                        <th></th>
                        <td>队伍名</td>
                        <td>作品名</td>
                        <td>组别</td>
                        <td>队长</td>
                        <td>评委名称</td>
                        <td>分数</td>
                        <td>时间</td>
                        <td>ID</td>
                    </tr>
                </tfoot>
            </table>
        </div>
    </div>
</template>