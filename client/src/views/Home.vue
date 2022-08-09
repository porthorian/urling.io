<template>
  <div class="home">
    <div class="box_container">
      <div class="item">
        <h1>Shorten Me Up Baby</h1>
      </div>
      <div v-if="shortenUrl != ''">
        {{ shortenUrl }}
      </div>
      <div class="item form_container">
        <form v-on:submit.prevent="submitForm">
          <div class="item">
            <input type="text" v-model="longUrl" placeholder="Make me shorter please." />
          </div>
          <div class="item form_button">
            <button>Shorten</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue } from 'vue-class-component'
import axios from 'axios'

export default class HomeView extends Vue {
  longUrl = ''
  shortenUrl = ''
  submitForm () {
    axios.post('https://api.urling.io/shorten', {
      long_url: this.longUrl,
      user_id: 'hello'
    })
      .then((res) => {
        console.log(res)
        this.shortenUrl = res.data.short_url
      })
      .catch((error) => {
        console.log(error)
      })
      .finally(() => {
        console.log('All done')
      })
  }
}
</script>

<style lang="scss">
  .home {
    display: flex;
    flex-wrap: wrap;
    margin: auto;
    height: 100%;
    width: 100%;
    justify-content: center;
    align-items: center;
  }

  .box_container {
    max-height: 400px;
    max-width: 800px;
    width: 100%;
    height: 100%;
    background-color: #002F5C;
    padding: 15px;
    border-radius: 10px;
    box-shadow: 0 0 3px 3px #005883;
  }

  .form_container {
    display: flex;
    height: 80%;
    justify-content: center;
    align-items: center;

    form {
      width: 100%;
    }

    .item {
      input {
        box-sizing: border-box;
        background-color: #050430;
        border-radius: 10px;
        padding: 20px;
        color: #ffffff;
        width: 80%;
        height: 100%;
      }
    }

    .form_button {
      margin-top: 20px;

      button {
        max-width: 200px;
        width: 100%;
        border-radius: 10px;
        padding: 10px;
      }
    }
  }

  .item {
    width: 100%;
  }
</style>
