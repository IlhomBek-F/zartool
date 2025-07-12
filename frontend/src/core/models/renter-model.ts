import type { BaseModel } from "./base-model";
import type { RentToolType } from "./rent-tool-model";

export type RentType = BaseModel & {
    full_name: string,
    address?: string,
    rent_tools: RentToolType[],
    phones: string[],
    pre_payment: string,
    active: boolean
}


export type CreateRentRequestType = Omit<RentType, 'id' | 'created_at' | 'updated_at'>

export type UpdateRentRequestType = RentType;

export type GetRentersResponseType = RentType[];
