<template>
    <div>
        <div class="login-wrapper">
            <div class="card" style="padding-bottom: 70px;">
                <div style="font-weight: 500; font-size: 40px; margin-top: 30px; text-align: center;">Link-API</div>
                <div style="margin-top: 40px;" class="form-input-contaner">
                    <div class="form-input-wrap">
                        <span style="width: 60px; text-align: center; line-height: 30px;">登录名</span>
                        <input class="form-input" style="padding-left: 10px;"/>
                    </div>

                </div>

                <div class="form-input-contaner" style="margin-top: 10px;">
                    <div class="form-input-wrap">
                        <span style="width: 60px; text-align: center; line-height: 30px;">密码</span>
                        <input class="form-input" style="padding-left: 10px;" />
                    </div>
                </div>

                <div style="margin: 0 auto; margin-top: 30px;" class="form-input-wrap">
                    <input class="form-input" type="button" value="登录" style="height: 35px; background-color: #0C4283; color: white; font-size: 20px;">
                </div>

                <div style="text-align: center; margin-top: 20px;">
                    <span>忘记密码？</span>
                </div>
            </div>
            <div style="margin-top: 10px;">注册新用户</div>
        </div>

    </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.login-wrapper {
    width: 100vw;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    background: linear-gradient(to right bottom, #E5E9EA, #d6d9df);

    .card {
        width: 80vw;
        /* height: 80vh; */
        /* background-color: red; */
        max-height: 550px;
        max-width: 770px;
        background-color: rgba(255, 255, 255, 0.6);

        border-width: 2px;
        border-style: solid;
        border-radius: 15px;
        border-color: rgba(111, 109, 109, 0.2);

        .form-input-contaner {
            width: 100%;
            display: flex;
            justify-content: center;


        }

    }

    .form-input-wrap {
        width: 50%;
        display: flex;

        .form-input {
            /* border: 0; */
            border-radius: 5px;
            border-width: 1px;
            width: 100%;
            height: 30px;
        }
    }

}
</style>