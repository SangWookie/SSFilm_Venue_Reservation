<script lang="ts">
    import Calendar from '$lib/components/ui/calendar.svelte';
    import type { Venue } from '$lib/interfaces/api';
    import type { MinimalCalendarUIItemWithHref } from '$lib/interfaces/calendar';
    import type { HourString } from '$lib/interfaces/date';
    import { untrack } from 'svelte';
    import CollapsibleBlock from '../../../collapsible-block.svelte';
    import InputBox from '../../../input-box.svelte';
    import SelectableList from '../../../selectable-list.svelte';
    import ValidateMessage from '../../../validate-message.svelte';
    import {
        generateSelectableHours,
        getCalendarPlaceholder,
        type FormData,
        type FormProps,
        type InternalStates,
        type Validations
    } from '../index.ts';
    const {
        form_data,
        validations,
        form_props,
        internal_states
    }: {
        form_data: FormData;
        validations: Validations;
        form_props: FormProps;
        internal_states: InternalStates;
    } = $props();

    internal_states.reservations = {
        selectable_venue: form_props.venue_list.map((venue) => {
            return {
                key: venue.venue,
                label: venue.venue,
                toggle: false
            };
        }),
        selectable_hour: generateSelectableHours(),
        selectable_hour_disabled: false,
        current_reservations_data: [],
        rendered_calendar: getCalendarPlaceholder()
    };

    const current_venue: Venue | undefined = $derived(
        form_props.venue_list.find((venue) => venue.venue == form_data.reservations.venue)
    );

    const date_click_callback = (date: MinimalCalendarUIItemWithHref | undefined) => {
        // When called, do following:
        // 1. Validate the date is clickable.
        // 2. Set the form data
        // 3. Clears the hours selectable state, then enable it.
        // 4. Resets the calendar state.
    };
    const selectable_venue_callback = () => {
        const query = internal_states.reservations.selectable_venue
            .find(i => i.toggle)
        if (query) form_data.reservations.venue = query.key as string; 
        
        // Changing venue requires new calendar, so request it.
        refresh();
    }

    const selectable_hour_callback = () => {
        form_data.reservations.hours = internal_states.reservations.selectable_hour.map(
            (i) => i.key as HourString
        );
    };
    
    
    const clear_date_selection = () => {
        internal_states.reservations.rendered_calendar
            .filter(i => i.mark?.selected)
            .forEach(i => i.mark!.selected = false);
        form_data.reservations.date = '';
    };
    
    const clear_hour_selection = () => {
        internal_states.reservations.selectable_hour = generateSelectableHours();
        form_data.reservations.hours = [];
        
        // then initialize into disabled state.
        internal_states.reservations.selectable_hour_disabled = true;
    }
    
    const refresh = () => {
        calendar_props.status = 'loading';
        clear_date_selection();
        clear_hour_selection();
        internal_states.reservations.current_reservations_data = [];
        form_props.getReservations(undefined, current_venue?.venue)
            .then(i => {
                // maybe we should fill from reservations data...
                internal_states.reservations.rendered_calendar = getCalendarPlaceholder();
                internal_states.reservations.current_reservations_data = i;
                calendar_props.status = 'available';
            })
    }

    const calendar_props = {
        data: {
            items: internal_states.reservations.rendered_calendar
        },
        status: 'disabled' as 'available' | 'loading' | 'disabled',
        onDateClick: date_click_callback
    };
    
    $effect(() => {
        // eslint-disable-next-line @typescript-eslint/no-unused-expressions
        internal_states.reservations.selectable_venue;
        untrack(() => selectable_venue_callback());
    })

    $effect(() => {
        // eslint-disable-next-line @typescript-eslint/no-unused-expressions
        internal_states.reservations.selectable_hour;
        untrack(() => selectable_hour_callback());
    });
</script>

<CollapsibleBlock open={internal_states.collapsed.reservations}>
    {#snippet header()}
        예약 선택
    {/snippet}

    <InputBox title="장소 선택">
        {#snippet custom()}
            <SelectableList
                bind:list={internal_states.reservations.selectable_venue}
                isRadio={true}
            />

            {#if current_venue}
                <!-- eslint-disable-next-line svelte/no-at-html-tags -->
                {@html current_venue.requirement}
            {/if}
        {/snippet}
    </InputBox>

    <InputBox title="일자 선택">
        {#snippet custom()}
            <Calendar {...calendar_props} />
        {/snippet}
    </InputBox>

    <InputBox title="시간 선택">
        {#snippet custom()}
            <SelectableList
                bind:list={internal_states.reservations.selectable_hour}
                disabled={internal_states.reservations.selectable_hour_disabled}
            />
        {/snippet}
    </InputBox>
</CollapsibleBlock>
