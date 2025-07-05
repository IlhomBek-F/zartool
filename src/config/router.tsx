import { createBrowserRouter } from "react-router-dom";
import { Login } from "../pages/Login";
import { MainLayout } from "../components/layout/MainLayout";


export const router = createBrowserRouter([
    {
        path: "/login",
        element: <Login />
    },
    {
        path: '/main',
        element: <MainLayout />,
        children: [
            {   
                path: 'report',
                element: <div>Report page</div>
            },
            {   
                path: 'order',
                element: <div>Order page</div>
            },
            {
                path: 'setting',
                element: <div>Setting Page</div>
            }
        ]    
    }
])