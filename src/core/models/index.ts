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

export type RentToolType = {
    id: number,
    name: string,
    size: string,
    amount: number
}

export type WareHouseToolType = Omit< RentToolType, 'amount'>