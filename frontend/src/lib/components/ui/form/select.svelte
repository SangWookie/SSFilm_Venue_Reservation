<script lang="ts">
    import type { FormSelectItem } from "$lib/interfaces/ui";

    // https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
    interface Props {
        disabled?: boolean,
        required?: boolean,
        options: FormSelectItem<unknown>[],
        /// The current value of the select element
        value: FormSelectItem<unknown> | undefined,
    }

    let {
        disabled = $bindable(false),
        required = $bindable(false),
        options = $bindable([]),
        value = $bindable(undefined),
    }: Props = $props();
    
    let bindedValue = $state();
    $effect(() => {
        value = options.find(opt => opt.key === bindedValue);
    });
    </script>

<select 
    class="ui-form-select"
    {disabled} 
    {required} 
    bind:value={bindedValue}
>
    {#each options as item (item.label)}
        <option value={item.key} selected={item.key === value?.key}>
            {#if item.labelSnippet}
                <!-- eslint-disable-next-line svelte/no-at-html-tags -->
                {@html item.labelSnippet()}
            {:else}
                {item.label}
            {/if}
        </option>
    {/each}
</select>

<style lang="sass">
select.ui-form-select
    padding: 12px
    border: 1px solid white
    border-radius: 8px
    font-size: 16px
    background: white
</style>