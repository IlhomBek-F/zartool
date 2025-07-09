import { Button, Form, Table, Tag } from 'antd';
import type { TableProps } from 'antd';
import { Modal } from '../shared/Modal';
import { ColumnActions } from '../components/ColumnActions';
import { useEffect, useState } from 'react';
import type { ColumnsType } from 'antd/es/table';
import { RentForm } from '../components/RentForm';
import { closeRent, createRent, deleteRent, getRenters, updateRent } from '../api';
import type { RentType } from '../core/models/rent-tool-model';


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
    render: (_, { tools }) => (
      <>
        {tools.map((tag, index) => {
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
    const [data, setData] = useState<RentType[]>([]);
    const [editRentId, setEditRentId] = useState<number | null>(null);
    const [form] = Form.useForm();

    useEffect(() => {
        getData()
    }, [])

    const handleCloseRent = (id: number) => {
       closeRent(id)
        .then(() => {
          getData();
        })
    }

    const handleEditRent = (id: number) => {
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
        await form.validateFields();
        const {phone_1, phone_2, date, ...rest} = form.getFieldsValue();
        const rent = {phones: [phone_1, phone_2], date: date.format('MM-DD-YYYY'), ...rest};
        
        if(editRentId) {
          await updateRent({id: editRentId, ...rent});
          setEditRentId(null)
        } else {
          await createRent(rent);
        }

        setOpenModal(false);
        getData();
    }

    const getData = () => {
        getRenters()
        .then((res: RentType[]) => {
            setData(() => res.map((r) => ({...r, key: r.id})));
            addActionColumnToTable();
        })
    }

    const addActionColumnToTable = () => {
      const tableColumns = columns as ColumnsType;
      tableColumns[tableColumns.length - 1].render = (_, rent) => <ColumnActions id={rent.id} {...{handleCloseRent, handleDeleteRent, handleEditRent}}/>
    }

    return (
         <div className="p-4">
            <h1 className="text-2xl font-bold mb-4">Ижарачилар</h1>
            <Button type="primary" className='!bg-green-600 mb-2' icon={<i className='pi pi-plus' />} onClick={() => setOpenModal(true)}>Янги ижара яратиш</Button>
            <Table<RentType> columns={columns} dataSource={data} key={1}/>
            <Modal isOpen={openModal} 
                   handleConfirm={handleConfirmModal} 
                   handleClose={() => setOpenModal(false)}>
                  <RentForm form={form}/>
            </Modal>
        </div>
    )
}

export { Renters }