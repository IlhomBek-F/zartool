import { Button, Form, Table, Tag } from 'antd';
import type { TableProps } from 'antd';
import { Modal } from '../shared/Modal';
import { ColumnActions } from '../components/ColumnActions';
import { useEffect, useState } from 'react';
import { RentForm } from '../components/RentForm';
import { completeRent, createRent, deleteRent, getRenters, updateRent } from '../api';
import type { RentType } from '../core/models/renter-model';


const columns: TableProps<RentType>['columns'] = [
  {
    title: 'Исм, фамилия',
    dataIndex: 'full_name',
    key: 'full_name',
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
    render: (_, { rent_tools }) => (
      <>
        {rent_tools.map((tag, index) => {
          return (
            <Tag color='green' key={index}>
              {tag.name.toUpperCase()}
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
    render: (_, {phones}) => <span>{phones[0]} {phones[1] && `| ${phones[1]}`}</span>,
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

function Renters() {
    const [openModal, setOpenModal] = useState(false);
    const [data, setData] = useState<any>([]);
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

    const handleEditRent = (id: number) => {
       const rental = data.find((rent: RentType) => rent.id === id);
       form.setFieldsValue(rental);
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
        const {phone_1, phone_2, date, ...rest} = await form.validateFields();
        const rent = {phones: [phone_1, phone_2], ...rest};
        
        if(editRentId) {
          await updateRent({id: editRentId, ...rent});
          setEditRentId(null)
        } else {
          await createRent(rent);
        }
        
        form.resetFields();
        setOpenModal(false);
        getData();
    }

    const getData = () => {
        getRenters()
        .then(res => {
            setData(res.data.map((r) => ({...r, key: r.id})));
        })
    }

    return (
         <div className="p-4">
            <h1 className="text-2xl font-bold mb-4">Ижарачилар</h1>
            <Button type="primary" className='!bg-green-600 mb-2' icon={<i className='pi pi-plus' />} onClick={() => setOpenModal(true)}>Янги ижара яратиш</Button>
            <Table<RentType> columns={columns?.concat({
                key: "action",
                render: (_, record) => <ColumnActions id={record.id} {...{handleCloseRent, handleEditRent, handleDeleteRent}}/>
              }
            )} dataSource={data} key={1}/>
            <Modal isOpen={openModal} 
                   handleConfirm={handleConfirmModal} 
                   handleClose={() => setOpenModal(false)}>
                  <RentForm form={form}/>
            </Modal>
        </div>
    )
}

export { Renters }