import type { ReservationByVenue, ReservationList, Venue } from '$lib/interfaces/api';

export const getUnavilableHours = (reservation: ReservationList, venue: Venue): number[] => {
    const data = reservation.venues.find((v) => v.venue === venue.venue);
    if (!data) return [];
    return getUnavilableHoursByVenue(data);
};

export const getUnavilableHoursByVenue = (venue_data: ReservationByVenue): number[] => {
    if (!venue_data) return [];

    return [
        ...venue_data.reservations.flatMap((r) => r.time),
        ...venue_data.unavailable_periods.flatMap((r) => r.time)
    ].sort((a, b) => a - b);
};
