import type { RequestNewReservationResponse } from '$lib/interfaces/api';
import type { DateString, HourString } from '$lib/interfaces/date';
import { fromDateString } from '$lib/utils/date';
import { zeroPad } from '$lib/utils/helper';
import { type SelectableItem } from '$lib/components/ui/form/selectable-list.svelte.ts';
import { postNewReservation } from '$lib/api/mock';

export interface FormData {
    requester_info: {
        name: string;
        school_id: string;
        email: DateString | '';
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
            email: ''
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
        email: boolean;
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
    };
    agreement: {
        // 유의사항에 동의하였는가?
        agreement: boolean;
    };
}

export const isAllValidated = (validations: Validations): boolean => {
    return is_valid(validations);
};

const is_valid = (obj: object): boolean =>
    Object.values(obj).every((i) => {
        if (typeof i === 'object') return is_valid(i);
        return i;
    });

/// Validates the form data, mustve not use heavy tasks.
export const validate = (form_data: FormData): Validations => {
    const requester_info = (() => {
        const data = form_data.requester_info;

        const name = data.name.length > 0 && data.name.length < 30;
        const school_id = !isNaN(parseInt(data.school_id));
        // validate email
        let email = data.email.length > 0 && data.email.includes('@');

        if (email) {
            const split = data.email.split('@');
            email = split.length === 2 && split[1].includes('.');
        }

        return { name, school_id, email };
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

        return { venue, date, hours, purpose, purpose_detail };
    })();

    const agreement = (() => {
        //const agreement = form_data.agreement.agreement;
        const agreement = true;
        return { agreement };
    })();

    return { requester_info, reservations, agreement };
};

export const generateSelectableHours = (): SelectableItem<HourString>[] =>
    Array.from(Array(24).keys())
        .map((i) => zeroPad(i.toString()))
        .map((i) => {
            return {
                value: i as HourString,
                key: i,
                label: `${i}시`,
                toggle: false
            };
        })

export const requestNewReservationFromData = async (
    data: FormData
): Promise<RequestNewReservationResponse> =>
    postNewReservation({
        name: data.requester_info.name,
        studentID: data.requester_info.school_id,
        email: data.requester_info.email,

        date: data.reservations.date as DateString,
        venue: data.reservations.venue,
        time: data.reservations.hours,

        category: data.reservations.purpose,
        purpose: data.reservations.purpose_detail,
        companions: data.reservations.companions
    });
