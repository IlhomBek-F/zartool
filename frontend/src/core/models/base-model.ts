
export type BaseModel = {
    id: number;
    created_at?: string;
    updated_at?: string;
}

export type ResponseMetaType = {
    page: number,
    total: number,
    per_page?: number
}

export type ResponseType<T = any> = {
    status: number;
    message: string;
    data: T;
    meta: ResponseMetaType
}