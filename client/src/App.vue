<template>
  <div class="container">
    <router-view/>
  </div>
</template>

<style lang="scss">
html {
  background-color: #050430;
  height: 100%;
  width: 100%;
}

body {
  padding: 0px;
  margin: 0px;
  height: 100%;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #ffffff;
  height: 100%;
}

.container {
  display: flex;
  margin-left: auto;
  margin-right: auto;
  max-width: 1800px;
  height: 100%;
}

.center-content {
    display: flex;
    flex-wrap: wrap;
    margin: auto;
    height: 100%;
    width: 100%;
    justify-content: center;
    align-items: center;
  }
</style>
