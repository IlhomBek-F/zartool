
export type BaseModel = {
    id: number;
    createdAt?: string;
    updatedAt?: string;
}

export type ResponseType<T> = {
    status: number;
    message: string;
    data: T
}