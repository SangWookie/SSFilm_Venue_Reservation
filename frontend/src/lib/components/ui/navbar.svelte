<script lang="ts">
    import type { NavbarItem } from '$lib/interfaces/ui';
    import { page } from '$app/state';
    import { Menu, X } from '@lucide/svelte';
    import { afterNavigate } from '$app/navigation';

    interface Props {
        items: NavbarItem[];
    }
    let props: Props = $props();

    const nav_state = $derived.by(() => {
        return props.items.map((item) => {
            return {
                text: item.label,
                href: item.href,
                enabled: item.checkEnabled(page.url.pathname)
            };
        });
    });
    const current = $derived(nav_state.find((item) => item.enabled));

    let isOepn = $state(false);

    const openToggleCallback = () => {
        isOepn = !isOepn;
    };
    afterNavigate(() => {
        isOepn = false;
    });
</script>

{#snippet renderItems()}
    {#each nav_state as item (item.text)}
        <li class="navbar-item" class:enabled={item.enabled}>
            <a class="navbar-link" href={item.href} title={item.text}>
                {item.text}
            </a>
        </li>
    {/each}
{/snippet}

<nav>
    <ul class="navbar-list">
        {@render renderItems()}
    </ul>
    <ul class="navbar-list-small">
        <button onclick={openToggleCallback}>
            <Menu />
        </button>
        {current?.text}
    </ul>
</nav>

<div class="navbar-list-small-opened" class:open={isOepn}>
    <button onclick={openToggleCallback}>
        <X />
    </button>
    <ul>
        {@render renderItems()}
    </ul>
</div>

<style lang="sass">
nav
    display: flex
    position: sticky
    top: 0
    background: rgba(255, 255, 255, .8)
    backdrop-filter: blur(12px)
    z-index: 99999

    ul.navbar-list
        padding: 16px 24px
        li.navbar-item
            // remove dot
            list-style-type: none
            font-size: 20px
            font-weight: 700
            border-radius: 9999px
            a.navbar-link
                display: block
                padding: 8px 16px
                text-decoration: none
                color: var(--color-text)
            &.enabled
                background-color: var(--color-bg)
                a.navbar-link
                    color: white
            &:hover
                background-color: var(--color-hover)
                a.navbar-link
                    color: white
            &:active
                background-color: var(--color-active)
                a.navbar-link
                    color: white

    ul.navbar-list-small
        padding: 16px 24px
        margin: 0
        gap: 12px
        align-items: center
        font-size: 20px
        font-weight: 500
        button
            display: flex
            //background: var(--color-slate-50)
            background: white
            border: none
            border-radius: 12px
            padding: 8px
            &:hover
                background: var(--color-slate-100)
            &:active
                background: var(--color-slate-200)

div.navbar-list-small-opened
    position: fixed
    top: 0
    z-index: 999999
    box-sizing: border-box
    width: 100%
    min-height: 100%
    background: rgba(255, 255, 255, .9)
    backdrop-filter: blur(12px)
    padding: 16px 24px

    display: none
    flex-direction: column
    gap: 64px

    &.open
        display: flex

    button
        display: flex
        background: var(--color-slate-50)
        border: none
        border-radius: 999px
        padding: 20px 24px
        &:hover
            background: var(--color-slate-100)
        &:active
            background: var(--color-slate-200)

    ul
        display: flex
        flex-direction: column
        padding: 0
        margin: 0
        li.navbar-item
            list-style-type: none
            font-size: 20px
            font-weight: 700
            border-radius: 9999px
            a.navbar-link
                padding: 20px 24px
                display: block
                width: 100%
                height: 100%
                text-decoration: none
                color: var(--color-text)
            &.enabled
                background-color: var(--color-bg)
                a.navbar-link
                    color: white
            &:hover
                background-color: var(--color-hover)
                a.navbar-link
                    color: white
            &:active
                background-color: var(--color-active)
                a.navbar-link
                    color: white
@media (min-width: 600px)
    nav
        ul.navbar-list
            display: flex
        ul.navbar-list-small
            display: none
    div.navbar-list-small-opened
        display: none

@media (max-width: 600px)
    nav
        ul.navbar-list
            display: none
        ul.navbar-list-small
            display: flex
    //div.navbar-list-small-opened
    //    display: flex
</style>
