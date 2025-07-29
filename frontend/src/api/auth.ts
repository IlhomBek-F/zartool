import { publicHttp } from "./http";

export async function login(login: string, password: string) {
    return publicHttp.post("/auth/login", {login, password})
}