<script module lang="ts">
    import { defineMeta } from '@storybook/addon-svelte-csf';
    import Template from './new-reservations.svelte';
    import { getReservations, getVenueList } from '$lib/api/mock';
    import type { Venue } from '$lib/interfaces/api';
    import type { MinimalCalendarUIItemMark, MinimalCalendarUIItemWithHref } from '$lib/interfaces/calendar';
    import { getTwoWeekRange } from '$lib/utils/calendar';
    import { DateTime } from 'luxon';

    const { Story } = defineMeta({
        title: 'UI/Form/Templates/New Reservations',
        component: Template
    });
</script>

<script lang="ts">
    let venue_list: Venue[] = $state([]);
    let calendar: MinimalCalendarUIItemWithHref[] = getTwoWeekRange().map(date => {
        const mark: MinimalCalendarUIItemMark = {};
        
        if (date.hasSame(DateTime.local(), 'day'))
            mark.today = true;
        if (date.diff(DateTime.local(), 'hours').hours < -1)
            mark.past = true;
        
        
        return {
            date, mark
        }
    })
    
    console.log(calendar.map(c => c.date.toISODate()))

    $effect(() => {
        (async () => {
            venue_list = await getVenueList();
            console.log(venue_list)
        })();
    });
</script>

<Story name="default">
    <Template {venue_list} {calendar} {getReservations} />
</Story>