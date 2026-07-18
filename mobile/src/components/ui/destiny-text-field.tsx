import { useState } from 'react';
import {
  StyleSheet,
  Text,
  TextInput,
  View,
  type TextInputProps,
} from 'react-native';

import { DestinyRadii } from '@/constants/destiny-theme';
import { Spacing } from '@/constants/theme';
import { useTheme } from '@/hooks/use-theme';

export type DestinyTextFieldProps = TextInputProps & {
  label: string;
  errorMessage?: string;
  successMessage?: string;
};

export function DestinyTextField({
  label,
  errorMessage,
  successMessage,
  editable = true,
  onFocus,
  onBlur,
  style,
  ...rest
}: DestinyTextFieldProps) {
  const theme = useTheme();
  const [focused, setFocused] = useState(false);

  const hasError = Boolean(errorMessage);
  const hasSuccess = Boolean(successMessage) && !hasError;
  const borderColor = hasError
    ? theme.error
    : hasSuccess
      ? theme.success
      : focused
        ? theme.focus
        : theme.border;

  return (
    <View style={styles.wrapper}>
      <Text style={[styles.label, { color: theme.textSecondary }]}>{label}</Text>
      <TextInput
        accessibilityLabel={label}
        editable={editable}
        placeholderTextColor={theme.textSecondary}
        style={[
          styles.input,
          {
            color: theme.text,
            backgroundColor: theme.backgroundElement,
            borderColor,
            opacity: editable ? 1 : 0.55,
          },
          style,
        ]}
        onBlur={(event) => {
          setFocused(false);
          onBlur?.(event);
        }}
        onFocus={(event) => {
          setFocused(true);
          onFocus?.(event);
        }}
        {...rest}
      />
      {hasError ? (
        <Text accessibilityLiveRegion="polite" style={[styles.helper, { color: theme.error }]}>
          {errorMessage}
        </Text>
      ) : null}
      {hasSuccess ? (
        <Text accessibilityLiveRegion="polite" style={[styles.helper, { color: theme.success }]}>
          {successMessage}
        </Text>
      ) : null}
    </View>
  );
}

const styles = StyleSheet.create({
  wrapper: {
    gap: Spacing.one,
  },
  label: {
    fontSize: 14,
    fontWeight: '600',
  },
  input: {
    minHeight: 48,
    borderWidth: 1,
    borderRadius: DestinyRadii.input,
    paddingHorizontal: Spacing.three,
    fontSize: 16,
  },
  helper: {
    fontSize: 13,
    lineHeight: 18,
  },
});
