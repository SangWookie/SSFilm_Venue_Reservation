<script lang="ts">
    import { ChevronDownIcon } from '@lucide/svelte';
    import type { Snippet } from 'svelte';

    let {
        open = $bindable(true),
        header,
        children
    }: {
        open: boolean;
        header?: Snippet;
        children?: Snippet;
    } = $props();

    let offsetHeight = $state(0);
</script>

<div class="ui-form-collapsible-block" class:open={open}>
    <summary>
        {@render header?.()}
        <button class="control" onclick={() => open = !open}>
            <ChevronDownIcon />
        </button>
    </summary>
    <div class="content-wrapper" style:--content-height={offsetHeight + 24 + 'px'}>
        <div class="content" bind:offsetHeight={offsetHeight}>
            {@render children?.()}
        </div>
    </div>
</div>

<style lang="sass">
.ui-form-collapsible-block
    display: flex
    flex-direction: column
    background: var(--color-slate-100)
    padding: 16px 24px
    border-radius: 12px
    summary
        font-size: 24px
        display: flex
        align-items: center
        justify-content: space-between
        .control
            background: none
            border: none
            transition: transform 0.2s
            display: flex

    &.open
        summary
            .control
                transform: rotate(180deg)
        .content-wrapper
            height: var(--content-height)
    &:not(.open)
        .content-wrapper
            opacity: 0
            overflow: hidden

    div.content
        display: flex
        flex-direction: column
        gap: 16px
        padding-top: 24px
    div.content-wrapper
        transition: all 0.3s
        height: 0px
        overflow: hidden
</style>
