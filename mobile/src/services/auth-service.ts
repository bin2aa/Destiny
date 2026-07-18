import { apiRequest } from '@/services/api-client';
import { clearTokens, getRefreshToken, saveTokens } from '@/lib/token-storage';
import type {
  LoginRequest,
  RegisterRequest,
  TokenResponse,
  User,
} from '@/types/api';

export async function login(input: LoginRequest): Promise<TokenResponse> {
  const tokens = await apiRequest<TokenResponse>('/auth/login', {
    method: 'POST',
    body: input,
    skipAuth: true,
  });
  await saveTokens(tokens.access_token, tokens.refresh_token);
  return tokens;
}

export async function register(input: RegisterRequest): Promise<TokenResponse> {
  const tokens = await apiRequest<TokenResponse>('/auth/register', {
    method: 'POST',
    body: input,
    skipAuth: true,
  });
  await saveTokens(tokens.access_token, tokens.refresh_token);
  return tokens;
}

export async function refreshSession(): Promise<TokenResponse | null> {
  const refreshToken = await getRefreshToken();
  if (!refreshToken) {
    return null;
  }

  try {
    const tokens = await apiRequest<TokenResponse>('/auth/refresh', {
      method: 'POST',
      body: { refresh_token: refreshToken },
      skipAuth: true,
    });
    await saveTokens(tokens.access_token, tokens.refresh_token);
    return tokens;
  } catch {
    await clearTokens();
    return null;
  }
}

export async function logout(): Promise<void> {
  const refreshToken = await getRefreshToken();

  try {
    if (refreshToken) {
      await apiRequest<{ message?: string }>('/auth/logout', {
        method: 'POST',
        body: { refresh_token: refreshToken },
      });
    }
  } finally {
    await clearTokens();
  }
}

export async function getProfile(): Promise<User> {
  return apiRequest<User>('/users/me');
}
