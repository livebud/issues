<script>
  export let post = {}
</script>

<h1>Showing Post</h1>

<table border="1" cellpadding="10">
  <thead>
    {#each Object.keys(post) as key}
      <th>{key}</th>
    {/each}
  </thead>
  <tr>
    {#each Object.keys(post) as key}
      <td>{post[key]}</td>
    {/each}
  </tr>
</table>

<br />

<a href="..">Back</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
