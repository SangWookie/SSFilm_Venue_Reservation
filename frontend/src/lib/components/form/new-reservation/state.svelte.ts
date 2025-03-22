import {
    SelectableListState,
    type SelectableItem
} from '$lib/components/ui/form/selectable-list.svelte.ts';
import type { ReservationList, Venue } from '$lib/interfaces/api';
import type { MinimalCalendarUIItemWithHref } from '$lib/interfaces/calendar';
import type { HourString } from '$lib/interfaces/date';
import { globalAppState } from '$lib/store.svelte';
import { getCalendarPlaceholderCustom, getNextTwoWeeks } from '$lib/utils/calendar';
import { intoDateString } from '$lib/utils/date';
import type { FormData } from '.';
import { untrack } from 'svelte';
import { generateSelectableHours } from '.';
import type { FormSelectItem } from '$lib/interfaces/ui';
import { getReservationByDate } from '$lib/api/api';
import { zeroPad } from '$lib/utils/helper';

export class ReservationSectionFormState {
    reservation: ReservationList | undefined = $state(undefined);
    loading_reservation: boolean = $state(false);
    venues: Venue[] = $state([]);
    purposes: string[] = $state([]);

    venue_selectable = new SelectableListState<Venue>();
    hour_selectable = new SelectableListState<HourString>();

    calendar = $state(getCalendarPlaceholderCustom(getNextTwoWeeks()));
    calendar_selected: MinimalCalendarUIItemWithHref[] = $state([]);

    category_selected: FormSelectItem<string> | undefined = $state(undefined);

    current_venue = $derived(this.venue_selectable.selected.at(0));
    current_date = $derived(this.calendar_selected.at(0));
    current_reservation = $derived.by(() => {
        if (!this.current_venue || !this.current_date) return;
        return this.reservation?.venues.find((i) => i.venue == this.current_venue?.value.venue);
    });

    unavailableHours = $derived(
        [
            ...(this.current_reservation?.reservations?.flatMap((r) => r.time) || []),
            ...(this.current_reservation?.unavailable_periods?.flatMap((r) => r.time) || [])
        ].map((r) => zeroPad(r.toString()) as HourString)
    );

    #form_data: FormData;
    constructor(form_data: FormData) {
        this.#form_data = form_data;
        globalAppState.subscribe((state) => {
            if (state?.venues && this.venues.length == 0) this.venues = state.venues;
            if (state?.purposes) this.purposes = state.purposes;
        });
        this.hour_selectable.list = generateSelectableHours();

        this.effectFeedVenueList();
        this.effectWriteUnableHoursToSelectable();
        this.effectWriteFormData();
        this.effectResetHourSelectedWhenVenueChanged();
    }

    effectFeedVenueList() {
        $effect(() => {
            this.venue_selectable.list = this.venues.map((venue) => {
                return {
                    value: venue,
                    key: venue.venue,
                    label: venue.venueKor,
                    disabled: false
                };
            });
        });
    }

    effectWriteUnableHoursToSelectable() {
        $effect(() => {
            console.log(this.unavailableHours);
            untrack(() => {
                this.hour_selectable.list.forEach((i) => {
                    const disabled = this.unavailableHours.includes(i.value);
                    if (disabled) this.hour_selectable.unselect(i);
                    i.disabled = disabled;
                });
            });
            this.unavailableHours;
        });
    }

    effectWriteFormData() {
        $effect(() => {
            this.#form_data.reservations.venue = this.current_venue?.value.venue || '';
            this.#form_data.reservations.date = this.current_date
                ? intoDateString(this.current_date.date)
                : '';
            this.#form_data.reservations.hours = this.hour_selectable.selected.map((i) => i.value);

            if (this.category_selected?.value) {
                this.#form_data.reservations.purpose = this.category_selected.value;
            }
        });
    }

    effectResetHourSelectedWhenVenueChanged() {
        $effect(() => {
            // Let's just clear the selected date just for now.
            this.hour_selectable.selected = [];

            // tried current_reservation, but the data might be same
            // thus just use current_venue for now.
            this.current_date;
            this.current_venue;
        });
    }

    fetchReservation() {
        if (this.loading_reservation || !this.current_date) return;

        this.loading_reservation = true;
        getReservationByDate(intoDateString(this.current_date.date)).then((reservation) => {
            // Clear hour selectable
            this.hour_selectable.selected = [];
            this.#form_data.reservations.hours = [];

            this.reservation = reservation;
            this.loading_reservation = false;

            reservation.venues.forEach((v) => {
                const venue = this.venues.find((i) => i.venue == v.venue);
                if (venue) venue.approval_mode = v.approval_mode;
            });
        });
    }

    hourSelectableClickCallback = (item: SelectableItem<HourString>) => {
        // Prevent click event if date is not selected.
        if (this.#form_data.reservations.date.length === 0) return;
        // Prevent click event if item is disabled.
        if (item.disabled) return;

        if (this.loading_reservation) return;

        const is_selected = this.hour_selectable.isSelected(item);
        const selected_len = this.hour_selectable.selected.length;

        // 만약 선택된 아이템이 1개라면, 그 아이템을 따라오는 아이템들을 선택한다.
        if (!is_selected && selected_len == 1) {
            const this_index = this.hour_selectable.list.findIndex((i) => i == item);
            const target_index = this.hour_selectable.list.findIndex(
                (i) => i == this.hour_selectable.selected.at(0)!
            );

            if (this_index == -1 || target_index == -1) {
                throw `hourSelectableClickCallback failed; ${this_index} vs ${target_index}`;
            }

            const selected: SelectableItem<HourString>[] = [];
            const plusminus = this_index < target_index ? -1 : 1;
            for (let i = target_index; i != this_index + plusminus; i += plusminus) {
                const item = this.hour_selectable.list[i];
                // 중간에 disabled 되어 있다면 중단한다.
                // 6시간 초과여도 중단한다.
                if (item.disabled || selected.length >= 6) {
                    break;
                }
                selected.push(item);
            }

            if (selected.length <= 1) {
                this.hour_selectable.selected = [item];
                return;
            }

            this.hour_selectable.selected = this.hour_selectable.list.filter((i) =>
                selected.includes(i)
            );
            return;
        }

        this.hour_selectable.selected = [item];
    };
}
