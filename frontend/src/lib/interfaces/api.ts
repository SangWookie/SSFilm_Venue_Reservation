import type { DateString } from './date';

/// The return data for `GET /reservation`. only intended for filtering dates.
export interface ReservationList {
    date: DateString;
    venues: ReservationByVenue[];
}

export interface ReservationByVenue {
    venue: string;
    venueKor: string;
    approval_mode: 'auto' | 'manual';
    reservations: {
        time: number[];
        name: string;
        purpose: string;
    }[];
    unavailable_periods: {
        time: number[];
        message?: string;
    }[];
}

export interface RequestNewReservationData {
    date: DateString;
    /// 장소
    venue: string;
    time: number[];

    name: string;
    email: string;
    /// 학번
    studentId: string;

    /// 목적 1차 카테고리
    category: string;
    /// 목적 2차 카테고리
    purpose?: string;
    /// 동료
    companion?: string;
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
