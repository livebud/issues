<script>
  export let academy = {}
</script>

<h1>Showing Academy</h1>

<table border="1" cellpadding="10">
  <thead>
    {#each Object.keys(academy) as key}
      <th>{key}</th>
    {/each}
  </thead>
  <tr>
    {#each Object.keys(academy) as key}
      <td>{academy[key]}</td>
    {/each}
  </tr>
</table>

<br />

<a href="/associations/:association_id/academies">Back</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
