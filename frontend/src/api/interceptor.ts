import { getToken } from "../utils/tokenUtil";
import { privateHttp } from "./http";

privateHttp.interceptors.request.use(function (config) {
    const token = getToken();

    if(!token) {
        window.location.replace("login");
    }

    config.headers.setAuthorization(`Bearer ${token}`);

    return config;
  }, (error) =>  Promise.reject(error));

privateHttp.interceptors.response.use((response) => response, function (error) {
    if(error.status === 401) {
       window.location.replace("login")
    }
    return Promise.reject(error);
});