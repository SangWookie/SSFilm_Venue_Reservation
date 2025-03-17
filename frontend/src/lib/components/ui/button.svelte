<script lang="ts">
    import type { Snippet } from 'svelte';
    import LoadingIcon from './loading_icon.svelte';
    export interface Props {
        children?: Snippet;
        state?: 'enabled' | 'disabled' | 'loading';
        onClick?: () => void;
    }

    let props: Props = $props();
</script>

<button class={['ui-button', props.state ?? 'enabled']} onclick={props.onClick}>
    <div class="button-loading">
        <LoadingIcon size={24} />
    </div>
    {@render props.children?.()}
</button>

<style lang="sass">
button.ui-button
    // should no gaps
    display: flex
    align-items: center
    justify-content: center
    padding: 12px 16px
    border: none
    border-radius: 12px
    color: white
    font-size: 20px
    font-weight: 500
    &.enabled
        background-color: var(--color-main-600)
        &:hover
            background-color: var(--color-main-500)
        &:active
            background-color: var(--color-main-700)
    &.disabled
        background-color: var(--color-slate-600)
        cursor: not-allowed
    .button-loading
        display: block
        opacity: 0
        width: 0
        margin-right: 0
        transition: opacity 0.3s, width 0.3s, marign-right 0.3s
    &.loading
        background-color: var(--color-main-800)
        cursor: wait
        .button-loading
            width: 24px
            opacity: 1
            margin-right: 12px
</style>
