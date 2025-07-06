import { Flex, Tooltip, Popconfirm, Button } from "antd";

type ColumnActionsProps = {
    handleDeleteRent: () => void;
    handleEditRent: () => void
    handleCloseRent: () => void
}

function ColumnActions({handleCloseRent, handleDeleteRent, handleEditRent}: ColumnActionsProps) {
    
    return <Flex gap="small" wrap>
        <Tooltip title="Ижарани ўчириш">
            <Popconfirm
          placement="topLeft"
          title={'Ҳақиқатдан ҳам ўчирилсинми ?'}
          description={'Диққат: ўчирилган ижара қайта тиклаб бўлмайди.'}
          okText="Ҳа"
          onConfirm={handleDeleteRent}
          cancelText="Йўқ"
        >
            <Button type="primary" danger shape="round" icon={<i className='pi pi-trash' />} />
        </Popconfirm>
         </Tooltip>
         <Tooltip title="Ижарани ўзгартириш">
            <Button type="primary" shape="round" icon={<i className='pi pi-pencil' />} onClick={handleEditRent}/>
         </Tooltip>
          <Tooltip title="Ижарани ёпиш">
            <Button type="primary" className='!bg-green-600' shape="round" icon={<i className='pi pi-lock' />} onClick={handleCloseRent}/>
         </Tooltip>
        </Flex>
}

export {ColumnActions}