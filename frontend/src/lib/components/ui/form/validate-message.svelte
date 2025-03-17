<script lang="ts">
    import { X, Check } from '@lucide/svelte';
    import type { Snippet } from 'svelte';

    let {
        isError = $bindable(true),
        message,
        children
    }: {
        isError: boolean;
        message?: string;
        children?: Snippet;
    } = $props();
</script>

<div class="ui-form-validate-message" class:error={isError}>
    {#if isError}
        <X size={16} color="red" />
    {:else}
        <Check size={16} color="green" />
    {/if}

    {@render children?.()}
    {message || ''}
</div>

<style lang="sass">
div.ui-form-validate-message
    display: flex
    gap: 4px
    text-size: 16px
</style>
