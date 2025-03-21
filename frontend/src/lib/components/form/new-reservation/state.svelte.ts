import { SelectableListState } from '$lib/components/ui/form/selectable-list.svelte.ts';
import type { ReservationList, Venue } from '$lib/interfaces/api';
import type { MinimalCalendarUIItemWithHref } from '$lib/interfaces/calendar';
import type { HourString } from '$lib/interfaces/date';
import { globalAppState } from '$lib/store.svelte';
import { getCalendarPlaceholder } from '$lib/utils/calendar';
import { intoDateString } from '$lib/utils/date';
import type { FormData } from '.';
import { untrack } from 'svelte';
import { generateSelectableHours } from '.';
import type { FormSelectItem } from '$lib/interfaces/ui';

export class ReservationSectionFormState {
    reservation: ReservationList | undefined = $state(undefined);
    loading_reservation: boolean = $state(false);
    venues: Venue[] = $state([]);
    purposes: string[] = $state([]);

    venue_selectable = new SelectableListState<Venue>();
    hour_selectable = new SelectableListState<HourString>();

    calendar = $state(getCalendarPlaceholder());
    calendar_selected: MinimalCalendarUIItemWithHref[] = $state([]);

    category_selected: FormSelectItem<string> | undefined = $state(undefined);

    current_venue = $derived(this.venue_selectable.selected.at(0));
    current_date = $derived(this.calendar_selected.at(0));
    current_reservation = $derived.by(() => {
        if (!this.current_venue || !this.current_date) return;
        return this.reservation?.venues.find((i) => i.venue == this.current_venue?.value.venue);
    });

    unavailableHours = $derived([
        ...(this.current_reservation?.reservations?.flatMap((r) => r.time) || []),
        ...(this.current_reservation?.unavailable_periods?.flatMap((r) => r.time) || [])
    ]);

    #form_data: FormData;
    constructor(form_data: FormData) {
        this.#form_data = form_data;
        globalAppState.subscribe((state) => {
            if (state?.venues) this.venues = state.venues;
            if (state?.purposes) this.purposes = state.purposes;
        });
        this.hour_selectable.list = generateSelectableHours();

        this.effectFeedVenueList();
        this.effectWriteUnableHoursToSelectable();
        this.effectWriteFormData();
    }

    effectFeedVenueList() {
        $effect(() => {
            this.venue_selectable.list = this.venues.map((venue) => {
                return {
                    value: venue,
                    key: venue.venue,
                    label: venue.venue,
                    disabled: false
                };
            });
        });
    }

    effectWriteUnableHoursToSelectable() {
        $effect(() => {
            this.hour_selectable.list = untrack(() =>
                this.hour_selectable.list.map((i) => {
                    return {
                        disabled: this.unavailableHours.includes(i.value),
                        ...i
                    };
                })
            );
            // eslint-disable-next-line @typescript-eslint/no-unused-expressions
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
}
