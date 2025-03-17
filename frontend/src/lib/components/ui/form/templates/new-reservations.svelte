<script lang="ts">
    import CollapsibleBlock from '../collapsible-block.svelte';
    import InputBox from '../input-box.svelte';
    import SelectableList from '../selectable-list.svelte';
    import ValidateMessage from '../validate-message.svelte';
    import Button from '../../button.svelte';
    import Calendar from '../../calendar.svelte';
    import type { SelectableItem } from '$lib/interfaces/ui';
    import { mergeReservationsIntoCalendar } from '$lib/utils/calendar';
    import { untrack } from 'svelte';
    import { intoDateString } from '$lib/utils/date';
    import type { MinimalCalendarUIItemWithHref } from '$lib/interfaces/calendar';
    import type { ReservationSingleResponse } from '$lib/interfaces/api';

    let { venue_list = [], calendar = [], getReservations } = $props();
    const data = $state({
        requester: {
            is_open: true,
            name: '',
            school_id: '',
            date_of_birth: ''
        },
        reservation: {
            is_open: true,
            venue: '',
            date: '',
            time: []
        },
        note: {
            is_open: true,
            agreement: false
        }
    });

    const validation_errors = $derived({
        requester: {
            name: data.requester.name.length <= 0,
            school_id: isNaN(parseInt(data.requester.school_id)),
            date_of_birth: data.requester.date_of_birth == ''
        }
    });

    let venue_seletable_list_items: SelectableItem[] = $state([]);
    $effect(() => {
        venue_seletable_list_items = venue_list.map((venue) => {
            return {
                key: venue.venue,
                label: venue.venue,
                toggle: false
            };
        });
    });

    // Saved Resrvations.
    let reservations: ReservationSingleResponse[] = [];
    $effect(() => {
        const current_venue = venue_seletable_list_items.find((v) => v.toggle)?.label || '';
        if (data.reservation.venue !== current_venue) {
            data.reservation.venue = current_venue;

            if (current_venue) {
                (async () => {
                    console.log('calling mergeReservationsIntoCalendar');
                    calendar_props.status = 'loading';
                    reservations = await getReservations(null, current_venue);
                    console.log('reservation', reservations);

                    calendar_props.data.items = mergeReservationsIntoCalendar(
                        reservations,
                        untrack(() => calendar_props.data.items)
                    );
                    untrack(() => {
                        data.reservation.date = '';
                        calendar_props.data.items
                            .filter((i) => i.mark.selected)
                            .forEach((i) => (i.mark.selected = false));
                    });
                    calendar_props.status = 'available';
                    time_selectable_props.disabled = true;
                    time_selectable_props.list.forEach((i) => {
                        i.toggle = false;
                        i.disabled = false;
                    });
                })();
            }
        }
    });

    const zeroPad = (text?: string) => text?.padStart(2, '0');
    const time_selectable_props = $state({
        list: Array(24)
            .keys()
            .map((i) => {
                return {
                    key: zeroPad(i.toString()),
                    label: `${zeroPad(i.toString())}시`,
                    toggle: false
                };
            })
            .toArray(),
        disabled: true
    });

    const calendar_props = $state({
        data: {
            items: calendar
        },
        status: 'disabled',
        onDateClick: (date: MinimalCalendarUIItemWithHref) => {
            if (calendar_props.status == 'loading' || calendar_props.status == 'disabled') return;
            console.log(calendar_props.status);
            if (date.mark?.past || date.mark?.unavailable) return;

            // unselect selected date
            untrack(() => {
                data.reservation.date = '';
                calendar_props.data.items
                    .filter((i) => i.mark.selected)
                    .forEach((i) => (i.mark.selected = false));
            });

            data.reservation.date = intoDateString(date.date);
            if (!date.mark) date.mark = {};
            date.mark.selected = true;

            const dateString = intoDateString(date.date);
            const unavailableHours = reservations
                .filter((r) => r.date == dateString)
                .flatMap((r) =>
                    r.reservations
                        .flatMap((r2) => r2.time)
                        .concat(r.unavailable_periods.flatMap((r2) => r2.time))
                );

            time_selectable_props.list.forEach((h) => {
                h.toggle = false;
                h.disabled = unavailableHours.includes(h.key);
            });
            time_selectable_props.disabled = false;
        }
    });

    $effect(() => {
        data.reservation.time = time_selectable_props.list
            .filter((i) => i.toggle)
            .map((i) => i.key);
    });
    
</script>

<CollapsibleBlock open={data.requester.is_open}>
    {#snippet header()}
        신청자 정보
    {/snippet}

    <InputBox title="이름" placeholder="여기에 이름 작성" bind:value={data.requester.name}>
        <ValidateMessage
            isError={validation_errors.requester.name}
            message="이름을 작성해 주세요."
        />
    </InputBox>

    <InputBox title="학번" placeholder="학교 학번" bind:value={data.requester.school_id}>
        <ValidateMessage
            isError={validation_errors.requester.school_id}
            message="올바른 학번을 작성해 주세요."
        />
    </InputBox>

    <InputBox title="생년월일" inputType="date" bind:value={data.requester.date_of_birth}>
        <ValidateMessage
            isError={validation_errors.requester.date_of_birth}
            message="생년월일을 입력해주세요."
        />
    </InputBox>

    올바르게 작성하지 않을 경우 불이익이 발생할 수 있습니다.
</CollapsibleBlock>

<CollapsibleBlock open={data.reservation.is_open}>
    {#snippet header()}
        장소 및 시간
    {/snippet}
    <InputBox title="장소 선택">
        {#snippet custom()}
            <SelectableList bind:list={venue_seletable_list_items} isRadio={true} />

            {#if data.reservation.venue != ''}
                <!-- eslint-disable-next-line svelte/no-at-html-tags -->
                {@html venue_list.find((v) => v.venue == data.reservation.venue).requirement}
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
            <SelectableList {...time_selectable_props} />
        {/snippet}
    </InputBox>
</CollapsibleBlock>

<CollapsibleBlock open={data.note.is_open}>
    {#snippet header()}
        유의사항
    {/snippet}

    유의사항 입력
</CollapsibleBlock>

<Button>예약하기</Button>
{JSON.stringify(data)}
