export interface ApiErrorBody {
  code: number;
  message: string;
  detail?: string;
}

export interface ApiResponse<T> {
  success: boolean;
  data?: T;
  message?: string;
  error?: ApiErrorBody;
}

export interface TokenResponse {
  access_token: string;
  refresh_token: string;
  expires_in: number;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  name: string;
  email: string;
  password: string;
}

export interface RefreshTokenRequest {
  refresh_token: string;
}

export interface User {
  id: string;
  name: string;
  email: string;
  avatar?: string | null;
  language: string;
  timezone: string;
  country?: string | null;
  is_premium: boolean;
  premium_until?: string | null;
  created_at: string;
  updated_at: string;
}

export class ApiClientError extends Error {
  readonly status: number;
  readonly code: number;

  constructor(status: number, message: string, code?: number) {
    super(message);
    this.name = 'ApiClientError';
    this.status = status;
    this.code = code ?? status;
  }
}
