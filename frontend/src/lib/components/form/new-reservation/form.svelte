<script lang="ts">
    import { type FormData, init_form_data, isAllValidated, validate, type Validations } from '.';

    import { goto } from '$app/navigation';

    import RequesterInfoSection from './sections/requester-info.svelte';
    import ReservationsSection from './sections/reservations.svelte';
    import Button from '$lib/components/ui/button.svelte';
    import AgreementSection from './sections/agreement.svelte';

    import { requestNewReservationFromData } from '.';
    import { CheckIcon } from '@lucide/svelte';

    let form_data: FormData = $state(init_form_data());
    let validations: Validations = $derived(validate(form_data));

    let submissionState: 'unavailable' | 'available' | 'waiting' | 'done' = $state('unavailable');
    let errorMessage: string | undefined = $state(undefined);

    let temp_prop_venue = $state(false);

    let validation_state = $state(false);
    $effect(() => {
        validation_state = isAllValidated(validations);
    });
    $effect(() => {
        if (validation_state && submissionState === 'unavailable') submissionState = 'available';
        else if (!validation_state && submissionState === 'available')
            submissionState = 'unavailable';
    });

    const button_state = $derived.by(() => {
        if (submissionState == 'available') return 'enabled';
        if (submissionState == 'waiting') return 'loading';
        return 'disabled';
    });

    const button_click_handler = () => {
        if (submissionState !== 'available') return;
        submissionState = 'waiting';

        requester_info_collapsible_open = false;
        reservations_collapsible_open = false;
        agreement_collapsible_open = false;

        requestNewReservationFromData(form_data)
            .then((response) => {
                submissionState = 'done';
                console.log(response);
                goto(`/form_done?venue_name=${form_data.reservations.venue}&manual=${temp_prop_venue}`);
            })
            .catch((e) => {
                console.error('Failed to request form', e);
                errorMessage = `서버와 통신 중 오류가 발생했습니다. 다시 요청이 가능합니다. ${e.message}`;
                alert(errorMessage);
                submissionState = 'available';
            });
    };

    let requester_info_collapsible_open = $state(true);
    let reservations_collapsible_open = $state(true);
    let agreement_collapsible_open = $state(true);
</script>

<RequesterInfoSection
    bind:form_data
    {validations}
    bind:collapsible_open={requester_info_collapsible_open}
    />
    <ReservationsSection
    bind:form_data
    {validations}
    bind:collapsible_open={reservations_collapsible_open}
    bind:temp_prop_venue
/>

<AgreementSection bind:form_data {validations} bind:collapsible_open={agreement_collapsible_open} />

{errorMessage ?? ''}
<Button state={button_state} onClick={button_click_handler}>
    {#if submissionState != 'done'}
        예약하기
    {:else}
        <CheckIcon /> &nbsp; 예약 요청 완료. 달력을 확인해주세요!
    {/if}
</Button>
