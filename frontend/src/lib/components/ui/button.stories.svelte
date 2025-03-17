<script module lang="ts">
    import { defineMeta, setTemplate, type Args } from '@storybook/addon-svelte-csf';
    import Button from './button.svelte';
    import { fn } from '@storybook/test';

    const { Story } = defineMeta({
        title: 'UI/Button',
        component: Button,
        args: {
            state: 'enabled',
            onClick: fn()
        },
        argTypes: {
            state: {
                control: 'radio',
                options: ['enabled', 'disabled', 'loading']
            }
        }
    });
</script>

<script lang="ts">
    setTemplate(template);
</script>

{#snippet template(args: Args<typeof Story>)}
    <Button {...args}>Click ME</Button>
{/snippet}

<Story name="default" />
<Story name="disabled" args={{ state: 'disabled' }} />
<Story name="loading" args={{ state: 'loading' }} />

<Story name="loading (interactive)">
    {#snippet children(args)}
        <Button
            {...args}
            onClick={fn(() => {
                args.state = 'loading';
                setTimeout(() => {
                    args.state = 'enabled';
                }, 2000);
            })}
        >
            Click ME
        </Button>
    {/snippet}
</Story>
