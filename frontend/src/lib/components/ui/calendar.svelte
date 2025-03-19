<script lang="ts">
    import type { DateTime } from 'luxon';
    import { ArrowLeft, ArrowRight } from '@lucide/svelte';

    import type { MinimalCalendarUIItemWithHref } from '$lib/interfaces/calendar';
    import { generateCalendarFromProps } from '$lib/utils/calendar';
    import { untrack } from 'svelte';

    interface Props {
        /// The dates to be displayed in the calendar
        /// IT must be starts with Sunday and ends with Saturday
        items: MinimalCalendarUIItemWithHref[];
        selected?: MinimalCalendarUIItemWithHref[];
        /**
         * The UI state of calendar, not the item.
         * available: The calendar is available
         * loading: The calendar is loading
         * disabled: The calendar is disabled. You cannot hover or click them.
         */
        status: 'available' | 'loading' | 'disabled';
        /**
         * callbacks when position change requests, like before/next week, or custom week search.
         */
        onPositionChangeRequest?: (start_week: number) => void;
        /**
         * callbacks when a date is clicked.
         */
        onDateClick?: (date: MinimalCalendarUIItemWithHref | undefined, event: MouseEvent) => void;
    }

    const props: Props = $props();

    let previous_data_items = props.items;
    let generated_calendars: (MinimalCalendarUIItemWithHref | undefined)[][] = $state([]);

    const generateDateText = (datetime?: DateTime) =>
        datetime?.toLocaleString({
            year: 'numeric',
            month: '2-digit',
            day: '2-digit'
        }) || '??';

    $effect(() => {
        // If generated calendars are not exist, or data is new, create it.
        if (
            generated_calendars == null ||
            generated_calendars.length === 0 ||
            previous_data_items != props.items
        ) {
            generated_calendars = generateCalendarFromProps(untrack(() => props.items));
        }
    });

    const derivation = $derived.by(() => {
        const start_day = props.items.at(0);
        const end_day = props.items.at(-1);

        const placeholders = {
            start_day: {
                text: generateDateText(start_day?.date),
                week: start_day?.date.weekNumber
            },
            end_day: {
                text: generateDateText(end_day?.date),
                week: end_day?.date.weekNumber
            }
        };

        const moveWeekDeltaCallback = (delta: number) =>
            placeholders.start_day.week &&
            props.onPositionChangeRequest?.(placeholders.start_day.week + delta);

        return { placeholders, moveWeekDeltaCallback };
    });

    const zeroPad = (text?: string) => text?.padStart(2, '0');
</script>

{#snippet dateBlock(data?: MinimalCalendarUIItemWithHref)}
    {@const number = data?.date?.day?.toString()}
    {@const weekNumber = data?.date?.weekday}
    <td
        class="date-block"
        class:sun={weekNumber === 7}
        class:reserved={data?.mark?.reserved}
        class:unavailable={data?.mark?.unavailable}
        class:past={data?.mark?.past}
        class:today={data?.mark?.today}
        class:selected={data && props.selected?.includes(data)}
        class:blank={!number}
    >
        {#if number}
        <a href={data?.href} onclick={(e) => props?.onDateClick?.(data, e)}>
                {zeroPad(number)}
                <div class="indicator-dot"></div>
            </a>
        {/if}
    </td>
{/snippet}

<div class={['calendar', props.status]}>
    {#if props.onPositionChangeRequest}
        <div class="header">
            <button
                onclick={() => derivation.moveWeekDeltaCallback(-4)}
                class:disabled={!props.onPositionChangeRequest}
                title="4주 전으로 이동"><ArrowLeft /></button
            >
            <button class="date">
                {derivation.placeholders.start_day.text} ~ {derivation.placeholders.end_day.text}
                <b>
                    Week {zeroPad(derivation.placeholders.start_day.week?.toString())} ~ {zeroPad(
                        derivation.placeholders.end_day.week?.toString()
                    )}
                </b>
            </button>
            <button
                onclick={() => derivation.moveWeekDeltaCallback(+4)}
                class:disabled={!props.onPositionChangeRequest}
                title="4주 후로 이동"><ArrowRight /></button
            >
        </div>
    {/if}
    <table class="data">
        <thead>
            <tr>
                <th>일</th>
                <th>월</th>
                <th>화</th>
                <th>수</th>
                <th>목</th>
                <th>금</th>
                <th>토</th>
            </tr>
        </thead>
        <tbody class:clickable={true}>
            <!-- eslint-disable-next-line svelte/require-each-key -->
            {#each generated_calendars as week}
                <tr class="week">
                    <!-- eslint-disable-next-line svelte/require-each-key -->
                    {#each week as day}
                        {@render dateBlock(day)}
                    {/each}
                </tr>
            {/each}
        </tbody>
    </table>
</div>

<style lang="sass">
div.calendar
    view-transition-name: calendar
    display: flex
    flex-direction: column
    max-width: 500px
    flex-grow: 1
    &.disabled
        opacity: .4
        cursor: not-allowed
        a
            cursor: not-allowed !important
    div.header
        display: flex
        gap: 12px
        justify-content: center
        button
            display: flex
            flex-direction: column
            align-items: center
            justify-content: center
            background: none
            border: none
            padding: 12px
            border-radius: 999px
            &:hover
                background: rgba(0, 0, 0, 0.1)
            &.date
                font-size: 12px
                padding: 12px 16px
                b
                    font-size: 24px
                    font-weight: 600
            &.disabled
                opacity: .4
    table.data
        thead
            color: var(--color-bg)
            th
                padding: 16px 0
        tbody
            text-align: center
            tr.week
                td.date-block
                    font-weight: 600
                    border-radius: 12px
                    border: 2px solid transparent;
                    box-sizing: border-box
                    &.reserved
                        a
                            .indicator-dot
                                opacity: 1
                                display: block
                    &.unavailable, &.past
                        a
                            color: #ccc !important
                            text-decoration: line-through
                    &.today
                        border: 2px solid var(--color-bg)
                    &.selected
                        background: var(--color-bg)
                        &:hover
                            background: var(--color-hover)
                        a
                            color: white !important
                            .indicator-dot
                                background-color: white
                    a
                        padding: 16px 8px
                        position: relative
                        display: flex
                        flex-direction: column
                        align-items: center
                        justify-content: center
                        color: black
                        text-decoration: none
                        cursor: pointer
                        div.indicator-dot
                            position: absolute
                            bottom: 8px
                            width: 6px
                            height: 6px
                            border-radius: 999px
                            background: var(--color-bg)
                            margin-top: 4px
                            display: none
                    &.sun
                        a
                            color: var(--color-bg)
                    
                    &.blank
                        a
                            display: none
                        &:hover, &:active, &:focus
                            background: none

                    &:hover
                        background: rgba(0, 0, 0, 0.1)

// loading animation
div.calendar.loading
    animation: pulse-loading 0.75s infinite ease-in-out
    transition: opacity 0.3s

@keyframes pulse-loading
    0%, 100%
        opacity: 0.2
    50%
        opacity: 0.4
</style>
