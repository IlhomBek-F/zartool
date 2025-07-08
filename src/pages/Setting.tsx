import { Button, Flex, Form, Input, Popconfirm, Table, Tooltip, type TableProps } from "antd";
import { Modal } from "../shared/Modal";
import { useState } from "react";

interface DataType {
  id: string;
  name: string;
  size: string
}

const data: DataType[] = [
  {
    id: '1',
    name: 'Мышалка',
    size: '1.8',
  },
  {
    id: '2',
    name: 'Опалавка',
    size: '2 x 3'
  },
  {
    id: '3',
    name: 'Леса',
    size: '1 x 1.5'
  },
];

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
    const [form] = Form.useForm();

    const handleConfirmModal = () => {
        setOpenModal(false)
    }

    const columns: TableProps<DataType>['columns'] = [
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
            render: (_, record) => (
                <Flex justify="end" gap={10}>
                    <Tooltip title="Ускунани ўзгартириш">
                        <Button type="primary" icon={<i className='pi pi-pencil' />} onClick={() => setOpenModal(true)}/>
                     </Tooltip>
                    <Tooltip title="Ускунани ўчириш">
                        <Popconfirm placement="topLeft"
                                title={'Ҳақиқатдан ҳам ўчирилсинми ?'}
                                okText="Ҳа"
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
            <Table<DataType> columns={columns} dataSource={data} />
            <Modal isOpen={openModal} handleClose={() => setOpenModal(false)} handleConfirm={handleConfirmModal}>
               <Form {...formItemLayout} layout='vertical' className='w-full' form={form}>
                   <Form.List name="tools" initialValue={[{ tool: '', size: ''}]}>
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
                                <Form.Item className='w-full'>
                                  <Button type="dashed" className='w-full !border-green-500' onClick={() => add()} block icon={<i className='pi pi-plus' />} />
                                </Form.Item>
                            </>
                        )}
                    </Form.List>
                </Form>
            </Modal>
        </div>
    );
}

export { Setting };