import { useState } from 'react';
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  UploadOutlined,
  UserOutlined,
  VideoCameraOutlined,
} from '@ant-design/icons';
import { Button, Layout, Menu, theme } from 'antd';
import { LOGO_SHORT, LOGO_TITLE } from '../../utils/constants';
import { Outlet, useNavigate } from 'react-router';

const { Header, Sider, Content } = Layout;



function MainLayout() {
  const [collapsed, setCollapsed] = useState(false);
  const { token: { colorBgContainer, borderRadiusLG }} = theme.useToken();
  const navigate = useNavigate();

  const menuItems = [
   {
      key: '1',
      icon: <UserOutlined />,
      label: 'nav 1',
      onClick: () => navigate('./report')
   },
   {
      key: '2',
      icon: <VideoCameraOutlined />,
      label: 'nav 2',
      onClick: () => navigate('./order')
   },
   {
      key: '3',
      icon: <UploadOutlined />,
      label: 'nav 3',
      onClick: () => navigate('./setting')
   },
]

  return (
    <Layout className='h-[100vh]'>
      <Sider trigger={null} collapsible collapsed={collapsed}>
        <div className="demo-logo-vertical bg-[#535672] m-2.5 text-center rounded-[4px] text-[20px]">{collapsed ? LOGO_SHORT : LOGO_TITLE}</div>
        <Menu
          theme="dark"
          mode="inline"
          defaultSelectedKeys={['1']}
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