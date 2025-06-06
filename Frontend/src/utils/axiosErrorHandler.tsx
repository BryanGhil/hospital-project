// utils/axiosErrorHandler.ts
import { AxiosError } from 'axios';

// Define the expected error response structure
interface BackendErrorItem {
  message: string;
}

interface BackendErrorResponse {
  error?: BackendErrorItem[];
  message?: string;
}

export const handleAxiosError = (error: unknown): string => {
  // Default message
  let message = "An unexpected error occurred";
  
  // Handle string errors
  if (typeof error === 'string') return error;
  
  // Handle Axios errors
  if (error instanceof AxiosError) {
    // Network error (no response)
    if (!error.response) {
      return "Network error - please check your connection";
    }

    const status = error.response.status;
    const data = error.response.data as BackendErrorResponse | undefined;

    // Handle specific status codes
    switch (status) {
      case 400:
        message = "Invalid request parameters";
        break;
      case 401:
      case 403:
        // Special handling for authentication errors
        if (data?.error?.length) {
          return data.error[0].message; // "invalid credentials"
        }
        return status === 401 ? "Invalid credentials" : "Access forbidden";
      case 404:
        message = "Resource not found";
        break;
      case 500:
        message = "Internal server error";
        break;
      default:
        message = `Request failed with status ${status}`;
    }

    // Try to extract message from response
    if (data?.error?.length) {
      return data.error[0].message;
    }
    if (data?.message) {
      return data.message;
    }
  }
  
  // Handle native JavaScript errors
  if (error instanceof Error) {
    return error.message;
  }
  
  return message;
};