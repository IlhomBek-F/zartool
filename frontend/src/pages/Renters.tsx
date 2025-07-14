import { Button, Form, Table } from 'antd';
import { Modal } from '../shared/Modal';
import { useEffect, useState } from 'react';
import { RentForm } from '../components/RentForm';
import { completeRent, createRent, deleteRent, getRenters, updateRent } from '../api';
import type { RentType } from '../core/models/renter-model';
import type { RentToolType } from '../core/models/rent-tool-model';
import type { ResponseMetaType } from '../core/models/base-model';
import { TABLE_PAGE_SIZE } from '../utils/constants';
import { renterTableColumns } from '../utils/tableUtil';

function Renters() {
    const [openModal, setOpenModal] = useState(false);
    const [data, setData] = useState<{meta: ResponseMetaType, rents: RentType[]}>();
    const [editableRent, setEditRent] = useState<RentType | null>(null);
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
       setEditRent({id, phones, ...rest});
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
        
        if(editableRent) {
          await updateRent({id: editableRent.id, ...rent, active: true, created_at: editableRent.created_at});
          setEditRent(null)
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
        .then(({meta, data}) => {
            setData({meta: meta, rents: data.map((r) => ({...r, key: r.id}))});
        })
    }

    return (
         <div className="p-4">
            <h1 className="text-2xl font-bold mb-4">Ижарачилар</h1>
            <Button type="primary" className='!bg-green-600 mb-2' icon={<i className='pi pi-plus' />} onClick={() => setOpenModal(true)}>Янги ижара яратиш</Button>
            <Table<RentType> pagination={{
                             pageSize: TABLE_PAGE_SIZE, 
                             onChange: (page) => getData(page), 
                             total: data?.meta.total}} 
                             columns={renterTableColumns({handleDeleteRent, handleEditRent, handleCloseRent})} 
                             dataSource={data?.rents} key={1}/>
            <Modal isOpen={openModal} 
                   handleConfirm={handleConfirmModal} 
                   handleClose={handleCloseModal}>
                  <RentForm form={form}/>
            </Modal>
        </div>
    )
}

export { Renters }