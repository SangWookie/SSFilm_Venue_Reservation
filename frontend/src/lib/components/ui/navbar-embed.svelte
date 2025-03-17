<script lang="ts">
    import { CalendarClock, SquareArrowOutUpRight } from '@lucide/svelte';
    const {
        href = '/new'
    }: {
        href?: string;
    } = $props();
</script>

<div class="header">
    <div class="left">
        <CalendarClock /> 예약 현황
    </div>
    <div class="right">
        <a {href}>예약하기 <SquareArrowOutUpRight size="12px" strokeWidth="3px" /></a>
    </div>
</div>

<style lang="sass">
div.header
    display: flex
    width: fill-available
    padding: 24px

    align-items: center
    justify-content: space-between
    
    div.left
        display: flex
        gap: 12px
        align-items: center

        font-size: 24px
        font-weight: 500

    div.right
        display: flex
        a
            display: flex
            align-items: center
            gap: 4px
            color: white
            text-decoration: none
            font-weight: 600
            background: var(--color-bg)
            padding: 12px 24px
            border-radius: 12px
            &:hover, &:focus
                filter: brightness(1.1)
            &:active
                filter: brightness(0.9)
</style>
