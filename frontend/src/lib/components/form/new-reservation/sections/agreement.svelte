<script lang="ts">
    import CollapsibleBlock from '$lib/components/ui/form/collapsible-block.svelte';
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    import InputBox from '$lib/components/ui/form/input-box.svelte';
    import ValidateMessage from '$lib/components/ui/form/validate-message.svelte';
    import type { FormData, Validations } from '../index.ts';
    const {
        form_data = $bindable(),
        validations = $bindable(),
        collapsible_open = $bindable(true)
    }: {
        form_data: FormData;
        validations: Validations;
        collapsible_open: boolean;
    } = $props();
</script>

<CollapsibleBlock open={collapsible_open}>
    {#snippet header()}
        유의사항
    {/snippet}

    <p class="note">올바르게 작성하지 않을 경우 불이익이 발생할 수 있습니다.</p>
</CollapsibleBlock>

<style lang="sass">
p.note
    text-align: center
    width: 100%
    opacity: .6
</style>
