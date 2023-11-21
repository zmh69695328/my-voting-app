<template>
    <div class="flex justify-center mt-8">
    
        <!-- Your HTML template code here -->
        <div class="w-1/2 content-center">
            <div class="flex flex-col w-full border-opacity-50">
                <!-- <div class="grid h-20 card bg-base-300 rounded-box place-items-center">第一轮</div> -->
                <button class="btn btn-wide" @click="$router.push('/round1')">第一轮</button>

                <div class="divider"></div>
                <!-- <div class="grid h-20 card bg-base-300 rounded-box place-items-center">第二轮</div> -->
                <button class="btn btn-wide" @click="$router.push('/round2')">第二轮</button>

            </div>

        </div>

    </div>
</template>
