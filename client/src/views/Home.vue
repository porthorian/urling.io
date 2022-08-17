<template>
  <div class="center-content">
    <div class="box_container">
      <div class="item">
        <h1 class="box-header">Shorten that URL</h1>
      </div>
      <div class="item form_container">
        <div class="item shorten_url" v-if="shortenUrl != ''">
          <span v-on:click="copy">{{ textDisplayCopy }}</span>
        </div>
        <form class="item" v-on:submit.prevent="submitForm">
          <div class="item">
            <input type="text" v-model="longUrl" placeholder="Make me shorter please." />
          </div>
          <div class="item form_button">
            <button>
              <spinner class="spinner-center" v-if="showSpinner"/>
              <span v-else>Shorten</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component'
import axios from 'axios'
import useClipboard from 'vue-clipboard3'
import spinner from '../components/atom_spinner.vue'

const { toClipboard } = useClipboard()

@Options({
  components: {
    spinner
  }
})

export default class HomeView extends Vue {
  longUrl = ''
  shortenUrl = ''
  textDisplayCopy = ''
  showSpinner = false

  submitForm () {
    this.showSpinner = true
    // @ts-expect-error: this.apiUrl is already defined in the vue prototype
    axios.post(`${this.apiUrl}shorten`, {
      long_url: this.longUrl,
      user_id: Math.random().toString(36).slice(2, 15)
    })
      .then((res) => {
        const shortUrlId = res.data.shorturl_id
        this.shortenUrl = `${window.location.origin}/${shortUrlId}`
        this.textDisplayCopy = this.shortenUrl
      })
      .catch((error) => {
        console.log(error)
      })
      .finally(() => {
        this.showSpinner = false
      })
  }

  async copy () {
    try {
      this.textDisplayCopy = 'Text Copied'
      await toClipboard(this.shortenUrl)
      setTimeout(() => {
        this.textDisplayCopy = this.shortenUrl
      }, 500)
    } catch (e) {
      console.error(e)
    }
  }
}
</script>

<style lang="scss">
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

  .box-header {
    margin: 50px 0px;
  }

  .form_container {
    display: flex;
    height: 48%;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    position: relative;

    .shorten_url {
      position: absolute;
      top: 0px;

      span:hover {
        cursor: copy;
      }
    }

    form {
      width: 100%;
      min-width: 100%;
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

      button:hover {
        background-color: #7ba1a9;
      }
    }
  }

  .item {
    width: 100%;
  }

  .spinner-center {
    margin-left: auto;
    margin-right: auto;
  }
</style>
