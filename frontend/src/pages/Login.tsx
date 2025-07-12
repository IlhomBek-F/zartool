import { LockOutlined, UserOutlined } from "@ant-design/icons"
import { Button, Card, Form, Input } from "antd"
import { useForm } from "antd/es/form/Form";
import { useNavigate } from "react-router-dom"
import { LOGO_TITLE, ROUTES_PATHS } from "../utils/constants";
import { login } from "../api";
import { setToken } from "../utils/tokenUtil";

function Login() {
    const [form] = useForm()
    const navigate = useNavigate();

    const _login = async () => {
        const credential = await form.validateFields();
        const {username , password} = credential;
        
        login(username, password)
        .then(({data}) => {
            setToken(data.access_token);
            navigate(`/${ROUTES_PATHS.RENTERS}`);
        })
    }

    return (
        <div className=" fixed left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 w-full max-w-[400px]">
           <h1 className="text-[25px] text-center mb-3 dark:text-white">{LOGO_TITLE}</h1>
            <Card className="shadow-md m-4 dark:bg-gray-900">
                <h1 className="text-[25px] text-center mb-3 dark:text-white">Login</h1>
                <Form name="validateOnly" layout="vertical" form={form}
                >
                    <Form.Item
                        name="login"
                        label="Login"
                        className="dark:!text-white"
                        rules={[{ required: true, message: '' }]}
                    >
                        <Input prefix={<UserOutlined />} placeholder="login" className="h-10"/>
                    </Form.Item>
                    <Form.Item
                        className="mb-0"
                        label="Password"
                        name="password"
                        rules={[{ required: true, message: '' }]}
                    >
                        <Input.Password prefix={<LockOutlined />} type="password" placeholder="Password" className="h-10"/>
                    </Form.Item>
                    <Button className='dark:text-white' block type="primary" htmlType="submit" onClick={_login}  >Log in</Button>
                </Form>
            </Card>
        </div>
    )
}

export { Login }