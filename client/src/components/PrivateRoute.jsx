import { Navigate, Outlet } from "react-router-dom";
import useLogin from "../stores/useStore";

const PrivateRoute = () => {
  const { isLoggedIn } = useLogin();
  return isLoggedIn ? <Outlet /> : <Navigate to="/posts" />;
};

export default PrivateRoute;
