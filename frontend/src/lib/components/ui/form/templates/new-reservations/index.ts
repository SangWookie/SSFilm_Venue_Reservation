import type { ReservationSingleResponse, Venue } from '$lib/interfaces/api';
import type {
    MinimalCalendarUIItem,
    MinimalCalendarUIItemWithHref
} from '$lib/interfaces/calendar';
import type { DateString, HourString } from '$lib/interfaces/date';
import type { SelectableItem } from '$lib/interfaces/ui';
import { getTwoWeekRange } from '$lib/utils/calendar';
import { zeroPad } from '$lib/utils/helper';
import { DateTime } from 'luxon';

export interface FormData {
    requester_info: {
        name: string;
        school_id: string;
        date_of_birth: DateString;
    };
    reservations: {
        venue: string;
        date: DateString | '';
        hours: HourString[];
    };
    agreement: {
        agreement: boolean;
    };
}

export interface Validations {
    requester_info: {
        name: boolean;
        school_id: boolean;
        date_of_birth: boolean;
    };
    reservations: {
        venue: boolean;
        date: {
            not_deadline: boolean;
            not_future: boolean;
        };
        hours: {
            // 연속으로 예약해야 합니다.
            should_sequence: boolean;
            // 6시간 초과 예약이 불가능합니다.
            less_then_6hours: boolean;
        };

        // 해당 예약은 비어있는가?
        is_free: boolean;
    };
    agreement: {
        // 유의사항에 동의하였는가?
        agreement: boolean;
    };
}

/// props from form component.
export interface FormProps {
    venue_list: Venue[];
    calendar: MinimalCalendarUIItemWithHref[];
    getReservations: (date?: DateString, venue?: string) => Promise<ReservationSingleResponse[]>;
    // submitForm: (data: FormData) => Promise<FormSubmissionResult>;

    /// update calendar into `internal.calendar`
    updateCalendar(calendar: MinimalCalendarUIItem): void;
}

export interface InternalStates {
    reservations: {
        selectable_venue: SelectableItem[];
        selectable_hour: SelectableItem[];
        selectable_hour_disabled: boolean;
        current_reservations_data: ReservationSingleResponse[];
        rendered_calendar: MinimalCalendarUIItemWithHref[];
    };
    collapsed: {
        requester_info: boolean;
        reservations: boolean;
        agreement: boolean;
    };
}

/// Validates the form data, mustve not use heavy tasks.
export const validate = (form_data: FormData): Validations => {
    const requester_info = (() => {
        const data = form_data.requester_info;

        const name = data.name.length > 0 || data.name.length < 30;
        const school_id = !isNaN(parseInt(data.school_id));
        const date_of_birth = data.date_of_birth.length > 0;

        return { name, school_id, date_of_birth };
    })();

    return { requester_info };
};

export const generateSelectableHours = (): SelectableItem[] =>
    Array(24)
        .keys()
        .map((i) => zeroPad(i.toString()))
        .map((i) => {
            return {
                key: i,
                label: i,
                toggle: false
            };
        })
        .toArray();

export const getCalendarPlaceholder = (): MinimalCalendarUIItemWithHref[] =>
    getTwoWeekRange().map((date) => {
        return {
            date,
            mark: {
                today: date.hasSame(DateTime.local(), 'day'),
                past: date.diff(DateTime.local(), 'hours').hours < -1
            }
        };
    });
