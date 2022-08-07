<script>
 import { onMount } from 'svelte';
 import Card from './lib/Card.svelte'
 import { isFetching, getAllGarments } from './store'

 let promise = getAllGarments();
 let garments = [];

 onMount(async () => {
   const newGarments = await promise;
   garments = [...newGarments.data]
 })

</script>

<main>
      {#each garments as garment}
        <Card {garment} />
      {/each}
</main>

<style>

 :global(:root) {
   --red: #e27d60;
   --cyan: #85DCB;
   --terracotta: #e8a87c;
   --dark-pink: #c38d9e;
   --teal: #41b3a3;
 }

 body {

   background: var(--cyan);
 }

 main {
   display: flex;
   flex-direction: row;
   flex-wrap: wrap;
 }
</style>
