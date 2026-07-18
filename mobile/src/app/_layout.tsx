import { DarkTheme, DefaultTheme, ThemeProvider } from 'expo-router';
import { Stack } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';
import { useEffect } from 'react';
import { useColorScheme } from 'react-native';

import { AnimatedSplashOverlay } from '@/components/animated-icon';
import { AuthProvider, useAuth } from '@/contexts/auth-context';
import { DestinyPalette } from '@/constants/destiny-theme';
import { useProtectedRoute } from '@/hooks/use-protected-route';

SplashScreen.preventAutoHideAsync();

const DestinyDarkTheme = {
  ...DarkTheme,
  colors: {
    ...DarkTheme.colors,
    primary: DestinyPalette.dark.accent,
    background: DestinyPalette.dark.background,
    card: DestinyPalette.dark.backgroundElement,
    text: DestinyPalette.dark.text,
    border: DestinyPalette.dark.border,
  },
};

const DestinyLightTheme = {
  ...DefaultTheme,
  colors: {
    ...DefaultTheme.colors,
    primary: DestinyPalette.light.accent,
    background: DestinyPalette.light.background,
    card: DestinyPalette.light.backgroundElement,
    text: DestinyPalette.light.text,
    border: DestinyPalette.light.border,
  },
};

function RootNavigator() {
  const { isLoading } = useAuth();

  useProtectedRoute();

  useEffect(() => {
    if (!isLoading) {
      SplashScreen.hideAsync();
    }
  }, [isLoading]);

  if (isLoading) {
    return null;
  }

  return (
    <>
      <AnimatedSplashOverlay />
      <Stack screenOptions={{ headerShown: false }}>
        <Stack.Screen name="(auth)" />
        <Stack.Screen name="(tabs)" />
      </Stack>
    </>
  );
}

export default function RootLayout() {
  const colorScheme = useColorScheme();

  return (
    <AuthProvider>
      <ThemeProvider value={colorScheme === 'dark' ? DestinyDarkTheme : DestinyLightTheme}>
        <RootNavigator />
      </ThemeProvider>
    </AuthProvider>
  );
}
