import type { ReservationSingleResponse } from '$lib/interfaces/api';
import type { MinimalCalendarUIItem } from '$lib/interfaces/calendar';
import { DateTime, Interval, Settings } from 'luxon';
import { intoDateString } from './date';

Settings.defaultZone = 'Asia/Seoul';
Settings.defaultWeekSettings = {
    firstDay: 7,
    minimalDays: 4,
    weekend: [6, 7]
};
/**
 * DateTime Array을 2차원 구조인 달력으로 만드는 함수
 * @param dates 연속된 DateTime. Sort 되어야만 함.
 * @returns 달력 Array. 해당 주에 없는 데이트가 존재할 경우 undefined로 저장됨.
 */
export const generateCalendar = <T extends DateTime>(dates: T[]): (T | undefined)[][] => {
    console.log(
        'generateCalendar',
        dates.map((date) => date.toISO())
    );
    const temp_total: (undefined | T)[][] = []; // total week.
    let temp_week: (undefined | T)[] = []; // single week
    let previous_date: undefined | T = undefined;
    dates.forEach((date) => {
        /// Pushs the current week into final data and then clear it.
        const pushCurrentWeek = () => {
            temp_week = temp_week.concat(
                Array(7 - temp_week.length)
                    .keys()
                    .map(() => undefined)
                    .toArray()
            );
            temp_total.push(temp_week);
            temp_week = [];
            previous_date = undefined;
        };

        // Clear if temp_week is full.
        if (temp_week.length === 7) pushCurrentWeek();

        if (previous_date) {
            // If delta week of previous date versus current date is not 0,
            // must generate new week instead.
            // NOTICE: this requires `Settings.defaultWeekSettings` to correctly set.
            if (!date.hasSame(previous_date, 'week', { useLocaleWeeks: true })) {
                // pushes previous week.
                pushCurrentWeek();
                console.log('pushCurrentWeek', date.weekdayShort);
                // the generation work will be continued on below.
            }
        }

        // If the date is not Monday and weeks are blank,
        // generate the previous days.
        // Fri(5) -> Generate five(sun.mon,tue,wed,thur)
        // Sun(7) -> Not in this case due to `date.weekday !== 7`
        if (date.weekday !== 7 && temp_week.length === 0) {
            temp_week = temp_week.concat(
                Array(date.weekday)
                    .keys()
                    .map(() => undefined)
                    .toArray()
            );
        }

        previous_date = date;
        temp_week.push(date);
    });
    if (temp_week.length > 0) temp_total.push(temp_week);
    return temp_total;
};

export const generateCalendarFromProps = <T extends DateTime, WrappedT extends { date: T }>(
    data: WrappedT[]
): (WrappedT | undefined)[][] => {
    const generated = generateCalendar(data.map((i) => i.date));

    return generated.map((week) =>
        week.map((day) => {
            if (day) return data.find((d) => d.date == day);
            return undefined;
        })
    );
};

export const getTwoWeekRange = () =>
    Interval.fromDateTimes(DateTime.local().minus({ days: 7 }), DateTime.local().plus({ days: 14 }))
        .splitBy({ day: 1 })
        .map((i) => i.start)
        .filter((i) => i != null);

export const mergeReservationsIntoCalendar = (
    reservations: ReservationSingleResponse[],
    calendar: MinimalCalendarUIItem[]
): MinimalCalendarUIItem[] => {
    console.log('mergeReservationsIntoCalendar', reservations, calendar);
    const reservedDays = reservations.filter((r) => r.reservations.length > 0).map((r) => r.date);
    const unavilableDays = reservations
        .filter((r) => r.unavailable_periods.length > 0)
        .map((r) => r.date);
    calendar.forEach((item) => {
        const dateString = intoDateString(item.date);

        if (item.mark == null) item.mark = {};
        
        item.mark.reserved = reservedDays.includes(dateString);
        item.mark.unavailable = unavilableDays.includes(dateString);
        if (item.mark.unavailable) console.log(item.mark.unavailable);
        
        
    });
    console.log('reservedDays', reservedDays, 'unavilableDays', unavilableDays, 'calendar', calendar);

    return calendar;
};
