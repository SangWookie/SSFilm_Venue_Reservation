<script lang="ts">
    import type { SelectableItem } from '$lib/interfaces/ui';

    let {
        list = $bindable([]),
        disabled = $bindable(false),
        isRadio = $bindable(false)
    }: {
        list: SelectableItem[];
        disabled?: boolean;
        isRadio?: boolean;
    } = $props();

    const clickHandler = (item: SelectableItem) => {
        if (disabled || item.disabled) return;
        if (isRadio) {
            list.filter((i) => i != item).forEach((i) => (i.toggle = false));
            item.toggle = true;
        } else {
            item.toggle = !item.toggle;
        }
    };
</script>

<div class="ui-form-selectable-list" class:disabled={disabled}>
    {#each list as item (item.label)}
        <button
            class="item"
            class:disabled={item.disabled}
            class:toggle={item.toggle}
            onclick={() => clickHandler(item)}
        >
            {item.label}
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
