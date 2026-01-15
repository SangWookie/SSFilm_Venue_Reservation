<script module lang="ts">
    import { defineMeta, type Args } from '@storybook/addon-svelte-csf';
    import Button, { type Props } from './button.svelte';

    const { Story } = defineMeta({
        title: 'UI/Button',
        component: Button,
        args: {
            state: 'enabled',
        },
        render: template,
        argTypes: {
            state: {
                control: 'radio',
                options: ['enabled', 'disabled', 'loading']
            }
        }
    });
</script>

{#snippet template(args: Props | undefined)}
    <Button {...args}>Click ME</Button>
{/snippet}

<Story name="default" />
<Story name="disabled" args={{ state: 'disabled' }} />
<Story name="loading" args={{ state: 'loading' }} />

