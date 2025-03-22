<script lang="ts">
    import './../../app.sass';
    import Calendar from '$lib/components/ui/calendar.svelte';
    import { getCalendarPlaceholder } from '$lib/utils/calendar';
    import type { ReservationList } from '$lib/interfaces/api';
    import type { MinimalCalendarUIItemWithHref } from '$lib/interfaces/calendar';
    import LoadingBox from '$lib/components/ui/loading_box.svelte';
    import { onMount } from 'svelte';
    import { getReservationByDate } from '$lib/api/api';
    import { intoDateString } from '$lib/utils/date';
    //import { getReservations } from '$lib/api/nonstate.mock';
    
    import NavbarEmbed from '$lib/components/ui/navbar-embed.svelte'
    
    import VenueHeaderInformationComponent from '$lib/components/ui/information/venue-header.svelte';
    import VenueInformationComponent from '$lib/components/ui/information/venue.svelte';
    
    let response: ReservationList | undefined = $state(undefined);

    const calendar_props = $state({
        items: [] as MinimalCalendarUIItemWithHref[],
        selected: [] as MinimalCalendarUIItemWithHref[],
        status: 'available' as 'available' | 'loading' | 'disabled',
        onPositionChangeRequest: undefined,
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        onDateClick: (date?: MinimalCalendarUIItemWithHref) => {}
    });
    
    calendar_props.onDateClick = (date?: MinimalCalendarUIItemWithHref) => {
        if (!date || loading_status) return;
        console.log('date clicked', date);
        loading_status = true;
        calendar_props.selected = [date];
        getReservationByDate(intoDateString(date.date))
            .then(res => {
                calendar_props.selected = [date];
                loading_status = false;
                response = res;
            })
    }

    let loading_status = $state(false);

    onMount(() => {
        // Performance issue
        // https://github.com/moment/luxon/issues/1130
        async function load() {
            calendar_props.items = getCalendarPlaceholder();
        }
        load();
    });
</script>


<NavbarEmbed href="https://ssfilm-demo-ilsubyeega.pages.dev/new"/>

<div class="page">
    <div class="calendar-wrapper">
        <Calendar {...calendar_props} />
    </div>
    <!-- current it breaks some centering issue however will fixed afterwards.-->
    <ul class="reservations">
        <LoadingBox enabled={loading_status} size={32} />
        {#if response}
            <div class="date-header">
                {response.date}
            </div>
            <VenueHeaderInformationComponent/>
            <div class="venues">
                {#each response.venues as venue (venue)}
                    <VenueInformationComponent data={venue} />
                {/each}
            </div>
        {/if}

    </ul>
</div>

<style lang="sass">
div.page
    display: flex
    flex-direction: column
    justify-content: center
    align-items: center
    gap: 12px
    overflow-x: hidden

    div.calendar-wrapper
        display: flex
        flex-direction: column
        justify-content: center
        flex-grow: 2
        min-width: 280px
        max-width: 500px
        width: 100%

    ul.reservations
        display: flex
        flex-direction: column
        gap: 16px
        max-width: 800px
        padding: 16px
        box-sizing: border-box
        flex-grow: 1
        max-width: 850px
        width: 100%
        
        div.date-header
            font-size: 48px
            font-weight: 300
            text-align: center
        
        div.venues
            display: flex
            flex-direction: column
            gap: 12px

</style>
