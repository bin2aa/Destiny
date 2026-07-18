import {
  ActivityIndicator,
  Pressable,
  StyleSheet,
  Text,
  type PressableProps,
  type StyleProp,
  type ViewStyle,
} from 'react-native';

import { DestinyRadii } from '@/constants/destiny-theme';
import { Spacing } from '@/constants/theme';
import { useTheme } from '@/hooks/use-theme';

type ButtonVariant = 'primary' | 'secondary' | 'ghost';
type ButtonState = 'default' | 'loading' | 'disabled' | 'error' | 'success';

export type DestinyButtonProps = Omit<PressableProps, 'children'> & {
  label: string;
  variant?: ButtonVariant;
  loading?: boolean;
  success?: boolean;
  error?: boolean;
  style?: StyleProp<ViewStyle>;
};

function resolveState({
  disabled,
  loading,
  success,
  error,
}: Pick<DestinyButtonProps, 'disabled' | 'loading' | 'success' | 'error'>): ButtonState {
  if (disabled) return 'disabled';
  if (loading) return 'loading';
  if (success) return 'success';
  if (error) return 'error';
  return 'default';
}

export function DestinyButton({
  label,
  variant = 'primary',
  loading = false,
  success = false,
  error = false,
  disabled,
  style,
  ...rest
}: DestinyButtonProps) {
  const theme = useTheme();
  const state = resolveState({ disabled, loading, success, error });
  const isDisabled = state === 'disabled' || state === 'loading';

  const backgroundColor =
    variant === 'primary'
      ? state === 'success'
        ? theme.success
        : state === 'error'
          ? theme.error
          : theme.accent
      : variant === 'secondary'
        ? theme.backgroundElement
        : 'transparent';

  const textColor =
    variant === 'primary'
      ? theme.accentForeground
      : state === 'success'
        ? theme.success
        : state === 'error'
          ? theme.error
          : theme.text;

  return (
    <Pressable
      accessibilityRole="button"
      accessibilityState={{ disabled: isDisabled, busy: loading }}
      disabled={isDisabled}
      style={({ pressed }) => [
        styles.base,
        variant === 'secondary' && {
          borderWidth: 1,
          borderColor: theme.border,
        },
        {
          backgroundColor,
          opacity: pressed && !isDisabled ? 0.88 : state === 'disabled' ? 0.5 : 1,
          transform: [{ scale: pressed && !isDisabled ? 0.985 : 1 }],
        },
        style,
      ]}
      {...rest}>
      {loading ? (
        <ActivityIndicator color={textColor} size="small" />
      ) : (
        <Text style={[styles.label, { color: textColor }]}>
          {success ? 'Hoàn tất' : error ? 'Thử lại' : label}
        </Text>
      )}
    </Pressable>
  );
}

const styles = StyleSheet.create({
  base: {
    minHeight: 48,
    borderRadius: DestinyRadii.button,
    alignItems: 'center',
    justifyContent: 'center',
    paddingHorizontal: Spacing.four,
  },
  label: {
    fontSize: 16,
    fontWeight: '600',
    letterSpacing: 0.2,
  },
});
