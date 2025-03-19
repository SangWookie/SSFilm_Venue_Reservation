<script lang="ts">
    import './../../app.sass';
    import Navbar from '$lib/components/ui/navbar.svelte';
    import { onNavigate } from '$app/navigation';
    import { globalAppState } from '$lib/store.svelte';

    let { children } = $props();
    const navbar_props = {
        items: [
            { label: '예약 현황', href: '/' },
            { label: '예약하기', href: '/new' }
            //{ label: '내 예약', href: '/my' }
        ].map((item) => {
            return {
                ...item,
                checkEnabled: (path: string) => path == item.href // FIXME: would be better to use state instead.
            };
        })
    };

    onNavigate((navigation) => {
        if (!document.startViewTransition) return;

        return new Promise((resolve) => {
            document.startViewTransition(async () => {
                resolve();
                await navigation.complete;
            });
        });
    });
</script>

<Navbar {...navbar_props} />

<div>
    {@render children()}
</div>

<style lang="sass">
div
    display: flex
    flex-direction: column
</style>
