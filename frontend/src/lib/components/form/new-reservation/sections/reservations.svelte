<script lang="ts">
    import Calendar from '$lib/components/ui/calendar.svelte';
    import type { Venue } from '$lib/interfaces/api';
    import CollapsibleBlock from '$lib/components/ui/form/collapsible-block.svelte';
    import LoadingBox from '$lib/components/ui/loading_box.svelte';
    import InputBox from '$lib/components/ui/form/input-box.svelte';
    import SelectableList from '$lib/components/ui/form/selectable-list.svelte';
    import ValidateMessage from '$lib/components/ui/form/validate-message.svelte';
    import Select from '$lib/components/ui/form/select.svelte';
    import { type FormData, type Validations } from '../index.ts';
    import { type SelectableItem } from '$lib/components/ui/form/selectable-list.svelte.ts';
    import { ReservationSectionFormState } from '../state.svelte.ts';
    const {
        form_data = $bindable(),
        validations = $bindable(),
        collapsible_open = $bindable(true)
    }: {
        form_data: FormData;
        validations: Validations;
        collapsible_open: boolean;
    } = $props();

    let data = new ReservationSectionFormState(form_data);
</script>

<CollapsibleBlock open={collapsible_open}>
    {#snippet header()}
        예약 선택
    {/snippet}

    <InputBox title="일자 선택" description={form_data.reservations.date}>
        {#snippet custom()}
            <span>최대 2주 이내의 일정까지 예약 가능합니다.</span>
            <div class="calendar-wrapper">
                <Calendar
                    items={data.calendar}
                    selected={data.calendar_selected}
                    status={data.loading_reservation ? 'loading' : 'available'}
                    onDateClick={(date) => {
                        if (date && !date.mark?.past && !data.loading_reservation) {
                            data.calendar_selected = [date];
                            data.fetchReservation();
                        }
                    }}
                />
            </div>
        {/snippet}
        <ValidateMessage
            isValid={validations.reservations.date.not_deadline}
            message="예약은 사용일 하루 전 17시까지 신청해야 합니다."
        />
        <ValidateMessage
            isValid={validations.reservations.date.not_past}
            message="과거에 대해 예약할 수 없습니다."
        />
    </InputBox>
    <LoadingBox enabled={data.loading_reservation} />

    <InputBox title="장소 선택" description={form_data.reservations.venue}>
        {#snippet custom()}
            <SelectableList data={data.venue_selectable} isRadio={true} disabled={data.loading_reservation || !(form_data.reservations.date.length > 0)}>
                {#snippet labelSnippet(item: SelectableItem<Venue>)}
                    {item.label}

                    {#if item.value?.approval_mode == 'manual'}
                        (수동)
                    {/if}
                {/snippet}
            </SelectableList>

            {#if data.current_venue?.value.approval_mode == 'manual'}
                <p style="color: red; font-weight: 600;">해당 공간의 예약은 관리자의 승인 후 확정됩니다. 승인 여부는 이메일로 알려드립니다.</p>
            {/if}
            {#if data.current_venue?.value?.requirement}
                <!-- eslint-disable-next-line svelte/no-at-html-tags -->
                {@html data.current_venue?.value?.requirement}
            {/if}
        {/snippet}

    </InputBox>

    <InputBox
        title="시간 선택"
        description={form_data.reservations.hours.map((i) => `${i}시`).join(', ')}
    >
        {#snippet custom()}
            <SelectableList
                data={data.hour_selectable}
                disabled={data.loading_reservation || form_data.reservations.date.length === 0 || form_data.reservations.venue.length === 0}
                clickHandler={data.hourSelectableClickCallback}
            />
        {/snippet}
        <ValidateMessage
            isValid={validations.reservations.hours.should_sequence}
            message="예약 시간은 연속으로 선택해야 합니다."
        />
        <ValidateMessage
            isValid={validations.reservations.hours.less_than_6hours}
            message="6시간 초과 예약이 불가능합니다."
        />
        <!--
        <ValidateMessage
            isValid={validations.reservations.is_free}
            message="비어있는 시간만 예약이 가능합니다."
        />
        --->
    </InputBox>

    <InputBox title="예약 목적" inputType="select" description={form_data.reservations.purpose}>
        {#snippet custom()}
            <Select
                options={data.purposes.map((i) => {
                    return {
                        value: i,
                        key: i,
                        label: i
                    };
                })}
                bind:value={data.category_selected}
            />
            <ValidateMessage
                isValid={validations.reservations.purpose}
                message="목적을 선택해 주세요."
            />
        {/snippet}
    </InputBox>

    <InputBox
        title="예약 상세 목적"
        placeholder="예약 상세 목적을 작성해 주세요."
        bind:value={form_data.reservations.purpose_detail}
    >
        <ValidateMessage
            isValid={validations.reservations.purpose_detail}
            message="상세 목적을 작성해 주세요."
        />
    </InputBox>

    <InputBox
        title="동행인 (선택사항)"
        placeholder="동행인이 있다면 작성해주세요."
        bind:value={form_data.reservations.companions}
    />
</CollapsibleBlock>

<style lang="sass">
.calendar-wrapper
    display: flex
    justify-content: center
</style>
