<script lang="ts">
    import { CheckCircle } from '@lucide/svelte';

    import { page } from '$app/state';
    import { onMount } from 'svelte';

    import { venues } from '$lib/mock.const.ts';

    let extra_information_html: string = $state('');
    onMount(() => {
        if (page.url.searchParams.has('venue_name')) {
            extra_information_html =
                venues.find((i) => i.venue == page.url.searchParams.get('venue_name'))?.requirement || '';
        }
    });
</script>

<div class="form-done-page">
    <div class="icon">
        <CheckCircle size={48} />
    </div>

    <h2>예약 신청이 완료되었습니다.</h2>

    <p>이메일로 예약 결과를 확인하거나, 이후 캘린더에서 확인이 가능합니다.</p>

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
