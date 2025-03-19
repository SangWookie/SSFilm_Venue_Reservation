<script module lang="ts">
    import { defineMeta } from '@storybook/addon-svelte-csf';
    import SelectableList from './selectable-list.svelte';
    import type { SelectableItem } from '$lib/interfaces/ui';

    const { Story } = defineMeta({
        title: 'UI/Form/Selectable List',
        component: SelectableList
    });

    const list: SelectableItem<string>[] = Array(10)
        .keys()
        .map((i) => {
            return {
                value: `${i}`,
                key: `${i}`,
                label: `#${i} Item`
            };
        })
        .toArray();
</script>

<Story name="default" args={{ list }}></Story>
