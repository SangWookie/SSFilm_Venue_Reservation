import type {
    ReservationList,
    RequestNewReservationData,
    RequestNewReservationResponse,
    AppState
} from '$lib/interfaces/api';
import type { DateString } from '$lib/interfaces/date';
import * as MockConst from '$lib/mock.const.ts';

const api_route = 'https://oxu8i5hf5h.execute-api.ap-northeast-2.amazonaws.com/dev';
export async function getReservationByDate(date: DateString): Promise<ReservationList> {
    return fetch(`${api_route}/reservations?date=${date}`)
        .then((res) => res.json());
}

export async function postNewReservation(
    body: RequestNewReservationData
): Promise<RequestNewReservationResponse> {
    console.log(body)
    return fetch(`${api_route}/reservations`, {
        method: 'POST',
        body: JSON.stringify(body),
        headers: {
            'content-type': 'application/json'
        }
    }).then((res) => res.json());
}

/// Returns by 200 or 404.
export async function getReservationStatus(reservationId: string): Promise<boolean> {
    return fetch(`${api_route}/reservations/check?reservationId=${reservationId}`)
        .then((res) => res.status === 200);
}

export async function getAppState(): Promise<AppState> {
    console.log('mock: getAppState');
    await new Promise((resolve) => setTimeout(resolve, 500));
    return {
        venues: MockConst.venues,
        purposes: MockConst.purposes
    };
}
