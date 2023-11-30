<template>
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
