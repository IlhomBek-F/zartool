import { Button, Flex, Popconfirm, Tag, Tooltip, type TableProps } from "antd"
import type { RentType } from "../core/models/renter-model"
import type { RentToolType } from "../core/models/rent-tool-model";
import { formatDate } from "./helper";
import type { WarehouseToolType } from "../core/models/warehouse-tool-model";
import { ColumnActions, type ColumnActionsProps } from "../components/ColumnActions";

export const renterTableColumns = ({
   handleDeleteRent,
    handleEditRent,
    handleCloseRent
}: Omit<ColumnActionsProps, 'item'>) =>  [
  {
    title: 'Исм, фамилия',
    dataIndex: 'full_name',
    key: 'full_name',
    render: (value: string, record: RentType) => <span className={!record.active && 'line-through' || ''}>{value}</span>
  },
  {
    title: 'Манзил',
    dataIndex: 'address',
    key: 'address',
    render: (value: string, record: RentType) => <span className={!record.active && 'line-through' || ''}>{value}</span>
  },
  {
    title: 'Ижарага берилган нарсалар',
    key: 'tags',
    dataIndex: 'tags',
    render: (_, { rent_tools }: RentType) => (
      <>
        {rent_tools.map((tool: RentToolType, index: number) => {
          return (
            <Tag color='green' key={index}>
              {tool.name.toUpperCase()} | {tool.size} | {tool.quantity}
            </Tag>
          );
        })}
      </>
    ),
  },
  {
    title: 'Телефон',
    dataIndex: 'phone',
    key: 'phone',
    render: (_, {phones, active}: RentType) => <span className={!active && 'line-through' || ''}>{phones[0]} {phones[1] && `| ${phones[1]}`}</span>,
  },
  {
    title: 'Сана',
    dataIndex: 'created_at',
    key: 'created_at',
    render: (value: string) => <span>{formatDate(value)}</span>
  },
  {
    title: 'Бошлангич тўлов',
    dataIndex: 'pre_payment',
    key: 'pre_payment',
    render: (text: string, record: RentType) => <span className={!record.active && 'line-through' || ''}>{text} сом</span>,
  },
  {
    key: "action",
    render: (_: any, record: RentType) => <ColumnActions item={record} {...{handleCloseRent, handleEditRent, handleDeleteRent}}/>
  }
];

export const warehouseTableColumns = (
  handleEditTool: (tool: WarehouseToolType) => void, 
  handleDeleteTool: (id: number) => void) => [
        {
            title: 'Ускуна',
            dataIndex: 'name',
            key: 'name',
        },
        {
            title: 'Размер',
            dataIndex: 'size',
            key: 'size',
        },
        {
            key: 'action',
            render: (_, tool: WarehouseToolType) => (
                <Flex justify="end" gap={10}>
                    <Tooltip title="Ускунани ўзгартириш">
                        <Button type="primary" icon={<i className='pi pi-pencil' />} onClick={() => handleEditTool(tool)}/>
                     </Tooltip>
                    <Tooltip title="Ускунани ўчириш">
                        <Popconfirm placement="topLeft"
                                    title={'Ҳақиқатдан ҳам ўчирилсинми ?'}
                                    okText="Ҳа"
                                    onConfirm={() => handleDeleteTool(tool.id)}
                                    cancelText="Йўқ">
                        <Button type="primary" danger icon={<i className='pi pi-trash' />} />
                       </Popconfirm>
                    </Tooltip>
                </Flex>
            ),
        },
];