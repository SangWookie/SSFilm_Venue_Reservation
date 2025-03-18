<script lang="ts">
    import { CalendarClock, X } from '@lucide/svelte';
    import NewReservationForm from '../../lib/components/ui/form/templates/new-reservation/form.svelte'
  
</script>

<div class="top-header">
    <div class="title">
        <CalendarClock size={24} /> 예약하기
    </div>
</div>

<div class="page">
    <NewReservationForm/>
</div>

<style lang="sass">
div.page
    display: flex
    flex-direction: column
    //align-items: center
    gap: 32px
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

div.input-container
    display: flex
    flex-direction: column
    gap: 24px
    margin-top: 32px
    padding: 16px 24px
    background: var(--color-slate-100)
    border-radius: 12px
    div.header
        display: flex
        gap: 12px
        font-weight: 500
        div.title
            color: rgba(0,0,0, .4)
        div.description

    div.input
        display: flex
        flex-direction: column
        gap: 8px
        div.title
            font-size: 24px
            font-weight: 500

        input
            padding: 12px
            border: 1px solid white
            border-radius: 8px
            font-size: 16px

        button
            padding: 12px
            border: none
            border-radius: 8px
            background: var(--color-main-500)
            color: white
            font-size: 16px
            font-weight: 500
            &:hover
                background: var(--color-main-400)

        div.error
            display: flex
            gap: 8px
            align-items: center

    div.extra-note
        font-size: 12px
        font-weight: 500
        opacity: .4
        text-align: center

div.confirm
    display: flex
    flex-direction: column
    gap: 12px
    padding: 16px 24px
    border-radius: 12px
    div.errors
        display: flex
        flex-direction: column
        gap: 8px
        div.error
            display: flex
            gap: 8px
            align-items: center
    button
        padding: 12px
        border: none
        border-radius: 8px
        background: var(--color-slate-200)
        color: white
        font-size: 16px
        font-weight: 500
        cursor: pointer
        &.disabled
            background: var(--color-slate-100)
            color: var(--color-slate-200)
            cursor: not-allowed
</style>
