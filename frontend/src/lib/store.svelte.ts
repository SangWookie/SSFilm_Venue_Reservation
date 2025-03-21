import { getAppState } from '$lib/api/mock';
import type { AppState } from '$lib/interfaces/api';
import { writable } from 'svelte/store';
import { type Writable } from 'svelte/store';

export const globalAppState: Writable<AppState | undefined> = writable(undefined);

// Load it on startup
getAppState().then((state) => {
    globalAppState.set(state);
    console.log('Loaded global app state', state);
});
