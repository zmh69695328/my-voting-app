<template>
<div id="scrollTopButton" class="sticky w-full flex pt-3 justify-end bottom-0 pb-3 pr-2 z-10 lg:pr-16 transition">
    <div class="bg-blue-300 bg-opacity-75 rounded-full cursor-pointer text-gray-800 hover:text-white transition">
        <button role="button" aria-label="scroll to top of the page" @click="scrollToTop">
            <i class="inline-block w-7 h-6">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M11.9997 10.8284L7.04996 15.7782L5.63574 14.364L11.9997 8L18.3637 14.364L16.9495 15.7782L11.9997 10.8284Z"></path></svg>
            </i>
        </button>
    </div>
</div>
</template>
<script setup>
import { defineEmits } from 'vue';
const emit = defineEmits(['scrollToTop'])

function scrollToTop(){
    // window.scrollTo({ top: 0, behavior: 'smooth' });
    emit('scrollToTop')
}
</script>