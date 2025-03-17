<script module lang="ts">
    import { defineMeta } from '@storybook/addon-svelte-csf';
    import SelectableList from './selectable-list.svelte';
    import type { SelectableItem } from '$lib/interfaces/ui';

    const { Story } = defineMeta({
        title: 'UI/Form/Selectable List',
        component: SelectableList
    });

    const list: SelectableItem[] = Array(10)
        .keys()
        .map((i) => {
            return {
                label: `#${i} Item`,
                toggle: false
            };
        })
        .toArray();
    
</script>

<Story name="default" args={{ list }}></Story>
