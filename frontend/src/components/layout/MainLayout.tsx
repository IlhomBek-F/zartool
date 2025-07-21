import { useState } from 'react';
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
} from '@ant-design/icons';
import { Button, Flex, Layout, Menu, theme } from 'antd';
import { LOGO_SHORT, LOGO_TITLE, ROUTES_PATHS } from '../../utils/constants';
import { Outlet, useNavigate } from 'react-router';
import { clearToken } from '../../utils/tokenUtil';

const { Header, Sider, Content } = Layout;


function MainLayout() {
  const [collapsed, setCollapsed] = useState(false);
  const { token: { colorBgContainer, borderRadiusLG }} = theme.useToken();
  const navigate = useNavigate();
  
  const logout = () => {
     clearToken()
     navigate(ROUTES_PATHS.LOGIN)
  }

  const menuItems = [
   {
      key: '1',
      icon: <i className='pi pi-chart-line' />,
      label: 'Кунлик хисобот',
      onClick: () => navigate(ROUTES_PATHS.REPORT),
   },
   {
      key: '2',
      icon: <i className='pi pi-users' />,
      label: 'Ижарачилар',
      onClick: () => navigate(ROUTES_PATHS.RENTERS)
   },
   {
      key: '3',
      icon: <i className='pi pi-cog' />,
      label: 'Омбор/склад',
      onClick: () => navigate(ROUTES_PATHS.Warehouse)
   },
]

  return (
    <Layout className='h-[100vh]'>
      <Sider trigger={null} collapsible collapsed={collapsed}>
        <div className="demo-logo-vertical bg-[#535672] m-2.5 text-center rounded-[4px] text-[20px]">{collapsed ? LOGO_SHORT : LOGO_TITLE}</div>
        <Menu
          theme="dark"
          mode="inline"
          defaultSelectedKeys={['2']}
          items={menuItems}
        />
      </Sider>
      <Layout>
        <Header style={{ padding: 0, background: colorBgContainer }}>
          <Flex align='center' justify='space-between' className='w-full'>
          <Button
            type="text"
            icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
            onClick={() => setCollapsed(!collapsed)}
            className='!text-[16px] !size-16'
          />
          <Button
            type="text"
            icon={<i className='pi pi-sign-out cursor-pointer '></i>}
            onClick={logout}
            className='!text-[16px] !size-16'
          />
          </Flex>
          
        </Header>
        <Content
          style={{
            margin: '24px 16px',
            overflowY: 'auto',
            padding: 24,
            minHeight: 280,
            background: colorBgContainer,
            borderRadius: borderRadiusLG,
          }}
        >
          <Outlet />
        </Content>
      </Layout>
    </Layout>
  );
};

export {MainLayout};