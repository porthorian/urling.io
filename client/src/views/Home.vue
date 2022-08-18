<template>
  <div class="center-content home">
    <div class="box_container">
      <div class="item">
        <h1 class="box-header">Shorten that URL</h1>
      </div>
      <div class="item form_container">
        <Transition name="fade">
          <div class="item shorten_url" v-if="shortenUrl != ''">
            <span v-on:click="copy">{{ textDisplayCopy }}</span>
          </div>
        </Transition>
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
        <button class="toggle_button" @click="showCode = !showCode">Toggle API Example</button>
      </div>
      <div class="item">
        <Transition name="fade">
          <div class="code_container" v-if="showCode">
            <pre>
#!/bin/bash

curl "{{ apiUrl() }}shorten" \
-H 'Content-Type: application/json;' \
-X POST --data-binary @- &lt;&lt; EOF
{
  "long_url": "https://www.youtube.com/watch?v=iik25wqIuFo",
  "user_id": "rick_rolled"
}
EOF
            </pre>
          </div>
        </Transition>
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

  showCode = false

  submitForm () {
    this.showSpinner = true
    // @ts-expect-error: this.apiUrl is already defined in the vue prototype
    axios.post(`${this.apiUrl()}shorten`, {
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
  .fade-enter-active,
  .fade-leave-active {
    transition: opacity 0.5s ease;
  }

  .fade-enter-from,
  .fade-leave-to {
    opacity: 0;
  }

  .home {
    max-height: 700px;
  }
  .box_container {
    max-width: 800px;
    width: 100%;
    background-color: #002F5C;
    padding: 15px 15px 20px 15px;
    border-radius: 10px;
    box-shadow: 0 0 3px 3px #005883;
  }

  .box-header {
    margin: 30px 0px;
  }

  .form_container {
    display: flex;
    height: 48%;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    position: relative;

    .shorten_url {
      margin: 30px 0px;

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
      margin: 30px 0px;

      button {
        max-width: 200px;
        width: 100%;
        border-radius: 10px;
        padding: 10px;

        &:hover {
          background-color: #7ba1a9;
        }
      }
    }

    .toggle_button {
      border-radius: 10px;
      padding: 10px;

      &:hover {
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

  .code_container {
    pre {
      background: #050430;
      box-shadow: 0 0 3px 3px #005883;
      page-break-inside: avoid;
      font-family: monospace;
      font-size: 15px;
      line-height: 1.6;
      margin-bottom: 1.6em;
      max-width: 100%;
      overflow: auto;
      padding: 1em 1.5em;
      display: block;
      word-wrap: break-word;
      text-align: initial;
    }
  }
</style>
