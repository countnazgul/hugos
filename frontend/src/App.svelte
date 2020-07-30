<script>
import Header from "./partials/header.svelte";
  import Cta from "./partials/cta.svelte";
  import Features from "./partials/features.svelte";
  import Stats from "./partials/stats.svelte";
  import Pricing from "./partials/pricing.svelte";

  import { onMount } from 'svelte';

  export let name;

  let message = '';

  function getMessage() {
    window.backend.basic().then(result => {
     message = result
    })
  }

  onMount(() => {
   getMessage();
  });
</script>

<main class="bg-gray-50 flex flex-col mx-auto">
<Header {message} {name}/>
<Features/>
<Cta/>
<Stats/>
<Pricing/>
</main>

<svelte:head>
<link rel="stylesheet" href="/fonts/inter.css" />
</svelte:head>
