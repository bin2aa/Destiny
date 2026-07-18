import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState,
  type ReactNode,
} from 'react';

import * as authService from '@/services/auth-service';
import { getAccessToken } from '@/lib/token-storage';
import { ApiClientError, type LoginRequest, type RegisterRequest, type User } from '@/types/api';

type AuthContextValue = {
  user: User | null;
  isLoading: boolean;
  isAuthenticated: boolean;
  login: (input: LoginRequest) => Promise<void>;
  register: (input: RegisterRequest) => Promise<void>;
  logout: () => Promise<void>;
  refreshProfile: () => Promise<void>;
};

const AuthContext = createContext<AuthContextValue | null>(null);

async function loadProfile(): Promise<User | null> {
  const token = await getAccessToken();
  if (!token) {
    return null;
  }

  try {
    return await authService.getProfile();
  } catch (error) {
    if (error instanceof ApiClientError && error.status === 401) {
      const refreshed = await authService.refreshSession();
      if (refreshed) {
        return authService.getProfile();
      }
    }
    await authService.logout();
    return null;
  }
}

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    let mounted = true;

    loadProfile()
      .then((profile) => {
        if (mounted) {
          setUser(profile);
        }
      })
      .finally(() => {
        if (mounted) {
          setIsLoading(false);
        }
      });

    return () => {
      mounted = false;
    };
  }, []);

  const refreshProfile = useCallback(async () => {
    const profile = await loadProfile();
    setUser(profile);
  }, []);

  const login = useCallback(async (input: LoginRequest) => {
    await authService.login(input);
    const profile = await authService.getProfile();
    setUser(profile);
  }, []);

  const register = useCallback(async (input: RegisterRequest) => {
    await authService.register(input);
    const profile = await authService.getProfile();
    setUser(profile);
  }, []);

  const logout = useCallback(async () => {
    await authService.logout();
    setUser(null);
  }, []);

  const value = useMemo(
    () => ({
      user,
      isLoading,
      isAuthenticated: Boolean(user),
      login,
      register,
      logout,
      refreshProfile,
    }),
    [user, isLoading, login, register, logout, refreshProfile],
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

export function useAuth(): AuthContextValue {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within AuthProvider');
  }
  return context;
}
