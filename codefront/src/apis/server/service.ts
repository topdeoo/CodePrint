// 封装axios
import axios from 'axios';

const baseUrl = 'http://localhost:8081/api/v1/client';
// 创建axios实例
const service = axios.create({
    baseURL: baseUrl, // api的base_url
    timeout: 5000 // 请求超时时间
});

// request拦截器
service.interceptors.request.use(
    config => {
        config.headers['Content-Type'] = 'application/json;charset=UTF-8';
        // 请求头携带token
        const token = JSON.parse(window.localStorage.getItem('teamInfo') || "{}").token;
        if (token) {
            config.headers['Authorization'] = token;
        }
        return config;
    },
    error => {
        // Do something with request error
        return Promise.reject(error);
    }
);

// respone拦截器
service.interceptors.response.use(
    response => {
        return response.data;
    },
    error => {
        return Promise.reject(error);
    }
);

export default service;