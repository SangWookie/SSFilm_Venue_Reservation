import { DateTime, Interval } from 'luxon';

/// Get days between start_time and end_time
export const fillDaysBetween = (start_time: DateTime, end_time: DateTime): DateTime[] =>
    Interval.fromDateTimes(start_time, end_time)
        .splitBy({ days: 1 })
        .map((interval) => interval.start)
        .filter((i) => i != undefined);

/// Get the week number of the year from date.
export const getWeekNumber = (date: DateTime): number => date.localWeekNumber;

/// Get days of the week from year and week number.
export const getWeek = (year: number, weeks: number): DateTime[] =>
    Interval.after(
        DateTime.local(year, 1, 1, { zone: 'Asia/Seoul' })
            .plus({ weeks: weeks - 1 })
            .set({ weekday: 7 }), // 7 = Sunday
        { days: 7 }
    )
        .splitBy({ days: 1 })
        .map((interval) => interval.start)
        .filter((i) => i != undefined);

export const getWeeksYearMinMaxRange = (year: number): number[] => [
    DateTime.local(year, 1, 1, { zone: 'Asia/Seoul' }).localWeekNumber,
    DateTime.local(year, 12, 31).localWeekNumber
];
