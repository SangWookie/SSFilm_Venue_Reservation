<script lang="ts">
    import { LoaderCircle } from '@lucide/svelte';

    let { size = 16 }: { size?: number } = $props();
</script>

<div class="loading-icon">
    <LoaderCircle {size} />
</div>

<style lang="sass">
div.loading-icon
    display: flex
    animation: spin 1s linear infinite

@keyframes spin
    0%
        transform: rotate(0deg)
    100%
        transform: rotate(360deg)
</style>
