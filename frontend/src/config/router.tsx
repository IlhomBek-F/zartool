import { createBrowserRouter, Navigate } from "react-router-dom";
import { Login } from "../pages/Login";
import { MainLayout } from "../components/layout/MainLayout";
import { ROUTES_PATHS } from "../utils/constants";
import { PrivateRoute } from "./PrivateRoute";

export const router = createBrowserRouter([
    {
        index: true,
        element: <PrivateRoute><Navigate to="/renters" replace/></PrivateRoute>
    },
    {
        path: `/${ROUTES_PATHS.LOGIN}`,
        element: <Login />
    },
    {
        path: `/`,
        element: <PrivateRoute><MainLayout /></PrivateRoute>,
        children: [
            {   
                path: ROUTES_PATHS.REPORT,
                lazy:() => import("../pages/Report").then(({Report}) => ({Component: Report})),
                errorElement: <div>Error page</div>
            },
            {   
                index: true,
                path: ROUTES_PATHS.RENTERS,
                lazy: () => import("../pages/Renters").then(({Renters}) => ({Component: Renters})),
                errorElement: <div>Error page</div>
            },
            {
                path: ROUTES_PATHS.SETTING,
                lazy: () => import("../pages/Setting").then(({Setting}) => ({Component: Setting})),
                errorElement: <div>Error page</div>
            }
        ]    
    },
    {
        path: "*",
        element: <Navigate to="/renters"/>
    }
])