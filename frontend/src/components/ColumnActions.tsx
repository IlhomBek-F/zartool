import { Flex, Tooltip, Popconfirm, Button } from "antd";
import type { RentType } from "../core/models/renter-model";

export type ColumnActionsProps = {
    item: RentType;
    handleDeleteRent: (id: number) => void;
    handleEditRent: (rent: RentType) => void;
    handleCloseRent: (id: number) => void;
}

function ColumnActions({handleCloseRent, handleDeleteRent, handleEditRent, item}: ColumnActionsProps) {
    
    return <Flex gap="small" wrap>
        <Tooltip title="Ижарани ўчириш">
            <Popconfirm placement="topLeft"
                        title={'Ҳақиқатдан ҳам ўчирилсинми ?'}
                        description={'Диққат: ўчирилган ижара қайта тиклаб бўлмайди.'}
                        okText="Ҳа"
                        onConfirm={() => handleDeleteRent(item.id)}
                        cancelText="Йўқ"
                    >
                    <Button type="primary" danger icon={<i className='pi pi-trash' />} />
            </Popconfirm>
        </Tooltip>
        <Tooltip title="Ижарани ўзгартириш">
            <Button type="primary" disabled={!item.active} icon={<i className='pi pi-pencil' />} onClick={() => handleEditRent(item)}/>
        </Tooltip>
        <Tooltip title="Ижарани ёпиш">
            <Button type="primary" disabled={!item.active} className='!bg-green-600' icon={<i className='pi pi-lock' />} onClick={() => handleCloseRent(item.id)}/>
        </Tooltip>
    </Flex>
}

export {ColumnActions}