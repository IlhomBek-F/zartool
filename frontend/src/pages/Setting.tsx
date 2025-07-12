import { Button, Flex, Form, Input, Popconfirm, Table, Tooltip, type TableProps } from "antd";
import { Modal } from "../shared/Modal";
import { useEffect, useState } from "react";
import { addNewTool, deleteTool, getRentTools, updateTool } from "../api";
import type { WarehouseToolType } from "../core/models/warehouse-tool-model";

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

function Setting() {
    const [openModal, setOpenModal] = useState(false);
    const [tools, setTools] = useState<WarehouseToolType[]>([]);
    const [editToolId, setEditToolId] = useState<number | null>(null)
    const [form] = Form.useForm();

     useEffect(() => {
       getTools();
    }, [])

    const getTools = () => {
       getRentTools()
        .then(({data}) => {
           setTools(data.map((tool) => ({...tool, key: tool.id})))
        })
    }

    const handleConfirmModal = async () => {
      const formData = await form.validateFields();
      const {tools} = formData;

      const action = editToolId ? updateTool({...tools[0], id: editToolId}) : addNewTool(tools)
      
      action.then(() => {
          getTools();
          setEditToolId(null);
          form.resetFields();
          setOpenModal(false);
      })
    }

    const handleDeleteTool = (id: number) => {
       deleteTool(id)
       .then(() => {
         getTools()
       })
    }

    const handleEditTool = (tool: WarehouseToolType) => {
       form.setFieldsValue({tools: [tool]})
       setEditToolId(tool.id);
       setOpenModal(true);
    }

    const columns: TableProps<WarehouseToolType>['columns'] = [
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
            render: (_, tool) => (
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

    return (
        <div className="p-4">
            <h1 className="text-2xl font-bold mb-4">Омбор/Склад</h1>
            <Button type="primary" 
                    className='!bg-green-600 mb-2' 
                    icon={<i className='pi pi-plus' />} 
                    onClick={() => setOpenModal(true)}
                    >Янги ускуна киритиш</Button>
            <Table<WarehouseToolType> columns={columns} dataSource={tools} />
            <Modal isOpen={openModal} handleClose={() => setOpenModal(false)} handleConfirm={handleConfirmModal}>
               <Form {...formItemLayout} layout='vertical' className='w-full' form={form}>
                   <Form.List name="tools" initialValue={[{ name: '', size: ''}]}>
                        {(fields, { add, remove }) => (
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
                                !editToolId && <Form.Item className='w-full'>
                                  <Button type="dashed" className='w-full !border-green-500' onClick={() => add()} block icon={<i className='pi pi-plus' />} />
                                </Form.Item>
                               } 
                            </>
                        )}
                    </Form.List>
                </Form>
            </Modal>
        </div>
    );
}

export { Setting };