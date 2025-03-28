<script lang="ts">
    import { X } from '@lucide/svelte';
    interface Props {
        /// false by default.
        hidden?: boolean;
        /// Will be displayed in the announcement if not undefined.
        text?: string;
    }
    let { hidden = false, text }: Props = $props();
    const clickHandler = () => {
        hidden = true;
    };
</script>

{#if !hidden && text}
    <div class="announcement">
        <b>알립니다 <button onclick={clickHandler}><X size="20px" /></button></b>
        <p>{text || '공지사항이 비어 있습니다.'}</p>
    </div>
{/if}

<style lang="sass">
div.announcement
    display: flex
    flex-direction: column
    justify-content: center
    padding: 16px 24px
    background: var(--color-bg)
    color: white
    b
        display: flex
        gap: 12px
        justify-content: space-between
        align-items: center
        font-size: 20px
        font-weight: 600
        button
            display: flex
            background: var(--color-bg)
            border: none
            border-radius: 999px
            color: white
            padding: 6px
            cursor: pointer
            &:hover
                background: var(--color-hover)
            &:active
                background: var(--color-active)
</style>