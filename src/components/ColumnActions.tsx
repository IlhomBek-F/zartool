import { Flex, Tooltip, Popconfirm, Button } from "antd";

type ColumnActionsProps = {
    id: number;
    handleDeleteRent: (id: number) => void;
    handleEditRent: (id: number) => void;
    handleCloseRent: (id: number) => void;
}

function ColumnActions({handleCloseRent, handleDeleteRent, handleEditRent, id}: ColumnActionsProps) {
    
    return <Flex gap="small" wrap>
        <Tooltip title="Ижарани ўчириш">
            <Popconfirm placement="topLeft"
                        title={'Ҳақиқатдан ҳам ўчирилсинми ?'}
                        description={'Диққат: ўчирилган ижара қайта тиклаб бўлмайди.'}
                        okText="Ҳа"
                        onConfirm={() => handleDeleteRent(id)}
                        cancelText="Йўқ"
                    >
                    <Button type="primary" danger icon={<i className='pi pi-trash' />} />
            </Popconfirm>
        </Tooltip>
        <Tooltip title="Ижарани ўзгартириш">
            <Button type="primary" icon={<i className='pi pi-pencil' />} onClick={() => handleEditRent(id)}/>
        </Tooltip>
        <Tooltip title="Ижарани ёпиш">
            <Button type="primary" className='!bg-green-600' icon={<i className='pi pi-lock' />} onClick={() => handleCloseRent(id)}/>
        </Tooltip>
    </Flex>
}

export {ColumnActions}