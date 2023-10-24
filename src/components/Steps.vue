<script setup>
const props = defineProps({
    teamVotesList: {
        type: Array,
        default:() => []
    }
})
</script>
<template>
    <div class="overflow-x-auto m-2 mb-2">
        <ul class="steps">
            <!-- <li class="step">start</li>
            <li class="step step-secondary">2</li>
            <li class="step step-secondary">3</li> -->
            <li v-for="item in props.teamVotesList" :key="item.id" class="step" :class="{ 'step-info': item.score!==undefined&&item.score!==null }">{{ item.id }}</li>

        </ul>
    </div>
</template>