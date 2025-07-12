import type { ResponseType } from "../core/models/base-model";
import type { CreateRentToolRequestType, RentToolType, UpdateRentToolRequestType } from "../core/models/rent-tool-model";
import type { CreateRentRequestType, RentType, UpdateRentRequestType } from "../core/models/renter-model";
import { privateHttp, publicHttp } from "./http";

let data: RentType[] = [];
let tools: RentToolType[] = [];

export async function login(login: string, password: string) {
    return publicHttp.post("/auth/login", {login, password})
}

export async function createRent(payload: CreateRentRequestType): Promise<RentType> {
  return new Promise((resolve, _) => {
     payload.id = Date.now();
     data.push(payload);
     resolve(payload)
  })
}

export async function getRenters(): Promise<RentType[]> {
    return new Promise((resolve, _) => {
       resolve(data)
    })
}

export async function updateRent(payload: UpdateRentRequestType): Promise<RentType> {
    return new Promise((resolve, _) => {
      data = data.map((rent) => rent.id === payload.id ? payload : rent);
      resolve(payload)  
    })
}

export async function deleteRent(id: number) {
    return new Promise((resolve, _) => {
        data = data.filter((rent) => rent.id !== id);
        resolve(id)
    })
}

export async function closeRent(id: number) {
    return new Promise((resolve, _) => {
        data = data.filter((rent) => rent.id === id ? {...rent, status: 'closed'} : rent);
        resolve(id)
    })
}

export async function getRentTools(): Promise<ResponseType<RentToolType[]>> {
    return privateHttp.get("/warehouse-tools")
}

export async function addNewTool(payload: CreateRentToolRequestType[]) {
 return privateHttp.post("/add-warehouses-tool", payload)
}

export async function updateTool(payload: UpdateRentToolRequestType) {
    return privateHttp.put("/warehouse-tool", payload)
}

export async function deleteTool(id: number) {
    return privateHttp.delete(`/warehouse-tool/${id}`)
}