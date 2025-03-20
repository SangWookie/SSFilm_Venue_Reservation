import type { DateString, HourString } from './date';

export interface ReservationItem {
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
    }[]
}

export interface ReservationItemCompact {
    date: DateString;
    reservations: {
        time: HourString[];
        //name: string;
        purpose?: string;
    }[];
    unavailable_periods: {
        time: HourString[];
        message?: string;
    }[];
}

/// 한 방의 예약 정보를 나타내는 인터페이스.
/// FIXME: 여러개 필요.
export interface ReservationSingleResponse extends ReservationItemCompact {
    venue: string;
    approval_mode: 'auto' | 'manual';
}

export interface ReservationRequest {
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

export type ReservationRequestDraft = Partial<ReservationRequest>;

export interface Venue {
    venue: string;
    venueKor: string;
    requirement?: string;
    approval_mode: 'auto' | 'manual';
}

export interface RequestNewReservationResponse {
    reservationId: string;
}

export interface AppState {
    venues: Venue[];
    purposes: string[];
}

export type LazyAppState = Partial<AppState>;
