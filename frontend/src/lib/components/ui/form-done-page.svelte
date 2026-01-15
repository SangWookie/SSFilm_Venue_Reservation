<script lang="ts">
    import { CheckCircle } from '@lucide/svelte';

    import { page } from '$app/state';
    import { onMount } from 'svelte';

    import { venues } from '$lib/mock.const.ts';

    let extra_information_html: string = $state('');
    let is_manual: boolean = $state(false);
    onMount(() => {
        if (page.url.searchParams.has('venue_name')) {
            extra_information_html =
                venues.find((i) => i.venue == page.url.searchParams.get('venue_name'))
                    ?.requirement || '';
            is_manual = page.url.searchParams.get('manual') == 'true';
        }
    });
</script>

<div class="form-done-page">
    <div class="icon">
        <CheckCircle size={48} />
    </div>

    <h2>예약 신청이 완료되었습니다.</h2>

    {#if !is_manual}
         <p>해당 공간은 신청 즉시 예약이 확정됩니다. 메인 화면에서 캘린더로 예약 정보를 확인해주세요.</p>
    {:else}
          <p>관리자가 검토 후 승인 예정입니다. 승인 시 이메일로 확정 메일이 전송됩니다.</p>
    {/if}

    {#if extra_information_html}
        <div>
            <!-- eslint-disable-next-line svelte/no-at-html-tags -->
            {@html extra_information_html}
        </div>
    {/if}
</div>

<style lang="sass">
.form-done-page
    display: flex
    flex-direction: column
    align-items: center
    gap: 12px

    .icon
        color: var(--color-green-500)
        display: flex
        justify-content: center
        align-items: center
        border-radius: 50%
        background: var(--color-green-100)
        width: 64px
        height: 64px

    h2
        font-size: 24px
        font-weight: 500

    p
        text-align: center
        opacity: .6
        width: 100%
</style>
