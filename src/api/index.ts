import type { RentType, WareHouseToolType } from "../core/models";

let data: RentType[] = [];
let tools: WareHouseToolType[] = [];

export async function createRent(payload: RentType) {
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

export async function updateRent(payload: RentType) {
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

export async function getRentTools(): Promise<WareHouseToolType[]> {
    return new Promise((resolve, _) => {
        resolve(tools)
    })
}

export async function addNewTool(payload: WareHouseToolType[]) {
 return new Promise((resolve, _) => {
    payload = payload.map((tool, index) => ({...tool, id: Date.now() + index}))
    tools.push(...payload);
    resolve(tools)
 })
}

export async function updateTool(payload: any) {
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