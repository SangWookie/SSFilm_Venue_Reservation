import type { ReservationItemCompact } from '$lib/interfaces/api';
import type { MinimalCalendarUIItem } from '$lib/interfaces/calendar';
import { fromDateString } from './date';

export const convertReservationItemCompactToMinimalCalendarItem = (
    item: ReservationItemCompact
): MinimalCalendarUIItem => {
    // selected, today mark will be calculated on ui side.
    return {
        date: fromDateString(item.date),
        mark: {
            reserved: item.reservations.length > 0,
            unavailable: item.unavailable_periods.length > 0
        }
    };
};
