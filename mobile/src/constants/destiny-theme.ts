/**
 * Hallmark · component: theme · genre: atmospheric · theme: Destiny Midnight
 * states: n/a (token palette)
 * contrast: pass
 *
 * Celestial palette for the Destiny astrology app — deep cosmic canvas,
 * warm gold accent, readable body type on dark surfaces.
 */

export const DestinyPalette = {
  light: {
    text: '#12141F',
    background: '#F4F0E8',
    backgroundElement: '#E8E4DC',
    backgroundSelected: '#D8D4CC',
    textSecondary: '#5C6370',
    accent: '#8B6914',
    accentForeground: '#FFFFFF',
    border: '#D0CCC4',
    error: '#B42318',
    errorSurface: '#FEE4E2',
    success: '#027A48',
    successSurface: '#D1FADF',
    focus: '#6B5A1E',
    glow: 'rgba(139, 105, 20, 0.12)',
    bloomPrimary: 'rgba(139, 105, 20, 0.18)',
    bloomSecondary: 'rgba(91, 74, 138, 0.12)',
  },
  dark: {
    text: '#F4F0E8',
    background: '#0B0D17',
    backgroundElement: '#151929',
    backgroundSelected: '#1E2438',
    textSecondary: '#9BA3B4',
    accent: '#D4AF37',
    accentForeground: '#0B0D17',
    border: '#2A3147',
    error: '#F97066',
    errorSurface: '#3B1C1C',
    success: '#6CE9A6',
    successSurface: '#0F2A1E',
    focus: '#E8C547',
    glow: 'rgba(212, 175, 55, 0.18)',
    bloomPrimary: 'rgba(212, 175, 55, 0.22)',
    bloomSecondary: 'rgba(91, 74, 138, 0.28)',
  },
} as const;

export type DestinyTheme = (typeof DestinyPalette)['light'];

export const DestinyRadii = {
  input: 12,
  button: 999,
  card: 20,
} as const;

export const DestinyMotion = {
  micro: 120,
  standard: 200,
} as const;
