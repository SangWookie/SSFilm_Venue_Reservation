<script lang="ts">
    import Calendar from '$lib/components/ui/calendar.svelte';
    import NavbarEmbed from '$lib/components/ui/navbar-embed.svelte';
    import { getCalendarPlaceholder, mergeReservationsIntoCalendar } from '$lib/utils/calendar';
    import { getReservations } from '$lib/api/mock';
    import type { ReservationSingleResponse } from '$lib/interfaces/api';
    import type { MinimalCalendarUIItemWithHref } from '$lib/interfaces/calendar';
    import LoadingBox from '$lib/components/ui/loading_box.svelte';
    //import { getReservations } from '$lib/api/nonstate.mock';

    const calendar_props = $state({
        items: getCalendarPlaceholder(),
        status: 'loading' as 'available' | 'loading' | 'disabled',
        onPositionChangeRequest: undefined
    });

    let loading_status = $state(false);
    let reservations: ReservationSingleResponse[] = $state([]);

    $effect(() => {
        loading_status = true;
        getReservations().then((res) => {
            reservations = res;
            calendar_props.items = mergeReservationsIntoCalendar(
                reservations,
                calendar_props.items,
                (date, item) => {
                    (item as MinimalCalendarUIItemWithHref).href = `#date-${date}`;
                }
            );
            calendar_props.status = 'available';
            loading_status = false;
        });
    });
</script>

<!--
<NavbarEmbed />
-->
<LoadingBox enabled={loading_status} size={32} />

<div class="page">
    <div class="calendar-wrapper">
        <Calendar {...calendar_props} />
    </div>
    <!-- current it breaks some centering issue however will fixed afterwards.-->
    <ul class="reservations">
        {#each reservations as reservation (reservation)}
            <li id={`date-${reservation.date}`}>
                <ul>
                    <li>기간: {reservation.date}</li>
                    <li>장소: {reservation.venue}</li>
                    <li>
                        예약:
                        <ul>
                            {#each reservation.reservations as r (r)}
                                <li>
                                    {r.time.map((i) => `${i}시`).join(', ')}
                                </li>
                            {/each}
                        </ul>
                    </li>
                </ul>
            </li>
        {/each}
    </ul>
</div>

<style lang="sass">
div.page
    display: flex
    flex-direction: row
    flex-wrap: wrap
    justify-content: center
    gap: 12px
    height: 100%

    div.calendar-wrapper
        display: flex
        flex-direction: column
        justify-content: center
        flex-grow: 2
        min-width: 350px
    
    ul.reservations
        overflow-y: scroll
        height: 100%
        flex-grow: 1

</style>
