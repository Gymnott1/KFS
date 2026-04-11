<template>
  <div class="min-h-screen bg-gray-100 flex flex-col items-center justify-center gap-6">
    <div class="bg-white rounded-xl shadow-md p-8 w-full max-w-sm text-center">
      <h1 class="text-3xl font-bold text-gray-800 mb-2">Welcome 👋</h1>
      <p class="text-gray-500 mb-6">You are signed in.</p>

      <div v-if="auth.user" class="text-left space-y-2 mb-6">
        <p class="text-sm text-gray-600"><span class="font-semibold text-gray-800">Name:</span> {{ auth.user.name }}</p>
        <p class="text-sm text-gray-600"><span class="font-semibold text-gray-800">Email:</span> {{ auth.user.email }}</p>
        <p class="text-sm text-gray-600"><span class="font-semibold text-gray-800">Role:</span>
          <span class="ml-1 px-2 py-0.5 rounded-full text-xs font-medium"
            :class="auth.user.role === 'admin' ? 'bg-indigo-100 text-indigo-700' : 'bg-gray-100 text-gray-600'">
            {{ auth.user.role }}
          </span>
        </p>
      </div>

      <button @click="handleLogout"
        class="w-full border border-indigo-600 text-indigo-600 rounded-lg py-2 text-sm font-semibold hover:bg-indigo-50">
        Logout
      </button>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const auth = useAuthStore()
const router = useRouter()

onMounted(async () => {
  if (!auth.isLoggedIn()) {
    router.push('/login')
    return
  }
  if (!auth.user) await auth.fetchMe()
})

async function handleLogout() {
  await auth.logout()
  router.push('/login')
}
</script>
