export interface SelectableItem<T> {
    value: T;
    key: string;
    label: string;
    disabled?: boolean;
}

export function createSelectableList<T>() {
    let list: SelectableItem<T>[] = $state([]);
    let selected: SelectableItem<T>[] = $state([]);
    let isRadio = $state(false);

    const select = (item?: SelectableItem<T>) => {
        if (!item) return;
        selected = list.filter((i) => selected.includes(i) || i === item);
    }

    const unselect = (item?: SelectableItem<T>) => {
        if (!item) return;
        selected = selected.filter((i) => i !== item);
    }

    const toggle = (item?: SelectableItem<T>) => {
        if (!item) return;
        if (selected.includes(item)) {
            unselect(item);
        } else {
            select(item);
        }
    }

    const getItemWithKey = (key: string) => list.find((i) => i.key === key);

    return {
        get list() { return list },
        set list(value: SelectableItem<T>[]) { list = value },
        get selected() { return selected },
        set selected(value: SelectableItem<T>[]) { selected = value },
        
        select, unselect, toggle,

        selectWithKey(key: string) { select(getItemWithKey(key)) },
        unselectWithKey(key: string) { unselect(getItemWithKey(key)) },
        toggleWithKey(key: string) { toggle(getItemWithKey(key)) },
        isSelected(item: SelectableItem<T>) { return selected.includes(item) },
    }
}