<script>
  export let academies = []
</script>

<h1>Academy Index</h1>

<table border="1" cellpadding="10">
  {#if academies.length > 0}
    <thead>
      {#each Object.keys(academies[0]) as key}
        <th>{key}</th>
      {/each}
    </thead>
  {/if}
  {#each academies as academy}
    <tr>
      {#each Object.keys(academy) as key}
        {#if key.toLowerCase() === "id"}
          <td><a href={`/associations/:association_id/academies/${academy[key]}`}>{academy[key]}</a></td>
        {:else}
          <td>{academy[key]}</td>
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
