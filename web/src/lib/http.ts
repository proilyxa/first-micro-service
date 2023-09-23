import axios from 'axios'

export const setAuthToken = (token: string): void => {
  localStorage.setItem('token', token)
  axios.defaults.headers.common['Authorization'] = 'Bearer ' + token
}

let token = localStorage.getItem('token')
if (token) {
  setAuthToken(token)
}

axios.defaults.headers.common['Content-Type'] = 'application/json'
axios.defaults.headers.common['Accept'] = 'application/json'

axios.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response.status === 401) {
      window.location.replace('/register')
      localStorage.removeItem('token')
      return
    }

    return Promise.reject(error)
  }
)

axios.defaults.baseURL = 'http://127.0.0.1:8082'
