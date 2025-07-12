import {
    Button,
  DatePicker,
  Flex,
  Form,
  Input,
  Select,
  type FormInstance,
} from 'antd';
import dayjs from 'dayjs';
import { useEffect, useState } from 'react';
import { getRentTools } from '../api';
import type { WarehouseToolType } from '../core/models/warehouse-tool-model';

const { Option } = Select;

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

function RentForm({form}: {form: FormInstance}) {
  const [tools, setTools] = useState<WarehouseToolType[]>([]);

  useEffect(() => {
    getRentTools()
     .then(({data}) => {
       setTools(data)
     })
  }, [])
 
  return <Form {...formItemLayout} layout='vertical' className='w-full' form={form}>
       <Flex className='w-full'>
           <Form.Item label="Исм, фамилия" name="full_name" className='w-full' hasFeedback  rules={[{ required: true, message: 'Илтимос исм ёки фамилияни киритинг!' }]}>
                <Input allowClear placeholder="ичарачини исм ёки фамилияси" className='w-full' />
            </Form.Item >
            <Form.Item label="Манзил" name="address" className='w-full' >
                <Input allowClear placeholder="ичарачини манзили" className='w-full'/>
            </Form.Item>
       </Flex>

       <Flex className='w-full'>
           <Form.Item name="phone_1" label="Тел 1" className='w-full' hasFeedback  rules={[{ required: true, message: 'Илтимос тел ракамини киритинг!' }]}>
                <Input addonBefore={"+992"} style={{ width: '100%' }} />
            </Form.Item>
            <Form.Item name="phone_2" label="Тел 2" className='w-full'>
                <Input addonBefore={"+992"} style={{ width: '100%' }} />
            </Form.Item>
       </Flex>

         <Flex className='w-full'>
           <Form.Item name="date" label="Сана" initialValue={dayjs()} className='w-full' hasFeedback rules={[{ required: true, message: 'Илтимос санани киритинг!' }]}>
               <DatePicker format={'MM-DD-YYYY'} className='w-full' />
            </Form.Item>
            <Form.Item name="initial_payment" initialValue={0} label="Бошлангич тўлов" className='w-full'>
                <Input style={{ width: '100%' }} type='number'/>
            </Form.Item>
        </Flex>
        
        <Form.List name="rent_tools" initialValue={[{ name: '', size: '', quantity: ''}]}>
            {(fields, { add, remove }) => (
            <>
             {fields.map((listItem, index) => (
                <Flex align='center' className='w-[98.5%]' key={index}>
                    <Form.Item name={[listItem.name, 'name']} label="Ускуна" hasFeedback className='w-full' rules={[{required: true, message: ''}]}>
                        <Select placeholder="ускуна" allowClear>
                           {tools.map((tool) => <Option value={tool.name}>{tool.name}</Option>)}
                        </Select>
                    </Form.Item>
                    <Form.Item name={[listItem.name, 'size']} label="Размер" hasFeedback className='w-full' rules={[{required: true, message: ''}]}>
                        <Select placeholder="размер" allowClear>
                           {tools.map((tool) => <Option value={tool.size}>{tool.size}</Option>)}
                        </Select>
                    </Form.Item>
                    <Form.Item name={[listItem.name, 'quantity']} label="Дона" className='w-full' hasFeedback rules={[{required: true, message: ''}]}>
                            <Input style={{ width: '100%' }} type='number' placeholder='дона' allowClear/>
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
}

export {RentForm}