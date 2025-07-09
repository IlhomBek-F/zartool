import type { CreateRentToolRequestType, RentToolType, UpdateRentToolRequestType } from "../core/models/rent-tool-model";
import type { CreateRentRequestType, RentType, UpdateRentRequestType } from "../core/models/renter-model";

let data: RentType[] = [];
let tools: RentToolType[] = [];

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

export async function getRentTools(): Promise<RentToolType[]> {
    return new Promise((resolve, _) => {
        resolve(tools)
    })
}

export async function addNewTool(payload: CreateRentToolRequestType[]) {
 return new Promise((resolve, _) => {
    payload = payload.map((tool, index) => ({...tool, id: Date.now() + index}))
    tools.push(...payload as RentToolType[]);
    resolve(tools)
 })
}

export async function updateTool(payload: UpdateRentToolRequestType) {
    return new Promise((resolve, _) => {
        tools = tools.map(tool => tool.id === payload.id ? payload : tool);
        resolve(payload)
    })
}

export async function deleteTool(id: number) {
    return new Promise((resolve, _) => {
        tools = tools.filter(tool => tool.id !== id);
        resolve(id)
    })
}