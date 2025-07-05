import { createBrowserRouter } from "react-router-dom";
import { Login } from "../pages/Login";
import { MainLayout } from "../components/layout/MainLayout";
import { ROUTES_PATHS } from "../utils/constants";
import { Report } from "../pages/Report";
import { Renters } from "../pages/Renters";
import { Setting } from "../pages/Setting";


export const router = createBrowserRouter([
    {
        path: `/${ROUTES_PATHS.LOGIN}`,
        element: <Login />
    },
    {
        path: `/${ROUTES_PATHS.MAIN}`,
        element: <MainLayout />,
        children: [
            {   
                path: ROUTES_PATHS.REPORT,
                element: <Report />
            },
            {   
                path: ROUTES_PATHS.RENTERS,
                element: <Renters />
            },
            {
                path: ROUTES_PATHS.SETTING,
                element: <Setting />
            }
        ]    
    }
])