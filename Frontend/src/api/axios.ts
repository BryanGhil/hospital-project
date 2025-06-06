import axios from "axios";
import { useAuthStore } from "../store/UseAuthStore";
import toast from "react-hot-toast";


const instance = axios.create({
  baseURL: "http://localhost:8000/api/v1",
  withCredentials: true,
});

// Add request interceptor to inject token
instance.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Add response interceptor to handle token expiration
instance.interceptors.response.use(
  response => response,
  error => {
    if (error.response?.status === 401) {
      // Token expired or invalid
      useAuthStore.getState().clearToken();
      toast.error("Token expired, please login again")
      // window.location.href = "/login";
    }
    return Promise.reject(error);
  }
);

export default instance;