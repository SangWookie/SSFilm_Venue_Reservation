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
</script>

<details bind:open class="ui-form-collapsible-block">
    <summary>
        {@render header?.()}
        <div class="control">
            <ChevronDownIcon />
        </div>
    </summary>

    <div class="content">
        {@render children?.()}
    </div>
</details>

<style lang="sass">
details
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
            transition: transform 0.2s
            display: flex

    &:open
        summary
            .control
                transform: rotate(180deg)

    div.content
        display: flex
        flex-direction: column
        gap: 16px
        padding-top: 24px
</style>
