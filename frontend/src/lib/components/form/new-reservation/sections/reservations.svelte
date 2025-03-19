<script lang="ts">
    import Calendar from '$lib/components/ui/calendar.svelte';
    import type { ReservationSingleResponse, Venue } from '$lib/interfaces/api';
    import type { HourString } from '$lib/interfaces/date';
    import { untrack } from 'svelte';
    import CollapsibleBlock from '$lib/components/ui/form/collapsible-block.svelte';
    import LoadingBox from '$lib/components/ui/loading_box.svelte';
    import InputBox from '$lib/components/ui/form/input-box.svelte';
    import SelectableList from '$lib/components/ui/form/selectable-list.svelte';
    import ValidateMessage from '$lib/components/ui/form/validate-message.svelte';
    import Select from '$lib/components/ui/form/select.svelte';
    import { intoDateString } from '$lib/utils/date.ts';
    import {
        type FormData,
        type FormProps,
        type InternalStates,
        type Validations
    } from '../index.ts';
    import { mergeReservationsIntoCalendar } from '$lib/utils/calendar.ts';
    import { type SelectableItem } from '$lib/interfaces/ui.ts';
    const {
        form_data = $bindable(),
        validations = $bindable(),
        form_props = $bindable(),
        internal_states = $bindable()
    }: {
        form_data: FormData;
        validations: Validations;
        form_props: FormProps;
        internal_states: InternalStates;
    } = $props();

    $effect(() => {
        form_data.reservations.venue =
            internal_states.reservations.selectable_venue_selected.at(0)?.value.venue || '';

        const selected_date = internal_states.reservations.calendar_selected.at(0)?.date;

        if (selected_date) form_data.reservations.date = intoDateString(selected_date);

        form_data.reservations.hours = internal_states.reservations.selectable_hour_selected.map(
            (i) => i.value
        );
    });

    const unavailableHours = $derived(
        internal_states.reservations.current_reservations_data
            .filter((r) => r.date == form_data.reservations.date)
            .flatMap((r2) =>
                r2.reservations
                    .flatMap((i) => i.time)
                    .concat(r2.unavailable_periods.flatMap((i) => i.time))
            )
    );

    // when unavailablehours, current_reservations_data, gets updated, update the selectable_hour state.
    $effect(() => {
        untrack(() => {
            internal_states.reservations.selectable_hour.forEach((i) => {
                i.disabled = unavailableHours.includes(i.value as HourString);
            });
        });

        // eslint-disable-next-line @typescript-eslint/no-unused-expressions
        unavailableHours;
    });

    $effect(() => {
        form_data.reservations.purpose =
            internal_states.reservations.selected_category?.value || '';
    });

    const cache: { [key: string]: ReservationSingleResponse[] | 'loading' } = {};
    // loading new reservation logic.
    $effect(() => {
        if (!form_data.reservations.venue) return;
        const venue = form_data.reservations.venue;
        if (cache[venue] === 'loading') return;
        if (cache[venue]) {
            internal_states.reservations.current_reservations_data = cache[venue];
            internal_states.reservations.selectable_hour_selected = [];
            return;
        }
        cache[venue] = 'loading';
        internal_states.reservations.current_reservations_data = [];
        console.log('loading reservations for ', venue);
        form_props
            .getReservations(undefined, venue)
            .then((reservations) => {
                console.log('loaded reservation', reservations);
                internal_states.reservations.current_reservations_data = reservations;
                cache[venue] = reservations;
            })
            .catch((e) => {
                delete cache[venue];
                console.error(e);
                throw e;
            });
    });

    $effect(() => {
        untrack(() => {
            internal_states.reservations.rendered_calendar = mergeReservationsIntoCalendar(
                internal_states.reservations.current_reservations_data,
                internal_states.reservations.rendered_calendar
            );
        });

        internal_states.reservations.current_reservations_data;
    });

    const calendar_status: 'available' | 'loading' | 'disabled' = $derived.by(() => {
        if (!form_data.reservations.venue || form_data.reservations.venue.length === 0)
            return 'disabled';
        if (
            internal_states.reservations.current_reservations_data.length === 0 &&
            cache[form_data.reservations.venue] === 'loading'
        )
            return 'loading';
        return 'available';
    });
</script>

<CollapsibleBlock open={internal_states.collapsed.reservations}>
    {#snippet header()}
        예약 선택
    {/snippet}

    <InputBox title="장소 선택" description={form_data.reservations.venue}>
        {#snippet custom()}
            <SelectableList
                bind:list={internal_states.reservations.selectable_venue}
                bind:selected={internal_states.reservations.selectable_venue_selected}
                isRadio={true}
            >
                {#snippet labelSnippet(item: SelectableItem<unknown>)}
                    {item.label}

                    {#if (item as SelectableItem<Venue>).value?.approval_mode == 'manual'}
                        (수동)
                    {/if}
                {/snippet}
            </SelectableList>

            {#if form_data.reservations.venue}
                <!-- eslint-disable-next-line svelte/no-at-html-tags -->
                {@html internal_states.reservations.selectable_venue_selected.at(0)?.value
                    ?.requirement}
            {/if}
        {/snippet}
        <ValidateMessage isValid={validations.reservations.venue} message="장소를 선택해 주세요." />
    </InputBox>

    <InputBox title="일자 선택" description={form_data.reservations.date}>
        {#snippet custom()}
            <LoadingBox enabled={calendar_status === 'loading'} />
            <div class="calendar-wrapper">
                <Calendar
                    items={internal_states.reservations.rendered_calendar}
                    selected={internal_states.reservations.calendar_selected}
                    status={calendar_status}
                    onDateClick={(date) => {
                        if (date && calendar_status == 'available')
                            internal_states.reservations.calendar_selected = [date];
                    }}
                />
            </div>
        {/snippet}
        <ValidateMessage
            isValid={validations.reservations.date.not_deadline}
            message="데드라인 (TODO 추가; 전날 12시까지)에 맞춰야 합니다."
        />
        <ValidateMessage
            isValid={validations.reservations.date.not_past}
            message="과거에 대해 예약할 수 없습니다."
        />
    </InputBox>

    <InputBox
        title="시간 선택"
        description={form_data.reservations.hours.map((i) => `${i}시`).join(', ')}
    >
        {#snippet custom()}
            <SelectableList
                bind:list={internal_states.reservations.selectable_hour}
                bind:selected={internal_states.reservations.selectable_hour_selected}
                disabled={form_data.reservations.date.length === 0}
            />
        {/snippet}
        <ValidateMessage
            isValid={validations.reservations.hours.should_sequence}
            message="예약 시간은 연속으로 선택해야 합니다."
        />
        <ValidateMessage
            isValid={validations.reservations.hours.less_than_6hours}
            message="6시간 초과 예약이 불가능합니다."
        />
        <ValidateMessage
            isValid={validations.reservations.is_free}
            message="비어있는 시간만 예약이 가능합니다."
        />
    </InputBox>

    <InputBox title="예약 목적" inputType="select" description={form_data.reservations.purpose}>
        {#snippet custom()}
            <Select
                options={form_props.purposes.map((i) => {
                    return {
                        value: i,
                        key: i,
                        label: i
                    };
                })}
                bind:value={internal_states.reservations.selected_category}
            />
            <ValidateMessage
                isValid={validations.reservations.purpose}
                message="목적을 선택해 주세요."
            />
        {/snippet}
    </InputBox>

    <InputBox
        title="예약 상세 목적"
        placeholder="예약 상세 목적을 작성해 주세요."
        bind:value={form_data.reservations.purpose_detail}
    >
        <ValidateMessage
            isValid={validations.reservations.purpose_detail}
            message="상세 목적을 작성해 주세요."
        />
    </InputBox>

    <InputBox
        title="동행인 (선택사항)"
        placeholder="동행인이 있다면 작성해주세요."
        bind:value={form_data.reservations.companions}
    />
</CollapsibleBlock>

<style lang="sass">
.calendar-wrapper
    display: flex
    justify-content: center
</style>
