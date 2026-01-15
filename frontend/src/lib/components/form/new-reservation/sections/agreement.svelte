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

    <p class="note">정확한 정보를 입력해 주시기 바랍니다.
    <br/>균등한 공간 사용을 위해, 사용할 만큼만 예약 바랍니다.
    <br/>타인 명의 도용 발각 시 불이익이 있습니다.</p>
</CollapsibleBlock>

<style lang="sass">
p.note
    text-align: center
    width: 100%
    opacity: .6
</style>
