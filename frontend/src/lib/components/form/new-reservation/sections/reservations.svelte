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

    <InputBox title="장소 선택" description={form_data.reservations.venue}>
        {#snippet custom()}
            <SelectableList data={data.venue_selectable} isRadio={true}>
                {#snippet labelSnippet(item: SelectableItem<Venue>)}
                    {item.label}

                    {#if item.value?.approval_mode == 'manual'}
                        (수동)
                    {/if}
                {/snippet}
            </SelectableList>

            {console.log(data.venue_selectable.selected)}
            {#if data.current_venue?.value?.requirement}
                <!-- eslint-disable-next-line svelte/no-at-html-tags -->
                {@html data.current_venue?.value?.requirement}
            {/if}
        {/snippet}
        <ValidateMessage isValid={validations.reservations.venue} message="장소를 선택해 주세요." />
    </InputBox>

    <InputBox title="일자 선택" description={form_data.reservations.date}>
        {#snippet custom()}
            <div class="calendar-wrapper">
                <Calendar
                    items={data.calendar}
                    selected={data.calendar_selected}
                    status="available"
                    onDateClick={(date) => {
                        if (date) data.calendar_selected = [date];
                    }}
                />
            </div>
        {/snippet}
        <ValidateMessage
            isValid={validations.reservations.date.not_deadline}
            message="데드라인 (TODO 추가; 전날 18시까지)에 맞춰야 합니다."
        />
        <ValidateMessage
            isValid={validations.reservations.date.not_past}
            message="과거에 대해 예약할 수 없습니다."
        />
    </InputBox>

    <LoadingBox enabled={data.loading_reservation} />
    <InputBox
        title="시간 선택"
        description={form_data.reservations.hours.map((i) => `${i}시`).join(', ')}
    >
        {#snippet custom()}
            <SelectableList
                data={data.hour_selectable}
                disabled={form_data.reservations.date.length === 0}
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
