<!--
<div class="hours-info">
    <div class="item">
        <div class="active">00</div>
        <div class="label">예약 중</div>
    </div>
    <div class="item">
        <div class="inactive">00</div>
        <div class="label">예약 가능</div>
    </div>
</div>

-->

<style lang="sass">
.hours-info
    display: flex
    justify-content: center
    gap: 32px

    .item
        display: flex
        align-items: center
        gap: 8px

        .active
            color: var(--color-bg)
            font-family: "Space Mono", monospace
            font-size: 16px
            font-weight: 600

        .inactive
            color: var(--color-bg)
            font-family: "Space Mono", monospace
            font-weight: 600
            font-size: 16px
            opacity: .2

        .label
            font-size: 16px
            opacity: .6
</style>
