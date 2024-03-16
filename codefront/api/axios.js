import axios from 'axios'
const http = axios.create({
    baseURL: "http://localhost:8080/",
    withCredentials: true,
})
const beforeRequest = config => {
    config.headers['content-type'] = 'application/json'
    return config
}

http.interceptors.request.use(beforeRequest)

const responseSuccess = response => {
    return Promise.resolve(response.data)
}

const responseFailed = error => {
    const { response } = error
    if (response) {
        return Promise.reject(response.data)
    }
    return Promise.reject(error)
}

http.interceptors.response.use(responseSuccess, responseFailed)

export default http;
