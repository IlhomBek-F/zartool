import React, { createContext, useContext } from "react";
import {login as _login} from "../api/index";
import { getToken } from "../utils/tokenUtil";

type AuthContextType = {
    isAuthenticated: () => boolean;
};

const authContext = createContext<AuthContextType>({
    isAuthenticated: () => false,
});

function AuthProvider({children}: {children: React.ReactNode}) {
    const isAuthenticated = () => !!getToken()

    return <authContext.Provider value={{isAuthenticated}}>
        {children}
    </authContext.Provider>
}

const useAuth = () => useContext(authContext);

export {AuthProvider, useAuth}