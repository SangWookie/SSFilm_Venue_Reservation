<script lang="ts">
    import Calendar from '$lib/components/ui/calendar.svelte';
    import { getWeek } from '$lib/calculator';
    import NavbarEmbed from '$lib/components/ui/navbar-embed.svelte';
    import type { MinimalCalendarUIItemWithHref } from '$lib/interfaces/calendar';
    //import { getReservations } from '$lib/api/nonstate.mock';

    const calendar_props = $state({
        data: {
            items: [
                ...getWeek(2025, 0).slice(2, 4),
                //...getWeek(2025, 0),
                ...getWeek(2025, 1),
                ...getWeek(2025, 2),
                ...getWeek(2025, 3),
                ...getWeek(2025, 4)
            ].map((item) => {
                return {
                    date: item,
                    mark: []
                };
            })
        },
        status: 'available' as 'available' | 'loading' | 'disabled',
        onPositionChangeRequest: undefined,
        onDateClick: (num: MinimalCalendarUIItemWithHref) => {
            console.log(num);
            num.mark =
                num.mark && num.mark.length > 0
                    ? []
                    : ['reserved', 'unavailable', 'today', 'selected'];
        }
    });

    $effect(() => {});
</script>

<NavbarEmbed />
<Calendar {...calendar_props} />
