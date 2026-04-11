import { defineStore } from 'pinia'
import { ref } from 'vue'
import { api } from '@/api'

export const useAuthStore = defineStore('auth', () => {
  const user = ref(null)
  const token = ref(localStorage.getItem('token'))

  async function login(email, password) {
    const data = await api.post('/auth/login', { email, password })
    token.value = data.token
    localStorage.setItem('token', data.token)
    await fetchMe()
  }

  async function fetchMe() {
    user.value = await api.get('/me')
  }

  async function logout() {
    try { await api.post('/auth/logout') } catch (_) {}
    token.value = null
    user.value = null
    localStorage.removeItem('token')
  }

  function isLoggedIn() {
    return !!token.value
  }

  return { user, token, login, logout, fetchMe, isLoggedIn }
})
