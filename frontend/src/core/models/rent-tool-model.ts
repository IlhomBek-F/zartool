export type RentToolType = {
    id: number,
    name: string,
    size: string,
    amount: number
}

export type CreateRentToolRequestType = Omit<RentToolType, 'amount' | 'id'>;
export type UpdateRentToolRequestType = Omit<RentToolType, 'amount'>;

