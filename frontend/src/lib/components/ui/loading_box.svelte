<script lang="ts">
    import LoadingIcon from '$lib/components/ui/loading_icon.svelte';

    let { size = 16, enabled = false }: { size?: number; enabled: boolean } = $props();
</script>

<div class="loading-box" class:enabled style:--size={`${size}px`}>
    <LoadingIcon {size} />
</div>

<style lang="sass">
div.loading-box
    display: flex
    justify-content: center
    align-items: center
    gap: 8px
    opacity: 0
    transition: all .2s
    pointer-events: none
    height: 0
    padding: 0

    &.enabled
        opacity: 1
        height: var(--size, 16)
        padding: 12px
</style>
