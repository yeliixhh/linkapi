import axios from 'axios'

const BASE_URL = 'http://127.0.0.1:8080/api/v1'

const intance = axios.create({
    baseURL: BASE_URL
})


// 写请求拦截器
intance.interceptors.response.use((res) => {

    if (res.data === undefined || res.data === null) {
        return res
    }

    return res.data

}, (rej) => {
    return rej
})

const post = function<T>(url: string, data: T) {
    return intance.post(url, data, {

    }) as Promise<any>
}

export {
    intance,
    post
}