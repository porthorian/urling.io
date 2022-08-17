import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

const mainApp = createApp(App).use(router)

let apiUrl = 'http://localhost:9808/'
if (process.env.NODE_ENV === 'production') {
  const regex = /^[^.]+\.([a-zA-Z0-9\-.]*)\/?/
  let match

  if ((match = regex.exec(window.location.host)) !== null) {
    apiUrl = 'https://api.' + match[1] + '/'
  }
}

mainApp.config.globalProperties.apiUrl = apiUrl

mainApp.mount('#app')
