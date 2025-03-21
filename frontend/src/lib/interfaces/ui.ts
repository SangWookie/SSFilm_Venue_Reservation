import type { Snippet } from 'svelte';

export interface NavbarItem {
    label: string;
    href: string;
    checkEnabled: (url: string) => boolean;
}

export interface FormSelectItem<T> {
    value: T;

    /// The option value in select element.
    key: string;
    label: string;
    labelSnippet?: Snippet;
}
