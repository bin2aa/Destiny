import { useState } from 'react';
import { ScrollView, StyleSheet, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { DestinyButton } from '@/components/ui/destiny-button';
import { DestinyRadii } from '@/constants/destiny-theme';
import { BottomTabInset, MaxContentWidth, Spacing } from '@/constants/theme';
import { useAuth } from '@/contexts/auth-context';
import { ApiClientError } from '@/types/api';

export default function ProfileScreen() {
  const { user, logout } = useAuth();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string>();

  async function handleLogout() {
    setLoading(true);
    setError(undefined);
    try {
      await logout();
    } catch (err) {
      const message =
        err instanceof ApiClientError ? err.message : 'Không thể đăng xuất. Vui lòng thử lại.';
      setError(message);
    } finally {
      setLoading(false);
    }
  }

  return (
    <ThemedView style={styles.container}>
      <SafeAreaView style={styles.safeArea}>
        <ScrollView contentContainerStyle={styles.content} showsVerticalScrollIndicator={false}>
          <View style={styles.header}>
            <ThemedText type="small" themeColor="textSecondary" style={styles.eyebrow}>
              Tài khoản
            </ThemedText>
            <ThemedText type="subtitle">{user?.name ?? '—'}</ThemedText>
            <ThemedText themeColor="textSecondary">{user?.email ?? '—'}</ThemedText>
          </View>

          <ThemedView type="backgroundElement" style={styles.card}>
            <InfoRow label="Gói" value={user?.is_premium ? 'Premium' : 'Miễn phí'} />
            <InfoRow label="Ngôn ngữ" value={user?.language ?? '—'} />
            <InfoRow label="Múi giờ" value={user?.timezone ?? '—'} />
          </ThemedView>

          {error ? (
            <ThemedText accessibilityLiveRegion="polite" themeColor="error" type="small">
              {error}
            </ThemedText>
          ) : null}

          <DestinyButton label="Đăng xuất" loading={loading} variant="secondary" onPress={handleLogout} />
        </ScrollView>
      </SafeAreaView>
    </ThemedView>
  );
}

function InfoRow({ label, value }: { label: string; value: string }) {
  return (
    <View style={styles.row}>
      <ThemedText themeColor="textSecondary" type="small">
        {label}
      </ThemedText>
      <ThemedText type="smallBold">{value}</ThemedText>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  safeArea: {
    flex: 1,
  },
  content: {
    flexGrow: 1,
    paddingHorizontal: Spacing.four,
    paddingTop: Spacing.four,
    paddingBottom: BottomTabInset + Spacing.four,
    gap: Spacing.four,
    maxWidth: MaxContentWidth,
    width: '100%',
    alignSelf: 'center',
  },
  header: {
    gap: Spacing.one,
  },
  eyebrow: {
    letterSpacing: 1.5,
    textTransform: 'uppercase',
  },
  card: {
    gap: Spacing.three,
    padding: Spacing.four,
    borderRadius: DestinyRadii.card,
  },
  row: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    gap: Spacing.two,
  },
});
