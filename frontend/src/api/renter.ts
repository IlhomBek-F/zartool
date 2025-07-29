import type { Query } from "../core/models/base-model"
import type { CreateRentRequestType, RentType, UpdateRentRequestType } from "../core/models/renter-model"
import type { ResponseType } from "../core/models/base-model"
import { privateHttp } from "./http"
import type { GetReportRentResponseType } from "../core/models/rent-report-model"

export async function createRent(payload: CreateRentRequestType): Promise<ResponseType> {
  return privateHttp.post("/rental/create", payload)
}

export async function getRenters(query: Query): Promise<ResponseType<RentType[]>> {
    return privateHttp.get("/rentals", {params: query})
}

export async function updateRent(payload: UpdateRentRequestType): Promise<ResponseType> {
    return privateHttp.patch("/rental/update", payload)
}

export async function deleteRent(id: number): Promise<ResponseType> {
    return privateHttp.delete(`/rental/delete/${id}`)
}

export async function completeRent(id: number): Promise<ResponseType> {
    return privateHttp.post(`/rental/complete/${id}`)
}

export async function getRentReport(query: Query): Promise<ResponseType<GetReportRentResponseType>> {
    return privateHttp.get("/rental/report", {params: query})
}