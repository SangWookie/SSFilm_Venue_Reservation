<script module lang="ts">
    import { defineMeta } from '@storybook/addon-svelte-csf';
    import CollapsibleBlock from './collapsible-block.svelte';

    const { Story } = defineMeta({
        title: 'UI/Form/Collapsible Block',
        component: CollapsibleBlock,
        args: {
            open: false
        },
        argTypes: {
            open: {
                control: 'boolean'
            }
        }
    });
</script>

<!-- FIXME: This results nested component building -->
<Story name="default" args={{ open: true }}>
    <div style="padding: 16px;">
        <CollapsibleBlock open={true}>
        {#snippet header()}
            Hello
        {/snippet}
        Content
    </CollapsibleBlock>
    </div>
</Story>
