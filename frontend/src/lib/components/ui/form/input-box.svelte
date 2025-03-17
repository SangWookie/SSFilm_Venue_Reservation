<script lang="ts">
    import type { Snippet } from 'svelte';

    let {
        title,
        description,
        placeholder,
        inputType = 'text',
        isError = $bindable(false),
        value = $bindable(),
        children,
        custom
    }: {
        title: string;
        description?: string;
        placeholder?: string;
        inputType?: string;
        isError?: boolean;
        value?: string;
        children?: Snippet;
        custom?: Snippet;
    } = $props();
</script>

<div class="ui-form-input-box" class:error={isError}>
    <div class="header">
        <div class="title">
            {title}
        </div>
        {#if description}
            <div class="description">
                {description}
            </div>
        {/if}
    </div>
    {#if !custom}
        <input type={inputType} bind:value {placeholder} />
    {:else}
        {@render custom()}
    {/if}

    {@render children?.()}
</div>

<style lang="sass">
div.ui-form-input-box
    display: flex
    flex-direction: column
    gap: 12px

    div.header
        display: flex
        align-items: center
        gap: 8px
        div.title
            font-size: 24px
        div.description
            font-size: 16px
            opacity: .8

    input
        padding: 12px
        border: 1px solid white
        border-radius: 8px
        font-size: 16px
</style>
