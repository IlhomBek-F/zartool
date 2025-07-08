import { useState } from 'react';
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  UploadOutlined,
  UserOutlined,
  VideoCameraOutlined,
} from '@ant-design/icons';
import { Button, Layout, Menu, theme } from 'antd';
import { LOGO_SHORT, LOGO_TITLE, ROUTES_PATHS } from '../../utils/constants';
import { Outlet, useNavigate } from 'react-router';

const { Header, Sider, Content } = Layout;



function MainLayout() {
  const [collapsed, setCollapsed] = useState(false);
  const { token: { colorBgContainer, borderRadiusLG }} = theme.useToken();
  const navigate = useNavigate();

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
      onClick: () => navigate(ROUTES_PATHS.SETTING)
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
          <Button
            type="text"
            icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
            onClick={() => setCollapsed(!collapsed)}
            style={{
              fontSize: '16px',
              width: 64,
              height: 64,
            }}
          />
        </Header>
        <Content
          style={{
            margin: '24px 16px',
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