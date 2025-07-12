import type { BaseModel } from "./base-model";

export type WarehouseToolType = BaseModel & {
  name: string;
  size: string;
}

export type CreateWarehouseToolRequestType = Omit<WarehouseToolType, "id" | "createdAt" | "updatedAt">;
export type UpdateWarehouseToolRequestType = BaseModel & WarehouseToolType;

