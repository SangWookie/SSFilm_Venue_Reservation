<script lang="ts">
    import type { ReservationByVenue } from '$lib/interfaces/api';
    import { getUnavilableHoursByVenue } from '$lib/utils/api';
    import { generateHours, getHourRangeString, splitArray, zeroPad } from '$lib/utils/helper';

    const { data }: { data: ReservationByVenue } = $props();

    const hours = generateHours();
    const unavailableHours = $derived(getUnavilableHoursByVenue(data));
    
    const has_item = $derived(unavailableHours.length > 0);
</script>

<div class="venue-item" class:no_item={!has_item}>
    <div class="venue-name">
        {data.venueKor}
        {data.approval_mode == 'manual' ? '(수동)' : ''}
    </div>
    {#if has_item}
        <!--
        * client request: remove this;
        <div class="hours-list">
            {#each splitArray(hours, 6) as hours_spllited (hours_spllited)}
                <div class="hours-inner">
                    {#each hours_spllited as hour (hour)}
                        {@const hourString = zeroPad(hour.toString())}
                        {@const is_unavailable = unavailableHours.includes(hour)}
                        <div
                            class="hour-single"
                            class:unavailable={is_unavailable}
                            title={is_unavailable ? '사용 불가능' : '사용 가능'}
                        >
                            {hourString}
                        </div>
                    {/each}
                </div>
            {/each}
        </div>
        -->
    {:else}
        예약이 비어있습니다.
    {/if}

    {#each data.unavailable_periods as item (item)}
        <div class="unavailable-period">
            <div class="label">
                예약 불가
            </div>
            <div class="message">
                {item.message}
            </div>
            <div class="hours">
                {getHourRangeString(item.time)} <div class="extra">시</div>
            </div>
        </div>
    {/each}

    {#each data.reservations as item (item)}
        <div class="reservation">
            <div class="label">예약</div>
            <div class="requester" title="예약자">
                {item.name}
            </div>
            <div class="purpose" title="사유">
                {item.purpose}
            </div>
            <div class="hours" title="예약한 시간">
                {getHourRangeString(item.time)} <div class="extra">시</div>
            </div>
        </div>
    {/each}
</div>

<style lang="sass">
div.venue-item
    display: flex
    flex-direction: column
    gap: 8px
    
    &.no_item
        flex-direction: row
        align-items: center
        opacity: .4
        div.venue-name
            font-size: 20px
    div.venue-name
        font-size: 30px
    div.hours-list
        display: flex
        flex-wrap: wrap
        font-size: 20px
        gap: 0 8px
        div.hours-inner
            display: flex
            flex-shrink: 0
            gap: 8px
            div.hour-single
                font-family: "Space Mono", monospace, "Pretendard"
                display: flex
                align-items: center
                justify-content: center

                opacity: .2
                color: var(--color-bg)
                font-weight: 600
                &.unavailable
                    opacity: 1

    div.unavailable-period
        font-size: 20px
        display: flex
        flex-wrap: wrap
        align-items: center
        gap: 4px 12px
        div.label
            font-size: 16px
            font-weight: 600
            color: var(--color-bg)
            flex-shrink: 0 
        div.message
            font-size: 16px
        div.hours
            font-family: "Space Mono", monospace, "Pretendard"
            display: flex
            align-items: center
            gap: 8px
            font-size: 20px
            color: var(--color-bg)
            font-weight: 600
            .extra
                font-size: 16px
                font-weight: 500
                opacity: .4

    div.reservation
        font-size: 20px
        display: flex
        flex-wrap: wrap
        align-items: center
        gap: 4px 12px
        div.label
            font-size: 16px
            font-weight: 600
            color: var(--color-bg)
            flex-shrink: 0
        div.requester
            font-weight: 500
        div.purpose
            font-size: 16px
        div.hours
            display: flex
            align-items: center
            gap: 8px
            font-size: 16px
            font-family: "Space Mono", monospace, "Pretendard"
            color: var(--color-bg)
            font-weight: 600
            .extra
                font-size: 16px
                font-weight: 500
                opacity: .4


</style>
