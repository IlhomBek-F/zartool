import { createBrowserRouter } from "react-router";
import { Login } from "../pages/Login";


export const router = createBrowserRouter([
    {
        element: <div>Home page</div>,
        index: true,
    },
    {
        path: "/login",
        element: <Login />
    }
])