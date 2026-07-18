import { Platform, StyleSheet } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { AnimatedIcon } from '@/components/animated-icon';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { WebBadge } from '@/components/web-badge';
import { BottomTabInset, MaxContentWidth, Spacing } from '@/constants/theme';
import { useAuth } from '@/contexts/auth-context';

export default function HomeScreen() {
  const { user } = useAuth();

  return (
    <ThemedView style={styles.container}>
      <SafeAreaView style={styles.safeArea}>
        <ThemedView style={styles.heroSection}>
          <AnimatedIcon />
          <ThemedText type="small" themeColor="textSecondary" style={styles.eyebrow}>
            Destiny
          </ThemedText>
          <ThemedText type="title" style={styles.title}>
            Xin chào, {user?.name?.split(' ')[0] ?? 'bạn'}
          </ThemedText>
          <ThemedText themeColor="textSecondary" style={styles.subtitle}>
            Khám phá tử vi, tarot và lộ trình cá nhân của bạn.
          </ThemedText>
        </ThemedView>

        <ThemedView type="backgroundElement" style={styles.card}>
          <ThemedText type="smallBold">Hôm nay</ThemedText>
          <ThemedText themeColor="textSecondary" type="small">
            Các tính năng chi tiết sẽ sớm có mặt. Tài khoản của bạn đã sẵn sàng.
          </ThemedText>
        </ThemedView>

        {Platform.OS === 'web' && <WebBadge />}
      </SafeAreaView>
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    flexDirection: 'row',
  },
  safeArea: {
    flex: 1,
    paddingHorizontal: Spacing.four,
    alignItems: 'center',
    gap: Spacing.three,
    paddingBottom: BottomTabInset + Spacing.three,
    maxWidth: MaxContentWidth,
  },
  heroSection: {
    alignItems: 'center',
    justifyContent: 'center',
    flex: 1,
    paddingHorizontal: Spacing.four,
    gap: Spacing.three,
  },
  eyebrow: {
    letterSpacing: 2,
    textTransform: 'uppercase',
  },
  title: {
    textAlign: 'center',
    fontSize: 40,
    lineHeight: 46,
  },
  subtitle: {
    textAlign: 'center',
    maxWidth: 320,
  },
  card: {
    alignSelf: 'stretch',
    gap: Spacing.two,
    paddingHorizontal: Spacing.three,
    paddingVertical: Spacing.four,
    borderRadius: Spacing.four,
  },
});
