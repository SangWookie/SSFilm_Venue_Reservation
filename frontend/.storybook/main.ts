import type { StorybookConfig } from '@storybook/sveltekit';

const config: StorybookConfig = {
    stories: ['../src/**/*.mdx', '../src/**/*.stories.@(js|ts|svelte)'],
    staticDirs: ['../static'],
    addons: [
        {
            name: '@storybook/addon-svelte-csf',
            options: {
                legacyTemplate: true
            },
        },
        '@chromatic-com/storybook',
    ],
    framework: {
        name: '@storybook/sveltekit',
        options: {}
    },
    core: {
        disableTelemetry: true,
        disableWhatsNewNotifications: true
    }
};
export default config;
