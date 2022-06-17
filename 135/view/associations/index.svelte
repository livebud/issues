<script>
  export let associations = []
</script>

<h1>Association Index</h1>

<table border="1" cellpadding="10">
  {#if associations.length > 0}
    <thead>
      {#each Object.keys(associations[0]) as key}
        <th>{key}</th>
      {/each}
    </thead>
  {/if}
  {#each associations as association}
    <tr>
      {#each Object.keys(association) as key}
        {#if key.toLowerCase() === "id"}
          <td><a href={`/associations/${association[key]}`}>{association[key]}</a></td>
        {:else}
          <td>{association[key]}</td>
        {/if}
      {/each}
    </tr>
  {/each}
</table>

<style>
  table {
    border-collapse: collapse;
  }
</style>
