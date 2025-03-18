<script lang="ts">
    import type { SelectableItem } from '$lib/interfaces/ui';
    import type { Snippet } from 'svelte';

    let {
        list = $bindable([]),
        selected = $bindable([]),
        disabled = $bindable(false),
        isRadio = $bindable(false),
        labelSnippet
    }: {
        list: SelectableItem<unknown>[];
        selected?: SelectableItem<unknown>[];
        disabled?: boolean;
        isRadio?: boolean;
        labelSnippet?: Snippet<[ item: SelectableItem<unknown> ]>;
    } = $props();

    const clickHandler = (item: SelectableItem<unknown>) => {
        if (disabled || item.disabled) return;
        if (isRadio) {
            if (selected.length > 0) selected = [];
            selected = [item];
        } else {
            if (selected.includes(item)) {
                selected = selected.filter((i) => i !== item);
            } else {
                selected = list.filter((i) => selected?.includes(i) || i == item);
            }
        }
    };
</script>

<div class="ui-form-selectable-list" class:disabled>
    {#each list as item (item.label)}
        <button
            class="item"
            class:disabled={item.disabled}
            class:toggle={selected?.includes(item)}
            onclick={() => clickHandler(item)}
        >
            {#if labelSnippet}
                {@render labelSnippet(item)}
            {:else}
                {item.label}
            {/if}
        </button>
    {/each}
</div>

<style lang="sass">
div.ui-form-selectable-list
    display: flex
    gap: 4px
    flex-wrap: wrap

    &.disabled
        opacity: .6

    button.item
        flex-shrink: 0
        border: 0
        border-radius: 99999px
        padding: 8px 16px
        font-size: 16px
        font-weight: 500
        background-color: var(--color-main-100)
        color: var(--color-main-500)

        &:hover
            background-color: var(--color-main-50)
        &:active
            background-color: var(--color-main-200)

        &.toggle
            background-color: var(--color-main-500)
            color: white
            &:hover
                background-color: var(--color-main-400)
            &:active
                background-color: var(--color-main-600)

        &.disabled
            background-color: var(--color-slate-600)
            color: var(--color-slate-50)
            text-decoration: line-through
            &:hover, &:active
                background-color: var(--color-slate-50)

</style>
