import { Button, Flex, Form, Input, Table} from "antd";
import { Modal } from "../shared/Modal";
import { useState } from "react";
import type { WarehouseToolType } from "../core/models/warehouse-tool-model";
import { TABLE_PAGE_SIZE } from "../utils/constants";
import { warehouseTableColumns } from "../utils/tableUtil";
import { useNotification } from "../hooks/useNotification";
import { useWarehouse } from "../hooks/useWarehouse";

const formItemLayout = {
  labelCol: {
    xs: { span: 24 },
    sm: { span: 10 },
  },
  wrapperCol: {
    xs: { span: 10 },
    sm: { span: 22 },
  },
};

function Warehouse() {
    const [openModal, setOpenModal] = useState(false);
    const {contextHolder, error} = useNotification();
    const {dataSource, getRentTools, updateTool, addNewTool, deleteTool} = useWarehouse();
    const [editToolId, setEditToolId] = useState<number | null>(null)
    const [form] = Form.useForm();

    const getTools = (page = 1) => {
       getRentTools(page, () => error("Error while getting rent tools"))
    }

    const handleConfirmModal = async () => {
      const formData = await form.validateFields();
      const {tools} = formData;

      const successResp = () => {
        getTools();
        setEditToolId(null);
        form.resetFields();
        setOpenModal(false);
      }
      
      const errorResp = () => error("Something went wrong. Please try again")
      
      if(editToolId) {
        updateTool({...tools[0], id: editToolId}, errorResp, successResp)
      } else {
        addNewTool(tools, errorResp, successResp)
      }
    }

    const handleDeleteTool = (id: number) => {
       deleteTool(id, () => error("Error while deleting rent tool"), getTools)
    }

    const handleEditTool = (tool: WarehouseToolType) => {
       form.setFieldsValue({tools: [tool]})
       setEditToolId(tool.id);
       setOpenModal(true);
    }

    return (
      <>
      {contextHolder}
       <div className="p-4">
            <h1 className="text-2xl font-bold mb-4">Омбор/Склад</h1>
            <Button type="primary" 
                    className='!bg-green-600 mb-2' 
                    icon={<i className='pi pi-plus' />} 
                    onClick={() => setOpenModal(true)}
                    >Янги ускуна киритиш</Button>
            <Table<WarehouseToolType> pagination={dataSource?.meta.total > 5 && {
                                      pageSize: TABLE_PAGE_SIZE, 
                                      onChange: (page: number) => getTools(page), 
                                      total: dataSource?.meta.total}} 
                                      columns={warehouseTableColumns(handleEditTool, handleDeleteTool)} 
                                      dataSource={dataSource?.data} />
            <Modal isOpen={openModal} 
                   handleClose={() => setOpenModal(false)} 
                   handleConfirm={handleConfirmModal}>
               <Form {...formItemLayout} 
                     layout='vertical' 
                     className='w-full' 
                     form={form}>
                   <Form.List name="tools" initialValue={[{ name: '', size: ''}]}>
                        {(fields, { add, remove }, index) => (
                          <>
                            {fields.map((listItem, index) => (
                                <Flex className='w-[98.5%]' key={index} align="center">
                                    <Form.Item label="Ускуна" name={[listItem.name, 'name']} className='w-full' hasFeedback  rules={[{ required: true, message: 'Ускуна номини киритинг!' }]}>
                                        <Input allowClear className='w-full' />
                                    </Form.Item >
                                    <Form.Item label="Размер" name={[listItem.name, 'size']} className='w-full' hasFeedback rules={[{ required: true, message: 'Ускуна размерини киритинг!' }]}>
                                        <Input allowClear className='w-full'/>
                                    </Form.Item>
                                    {index > 0 && <i className='pi pi-trash cursor-pointer text-red-500' onClick={() => remove(+listItem.name)} />}
                                </Flex>
                            ))}
                               {
                                !editToolId && <Form.Item className='w-full' >
                                  <Button type="dashed" className='w-full !border-green-500' onClick={() => add()} block icon={<i className='pi pi-plus' />} />
                                </Form.Item>
                               } 
                            </>
                        )}
                    </Form.List>
                </Form>
            </Modal>
        </div>
      </>
    );
}

export { Warehouse };