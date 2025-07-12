import { useAuth } from "../contexts/Auth";

export function PrivateRoute({children}: {children: React.ReactNode}) {
    const { isAuthenticated } = useAuth();

    if (!isAuthenticated()) {
        window.location.replace("/login")
    }

    return children
}
