<script>
 import { onMount } from 'svelte';
 import Card from './lib/Card.svelte'
 import FormCard from './lib/FormCard.svelte'
 import { isFetching, getAllGarments } from './store'

 let promise = getAllGarments();
 let garments = [];
 let isHidden = true;
 let orphans = []

 onMount(async () => {
   const res = await promise;
   const newGarments = res.data.filter((g) => !g.is_orphan)
   const newOrphans = res.data.filter((g) => g.is_orphan)
   garments = [...newGarments]
   orphans = [...newOrphans]
 })

</script>

<main>
    <h1>haberdashery</h1>
    {#if orphans.length > 0}
    <button on:click="{() => isHidden = !isHidden}">
        <span>Lägg till </span>+
    </button>
    {/if}
    <div class="cards">
    <FormCard {isHidden} {orphans} />
      {#each garments as garment}
        <Card {garment} />
      {/each}
    </div>
</main>

<style>

 h1 {
     font-size: 5rem;
     font-family: 'Bebas Neue', cursive;
     color: rgb(var(--lemon));
 }

 .cards {
     display: flex;
     flex-direction: row;
     flex-wrap: wrap;
     margin: 7rem auto;
     justify-content: space-between;
     gap: 3rem;
 }

 .hidden {
     display: none;
 }

 button {
     position: absolute;
     right: 0px;
     top: 0px;
     margin: 15px;
     vertical-align: top;
     background: var(--lemon-gradient);
     transition: all .2s ease-in-out;
     max-width: 3.5rem;
     border-radius: 35px;
     border: none;

     box-shadow: 0 0 0 0 rgba(var(--lemon), 1);
     transform: scale(1);
     animation: pulse 2s infinite;
 }

 button:hover {
     border-radius: 15px;
     max-width: 11rem;
     border: none;
     animation: lowPulse 2s infinite;
 }

 button span {
     max-width: 0;
     transition: max-width .2s ease-in-out;
     display: inline-block;
     vertical-align: top;
     white-space: pre;
     overflow: hidden;
 }

 button:hover span {
     max-width: 5rem;
 }

 @keyframes lowPulse {
     0% {
         transform: scale(0.95);
         box-shadow: 0 0 0 0 rgba(var(--lemon), 0.7);
     }

     70% {
         transform: scale(1);
         box-shadow: 0 0 0 10px rgba(var(--lemon), 0);
     }

     100% {
         transform: scale(0.95);
         box-shadow: 0 0 0 0 rgba(var(--lemon), 0);
     }
 }

 @keyframes pulse {
     0% {
         transform: scale(0.85);
         box-shadow: 0 0 0 0 rgba(var(--lemon), 0.7);
     }

     70% {
         transform: scale(1);
         box-shadow: 0 0 0 20px rgba(var(--lemon), 0);
     }

     100% {
         transform: scale(0.85);
         box-shadow: 0 0 0 0 rgba(var(--lemon), 0);
     }
 }

</style>
