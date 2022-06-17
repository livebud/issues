<script>
  export let association = {}
</script>

<h1>Showing Association</h1>

<table border="1" cellpadding="10">
  <thead>
    {#each Object.keys(association) as key}
      <th>{key}</th>
    {/each}
  </thead>
  <tr>
    {#each Object.keys(association) as key}
      <td>{association[key]}</td>
    {/each}
  </tr>
</table>

<br />

<a href="/associations">Back</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
