<script lang="ts">
    import { CalendarClock, X } from '@lucide/svelte';
    import Button from '$lib/components/ui/button.svelte';
    import type { ReservationRequestDraft } from '$lib/interfaces/api';

    const draft: ReservationRequestDraft = $state({});
</script>

<div class="top-header">
    <div class="title">
        <CalendarClock size={24} /> 예약하기
    </div>
</div>

<div class="page">
    <div class="input-container">
        <div class="header">
            <div class="title">신청자 정보</div>
            <div class="description">신청자 정보를 입력해주세요.</div>
        </div>
        <div class="input">
            <div class="title">이름</div>
            <input type="text" placeholder="여기에 작성" bind:value={draft.name} />
            <div class="error">
                <X size={16} color="red" /> 이름을 입력해주세요.
            </div>
        </div>

        <div class="input">
            <div class="title">학번</div>
            <input type="text" placeholder="여기에 작성" bind:value={draft.studentID} />
            <div class="error">
                <X size={16} color="red" /> 올바른 학번을 입력하세요.
            </div>
        </div>

        <div class="input">
            <div class="title">이메일</div>
            <input type="email" placeholder="여기에 작성" bind:value={draft.email} />
            <div class="error">
                <X size={16} color="red" /> 올바른 이메일을 작성해주세요.
            </div>
        </div>

        <div class="input">
            <div class="title">생년월일</div>
            <input type="date" />
            <div class="error">
                <X size={16} color="red" /> 생년월일을 입력해주세요.
            </div>
        </div>

        <!--
    <div class="input">
      <div class="title">인증</div>
      <button>인증하기</button>
      <div class="error">
        <X size={16} color='red' /> 인증을 완료해주세요.
      </div>
    </div>
    -->

        <div class="extra-note">불이익이 발생할 수 있으니 올바르게 작성해주세요.</div>
    </div>

    <div class="confirm">
        <div class="errors">
            <div class="error">
                <X size={16} color="red" /> 이름을 입력해주세요.
            </div>
            <div class="error">
                <X size={16} color="red" /> 생년월일을 입력해주세요.
            </div>
        </div>
        <button class="disabled"> 예약하기 </button>
        <Button>예약하기</Button>

        {JSON.stringify(draft)}
    </div>
</div>

<style lang="sass">
div.page
    display: flex
    flex-direction: column
    //align-items: center
    gap: 32px
div.top-header
    display: flex
    padding: 32px 0
    flex-direction: column
    width: 100%
    align-items: center

    div.title
        display: flex
        gap: 12px
        align-items: center
        justify-content: center

        font-size: 32px
        font-weight: 500

div.input-container
    display: flex
    flex-direction: column
    gap: 24px
    margin-top: 32px
    padding: 16px 24px
    background: var(--color-slate-100)
    border-radius: 12px
    div.header
        display: flex
        gap: 12px
        font-weight: 500
        div.title
            color: rgba(0,0,0, .4)
        div.description

    div.input
        display: flex
        flex-direction: column
        gap: 8px
        div.title
            font-size: 24px
            font-weight: 500

        input
            padding: 12px
            border: 1px solid white
            border-radius: 8px
            font-size: 16px

        button
            padding: 12px
            border: none
            border-radius: 8px
            background: var(--color-main-500)
            color: white
            font-size: 16px
            font-weight: 500
            &:hover
                background: var(--color-main-400)

        div.error
            display: flex
            gap: 8px
            align-items: center

    div.extra-note
        font-size: 12px
        font-weight: 500
        opacity: .4
        text-align: center

div.confirm
    display: flex
    flex-direction: column
    gap: 12px
    padding: 16px 24px
    border-radius: 12px
    div.errors
        display: flex
        flex-direction: column
        gap: 8px
        div.error
            display: flex
            gap: 8px
            align-items: center
    button
        padding: 12px
        border: none
        border-radius: 8px
        background: var(--color-slate-200)
        color: white
        font-size: 16px
        font-weight: 500
        cursor: pointer
        &.disabled
            background: var(--color-slate-100)
            color: var(--color-slate-200)
            cursor: not-allowed
</style>
