import Constants from 'expo-constants';

const DEFAULT_API_URL = 'http://localhost:8080/api/v1';

function resolveApiBaseUrl(): string {
  const fromEnv = process.env.EXPO_PUBLIC_API_URL;
  if (fromEnv) {
    return fromEnv.replace(/\/$/, '');
  }

  const hostUri = Constants.expoConfig?.hostUri;
  if (hostUri) {
    const host = hostUri.split(':')[0];
    return `http://${host}:8080/api/v1`;
  }

  return DEFAULT_API_URL;
}

export const API_BASE_URL = resolveApiBaseUrl();
