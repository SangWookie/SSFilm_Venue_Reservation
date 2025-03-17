import type { DateString, HourString } from '$lib/interfaces/date';
import type { ReservationSingleResponse, Venue } from '$lib/interfaces/api';

import Data from '../../mock_data.json';

export const getReservations = async (
    date?: DateString,
    venue?: string
): Promise<ReservationSingleResponse[]> => {
    const reservations = Object.groupBy(
        Data.reservations.filter((i) => {
            if (date && i.date !== date) return false;
            if (venue && i.venue !== venue) return false;
            return true;
        }),
        ({ date, venue }) => `${date}-${venue}`
    );
    return Object.entries(reservations)
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        .map(([_key, items]) => {
            if (items === undefined || items.length === 0) return;
            return {
                date: items.at(0)!.date as unknown as DateString,
                reservations: items.map((i) => {
                    return {
                        time: i.time as unknown as HourString[]
                    };
                }),
                unavailable_periods: [],
                venue: items.at(0)!.venue,
                approval_mode: 'auto'
            } as ReservationSingleResponse;
        })
        .filter((i) => i !== undefined);
};

export const getVenueList = async () => Data.venue as Venue[];