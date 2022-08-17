<template>
  <div class="center-content">
    <div v-if="redirectUrl != ''">
      <h1>Redirecting you to</h1>
      <h2> {{ redirectUrl }}</h2>
      <h3>{{ counter }} seconds...</h3>
      <h6>Press Ctrl+C to stop</h6>
    </div>
    <spinner v-if="showSpinner"></spinner>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component'
import axios from 'axios'
import spinner from '../components/atom_spinner.vue'

@Options({
  components: {
    spinner
  }
})
export default class RedirectLoader extends Vue {
  showSpinner = true
  redirectUrl = ''
  counter = 5

  created () {
    const timeoutInterval: number[] = []
    window.addEventListener('keyup', (e) => {
      if (e.ctrlKey && e.keyCode === 67 && timeoutInterval.length !== 0) {
        for (const interval in timeoutInterval) {
          if (timeoutInterval[interval] !== undefined) {
            clearTimeout(timeoutInterval[interval])
          }
        }
      }
    })

    axios.request({
      method: 'HEAD',
      // @ts-expect-error: this.apiUrl is already defined in the vue prototype
      url: this.apiUrl() + this.$route.params.shorturl_id
    })
      .then((res) => {
        this.redirectUrl = res.headers['x-redirecturl']
        for (let i = 1; i <= 5; i++) {
          timeoutInterval.push(setTimeout(() => {
            this.counter--
            if (i === 5) {
              window.location.href = this.redirectUrl
            }
          }, i * 1000))
        }
      })
      .catch((error) => {
        console.log(error)
      })
      .finally(() => {
        this.showSpinner = false
      })
  }
}
</script>
