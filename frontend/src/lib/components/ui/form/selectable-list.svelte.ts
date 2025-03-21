export interface SelectableItem<T> {
    value: T;
    key: string;
    label: string;
    disabled?: boolean;
}

export class SelectableListState<T> {
    list: SelectableItem<T>[] = $state([]);
    selected: SelectableItem<T>[] = $state([]);

    constructor(list?: SelectableItem<T>[]) {
        if (list) this.list = list;
    }

    select(item: SelectableItem<T>) {
        this.selected = this.list.filter((i) => i == item || this.selected.includes(i));
    }

    unselect(item: SelectableItem<T>) {
        this.selected = this.selected.filter((i) => i != item);
    }

    toggle(item: SelectableItem<T>) {
        if (this.isSelected(item)) {
            this.unselect(item);
        } else {
            this.select(item);
        }
    }

    getItemWithKey(key: string) {
        return this.list.find((i) => i.key === key);
    }

    isSelected(item: SelectableItem<T>) {
        return this.selected.includes(item);
    }
}
