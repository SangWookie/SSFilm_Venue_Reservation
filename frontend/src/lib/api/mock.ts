import type {
    ReservationList,
    RequestNewReservationData,
    RequestNewReservationResponse,
    AppState
} from '$lib/interfaces/api';
import type { DateString } from '$lib/interfaces/date';
import * as MockConst from '$lib/mock.const.ts';

export async function getReservationByDate(date: DateString): Promise<ReservationList> {
    console.log('mock: getReservationByDate', date);
    await new Promise((resolve) => setTimeout(resolve, 500));
    return {
        date,
        venues: []
    };
}

export async function postNewReservation(
    body: RequestNewReservationData
): Promise<RequestNewReservationResponse> {
    console.log('mock: postNewReservation', body);
    await new Promise((resolve) => setTimeout(resolve, 2000));
    return {
        reservationId: 'reservationId'
    };
}

/// Returns by 200 or 404.
export async function getReservationStatus(reservationId: string): Promise<boolean> {
    console.log('mock: getReservationStatus', reservationId);
    return true;
}

export async function getAppState(): Promise<AppState> {
    console.log('mock: getAppState');
    await new Promise((resolve) => setTimeout(resolve, 500));
    return {
        venues: MockConst.venues,
        purposes: MockConst.purposes
    };
}
