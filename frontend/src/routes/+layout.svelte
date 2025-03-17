<script lang="ts">
    import './../app.sass';
    import Navbar from '$lib/components/ui/navbar.svelte';
    import Announcement from '$lib/components/ui/announcement.svelte';

    let { children } = $props();
    const announcement_props = {
        text: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.'
    };
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
</script>

<Announcement {...announcement_props} hidden={true} />
<Navbar {...navbar_props} />

<div>
    {@render children()}
</div>

<style lang="sass">
div
    display: flex
    flex-direction: column
</style>
