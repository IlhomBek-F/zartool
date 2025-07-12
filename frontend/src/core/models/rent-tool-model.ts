import type { BaseModel } from "./base-model";

export type RentToolType = BaseModel & {
    name: string,
    size: string,
    quantity: number
}

export type CreateRentToolRequestType = Pick<RentToolType, "name" | 'size' | "quantity">;
export type UpdateRentToolRequestType = Omit<RentToolType, 'quantity'>;

