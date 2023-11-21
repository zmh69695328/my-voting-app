<template>
    <div class="flex justify-center mt-8 mb-8">
    
        <!-- Your HTML template code here -->
        <div class="w-1/2 flexjustify-center">
            <div class="flex flex-col justify-center w-full">
                <!-- <div class="grid h-20 card bg-base-300 rounded-box place-items-center">第一轮</div> -->
                <button class="btn" @click="$router.push('/round1')">第一轮</button>

                <div class="divider"></div>
                <!-- <div class="grid h-20 card bg-base-300 rounded-box place-items-center">第二轮</div> -->
                <button class="btn" @click="$router.push('/round2')">第二轮</button>

            </div>

        </div>
    </div>
    <div class="h-3"></div>
</template>
