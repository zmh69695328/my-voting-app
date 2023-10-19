<template>
    <dialog id="my_modal_2" class="modal">
        <div class="modal-box">
            <form class="card-body">
                            <div class="form-control">
                                <label class="label">
                                    <span class="label-text">用户名</span>
                                </label>
                                <input type="text" placeholder="请输入您的姓名" class="input input-bordered" required />
                            </div>
                            <div class="form-control">
                                <label class="label">
                                    <span class="label-text">密码</span>
                                </label>
                                <input type="password" placeholder="请输入任意密码" class="input input-bordered" required />
                                <!-- <label class="label">
                                    <a href="#" class="label-text-alt link link-hover">Forgot password?</a>
                                </label> -->
                            </div>
                            <div class="form-control mt-6">
                                <button class="btn btn-primary text-lg">登录</button>
                            </div>
                        </form>
        </div>
        <form method="dialog" class="modal-backdrop">
            <button>close</button>
        </form>
    </dialog>
</template>