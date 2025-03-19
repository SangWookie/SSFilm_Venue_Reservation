<script lang="ts">
    import { CalendarClock, X } from '@lucide/svelte';
    import NewReservationForm from '$lib/components/form/new-reservation/form.svelte';
</script>

<div class="top-header">
    <div class="title">
        <CalendarClock size={24} /> 예약하기
    </div>
</div>

<div class="page">
    <div class="layout">
        <NewReservationForm />
    </div>
</div>

<style lang="sass">
div.top-header
    display: flex
    padding: 32px 0
    flex-direction: column
    width: 100%
    align-items: center

    div.title
        display: flex
        gap: 12px
        align-items: center
        justify-content: center

        font-size: 32px
        font-weight: 500

div.page
    display: flex
    flex-direction: column
    align-items: center
    gap: 32px
    padding: 12px
    div.layout
        display: flex
        gap: 32px
        flex-direction: column
        max-width: 800px
        width: 100%
</style>
