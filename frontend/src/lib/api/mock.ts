import type { DateString, HourString } from '$lib/interfaces/date';
import type {
    AppState,
    RequestNewReservationResponse,
    ReservationRequest,
    ReservationSingleResponse,
    Venue
} from '$lib/interfaces/api';

import Data from '../../mock_data.json';

export const getReservations = async (
    date?: DateString,
    venue?: string
): Promise<ReservationSingleResponse[]> => {
    // 500ms delay
    await new Promise((resolve) => setTimeout(resolve, 500));
    const reservations = Object.groupBy(
        Data.reservations.filter((i) => {
            if (date && i.date !== date) return false;
            if (venue && i.venue !== venue) return false;
            return true;
        }),
        ({ date, venue }) => `${date}-${venue}`
    );
    return (
        Object.entries(reservations)
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
            .filter((i) => i !== undefined)
    );
};

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export const requestNewReservation = async (
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    reservation: ReservationRequest
): Promise<RequestNewReservationResponse> => {
    console.log('requestNewReservation', reservation);
    // 3000ms delay
    await new Promise((resolve) => setTimeout(resolve, 3000));
    return {
        success: true,
        message: '예약 요청이 완료되었습니다.'
    };
};

export const getAppState = async (): Promise<AppState> => {
    // 500ms delay
    await new Promise((resolve) => setTimeout(resolve, 500));
    return {
        venues: Data.venue as Venue[],
        api_url: '/api',
        purposes: ['목적 A', '목적 B', '목적 C', '기타']
    }
};
