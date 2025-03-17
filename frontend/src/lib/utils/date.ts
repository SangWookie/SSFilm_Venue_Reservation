import type { DateString, YearString, MonthString, DayString } from '$lib/interfaces/date';
import { DateTime } from 'luxon';

/// Extracts the year, month, and day from a date string.
export const extractDateString = (date: DateString): [YearString, MonthString, DayString] => [
    date.slice(0, 4) as YearString,
    date.slice(5, 7) as MonthString,
    date.slice(8, 10) as DayString
];

export const intoDateString = (date: DateTime): DateString =>
    date.toFormat('yyyy-LL-dd') as DateString;

export const fromDateString = (date: DateString): DateTime => {
    const extract = extractDateString(date);
    return DateTime.local().set({
        year: parseInt(extract[0]),
        month: parseInt(extract[1]),
        day: parseInt(extract[2])
    });
};
