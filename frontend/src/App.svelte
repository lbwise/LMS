<script>
  export let clicks = 0
  export const increment = () => {
    clicks++
  }
</script>

<main>
  <h2>
    This is the title
  </h2>
  <p>This has been clicked { clicks } times</p>
  <button on:click={increment}>Increment</button>
</main>

<style>
  main {
    padding-top: 10vh;
    text-align: center;
  }

  h2 {
    font-size: 2.5em;
  }

  button {
    margin: 5vh 0;
    height: 5vh;
    width: 15%;
    background-color: rgb(113, 113, 207);
    font-size: 1.5em;
    font-weight: 700;
    color: white;
    border-radius: 10px;
  }
</style>
