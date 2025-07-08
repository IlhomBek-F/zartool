
import { Card, Col, Input, Row, Space, Statistic, Table, Tag, type TableProps } from 'antd';

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

const { Search } = Input;

function Report() {
    const date = new Date();
    const [day, month, year] = [`${date.getDate()}`.padStart(2, '0'),
                                `${date.getMonth() + 1}`.padStart(2, '0'),
                                `${date.getFullYear()}`.padStart(2, '0')];
    return (
        <div className="p-4">
            <h1 className="text-2xl font-bold mb-4">Кунлик хисобот | {`${day}.${month}.${year}`}</h1>
            <Row gutter={16} className='mb-4'>
              <Col span={12}>
                <Card variant="outlined">
                    <Statistic title="Умумий берилган ижаралар"
                               value={11}
                               valueStyle={{ color: '#3f8600' }}
                               prefix={<i className='pi pi-address-book mr-2' />}
                    />
                </Card>
              </Col>
             <Col span={12}>
                <Card variant="outlined">
                    <Statistic title="Умумий ёпилган ижаралар"
                               value={9.3}
                               valueStyle={{ color: '#cf1322' }}
                               prefix={<i className='pi pi-lock mr-2' />}
                    />
                </Card>
            </Col>
            </Row>
            <Space direction='horizontal' className='mb-4'>
               <Search placeholder="input search text" allowClear style={{ width: 200 }} />
            </Space>
            <Table<DataType> columns={columns} dataSource={data} />
        </div>
    );
}

export { Report };