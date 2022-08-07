import { writable } from 'svelte/store';

export const garments = writable([]);
export const isFetching = writable(false);

export const getAllGarments = () => {
  isFetching.set(true);
  return fetch('http://localhost:3000/v1/garments').then((x) => {
    isFetching.set(false);
    return x.json();
  });
}
