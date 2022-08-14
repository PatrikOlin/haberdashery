<script>
 import { isFetching, submitGarment } from '../store'

 export let isHidden = true;
 export let orphans = []
 let selectedGarment = {};

 const handleSubmit = () => {
     let purchaseDate = new Date(selectedGarment.purchased_at).toISOString()
     let offset = new Date(selectedGarment.purchased_at).getTimezoneOffset()
     let sign = offset > 0 ? '+' : '-'
     purchaseDate.replace('Z', `${sign}0${offset}:00`)
     selectedGarment.purchased_at = purchaseDate
     selectedGarment.price = +selectedGarment.price
     console.log(selectedGarment.purchased_at)
     submitGarment(selectedGarment)
 }
</script>
        <div class="{ isHidden ? 'formCard hidden' : 'formCard' }">
            <form on:submit|preventDefault={handleSubmit}>
                <label>ID
                    <select bind:value={selectedGarment}>
                        {#each orphans as orphan}
                            <option value={orphan}>
                                {orphan.id}
                            </option>
                        {/each}
                    </select>
                </label>
                <label>Färg
                    <input type="text" bind:value={selectedGarment.color}/>
                </label>
                <label>Märke
                    <input type="text" bind:value={selectedGarment.brand}/>
                </label>
                <label>Köpdatum
                    <input type="date" bind:value={selectedGarment.purchased_at}/>
                </label>
                <label>Pris
                    <input type="text" bind:value={selectedGarment.price}/>
                </label>
                <button type="submit">Spara</button>
            </form>
        </div>

<style>
 .formCard {
     display: flex;
     background: #ffffff;
     width: 18rem;
     min-width: 15rem;
     height: 27rem;
     border-radius: 5px;
     overflow: hidden;
     box-shadow: 0px 4px 20px -4px rgba(0,0,0,.35);
     margin: auto;
 }

 .hidden {
     display: none;
 }

 form {
     padding: 1rem 2rem;
     width: 100%;
 }

 label {
     display: flex;
     flex-direction: column;
     color: #000;
     align-items: start;
     width: 100%;
     margin-bottom: .5rem;
 }

 input, select {
     width: 100%;
     padding: .45rem;
     padding-right: 0;
     box-sizing: border-box;
 }

 button {
     background: var(--lemon-gradient);
     width: 100%;
     margin-top: .5rem;
 }
</style>
