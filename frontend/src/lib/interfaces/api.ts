import type { DateString, HourString } from './date';

/// The return data for `GET /reservation`. only intended for filtering dates.
export interface ReservationList {
    date: DateString;
    venues: {
        venue: string;
        venueKor: string;
        allowPolicy: 'auto' | 'manual';
        reservations: {
            time: HourString[];
            name: string;
            purpose: string;
        }[];
        unavailable_periods: {
            time: HourString[];
            message?: string;
        }[];
    }[];
}

export interface RequestNewReservationData {
    date: DateString;
    /// 장소
    venue: string;
    time: HourString[];

    name: string;
    email: string;
    /// 학번
    studentID: string;

    /// 목적 1차 카테고리
    category: string;
    /// 목적 2차 카테고리
    purpose?: string;
    /// 동료
    companions?: string;
}

export type RequestNewReservationDataDraft = Partial<RequestNewReservationData>;

export interface Venue {
    venue: string;
    venueKor: string;
    requirement?: string;
    approval_mode: 'auto' | 'manual';
}

export interface RequestNewReservationResponse {
    // 요청은 되었지만 확정인지는 모름
    reservationId: string;
}

export interface AppState {
    venues: Venue[];
    purposes: string[];
}

export type LazyAppState = Partial<AppState>;
