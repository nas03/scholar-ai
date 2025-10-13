/**
 * Authentication types and interfaces
 */

export interface SignUpFormData {
  username: string;
  email: string;
  phone: string;
  password: string;
  confirmPassword: string;
}

export interface SignInFormData {
  email: string;
  password: string;
}

export interface FormErrors {
  username?: string;
  email?: string;
  phone?: string;
  password?: string;
  confirmPassword?: string;
}

export interface AuthResponse {
  success: boolean;
  message: string;
  data?: {
    userId: string;
    email: string;
    token?: string;
  };
}

export interface User {
  id: string;
  username: string;
  email: string;
  phone: string;
  isVerified: boolean;
  createdAt: string;
  updatedAt: string;
}
