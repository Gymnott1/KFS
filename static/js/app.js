// ── Modal ────────────────────────────────────────────────
const modal = document.getElementById('modal')
const modalContent = document.getElementById('modal-content')

function openModal(html) {
  modalContent.innerHTML = html
  modal.classList.remove('hidden')
}
function closeModal() {
  modal.classList.add('hidden')
  modalContent.innerHTML = ''
}
document.getElementById('modal-close').addEventListener('click', closeModal)
modal.addEventListener('click', e => { if (e.target === modal) closeModal() })

// ── API helper ───────────────────────────────────────────
async function api(method, path, body) {
  const res = await fetch(`/api/v1${path}`, {
    method,
    headers: { 'Content-Type': 'application/json' },
    body: body ? JSON.stringify(body) : undefined,
    credentials: 'include',
  })
  const data = await res.json()
  if (!res.ok) throw new Error(data.error || 'Request failed')
  return data
}

// ── Auth helpers ─────────────────────────────────────────
function getToken() { return localStorage.getItem('token') }
function setToken(t) { localStorage.setItem('token', t) }
function clearToken() { localStorage.removeItem('token') }

async function logout() {
  try { await api('POST', '/auth/logout') } catch (_) {}
  clearToken()
  window.location.href = '/login'
}

// ── Nav: show logout if logged in ────────────────────────
const navAuth = document.getElementById('nav-auth')
if (navAuth) {
  if (getToken()) {
    navAuth.innerHTML = `<a href="#" onclick="logout();return false;">Logout</a>`
  } else {
    navAuth.innerHTML = `<a href="/login">Login</a>`
  }
}

// ── Login page ───────────────────────────────────────────
const loginForm = document.getElementById('login-form')
if (loginForm) {
  // Redirect if already logged in
  if (getToken()) window.location.href = '/'

  loginForm.addEventListener('submit', async e => {
    e.preventDefault()
    const errEl = document.getElementById('login-error')
    errEl.classList.add('hidden')

    const email = document.getElementById('email').value
    const password = document.getElementById('password').value

    try {
      const data = await api('POST', '/auth/login', { email, password })
      setToken(data.token)
      window.location.href = '/'
    } catch (err) {
      errEl.textContent = err.message
      errEl.classList.remove('hidden')
    }
  })
}

// ── Dashboard page ───────────────────────────────────────
const userInfo = document.getElementById('user-info')
if (userInfo) {
  if (!getToken()) {
    window.location.href = '/login'
  } else {
    api('GET', '/me').then(u => {
      userInfo.innerHTML = `
        <p><strong>Name:</strong> ${u.name}</p>
        <p><strong>Email:</strong> ${u.email}</p>
        <p><strong>Role:</strong> ${u.role}</p>
      `
    }).catch(() => {
      clearToken()
      window.location.href = '/login'
    })
  }
}
