<script lang="ts">
    import CollapsibleBlock from '../../../collapsible-block.svelte';
    import InputBox from '../../../input-box.svelte';
    import ValidateMessage from '../../../validate-message.svelte';
    import type { FormData, FormProps, InternalStates, Validations } from '../form.ts';
    const {
        form_data,
        validations,
        form_props,
        internal_states
    }: {
        form_data: FormData;
        validations: Validations;
        form_props: FormProps;
        internal_states: InternalStates;
    } = $props();
</script>

<CollapsibleBlock open={internal_states.collapsed.requester_info}>
    {#snippet header()}
        신청자 정보
    {/snippet}

    <InputBox
        title="이름"
        placeholder="여기에 이름 작성"
        bind:value={form_data.requester_info.name}
    >
        <ValidateMessage
            isError={!validations.requester_info.name}
            message="이름을 작성해 주세요."
        />
    </InputBox>

    <InputBox 
        title="학번" 
        placeholder="학교 학번" 
        bind:value={form_data.requester_info.school_id}>
        <ValidateMessage
            isError={!validations.requester_info.school_id}
            message="올바른 학번을 작성해 주세요."
        />
    </InputBox>

    <InputBox title="생년월일" inputType="date" bind:value={form_data.requester_info.date_of_birth}>
        <ValidateMessage
            isError={!validations.requester_info.date_of_birth}
            message="생년월일을 입력해주세요."
        />
    </InputBox>

    <p class="note">올바르게 작성하지 않을 경우 불이익이 발생할 수 있습니다.</p>
</CollapsibleBlock>

<style lang="sass">
p.note
    text-align: center
    width: 100%
    opacity: .6
</style>