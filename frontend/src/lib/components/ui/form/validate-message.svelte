<script lang="ts">
    import { X, Check } from '@lucide/svelte';
    import type { Snippet } from 'svelte';

    let {
        isValid = $bindable(false),
        message,
        children
    }: {
        isValid: boolean;
        message?: string;
        children?: Snippet;
    } = $props();
</script>

<div class="ui-form-validate-message" class:error={isValid}>
    {#if isValid}
        <Check size={16} color="green" />
    {:else}
        <X size={16} color="red" />
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
