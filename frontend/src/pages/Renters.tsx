import { Button, Form, Table, Tag } from 'antd';
import type { TableProps } from 'antd';
import { Modal } from '../shared/Modal';
import { ColumnActions } from '../components/ColumnActions';
import { useEffect, useState } from 'react';
import { RentForm } from '../components/RentForm';
import { completeRent, createRent, deleteRent, getRenters, updateRent } from '../api';
import type { RentType } from '../core/models/renter-model';
import type { RentToolType } from '../core/models/rent-tool-model';
import { formatDate } from '../utils/helper';
import type { ResponseMetaType } from '../core/models/base-model';
import { TABLE_PAGE_SIZE } from '../utils/constants';

const columns: TableProps<RentType>['columns'] = [
  {
    title: 'Исм, фамилия',
    dataIndex: 'full_name',
    key: 'full_name',
    render: (value, record) => <span className={!record.active && 'line-through' || ''}>{value}</span>
  },
  {
    title: 'Манзил',
    dataIndex: 'address',
    key: 'address',
    render: (value, record) => <span className={!record.active && 'line-through' || ''}>{value}</span>
  },
  {
    title: 'Ижарага берилган нарсалар',
    key: 'tags',
    dataIndex: 'tags',
    render: (_, { rent_tools }) => (
      <>
        {rent_tools.map((tool: RentToolType, index) => {
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
    render: (_, {phones, active}) => <span className={!active && 'line-through' || ''}>{phones[0]} {phones[1] && `| ${phones[1]}`}</span>,
  },
  {
    title: 'Сана',
    dataIndex: 'created_at',
    key: 'created_at',
    render: (value) => <span>{formatDate(value)}</span>
  },
  {
    title: 'Бошлангич тўлов',
    dataIndex: 'pre_payment',
    key: 'pre_payment',
    render: (text, record) => <span className={!record.active && 'line-through' || ''}>{text} сом</span>,
  },
];

function Renters() {
    const [openModal, setOpenModal] = useState(false);
    const [data, setData] = useState<{meta: ResponseMetaType, rents: RentType[]}>();
    const [editRentId, setEditRentId] = useState<number | null>(null);
    const [form] = Form.useForm();

    useEffect(() => {
        getData()
    }, [])

    const handleCloseRent = (id: number) => {
       completeRent(id)
        .then(() => {
          getData();
        })
    }

    const handleEditRent = ({id, phones, ...rest}: RentType) => {
       form.setFieldsValue(rest);
       form.setFieldValue('phone_1', phones[0])
       form.setFieldValue('phone_2', phones[1])
       setEditRentId(id);
       setOpenModal(true);
    }

    const handleDeleteRent = (id: number) => {
       deleteRent(id)
       .then(() => {
         getData();
       })
    }

    const handleConfirmModal = async () => {
        const {phone_1, phone_2, date, rent_tools, ...rest} = await form.validateFields();
        const toolQuantityToNumber = rent_tools.map((tool: RentToolType) => ({...tool, quantity: +tool.quantity}))
        const rent = {phones: [phone_1, phone_2], ...rest, rent_tools: toolQuantityToNumber, pre_payment: +rest.pre_payment};
        
        if(editRentId) {
          await updateRent({id: editRentId, ...rent, active: true});
          setEditRentId(null)
        } else {
          await createRent(rent);
        }
        
        form.resetFields();
        setOpenModal(false);
        getData();
    }

    const handleCloseModal = () => {
       setOpenModal(false)
       form.resetFields();
    }

    const getData = (page = 1) => {
        getRenters(page)
        .then(res => {
            setData({meta: res.meta, rents: res.data});
        })
    }

    return (
         <div className="p-4">
            <h1 className="text-2xl font-bold mb-4">Ижарачилар</h1>
            <Button type="primary" className='!bg-green-600 mb-2' icon={<i className='pi pi-plus' />} onClick={() => setOpenModal(true)}>Янги ижара яратиш</Button>
            <Table<RentType> pagination={{pageSize: TABLE_PAGE_SIZE, onChange: (page) => getData(page), total: data?.meta.total}} columns={columns?.concat({
                key: "action",
                render: (_, record) => <ColumnActions item={record} {...{handleCloseRent, handleEditRent, handleDeleteRent}}/>
              }
            )} dataSource={data?.rents} key={1}/>
            <Modal isOpen={openModal} 
                   handleConfirm={handleConfirmModal} 
                   handleClose={handleCloseModal}>
                  <RentForm form={form}/>
            </Modal>
        </div>
    )
}

export { Renters }