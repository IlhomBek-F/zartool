import { Button, Flex, Popconfirm, Table, Tag, Tooltip } from 'antd';
import type { TableProps } from 'antd';
import { Modal } from '../shared/Modal';
import { ColumnActions } from '../components/ColumnActions';
import { useEffect, useState } from 'react';
import type { ColumnsType } from 'antd/es/table';
import { RentForm } from '../components/RentForm';
import { useForm } from 'antd/es/form/Form';


interface DataType {
  key: string;
  name: string;
  address: string;
  tags: string[];
  phone: string,
  date: string,
  initial_payment: string
}

const columns: TableProps<DataType>['columns'] = [
  {
    title: 'Исм, фамилия',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: 'Манзил',
    dataIndex: 'address',
    key: 'address',
  },
  {
    title: 'Ижарага берилган нарсалар',
    key: 'tags',
    dataIndex: 'tags',
    render: (_, { tags }) => (
      <>
        {tags.map((tag) => {
          let color = tag.length > 5 ? 'geekblue' : 'green';
          if (tag === 'loser') {
            color = 'volcano';
          }
          return (
            <Tag color={color} key={tag}>
              {tag.toUpperCase()}
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
  },
  {
    title: 'Сана',
    dataIndex: 'date',
    key: 'date',
  },
  {
    title: 'Бошлангич тўлов',
    dataIndex: 'initial_payment',
    key: 'initial_payment',
    render: (text) => <span>{text} сом</span>,
  },
  {
    key: 'action',
  },
];

const data: DataType[] = [
  {
    key: '1',
    name: 'John Brown',
    address: 'New York No. 1 Lake Park',
    tags: ['мишалка 2', 'леса 20'],
    phone: '123-456-7890',
    date: '2023-10-01',
    initial_payment: '100'
  },
  {
    key: '2',
    name: 'Jim Green',
    address: 'London No. 1 Lake Park',
    tags: ['опаловка | 2 x 190 | 10'],
    phone: '123-456-7890',
    date: '2023-10-01',
    initial_payment: '300'
  },
  {
    key: '3',
    name: 'Joe Black',
    address: 'Sydney No. 1 Lake Park',
    tags: ['леса', 'опаловка'],
    phone: '123-456-7890',
    date: '2023-10-01',
    initial_payment: '500'
  },
];

function Renters() {
    const [openModal, setOpenModal] = useState(false);
    const [form] = useForm();

    useEffect(() => {
        addActionColumnToTable()
    }, [])

    const handleCloseRent = () => {

    }

    const handleEditRent = () => {
       setOpenModal(true)
    }

    const handleDeleteRent = () => {

    }

    const handleAddNewRent = () => {
        setOpenModal(true)
    }

    const handleConfirmModal = (value: any) => {
      console.log(form.getFieldsValue())
        setOpenModal(false)
    }

    const addActionColumnToTable = () => {
      const tableColumns = columns as ColumnsType;
      tableColumns[tableColumns.length - 1].render = (_, record) => <ColumnActions {...{handleCloseRent, handleDeleteRent, handleEditRent}}/>
    }

    return (
         <div className="p-4">
            <h1 className="text-2xl font-bold mb-4">Ижарачилар</h1>
            <Button type="primary" className='!bg-green-600 mb-2' icon={<i className='pi pi-plus' />} onClick={handleAddNewRent}>Янги ижара яратиш</Button>
            <Table<DataType> columns={columns} dataSource={data} />
            <Modal isOpen={openModal} handleConfirm={handleConfirmModal} handleClose={() => setOpenModal(false)}>
                <RentForm form={form}/>
            </Modal>
        </div>
    )
}

export { Renters }