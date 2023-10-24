<template>
    <dialog id="my_modal_2" class="modal" :class="{ 'modal-open': showModal }">
        <div class="modal-box p-2 w-auto">
            <form class="card-body">
                <div class="form-control">
                    <label class="label">
                        <span class="label-text text-lg">评委姓名</span>
                    </label>
                    <input v-model="username" type="text" placeholder="请输入您的姓名" class="input input-bordered" required />
                </div>
                <!-- <div class="form-control">
                                <label class="label">
                                    <span class="label-text">密码</span>
                                </label>
                                <input type="password" placeholder="请输入任意密码" class="input input-bordered" required />
                                <label class="label">
                                    <a href="#" class="label-text-alt link link-hover">Forgot password?</a>
                                </label>
                            </div> -->
                <div class="form-control mt-6">
                    <button class="btn btn-primary text-lg" @click="login">登录</button>
                </div>
            </form>
        </div>
        <form method="dialog" class="modal-backdrop" @click="showModal = false">
            <button>close</button>
        </form>
    </dialog>
</template>

<script setup>
import { ref } from 'vue'
import { useUsernameStore } from '@/stores/username.js'
const showModal = ref(false)
defineExpose({
    showModal
})
const store = useUsernameStore()
let username = ref('')
function login() {
    store.username = username.value
    sessionStorage.setItem('username', username.value)
    showModal.value = false
}
</script>