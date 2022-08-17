import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

const mainApp = createApp(App).use(router)

const apiUrl = () => {
  let url = 'http://localhost:9808/'
  if (process.env.NODE_ENV === 'production') {
    const regex = /^[^.]+\.([a-zA-Z0-9\-.]*\.[a-zA-Z]*)\/?/
    let match
    url = `https://api.${window.location.host}/`
    if ((match = regex.exec(window.location.host)) !== null) {
      url = 'https://api.' + match[1] + '/'
    }
  }

  return url
}

mainApp.config.globalProperties.apiUrl = apiUrl

mainApp.mount('#app')
