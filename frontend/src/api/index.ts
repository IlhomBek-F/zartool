import type { ResponseType } from "../core/models/base-model";
import type { CreateRentToolRequestType, RentToolType, UpdateRentToolRequestType } from "../core/models/rent-tool-model";
import type { CreateRentRequestType, RentType, UpdateRentRequestType } from "../core/models/renter-model";
import { privateHttp, publicHttp } from "./http";

export async function login(login: string, password: string) {
    return publicHttp.post("/auth/login", {login, password})
}

export async function createRent(payload: CreateRentRequestType): Promise<ResponseType> {
  return privateHttp.post("/rental/create", payload)
}

export async function getRenters(): Promise<ResponseType<RentType[]>> {
    return privateHttp.get("/rentals")
}

export async function updateRent(payload: UpdateRentRequestType): Promise<ResponseType> {
    return privateHttp.put("/rental/update", payload)
}

export async function deleteRent(id: number): Promise<ResponseType> {
    return privateHttp.delete(`/rental/delete/${id}`)
}

export async function completeRent(id: number): Promise<ResponseType> {
    return privateHttp.post(`/rental/complete/${id}`)
}

export async function getRentTools(): Promise<ResponseType<RentToolType[]>> {
    return privateHttp.get("/warehouse-tools")
}

export async function addNewTool(payload: CreateRentToolRequestType[]) {
 return privateHttp.post("/warehouse-tool/create", payload)
}

export async function updateTool(payload: UpdateRentToolRequestType) {
    return privateHttp.put("/warehouse-tool/update", payload)
}

export async function deleteTool(id: number) {
    return privateHttp.delete(`/warehouse-tool/delete/${id}`)
}