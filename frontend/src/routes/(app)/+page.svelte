<script lang="ts">
    import Calendar from '$lib/components/ui/calendar.svelte';
    import { getCalendarPlaceholder } from '$lib/utils/calendar';
    import type { ReservationList } from '$lib/interfaces/api';
    import type { MinimalCalendarUIItemWithHref } from '$lib/interfaces/calendar';
    import LoadingBox from '$lib/components/ui/loading_box.svelte';
    import { onMount } from 'svelte';
    //import { getReservations } from '$lib/api/nonstate.mock';

    const calendar_props = $state({
        items: [] as MinimalCalendarUIItemWithHref[],
        status: 'available' as 'available' | 'loading' | 'disabled',
        onPositionChangeRequest: undefined
    });

    let loading_status = $state(false);
    let reservations: ReservationList[] = $state([]);

    onMount(() => {
        // Performance issue
        // https://github.com/moment/luxon/issues/1130
        (async () => (calendar_props.items = getCalendarPlaceholder()))();
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
