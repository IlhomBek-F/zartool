import { Button, Form, Input, Space, Table } from 'antd';
import { Modal } from '../shared/Modal';
import { useRef, useState, type ChangeEvent } from 'react';
import { RentForm } from '../components/RentForm';
import type { RentType } from '../core/models/renter-model';
import type { RentToolType } from '../core/models/rent-tool-model';
import type { Query } from '../core/models/base-model';
import { TABLE_PAGE_SIZE } from '../utils/constants';
import { renterTableColumns } from '../utils/tableUtil';
import dayjs from 'dayjs';
import { useNotification } from '../hooks/useNotification';
import { useRenters } from '../hooks/useRenters';

const { Search } = Input;

function Renters() {
    const [openModal, setOpenModal] = useState(false);
    const {contextHolder, error, success} = useNotification();
    const {getData, createRent, updateRent, completeRent, deleteRent, data} = useRenters();
    const [editableRent, setEditRent] = useState<RentType | null>(null);
    const queryRef = useRef<Query>({page: 1, q: '', page_size: TABLE_PAGE_SIZE});
    const [form] = Form.useForm();

    const handleCloseRent = (id: number) => {
       completeRent(id, () => error("Error while closing rent"), () => {
          getData(queryRef.current);
          success("Rent closed successfully")
       })
    }

    const handleEditRent = ({id, phones, date, ...rest}: RentType) => {
       form.setFieldsValue(rest);
       form.setFieldValue('phone_1', phones[0])
       form.setFieldValue('phone_2', phones[1])
       form.setFieldValue('date', dayjs(date, "DD-MM-YYYY"))
       setEditRent({id, phones,date, ...rest});
       setOpenModal(true);
    }

    const handleDeleteRent = (id: number) => {
       deleteRent(id, () => error("Error while deleting rent"), () => getData(queryRef.current))
    }

    const handleConfirmModal = async () => {
        const {phone_1, phone_2, date, rent_tools, ...rest} = await form.validateFields();
        const toolQuantityToNumber = rent_tools.map((tool: RentToolType) => ({...tool, quantity: +tool.quantity}))
        const rent = {phones: 
                      [phone_1, phone_2], 
                      ...rest, 
                      rent_tools: toolQuantityToNumber, 
                      pre_payment: +rest.pre_payment, 
                      date: `${dayjs(date).format("DD-MM-YYYY")} ${dayjs(new Date()).format("HH:mm")}`};
        
        const errorResp = () => error("Error while creating new rent");
        const successResp = () => {
             form.resetFields();
             setOpenModal(false);
             getData(queryRef.current);
        };

        if(editableRent) {
          updateRent({id: editableRent.id, ...rent, active: true, created_at: editableRent.created_at}, errorResp, () => {
              successResp()
              setEditRent(null)
          })
        } else {
           createRent(rent, () => error("Error while creating new rent"), successResp)
        }
    }

    const handleCloseModal = () => {
       setOpenModal(false);
       setEditRent(null);
       form.resetFields();
    }

    const handleSearchChange = (e: ChangeEvent<HTMLInputElement>) => {
      queryRef.current.page = 1;
      queryRef.current.q = e.target.value;
      getData(queryRef.current, () => error("Error while fetching rents"));
    }

    const handlePageChange = (page: number) => {
      queryRef.current.page = page;
      getData(queryRef.current, () => error("Error while fetching rents"));
    }

    return (
         <>
          {contextHolder}
          <div className="p-4">
            <h1 className="text-2xl font-bold mb-4">Ижарачилар</h1>
             <Space direction='horizontal' className='mb-4'>
               <Button type="primary" className='!bg-green-600' icon={<i className='pi pi-plus' />} onClick={() => setOpenModal(true)}>Янги ижара яратиш</Button>
               <Search placeholder="Исм, фамилия, Телефон" allowClear onChange={handleSearchChange} style={{ width: 200 }} />
            </Space>
            <Table<RentType> pagination={{
                             pageSize: TABLE_PAGE_SIZE, 
                             onChange: (page) => handlePageChange(page), 
                             total: data?.meta.total}} 
                             columns={renterTableColumns({handleDeleteRent, handleEditRent, handleCloseRent})} 
                             dataSource={data?.rents} key={1}/>
            <Modal isOpen={openModal} 
                   handleConfirm={handleConfirmModal} 
                   handleClose={handleCloseModal}>
                  <RentForm form={form}/>
            </Modal>
        </div>
         </>
    )
}

export { Renters }