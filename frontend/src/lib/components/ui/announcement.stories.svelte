<script module lang="ts">
    import { defineMeta } from '@storybook/addon-svelte-csf';
    import Announcement from './announcement.svelte';

    const { Story } = defineMeta({
        title: 'UI/Announcement',
        component: Announcement,
        args: {
            hidden: false,
            text: '공지사항을 입력해주세요.'
        },
        argTypes: {
            hidden: {
                control: 'boolean'
            }
        }
    });
</script>

<Story name="default" />
<Story name="Long text" args={{ text: '안녕하세요'.repeat(100) }} />
