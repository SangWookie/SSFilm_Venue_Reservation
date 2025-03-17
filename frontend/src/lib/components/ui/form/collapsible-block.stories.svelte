<script module lang="ts">
    import { defineMeta } from '@storybook/addon-svelte-csf';
    import CollapsibleBlock from './collapsible-block.svelte';
    import type { ComponentProps } from 'svelte';

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

<Story name="default" args={{ open: true }}>
    {#snippet children(args)}
        <!-- this is dumb idea to converting the types... -->
        <CollapsibleBlock {...args as ComponentProps<typeof CollapsibleBlock>}>
            {#snippet header()}
                Hello
            {/snippet}
            Content
        </CollapsibleBlock>
    {/snippet}
</Story>
