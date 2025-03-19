import { requestNewReservation } from '$lib/api/mock';
import type {
    RequestNewReservationResponse,
    ReservationSingleResponse,
    Venue
} from '$lib/interfaces/api';
import type { MinimalCalendarUIItemWithHref } from '$lib/interfaces/calendar';
import type { DateString, HourString } from '$lib/interfaces/date';
import type { FormSelectItem, SelectableItem } from '$lib/interfaces/ui';
import { getCalendarPlaceholder, getTwoWeekRange } from '$lib/utils/calendar';
import { fromDateString } from '$lib/utils/date';
import { zeroPad } from '$lib/utils/helper';
import { DateTime } from 'luxon';

export interface FormData {
    requester_info: {
        name: string;
        school_id: string;
        date_of_birth: DateString | '';
    };
    reservations: {
        venue: string;
        date: DateString | '';
        hours: HourString[];
        purpose: string;
        purpose_detail: string;
        companions?: string;
    };
    agreement: {
        agreement: boolean;
    };
}

export const init_form_data = (): FormData => {
    return {
        requester_info: {
            name: '',
            school_id: '',
            date_of_birth: ''
        },
        reservations: {
            venue: '',
            date: '',
            hours: [],
            purpose: '',
            purpose_detail: '',
            companions: ''
        },
        agreement: {
            agreement: false
        }
    };
};

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
            not_past: boolean;
        };
        hours: {
            // 연속으로 예약해야 합니다.
            should_sequence: boolean;
            // 6시간 초과 예약이 불가능합니다.
            less_than_6hours: boolean;
        };

        purpose: boolean;
        purpose_detail: boolean;

        // 해당 예약은 비어있는가?
        is_free: boolean;
    };
    agreement: {
        // 유의사항에 동의하였는가?
        agreement: boolean;
    };
}

export const isAllValidated = (validations: Validations): boolean => {
    console.log(validations)
    return is_valid(validations);
};

const is_valid = (obj: object): boolean => Object.values(obj).every(i => {
    if (typeof i === 'object') return is_valid(i);
    return i;
})

/// props from form component.
export interface FormProps {
    /// Default calendar for usage.
    calendar: MinimalCalendarUIItemWithHref[];
    purposes: string[];
    getReservations: (date?: DateString, venue?: string) => Promise<ReservationSingleResponse[]>;
    // submitForm: (data: FormData) => Promise<FormSubmissionResult>;
}

export interface InternalStates {
    reservations: {
        selectable_venue: SelectableItem<Venue>[];
        selectable_venue_selected: SelectableItem<Venue>[];

        selectable_hour: SelectableItem<HourString>[];
        selectable_hour_selected: SelectableItem<HourString>[];

        // The reservations data from venue for entire month.
        // FIXME: maybe a single day?
        current_reservations_data: ReservationSingleResponse[];
        rendered_calendar: MinimalCalendarUIItemWithHref[];
        calendar_selected: MinimalCalendarUIItemWithHref[];

        selected_category?: FormSelectItem<string>;
    };
    collapsed: {
        requester_info: boolean;
        reservations: boolean;
        agreement: boolean;
    };
}

export const init_internal_states = (): InternalStates => {
    return {
        reservations: {
            selectable_venue: [],
            selectable_venue_selected: [],
            selectable_hour: generateSelectableHours(),
            selectable_hour_selected: [],
            current_reservations_data: [],
            rendered_calendar: getCalendarPlaceholder(),
            calendar_selected: [],

            selected_category: undefined
        },
        collapsed: {
            requester_info: true,
            reservations: true,
            agreement: true
        }
    };
};

export const feedVenueData = (venues: Venue[], internal_states: InternalStates) => {
    internal_states.reservations.selectable_venue = venues.map((venue) => {
        return {
            value: venue,
            key: venue.venue,
            label: venue.venue,
            disabled: false
        };
    });
};

/// Validates the form data, mustve not use heavy tasks.
export const validate = (form_data: FormData, internal_states: InternalStates): Validations => {
    const requester_info = (() => {
        const data = form_data.requester_info;

        const name = data.name.length > 0 && data.name.length < 30;
        const school_id = !isNaN(parseInt(data.school_id));
        const date_of_birth = data.date_of_birth.length > 0;

        return { name, school_id, date_of_birth };
    })();

    const reservations = (() => {
        const venue = form_data.reservations.venue.length > 0;
        const date = (() => {
            const date = form_data.reservations.date;
            if (!date) return { not_deadline: false, not_past: false };
            const parsed = fromDateString(date);
            const not_past = parsed.diffNow().milliseconds > 0;
            // deadline: 전날 18시 전
            const not_deadline =
                parsed.minus({ day: 1 }).set({ hour: 18 }).diffNow().milliseconds > 0;
            return { not_deadline, not_past };
        })();

        const hours = (() => {
            const hours = form_data.reservations.hours;
            if (hours.length === 0) return { should_sequence: false, less_than_6hours: false };
            const should_sequence = hours
                .map((i) => parseInt(i))
                .sort((a, b) => a - b)
                .every((value, index, array) => index == 0 || array[index - 1] + 1 == value);
            const less_than_6hours = hours.length <= 6;
            return { should_sequence, less_than_6hours };
        })();

        const purpose = form_data.reservations.purpose.length > 0;
        const purpose_detail = form_data.reservations.purpose_detail.length > 0;

        const is_free = (() => {
            const hour_check = internal_states.reservations.selectable_hour
                .filter((i) => i.disabled)
                .every((i) => !form_data.reservations.hours.includes(i.value));

            return hour_check;
        })();

        return { venue, date, hours, purpose, purpose_detail, is_free };
    })();

    const agreement = (() => {
        //const agreement = form_data.agreement.agreement;
        const agreement = true;
        return { agreement };
    })();

    return { requester_info, reservations, agreement };
};

export const generateSelectableHours = (): SelectableItem<HourString>[] =>
    Array(24)
        .keys()
        .map((i) => zeroPad(i.toString()))
        .map((i) => {
            return {
                value: i as HourString,
                key: i,
                label: `${i}시`,
                toggle: false
            };
        })
        .toArray();

export const requestNewReservationFromData = async (
    data: FormData
): Promise<RequestNewReservationResponse> =>
    requestNewReservation({
        name: data.requester_info.name,
        //email: data.requester_info.email,
        studentID: data.requester_info.school_id,
        birthday: data.requester_info.date_of_birth as string,

        date: data.reservations.date as DateString,
        venue: data.reservations.venue,
        time: data.reservations.hours,

        // FIXME
        category: data.reservations.purpose,
        purpose: data.reservations.purpose_detail,
        companions: data.reservations.companions
    });
