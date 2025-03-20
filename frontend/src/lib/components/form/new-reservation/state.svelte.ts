import { createSelectableList } from "$lib/components/ui/form/selectable-list.svelte.ts";
import type { ReservationItem, Venue } from "$lib/interfaces/api";
import type { MinimalCalendarUIItemWithHref } from "$lib/interfaces/calendar";
import type { HourString } from "$lib/interfaces/date";
import { globalAppState } from "$lib/store.svelte";
import { getCalendarPlaceholder } from "$lib/utils/calendar";
import { intoDateString } from "$lib/utils/date";
import type { FormData } from ".";
import { untrack } from "svelte";
import { type Validations, type FormProps, type InternalStates, generateSelectableHours } from ".";
import type { FormSelectItem } from "$lib/interfaces/ui";

type Wrapped<T> = { value: T}
const wrapped = <T>(value: T): Wrapped<T> => {
    return {
        value
    }
}
export const createReservationSectionForm = (
    form_data: FormData,
    validations: Validations,
    form_props: FormProps,
    internal_states: InternalStates,
) => {
    let reservation: Wrapped<ReservationItem | undefined> = $state(wrapped(undefined));
    let loading_reservation = $state(wrapped(false));
    let venues: Wrapped<Venue[]> = $state(wrapped([]));
    let purposes: Wrapped<string[]> = $state(wrapped([]));
    const venue_selectable = createSelectableList<Venue>();
    const hour_selectable = createSelectableList<HourString>();
    hour_selectable.list = generateSelectableHours();
    let calendar = $state(getCalendarPlaceholder());
    let calendar_selected: Wrapped<MinimalCalendarUIItemWithHref[]> = $state(wrapped([]));
    let selected_category: Wrapped<FormSelectItem<string> | undefined> = $state(wrapped(undefined));
    // Feeding venues and purposes.
    globalAppState.subscribe((state) => {
        if (state?.venues) venues.value = state.venues;
        if (state?.purposes) purposes.value = state.purposes;
    });
    $effect(() => {
        venue_selectable.list = venues.value.map((venue) => {
            return {
                value: venue,
                key: venue.venue,
                label: venue.venue,
                disabled: false
            };
        });
    })

    const selected_venue = $derived(venue_selectable.selected.at(0));
    let selected_date = $derived(calendar_selected.value.at(0));
    const current_reservation = $derived.by(() => {
        if (!selected_venue || !selected_date) return;
        return reservation.value?.venues
            .find(i => i.venue == selected_venue.value.venue)
    })

    const hour_selectable_disabled_state = $derived((selected_venue && selected_date))
    const unavailableHours = $derived.by(() => {
        return [
            ...current_reservation?.reservations?.flatMap(r => r.time) || [],
            ...current_reservation?.unavailable_periods?.flatMap(r => r.time) || []
        ];
    })

    // Form
    $effect(() => {
        form_data.reservations.venue = selected_venue?.value.venue || '';
        form_data.reservations.date = selected_date ? intoDateString(selected_date.date) : '';
        form_data.reservations.hours = hour_selectable.selected.map(i => i.value);

    })
    $effect(() => {
        hour_selectable.list = untrack(() => hour_selectable.list.map(i => {
            return {
                disabled: unavailableHours.includes(i.value),
                 ...i
            }
        }))
        unavailableHours;
    })
    $effect(() => {
        if (selected_category.value?.value)
            form_data.reservations.purpose = selected_category.value?.value;
    })

    // FIXME: fetch reservations


    
    return {
        purposes, loading_reservation,
        venue_selectable,
        hour_selectable,
        selected_category,
        calendar,
        get calendar_selected() { return calendar_selected },
        set calendar_selected(val) { calendar_selected = val},
        selected_venue, selected_date, current_reservation,
        hour_selectable_disabled_state
    }
}