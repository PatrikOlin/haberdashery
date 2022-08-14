import { writable } from 'svelte/store';

export const garments = writable([]);
export const isFetching = writable(false);

const url = 'http://haberdashery.home.lan:4040/v1'

export const getAllGarments = () => {
  isFetching.set(true);
  return fetch(`${url}/garments?includeOrphans=true`).then((x) => {
    isFetching.set(false);
    return x.json();
  });
}
