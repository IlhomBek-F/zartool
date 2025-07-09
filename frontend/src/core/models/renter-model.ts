import type { RentToolType } from "./rent-tool-model";

export type RentType = {
    id: number,
    full_name: string,
    address?: string,
    tools: RentToolType[],
    phones: string[],
    date: string,
    pre_payment: string,
    status: 'closed' | 'open'
}


export type CreateRentRequestType = Omit<RentType, 'id'>

export type UpdateRentRequestType = RentType;

export type GetRentersResponseType = RentType[];