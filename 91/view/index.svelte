<script>
  export let hellos = []
</script>

<h1>Hello Index</h1>

<table border="1" cellpadding="10">
  {#if hellos.length > 0}
    <thead>
      {#each Object.keys(hellos[0]) as key}
        <th>{key}</th>
      {/each}
    </thead>
  {/if}
  {#each hellos as hello}
    <tr>
      {#each Object.keys(hello) as key}
        {#if key === "id"}
          <td><a href={`/${hello[key]}`}>{hello[key]}</a></td>
        {:else}
          <td>{hello[key]}</td>
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
