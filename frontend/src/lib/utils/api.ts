import Reservations from '$lib/components/form/new-reservation/sections/reservations.svelte';
import type { ReservationItem, ReservationItemCompact, Venue } from '$lib/interfaces/api';
import type { MinimalCalendarUIItem } from '$lib/interfaces/calendar';
import type { HourString } from '$lib/interfaces/date';
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

export const getUnavilableHours = (reservation: ReservationItem, venue: Venue): HourString[] => {
    const data = reservation.venues.find((v) => v.venue === venue.venue);
    if (!data) return [];
    return [
        ...data.reservations.flatMap((r) => r.time),
        ...data.unavailable_periods.flatMap((r) => r.time)
    ];
};
