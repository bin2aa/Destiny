import { Link, router } from 'expo-router';
import { useState } from 'react';
import { Pressable, StyleSheet, View } from 'react-native';

import { AuthLayout } from '@/components/auth/auth-layout';
import { ThemedText } from '@/components/themed-text';
import { DestinyButton } from '@/components/ui/destiny-button';
import { DestinyTextField } from '@/components/ui/destiny-text-field';
import { Spacing } from '@/constants/theme';
import { useAuth } from '@/contexts/auth-context';
import { ApiClientError } from '@/types/api';

export default function LoginScreen() {
  const { login } = useAuth();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [emailError, setEmailError] = useState<string>();
  const [passwordError, setPasswordError] = useState<string>();
  const [formError, setFormError] = useState<string>();
  const [loading, setLoading] = useState(false);

  function validate(): boolean {
    let valid = true;
    setEmailError(undefined);
    setPasswordError(undefined);
    setFormError(undefined);

    if (!email.trim()) {
      setEmailError('Vui lòng nhập email');
      valid = false;
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.trim())) {
      setEmailError('Email không hợp lệ');
      valid = false;
    }

    if (!password) {
      setPasswordError('Vui lòng nhập mật khẩu');
      valid = false;
    }

    return valid;
  }

  async function handleSubmit() {
    if (!validate()) {
      return;
    }

    setLoading(true);
    try {
      await login({ email: email.trim(), password });
      router.replace('/');
    } catch (error) {
      const message =
        error instanceof ApiClientError
          ? error.message
          : 'Không thể đăng nhập. Vui lòng thử lại.';
      setFormError(message);
    } finally {
      setLoading(false);
    }
  }

  return (
    <AuthLayout
      footer={
        <View style={styles.footer}>
          <ThemedText themeColor="textSecondary">Chưa có tài khoản?</ThemedText>
          <Link href="/register" asChild>
            <Pressable accessibilityRole="link">
              <ThemedText type="linkPrimary">Đăng ký</ThemedText>
            </Pressable>
          </Link>
        </View>
      }>
      <View style={styles.hero}>
        <ThemedText type="small" themeColor="textSecondary" style={styles.eyebrow}>
          Destiny
        </ThemedText>
        <ThemedText type="subtitle" style={styles.title}>
          Khám phá vận mệnh của bạn
        </ThemedText>
        <ThemedText themeColor="textSecondary">
          Đăng nhập để xem tử vi, tarot và lộ trình cá nhân.
        </ThemedText>
      </View>

      <View style={styles.form}>
        <DestinyTextField
          autoCapitalize="none"
          autoComplete="email"
          keyboardType="email-address"
          label="Email"
          value={email}
          onChangeText={setEmail}
          errorMessage={emailError}
        />
        <DestinyTextField
          autoComplete="password"
          label="Mật khẩu"
          secureTextEntry
          value={password}
          onChangeText={setPassword}
          errorMessage={passwordError}
        />
        {formError ? (
          <ThemedText accessibilityLiveRegion="polite" themeColor="error" type="small">
            {formError}
          </ThemedText>
        ) : null}
        <DestinyButton label="Đăng nhập" loading={loading} onPress={handleSubmit} />
      </View>
    </AuthLayout>
  );
}

const styles = StyleSheet.create({
  hero: {
    gap: Spacing.two,
  },
  eyebrow: {
    letterSpacing: 2,
    textTransform: 'uppercase',
  },
  title: {
    maxWidth: 320,
  },
  form: {
    gap: Spacing.three,
  },
  footer: {
    marginTop: Spacing.five,
    flexDirection: 'row',
    justifyContent: 'center',
    gap: Spacing.one,
    alignItems: 'center',
  },
});
