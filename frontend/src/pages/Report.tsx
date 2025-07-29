
import { Card, Col, Input, Row, Space, Statistic, Table } from 'antd';
import dayjs from 'dayjs';
import type { RentType } from '../core/models/renter-model';
import { TABLE_PAGE_SIZE } from '../utils/constants';
import { reportTableColumns } from '../utils/tableUtil';
import { useRef } from 'react';
import type { Query } from '../core/models/base-model';
import { useNotification } from '../hooks/useNotification';
import { useReport } from '../hooks/useReport';

const { Search } = Input;

function Report() {
  const {report, getData} = useReport();
  const queryRef = useRef<Query>({page: 1, q: '', page_size: TABLE_PAGE_SIZE});
  const {contextHolder, error} = useNotification();

  const handleSearch = (q: string) => {
    queryRef.current.page = 1;
    queryRef.current.q = q;
    getData(queryRef.current, () => error("Error while getting report"))
  }

  const handlePageChange = (page: number) => {
    queryRef.current.page = page
    getData(queryRef.current, () => error("Error while getting report"))
  }

    return (
      <>
       {contextHolder}
       <div className="p-4">
            <h1 className="text-2xl font-bold mb-4">Кунлик хисобот | {dayjs(new Date()).format("DD.MM.YYYY")}</h1>
            <Row gutter={16} className='mb-4'>
              <Col span={12}>
                <Card variant="outlined">
                    <Statistic title="Умумий берилган ижаралар"
                               value={report?.reportData.total_created_rent}
                               valueStyle={{ color: '#3f8600' }}
                               prefix={<i className='pi pi-address-book mr-2' />}
                    />
                </Card>
              </Col>
             <Col span={12}>
                <Card variant="outlined">
                    <Statistic title="Умумий ёпилган ижаралар"
                               value={report?.reportData.total_completed_rent}
                               valueStyle={{ color: '#cf1322' }}
                               prefix={<i className='pi pi-lock mr-2' />}
                    />
                </Card>
            </Col>
            </Row>
            <Space direction='horizontal' className='mb-4'>
               <Search placeholder="input search text" allowClear style={{ width: 200 }} onChange={(e) => handleSearch(e.target.value)}/>
            </Space>
            <Table<RentType> pagination={{
                             pageSize: TABLE_PAGE_SIZE, 
                             onChange: handlePageChange, 
                             total: report?.meta.total}} 
                             columns={reportTableColumns} 
                             dataSource={report?.reportData.rents} key={1}/>
        </div>
      </>
        
    );
}

export { Report };