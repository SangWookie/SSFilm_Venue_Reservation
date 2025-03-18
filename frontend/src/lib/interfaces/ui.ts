import type { Snippet } from "svelte";

export interface NavbarItem {
    label: string;
    href: string;
    checkEnabled: (url: string) => boolean;
}

export interface SelectableItem<T> {
    value: T;
    label: string;
    // toggle: boolean;
    // use selected props on selectable-list instead.
    disabled?: boolean;
}

export interface FormSelectItem<T> {
    value: T;
    
    /// The option value in select element.
    key: string;
    label: string;
    labelSnippet?: Snippet;
}