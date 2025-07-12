import { getToken } from "../utils/tokenUtil";
import { privateHttp, publicHttp } from "./http";

publicHttp.interceptors.response.use((response) => response.data, (error) => Promise.reject(error));

privateHttp.interceptors.request.use(function (config) {
    const token = getToken();

    if(!token) {
        window.location.replace("login");
    }

    config.headers.setAuthorization(`Bearer ${token}`);

    return config;
  }, (error) =>  Promise.reject(error));

privateHttp.interceptors.response.use((response) => response.data, function (error) {
    if(error.status === 401) {
       window.location.replace("login")
    }
    return Promise.reject(error);
});