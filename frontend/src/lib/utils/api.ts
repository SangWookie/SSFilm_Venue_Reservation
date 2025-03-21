import type { ReservationList, Venue } from '$lib/interfaces/api';
import type { HourString } from '$lib/interfaces/date';

export const getUnavilableHours = (reservation: ReservationList, venue: Venue): HourString[] => {
    const data = reservation.venues.find((v) => v.venue === venue.venue);
    if (!data) return [];
    return [
        ...data.reservations.flatMap((r) => r.time),
        ...data.unavailable_periods.flatMap((r) => r.time)
    ];
};
