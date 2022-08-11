import { writable } from 'svelte/store';

export const garments = writable([]);
export const isFetching = writable(false);

export const getAllGarments = () => {
  isFetching.set(true);
  return fetch('http://localhost:4040/v1/garments?includeOrphans=true').then((x) => {
    isFetching.set(false);
    return x.json();
  });
}

export const getAllOrphans = () => {
  isFetching.set(true);
  return fetch("http://localhost:4040/v1/garments?orphans=true").then((x) => {
    isFetching.set(false);
    return x.json();
  })
}
