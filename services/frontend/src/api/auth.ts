import { apiClient } from './client';
import type { AuthResponse, LoginRequest, SignupRequest, User } from '../types';

export const login = (data: LoginRequest): Promise<AuthResponse> =>
  apiClient('/api/auth/login', {
    method: 'POST',
    body: JSON.stringify(data),
  });

export const signup = (data: SignupRequest): Promise<AuthResponse> =>
  apiClient('/api/auth/signup', {
    method: 'POST',
    body: JSON.stringify(data),
  });

export const getMe = (): Promise<User> =>
  apiClient('/api/me');
