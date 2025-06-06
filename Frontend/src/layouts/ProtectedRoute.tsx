import { useAuthStore } from '../store/UseAuthStore';
import { Navigate, Outlet } from 'react-router';

export const ProtectedRoute = () => {
  const isAuthenticated = useAuthStore(state => state.isAuthenticated);
  
  if (!isAuthenticated) {
    return <Navigate to="/login" replace />;
  }
  
  return <Outlet />;
};