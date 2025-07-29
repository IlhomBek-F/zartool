import type { CreateRentToolRequestType, RentToolType, UpdateRentToolRequestType } from "../core/models/rent-tool-model"
import type { ResponseType } from "../core/models/base-model"
import { privateHttp } from "./http"
import { TABLE_PAGE_SIZE } from "../utils/constants"

export async function getRentTools(page = 1): Promise<ResponseType<RentToolType[]>> {
    return privateHttp.get("/warehouse-tools", {params: {page, page_size: TABLE_PAGE_SIZE}})
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