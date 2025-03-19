<script lang="ts">
    import CollapsibleBlock from '$lib/components/ui/form/collapsible-block.svelte';
    import InputBox from '$lib/components/ui/form/input-box.svelte';
    import ValidateMessage from '$lib/components/ui/form/validate-message.svelte';
    import type { FormData, InternalStates, Validations } from '../index.ts';
    const {
        form_data = $bindable(),
        validations = $bindable(),
        internal_states = $bindable()
    }: {
        form_data: FormData;
        validations: Validations;
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
            isValid={validations.requester_info.name}
            message="이름을 작성해 주세요."email
        />
    </InputBox>

    <InputBox title="학번" placeholder="학교 학번" bind:value={form_data.requester_info.school_id}>
        <ValidateMessage
            isValid={validations.requester_info.school_id}
            message="올바른 학번을 작성해 주세요."
        />
    </InputBox>

    <InputBox title="이메일" bind:value={form_data.requester_info.email} description="예약 완료 시 이메일로 알려드립니다.">
        <ValidateMessage
            isValid={validations.requester_info.email}
            message="이메일을 입력해주세요."
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
