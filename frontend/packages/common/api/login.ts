import { post } from '../utils/request'

// 登录方法
const loginAPI = function(username: string, password: string) {
    return post('auth/login', {username, password})
}

export {
    loginAPI
}