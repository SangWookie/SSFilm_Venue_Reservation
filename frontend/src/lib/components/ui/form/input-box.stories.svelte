<script module lang="ts">
    import { defineMeta } from '@storybook/addon-svelte-csf';
    import InputBox from './input-box.svelte';

    const { Story } = defineMeta({
        title: 'UI/Form/Input Box',
        component: InputBox
    });
</script>

<Story
    name="default"
    args={{
        title: 'Title',
        description: 'Description',
        placeholder: 'Placeholder'
    }}
/>
