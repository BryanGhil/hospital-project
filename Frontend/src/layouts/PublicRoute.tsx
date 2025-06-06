import { useAuthStore } from '../store/UseAuthStore';
import { Navigate, Outlet } from 'react-router';

export const PublicRoute = () => {
  const isAuthenticated = useAuthStore(state => state.isAuthenticated);
  
  // Redirect to home if user is already authenticated
  if (isAuthenticated) {
    return <Navigate to="/" replace />;
  }
  
  return <Outlet />;
};