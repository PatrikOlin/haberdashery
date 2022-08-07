<script>
 import placeholder from '../assets/placeholder.png'

 export let garment;

 const fmtDate = (datetime) => {
    const date = new Date(datetime)
    const formatter = new Intl.DateTimeFormat('sv-SE');
    return formatter.format(date)
 }

 $: console.log(garment)
</script>

<div class="container">
    <span class="overlay"></span>
    <img src={placeholder} />
    <div class="garment-info">
        <div class="row">
            <span class="segment">
                <h3>ID</h3>
                <p>{garment.id}</p>
            </span>
            <span class="segment">
                <h3>Använd</h3>
                <p>{garment.times_worn} ggr</p>
            </span>
            <span class="segment">
                <h3>Tvättad</h3>
                <p>{garment.times_washed} ggr</p>
            </span>
        </div>
        <div class="row">
            <span class="segment">
                <h3>Senast använd</h3>
                <p>{fmtDate(garment.last_worn)}</p>
            </span>
            <span class="segment">
                <h3>Senast tvättad</h3>
                <p>{fmtDate(garment.last_washed)}</p>
            </span>
        </div>
        <div class="row">
            <span class="segment">
                <h3>Märke</h3>
                <p>{garment.brand}</p>
            </span>
            <span class="segment">
                <h3>Pris</h3>
                <p>{garment.price}</p>
            </span>
            <span class="segment">
                <h3>Inköpt</h3>
                <p>{fmtDate(garment.purchased_at)}</p>
            </span>
        </div>
    </div>
</div>

<style>
 img {
     height: 15rem;
     padding: .75rem;
     position: absolute;
     top: 5px;
     left: 0;
     margin-left: auto;
     right: 0;
     margin-right: auto;
     transition: all .2s ease-in-out;
 }

 .overlay {
     height: 15rem;
     padding: .75rem;
     position: absolute;
     top: 0;
     left: 0;
     right: 0;
     background: var(--teal);
     opacity: 0%;
     transition: all .2s ease-in-out;
     z-index:2;
 }

 .container {
     position: relative;
     display: flex;
     flex-direction: column-reverse;
     justify-content: space-between;
     background: #ffffff;
     width: 18rem;
     min-width: 15rem;
     height: 27rem;
     margin: 20px 15px;
     border-radius: 5px;
     overflow: hidden;
     box-shadow: 0px 4px 20px -4px rgba(0,0,0,.35);
     z-index: 100;
 }

 .garment-info {
     display:flex;
     flex-direction: column;
     color: #000;
     background: var(--lemon);
     height: 8.5rem;
     margin-top: 2rem;
     padding: 0px 20px;
     z-index: 2;
     transition: all .2s ease-in-out;
     overflow: hidden;
 }

 .container:hover .garment-info {
     height: 13rem;
 }

 .container:hover .overlay {
     opacity: .75;
 }

 .row:first-child {
     margin-top: 1rem;
 }

 .row {
     display: flex;
     flex-direction: row;
     justify-content: space-between;
     width: 100%;
     padding: 0px;
     margin: 5px 0px;
 }

 .segment {
     text-align: left;
     color: var(--navy);
 }

 .segment h3 {
     font-size: 1rem;
     margin: 0px;
 }

 .segment p {
     font-size: 1rem;
     margin: 0px;
 }
</style>
