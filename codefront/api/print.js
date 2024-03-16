import http from "./axios";


export const sendToPrinter = (params) => {
    return http({
        url: "/print",
        method: "post",
        data: params
    })
}