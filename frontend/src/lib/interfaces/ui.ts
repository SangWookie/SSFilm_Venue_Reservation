export interface NavbarItem {
    label: string;
    href: string;
    checkEnabled: (url: string) => boolean;
}

export interface SelectableItem {
    key: unknown;
    label: string;
    toggle: boolean;
    disabled?: boolean;
}
