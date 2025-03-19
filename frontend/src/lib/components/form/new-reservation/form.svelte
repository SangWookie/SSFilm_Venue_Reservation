<script lang="ts">
    import { untrack } from 'svelte';
    import {
        feedVenueData,
        type FormData,
        type FormProps,
        init_form_data,
        init_internal_states,
        type InternalStates,
        isAllValidated,
        validate,
        type Validations
    } from '.';

    import RequesterInfoSection from './sections/requester-info.svelte';
    import ReservationsSection from './sections/reservations.svelte';
    import Button from '$lib/components/ui/button.svelte';
    import { globalAppState } from '$lib/store.svelte.ts';
    //import AgreementSection from './sections/agreement.svelte'

    import { getReservations } from '$lib/api/mock';
    import { requestNewReservationFromData } from '.';
    import { CheckIcon } from '@lucide/svelte';

    let form_data: FormData = $state(init_form_data());
    let internal_states: InternalStates = $state(init_internal_states());
    let form_props: FormProps = $state({
        calendar: [],
        purposes: [],
        getReservations
    });
    let validations: Validations = $derived(validate(form_data, internal_states));

    globalAppState.subscribe((app) => {
        if (!app) return;
        if (app.venues) feedVenueData(app.venues, internal_states);
        if (app.purposes) form_props.purposes = app.purposes;
    });

    let submissionState: 'unavailable' | 'available' | 'waiting' | 'done' = $state('unavailable');
    let errorMessage: string | undefined = $state(undefined);

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

        internal_states.collapsed.requester_info = false;
        internal_states.collapsed.reservations = false;
        internal_states.collapsed.agreement = false;

        requestNewReservationFromData(form_data)
            .then((response) => {
                if (response.success) submissionState = 'done';
                else {
                    submissionState = 'available';
                    errorMessage = response.message;
                }
            })
            .catch((e) => {
                console.error('Failed to request form', e);
                errorMessage = `서버와 통신 중 오류가 발생했습니다. 다시 요청이 가능합니다. ${e.message}`;
                alert(errorMessage);
                submissionState = 'available';
            });
    };
</script>

<RequesterInfoSection bind:form_data {validations} bind:internal_states />
<ReservationsSection bind:form_data {validations} bind:internal_states bind:form_props />

{errorMessage ?? ''}
<Button state={button_state} onClick={button_click_handler}>
    {#if submissionState != 'done'}
        예약하기
    {:else}
        <CheckIcon /> &nbsp; 예약 요청 완료. 달력을 확인해주세요!
    {/if}
</Button>
