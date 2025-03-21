<script lang="ts">
    import type { SelectableItem, SelectableListState } from './selectable-list.svelte.ts';
    import type { Snippet } from 'svelte';
    import { Check } from '@lucide/svelte';

    const defaultClickHandler = (item: SelectableItem<T>) => {
        if (disabled || item.disabled) return;
        if (isRadio) {
            data.selected = [item];
        } else {
            data.toggle(item);
        }
    };

    type T = $$Generic;

    let {
        data = $bindable(),
        disabled = $bindable(false),
        isRadio = $bindable(false),
        clickHandler = defaultClickHandler,
        labelSnippet
    }: {
        data: SelectableListState<T>;
        disabled?: boolean;
        isRadio?: boolean;
        clickHandler?: (item: SelectableItem<T>) => void;
        labelSnippet?: Snippet<[item: SelectableItem<T>]>;
    } = $props();
</script>

<div class="ui-form-selectable-list" class:disabled>
    {#each data.list as item (item.key)}
        <button
            class="item"
            class:disabled={item.disabled}
            class:selected={data.isSelected(item)}
            onclick={() => clickHandler(item)}
        >
            <div class="check">
                <Check size={20}/>
            </div>
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
        filter: grayscale(1)
        button.item
            &:hover, &:active, &:focus
                background-color: var(--color-main-100) !important

    button.item
        display: flex
        align-items: center
        flex-direction: row
        flex-shrink: 0
        border: 0
        border-radius: 99999px
        padding: 8px 16px
        font-size: 16px
        font-weight: 500
        background-color: var(--color-main-100)
        color: var(--color-main-500)
        transition: background-color 0.2s, color 0.2s
        
        div.check
            display: flex
            align-items: center
            width: 0
            opacity: 0
            transition: all 0.3s

        &:hover
            background-color: var(--color-main-50)
        &:active
            background-color: var(--color-main-200)

        &.selected
            background-color: var(--color-main-500)
            color: white
            &:hover
                background-color: var(--color-main-400)
            &:active
                background-color: var(--color-main-600)
            div.check
                width: 24px
                opacity: 1

        &.disabled
            background: none !important
            color: var(--color-slate-600)
            text-decoration: line-through
            opacity: .1
            cursor: not-allowed
            &:hover, &:active, &:focus
                background: none !important

</style>
