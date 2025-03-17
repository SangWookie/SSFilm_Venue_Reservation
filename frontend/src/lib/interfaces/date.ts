//
// Related utils are in ../utils/date.ts
//

type Digit = '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9';
export type YearString = `19${Digit}${Digit}` | `20${Digit}${Digit}`;
export type MonthString = `0${Digit}` | `1${'0' | '1' | '2'}`;
export type DayString = `0${Digit}` | `1${Digit}` | `2${Digit}` | `3${'0' | '1'}`;

/// A string representing a time in the format `00` to `24`.
export type HourString = `0${Digit}` | `1${Digit}` | `2${'0' | '1' | '2' | '3'}`;

/// A string representing a date in the format `YYYY-MM-DD`.
export type DateString = `${YearString}-${MonthString}-${DayString}`;
