import { createBrowserRouter, Navigate } from "react-router-dom";
import { Login } from "../pages/Login";
import { MainLayout } from "../components/layout/MainLayout";
import { ROUTES_PATHS } from "../utils/constants";
import { Report } from "../pages/Report";
import { Renters } from "../pages/Renters";
import { Setting } from "../pages/Setting";


export const router = createBrowserRouter([
    {
        index: true,
        element: <Navigate to="/renters" replace/>
    },
    {
        path: `/${ROUTES_PATHS.LOGIN}`,
        element: <Login />
    },
    {
        path: `/`,
        element: <MainLayout />,
        children: [
            {   
                path: ROUTES_PATHS.REPORT,
                element: <Report />,
            },
            {   
                index: true,
                path: ROUTES_PATHS.RENTERS,
                element: <Renters />,
                errorElement: <div>Error page</div>
            },
            {
                path: ROUTES_PATHS.SETTING,
                element: <Setting />
            }
        ]    
    }
])