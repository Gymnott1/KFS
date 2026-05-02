<template>
  <main class="theme-shell min-h-screen" :class="theme">
    <div class="relative mx-auto flex min-h-screen w-full max-w-[1600px] flex-col lg:flex-row">
      <aside
        class="sidebar left-sidebar fixed inset-y-0 left-0 z-[900] w-[332px] max-w-[86vw] border-r lg:static lg:z-auto lg:max-w-none"
        :class="leftSidebarOpen ? 'translate-x-0' : '-translate-x-full lg:hidden'"
      >
        <div class="flex h-full flex-col">
          <div class="panel-header flex items-start justify-between gap-3 border-b px-5 py-5">
            <div>
              <p class="eyebrow">Kenya From Space</p>
              <h1 class="mt-1 text-2xl font-black">Map OS</h1>
              <p class="mt-2 text-sm font-semibold leading-5 muted">
                Satellite-first Kenya intelligence, grounded by live OpenStreetMap data.
              </p>
            </div>
            <button class="icon-button" @click="leftSidebarOpen = false" title="Hide layers sidebar">L</button>
          </div>

          <div class="space-y-4 overflow-y-auto px-4 py-4">
            <section class="surface rounded-lg border p-3">
              <label class="eyebrow muted">Search loaded data</label>
              <div class="field mt-2 flex items-center gap-2 rounded-md border px-3 py-2">
                <span class="text-sm font-black accent-text">Q</span>
                <input
                  v-model="search"
                  class="w-full bg-transparent text-sm font-semibold outline-none"
                  placeholder="KICC, hospital, route, Kibera"
                />
              </div>
            </section>

            <section class="surface rounded-lg border p-3">
              <div class="flex items-center justify-between">
                <h2 class="text-sm font-black uppercase tracking-[0.12em]">System Theme</h2>
                <span class="text-xs font-bold muted">{{ themeLabel }}</span>
              </div>
              <div class="mt-3 grid grid-cols-2 gap-2">
                <button
                  v-for="option in themeOptions"
                  :key="option.id"
                  class="segmented"
                  :class="theme === option.id ? 'is-active' : ''"
                  @click="theme = option.id"
                >
                  {{ option.label }}
                </button>
              </div>
            </section>

            <section class="surface rounded-lg border p-3">
              <div class="flex items-center justify-between">
                <h2 class="text-sm font-black uppercase tracking-[0.12em]">Map View</h2>
                <span class="text-xs font-bold muted">{{ mapViewLabel }}</span>
              </div>
              <div class="mt-3 grid grid-cols-2 gap-2">
                <button
                  v-for="view in mapViews"
                  :key="view.id"
                  class="segmented"
                  :class="baseMap === view.id ? 'is-active' : ''"
                  @click="baseMap = view.id"
                >
                  {{ view.label }}
                </button>
              </div>
            </section>

            <section>
              <div class="flex items-center justify-between">
                <h2 class="text-sm font-black uppercase tracking-[0.12em]">Categories</h2>
                <span class="text-xs font-bold muted">{{ activeLayers.length }} active</span>
              </div>
              <div class="mt-3 space-y-2">
                <button
                  v-for="layer in layers"
                  :key="layer.id"
                  @click="toggleLayer(layer.id)"
                  class="layer-row flex w-full items-center gap-3 rounded-lg border px-3 py-3 text-left transition"
                  :class="layer.enabled ? 'is-enabled' : ''"
                >
                  <span
                    class="grid h-9 w-9 shrink-0 place-items-center rounded-md text-xs font-black text-white"
                    :style="{ backgroundColor: layer.color }"
                  >
                    {{ layer.code }}
                  </span>
                  <span class="min-w-0 flex-1">
                    <span class="block text-sm font-black">{{ layer.name }}</span>
                    <span class="block truncate text-xs font-semibold muted">{{ layer.detail }}</span>
                  </span>
                  <span class="switch" :class="layer.enabled ? 'is-on' : ''">
                    <span />
                  </span>
                </button>
              </div>
            </section>

            <section class="space-panel rounded-lg border p-4">
              <p class="eyebrow">Live Data</p>
              <div class="mt-3 space-y-3 text-sm font-semibold">
                <p>Base maps: OSM, Esri satellite, CARTO dark, OpenTopoMap.</p>
                <p>Live pull: Overpass OSM amenities in Nairobi.</p>
                <p>Product data: real coordinates, ready for backend feeds.</p>
              </div>
              <button class="primary-button mt-4 w-full" @click="loadOsmAmenities" :disabled="osmLoading">
                {{ osmLoading ? 'Loading OSM...' : 'Refresh OSM Data' }}
              </button>
            </section>
          </div>
        </div>
      </aside>

      <section class="flex min-w-0 flex-1 flex-col">
        <header class="app-header z-[700] border-b px-3 py-3 backdrop-blur lg:px-5">
          <div class="flex items-center justify-between gap-2">
            <div class="flex min-w-0 items-center gap-2">
              <button class="icon-button shrink-0" @click="toggleLeftSidebar" title="Toggle layers sidebar">L</button>
              <div class="min-w-0">
                <p class="eyebrow">All-in-one Kenya map platform</p>
                <h1 class="truncate text-xl font-black sm:text-2xl">KENYA FROM SPACE</h1>
              </div>
            </div>
            <div class="flex shrink-0 items-center gap-2">
              <button class="header-button hidden sm:block" @click="toggleTheme">{{ themeToggleLabel }}</button>
              <button class="header-button hidden md:block" @click="cycleMapView">{{ mapViewLabel }}</button>
              <button class="danger-button" @click="isAlertOpen = true">SOS</button>
              <button class="icon-button" @click="toggleRightSidebar" title="Toggle data sidebar">R</button>
            </div>
          </div>

          <nav class="mt-3 flex gap-2 overflow-x-auto pb-1 lg:hidden">
            <button
              v-for="tab in mobileTabs"
              :key="tab.id"
              @click="activeTab = tab.id"
              class="tab-button shrink-0 rounded-full border px-4 py-2 text-sm font-black"
              :class="activeTab === tab.id ? 'is-active' : ''"
            >
              {{ tab.label }}
            </button>
          </nav>
        </header>

        <div class="grid flex-1 grid-cols-1" :class="rightSidebarOpen ? 'lg:grid-cols-[minmax(0,1fr)_390px]' : 'lg:grid-cols-1'">
          <section
            class="relative min-h-[calc(100vh-120px)] lg:min-h-0"
            :class="activeTab !== 'map' ? 'hidden lg:block' : ''"
          >
            <div ref="mapEl" class="absolute inset-0 z-0" />

            <div class="pointer-events-none absolute left-3 right-3 top-3 z-[500] grid gap-3 sm:left-5 sm:right-auto sm:w-[420px]">
              <section class="glass-panel pointer-events-auto rounded-lg border p-3 shadow-sm">
                <div class="flex items-start justify-between gap-3">
                  <div>
                    <p class="eyebrow muted">Map Focus</p>
                    <h2 class="text-lg font-black">{{ selectedPin.title }}</h2>
                  </div>
                  <span class="rounded-full px-3 py-1 text-xs font-black text-white" :style="{ backgroundColor: selectedLayerColor }">
                    {{ selectedPin.category }}
                  </span>
                </div>
                <p class="mt-2 text-sm font-semibold leading-5 muted-strong">{{ selectedPin.body }}</p>
                <div class="mt-3 grid grid-cols-3 gap-2 text-center">
                  <div class="metric rounded-md p-2">
                    <p class="text-lg font-black">{{ visiblePins.length }}</p>
                    <p class="text-[11px] font-bold muted">Visible</p>
                  </div>
                  <div class="metric rounded-md p-2">
                    <p class="text-lg font-black">{{ osmAmenities.length }}</p>
                    <p class="text-[11px] font-bold muted">OSM live</p>
                  </div>
                  <div class="metric rounded-md p-2">
                    <p class="text-lg font-black">{{ selectedPin.score }}</p>
                    <p class="text-[11px] font-bold muted">Score</p>
                  </div>
                </div>
              </section>
            </div>

            <div class="absolute bottom-4 left-3 right-3 z-[500] sm:left-5 sm:right-auto sm:w-[450px]">
              <article class="glass-panel rounded-lg border p-4 shadow-xl">
                <div class="flex items-start justify-between gap-3">
                  <div>
                    <p class="eyebrow accent-text">{{ selectedPin.source }}</p>
                    <h2 class="mt-1 text-xl font-black">{{ selectedPin.title }}</h2>
                  </div>
                  <span class="pill">
                    {{ selectedPin.lat.toFixed(3) }}, {{ selectedPin.lng.toFixed(3) }}
                  </span>
                </div>
                <p class="mt-2 text-sm font-semibold leading-6 muted-strong">{{ selectedPin.body }}</p>
                <div class="mt-4 flex flex-wrap gap-2">
                  <button class="primary-button">Open details</button>
                  <button class="secondary-button">Directions</button>
                  <button class="secondary-button">Share</button>
                </div>
              </article>
            </div>
          </section>

          <aside
            class="sidebar right-sidebar border-l lg:block"
            :class="[
              activeTab === 'map' ? 'hidden lg:block' : '',
              rightSidebarOpen ? '' : 'lg:hidden',
            ]"
          >
            <div class="flex items-center justify-between border-b px-4 py-3">
              <div>
                <p class="eyebrow muted">Intelligence Panel</p>
                <h2 class="text-lg font-black">Data & Layers</h2>
              </div>
              <button class="icon-button" @click="rightSidebarOpen = false" title="Hide data sidebar">R</button>
            </div>

            <div class="h-full overflow-y-auto px-4 py-4 lg:max-h-[calc(100vh-126px)]">
              <section v-if="activeTab === 'feed' || desktopPanel" class="space-y-3">
                <div class="flex items-center justify-between">
                  <h2 class="text-lg font-black">Real Map Data</h2>
                  <span class="text-xs font-bold muted">Nairobi</span>
                </div>
                <button
                  v-for="pin in filteredPins"
                  :key="pin.id"
                  @click="focusPin(pin)"
                  class="data-card w-full rounded-lg border p-3 text-left transition"
                >
                  <div class="flex items-start gap-3">
                    <span
                      class="grid h-10 w-10 shrink-0 place-items-center rounded-md text-xs font-black text-white"
                      :style="{ backgroundColor: categoryColor(pin.category) }"
                    >
                      {{ categoryCode(pin.category) }}
                    </span>
                    <span class="min-w-0 flex-1">
                      <span class="flex items-center justify-between gap-2">
                        <span class="truncate text-sm font-black">{{ pin.title }}</span>
                        <span class="shrink-0 text-xs font-bold muted">{{ pin.source }}</span>
                      </span>
                      <span class="mt-1 block text-sm font-semibold leading-5 muted-strong">{{ pin.body }}</span>
                      <span class="mt-2 block text-xs font-black muted">{{ pin.lat }}, {{ pin.lng }}</span>
                    </span>
                  </div>
                </button>
              </section>

              <section class="mt-5 space-y-3" :class="activeTab !== 'layers' ? 'lg:block hidden' : ''">
                <div class="flex items-center justify-between">
                  <h2 class="text-lg font-black">Layer Strategy</h2>
                  <span class="pill">MVP</span>
                </div>
                <article class="data-card rounded-lg border p-3">
                  <h3 class="text-sm font-black">Do not launch all 17 categories</h3>
                  <p class="mt-1 text-sm font-semibold leading-6 muted-strong">
                    Start with Transport, Safety, Business, and Community. Keep Agriculture, Real Estate, Climate, and B2B APIs as expansion modules.
                  </p>
                </article>
                <article v-for="module in roadmap" :key="module.name" class="data-card rounded-lg border p-3">
                  <div class="flex items-center justify-between">
                    <h3 class="text-sm font-black">{{ module.name }}</h3>
                    <span class="text-xs font-black accent-text">{{ module.phase }}</span>
                  </div>
                  <p class="mt-1 text-sm font-semibold leading-5 muted-strong">{{ module.detail }}</p>
                </article>
              </section>

              <section class="mt-5 grid grid-cols-2 gap-3" :class="activeTab !== 'ops' ? 'lg:grid hidden' : ''">
                <div class="space-panel col-span-2 rounded-lg border p-4">
                  <p class="eyebrow">Operations</p>
                  <div class="mt-2 flex items-end justify-between">
                    <h2 class="text-4xl font-black">5</h2>
                    <span class="pill">map views</span>
                  </div>
                </div>
                <div class="data-card rounded-lg border p-3">
                  <p class="text-2xl font-black">OSM</p>
                  <p class="text-xs font-bold muted">Base truth</p>
                </div>
                <div class="data-card rounded-lg border p-3">
                  <p class="text-2xl font-black">API</p>
                  <p class="text-xs font-bold muted">Ready model</p>
                </div>
                <div class="data-card col-span-2 rounded-lg border p-3">
                  <h2 class="text-sm font-black">Build Priorities</h2>
                  <div class="mt-3 space-y-3">
                    <ProgressRow label="Real map engine" value="100" />
                    <ProgressRow label="Theme system" value="100" />
                    <ProgressRow label="Community reports backend" value="20" />
                  </div>
                </div>
              </section>
            </div>
          </aside>
        </div>
      </section>
    </div>

    <div
      v-if="isAlertOpen"
      class="fixed inset-0 z-[1000] grid place-items-end bg-black/55 p-3 sm:place-items-center"
      @click.self="isAlertOpen = false"
    >
      <section class="modal w-full max-w-md rounded-lg border p-4 shadow-2xl">
        <div class="flex items-center justify-between">
          <div>
            <p class="eyebrow text-[#ef4444]">Emergency Broadcast</p>
            <h2 class="text-xl font-black">Alert nearby users</h2>
          </div>
          <button class="icon-button" @click="isAlertOpen = false" title="Close alert dialog">X</button>
        </div>
        <div class="mt-4 grid grid-cols-2 gap-2">
          <button v-for="type in alertTypes" :key="type" class="secondary-button py-3 text-sm">
            {{ type }}
          </button>
        </div>
        <textarea
          class="field mt-3 h-24 w-full rounded-md border p-3 text-sm font-semibold outline-none"
          placeholder="Describe what is happening and exact location."
        />
        <button class="danger-button mt-3 w-full px-4 py-3">Send to selected radius</button>
      </section>
    </div>
  </main>
</template>

<script setup>
import L from 'leaflet'
import { computed, defineComponent, h, nextTick, onBeforeUnmount, onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const auth = useAuthStore()
const router = useRouter()

const mapEl = ref(null)
const map = ref(null)
const markerLayer = ref(null)
const osmAmenities = ref([])
const osmLoading = ref(false)
const search = ref('')
const activeTab = ref('map')
const baseMap = ref('satellite')
const theme = ref('dark')
const leftSidebarOpen = ref(true)
const rightSidebarOpen = ref(true)
const isAlertOpen = ref(false)

let activeBaseLayers = []

const themeOptions = [
  { id: 'dark', label: 'Dark' },
  { id: 'light', label: 'Light' },
]

const mapViews = [
  { id: 'street', label: 'Street' },
  { id: 'satellite', label: 'Space' },
  { id: 'hybrid', label: 'Hybrid' },
  { id: 'dark', label: 'Dark Map' },
  { id: 'terrain', label: 'Terrain' },
]

const mobileTabs = [
  { id: 'map', label: 'Map' },
  { id: 'feed', label: 'Data' },
  { id: 'layers', label: 'Layers' },
  { id: 'ops', label: 'Ops' },
]

const layers = ref([
  { id: 'transport', name: 'Transport', detail: 'Matatus, routes, stages, disruptions', code: 'TR', color: '#3b82f6', enabled: true },
  { id: 'safety', name: 'Safety', detail: 'Police, hospitals, incidents, SOS', code: 'SF', color: '#ef4444', enabled: true },
  { id: 'business', name: 'Business', detail: 'Markets, malls, services, fuel', code: 'BZ', color: '#f59e0b', enabled: true },
  { id: 'community', name: 'Community', detail: 'Events, chats, user reports', code: 'CM', color: '#8b5cf6', enabled: true },
  { id: 'climate', name: 'Climate', detail: 'Flood risk, rainfall, heat, air', code: 'CL', color: '#06b6d4', enabled: false },
  { id: 'agriculture', name: 'Agriculture', detail: 'Crop health, markets, buyers', code: 'AG', color: '#a16207', enabled: false },
  { id: 'real-estate', name: 'Real Estate', detail: 'Plots, zoning, growth trends', code: 'RE', color: '#ec4899', enabled: false },
  { id: 'infrastructure', name: 'Infrastructure', detail: 'Roads, rail, projects, utilities', code: 'IN', color: '#64748b', enabled: false },
])

const curatedPins = [
  { id: 'kicc', category: 'business', title: 'KICC, Nairobi CBD', body: 'Major landmark and event venue. Good anchor point for city-scale discovery and tourism layers.', source: 'Curated real coordinate', lat: -1.2886, lng: 36.8231, score: 98 },
  { id: 'railways', category: 'transport', title: 'Nairobi Railway Station', body: 'Central rail and matatu movement node. Useful for fare, route, and crowd reporting.', source: 'Curated real coordinate', lat: -1.2921, lng: 36.8294, score: 94 },
  { id: 'knh', category: 'safety', title: 'Kenyatta National Hospital', body: 'Emergency services anchor for hospital discovery and SOS routing.', source: 'Curated real coordinate', lat: -1.3015, lng: 36.8063, score: 97 },
  { id: 'central-police', category: 'safety', title: 'Central Police Station', body: 'Police station layer item for safety routing and emergency awareness.', source: 'Curated real coordinate', lat: -1.2865, lng: 36.8247, score: 92 },
  { id: 'gikomba', category: 'business', title: 'Gikomba Market', body: 'High-value informal economy zone for vendors, logistics, and local price intelligence.', source: 'Curated real coordinate', lat: -1.2834, lng: 36.8427, score: 90 },
  { id: 'kibera', category: 'community', title: 'Kibera Community Mapping Zone', body: 'Priority zone for informal settlement addressing, water points, clinics, and NGO resource mapping.', source: 'Curated real coordinate', lat: -1.3133, lng: 36.7892, score: 95 },
  { id: 'mathare', category: 'community', title: 'Mathare Mapping Zone', body: 'Informal addressing and essential services mapping candidate for community mapper programs.', source: 'Curated real coordinate', lat: -1.2582, lng: 36.8587, score: 88 },
  { id: 'jkia', category: 'logistics', title: 'Jomo Kenyatta International Airport', body: 'National logistics and tourism gateway. Future fit for delivery, cargo, and mobility intelligence.', source: 'Curated real coordinate', lat: -1.3192, lng: 36.9278, score: 96 },
  { id: 'ngong-river', category: 'climate', title: 'Ngong River flood watch corridor', body: 'Climate layer candidate for flood-prone zone reporting and rainfall-triggered alerts.', source: 'Curated real coordinate', lat: -1.316, lng: 36.786, score: 74 },
  { id: 'karura', category: 'climate', title: 'Karura Forest environment zone', body: 'Environmental monitoring candidate for air quality, heat, recreation, and conservation alerts.', source: 'Curated real coordinate', lat: -1.2345, lng: 36.8338, score: 85 },
  { id: 'limuru', category: 'agriculture', title: 'Limuru agriculture edge', body: 'Expansion module for crop health, market access, and weather-based farmer advice.', source: 'Curated real coordinate', lat: -1.1058, lng: 36.6424, score: 82 },
  { id: 'tatu-city', category: 'real-estate', title: 'Tatu City growth zone', body: 'Real estate intelligence candidate for growth trends, pricing, zoning, and nearby development tracking.', source: 'Curated real coordinate', lat: -1.1489, lng: 36.964, score: 86 },
]

const selectedPin = ref(curatedPins[0])

const roadmap = [
  { name: 'Phase 1: Transport + Safety', phase: 'Now', detail: 'Matatu stages, hospitals, police, SOS alerts, and incident reporting.' },
  { name: 'Phase 2: Business + Community', phase: 'Next', detail: 'Verified shops, fundis, zone chats, events, and user contributions.' },
  { name: 'Phase 3: Climate + Agriculture', phase: 'Later', detail: 'Flood risk, weather advice, crop health, prices, and buyers.' },
  { name: 'Phase 4: Real Estate + APIs', phase: 'B2B', detail: 'Growth modeling, listings, risk scoring, and paid data access.' },
]

const alertTypes = ['Danger', 'Medical', 'Accident', 'Flood', 'Protest', 'Fire']

const activeLayers = computed(() => layers.value.filter((layer) => layer.enabled))
const enabledLayerIds = computed(() => new Set(activeLayers.value.map((layer) => layer.id)))
const allPins = computed(() => [...curatedPins, ...osmAmenities.value])
const visiblePins = computed(() => allPins.value.filter((pin) => enabledLayerIds.value.has(pin.category)))
const filteredPins = computed(() => {
  const term = search.value.trim().toLowerCase()
  if (!term) return visiblePins.value
  return visiblePins.value.filter((pin) => `${pin.title} ${pin.body} ${pin.category}`.toLowerCase().includes(term))
})
const desktopPanel = computed(() => activeTab.value === 'map')
const themeLabel = computed(() => (theme.value === 'dark' ? 'Dark system' : 'Light system'))
const themeToggleLabel = computed(() => (theme.value === 'dark' ? 'Light' : 'Dark'))
const mapViewLabel = computed(() => mapViews.find((view) => view.id === baseMap.value)?.label || 'Map')
const selectedLayerColor = computed(() => categoryColor(selectedPin.value.category))

const ProgressRow = defineComponent({
  props: {
    label: { type: String, required: true },
    value: { type: String, required: true },
  },
  setup(props) {
    return () =>
      h('div', [
        h('div', { class: 'flex items-center justify-between text-xs font-black' }, [
          h('span', props.label),
          h('span', `${props.value}%`),
        ]),
        h('div', { class: 'progress-track mt-1 h-2 rounded-full' }, [
          h('div', {
            class: 'h-2 rounded-full bg-[#f59e0b]',
            style: { width: `${props.value}%` },
          }),
        ]),
      ])
  },
})

onMounted(async () => {
  if (auth.isLoggedIn() && !auth.user) {
    try {
      await auth.fetchMe()
    } catch (_) {
      await auth.logout()
    }
  }

  await nextTick()
  initMap()
  renderMarkers()
  loadOsmAmenities()
})

onBeforeUnmount(() => {
  if (map.value) {
    map.value.remove()
    map.value = null
  }
})

watch([visiblePins, search], () => renderMarkers(), { deep: true })
watch(baseMap, () => setBaseMap())
watch([leftSidebarOpen, rightSidebarOpen, activeTab], () => {
  setTimeout(() => map.value?.invalidateSize(), 180)
})

function initMap() {
  if (!mapEl.value || map.value) return

  map.value = L.map(mapEl.value, {
    zoomControl: false,
    attributionControl: true,
  }).setView([-1.2864, 36.8172], 12)

  L.control.zoom({ position: 'bottomright' }).addTo(map.value)
  markerLayer.value = L.layerGroup().addTo(map.value)
  setBaseMap()
}

function setBaseMap() {
  if (!map.value) return

  activeBaseLayers.forEach((layer) => layer.remove())
  activeBaseLayers = []

  const street = L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '&copy; OpenStreetMap contributors',
  })

  const satellite = L.tileLayer(
    'https://server.arcgisonline.com/ArcGIS/rest/services/World_Imagery/MapServer/tile/{z}/{y}/{x}',
    {
      maxZoom: 19,
      attribution: 'Tiles &copy; Esri',
    },
  )

  const labels = L.tileLayer('https://{s}.basemaps.cartocdn.com/light_only_labels/{z}/{x}/{y}{r}.png', {
    maxZoom: 19,
    attribution: '&copy; CARTO &copy; OpenStreetMap contributors',
  })

  const dark = L.tileLayer('https://{s}.basemaps.cartocdn.com/dark_all/{z}/{x}/{y}{r}.png', {
    maxZoom: 20,
    attribution: '&copy; CARTO &copy; OpenStreetMap contributors',
  })

  const terrain = L.tileLayer('https://{s}.tile.opentopomap.org/{z}/{x}/{y}.png', {
    maxZoom: 17,
    attribution: '&copy; OpenTopoMap &copy; OpenStreetMap contributors',
  })

  if (baseMap.value === 'satellite') activeBaseLayers = [satellite.addTo(map.value)]
  else if (baseMap.value === 'hybrid') activeBaseLayers = [satellite.addTo(map.value), labels.addTo(map.value)]
  else if (baseMap.value === 'dark') activeBaseLayers = [dark.addTo(map.value)]
  else if (baseMap.value === 'terrain') activeBaseLayers = [terrain.addTo(map.value)]
  else activeBaseLayers = [street.addTo(map.value)]
}

function renderMarkers() {
  if (!markerLayer.value) return
  markerLayer.value.clearLayers()

  filteredPins.value.forEach((pin) => {
    const marker = L.marker([pin.lat, pin.lng], {
      icon: markerIcon(pin),
      title: pin.title,
    })
      .bindPopup(`<strong>${escapeHtml(pin.title)}</strong><br>${escapeHtml(pin.category)}<br>${escapeHtml(pin.source)}`)
      .on('click', () => {
        selectedPin.value = pin
      })

    markerLayer.value.addLayer(marker)
  })
}

async function loadOsmAmenities() {
  if (osmLoading.value) return
  osmLoading.value = true

  const query = `
    [out:json][timeout:18];
    (
      node["amenity"~"hospital|police|fire_station|fuel"](-1.34,36.76,-1.23,36.91);
      way["amenity"~"hospital|police|fire_station|fuel"](-1.34,36.76,-1.23,36.91);
    );
    out center 45;
  `

  try {
    const response = await fetch('https://overpass-api.de/api/interpreter', {
      method: 'POST',
      body: new URLSearchParams({ data: query }),
    })
    const data = await response.json()

    osmAmenities.value = data.elements
      .map((item) => {
        const amenity = item.tags?.amenity
        const lat = item.lat ?? item.center?.lat
        const lng = item.lon ?? item.center?.lon
        if (!amenity || !lat || !lng) return null

        const category = amenity === 'fuel' ? 'business' : 'safety'
        return {
          id: `osm-${item.type}-${item.id}`,
          category,
          title: item.tags?.name || readableAmenity(amenity),
          body: `${readableAmenity(amenity)} from OpenStreetMap around Nairobi.`,
          source: 'OpenStreetMap live',
          lat,
          lng,
          score: 80,
        }
      })
      .filter(Boolean)
  } catch (_) {
    osmAmenities.value = []
  } finally {
    osmLoading.value = false
  }
}

function toggleLayer(layerId) {
  const layer = layers.value.find((item) => item.id === layerId)
  if (!layer) return
  layer.enabled = !layer.enabled

  if (!enabledLayerIds.value.has(selectedPin.value.category)) {
    selectedPin.value = visiblePins.value[0] || curatedPins[0]
  }
}

function focusPin(pin) {
  selectedPin.value = pin
  activeTab.value = 'map'
  if (map.value) map.value.setView([pin.lat, pin.lng], 15, { animate: true })
}

function toggleTheme() {
  theme.value = theme.value === 'dark' ? 'light' : 'dark'
}

function cycleMapView() {
  const currentIndex = mapViews.findIndex((view) => view.id === baseMap.value)
  baseMap.value = mapViews[(currentIndex + 1) % mapViews.length].id
}

function toggleLeftSidebar() {
  leftSidebarOpen.value = !leftSidebarOpen.value
}

function toggleRightSidebar() {
  rightSidebarOpen.value = !rightSidebarOpen.value
}

function categoryColor(category) {
  return layers.value.find((layer) => layer.id === category)?.color || '#111827'
}

function categoryCode(category) {
  return layers.value.find((layer) => layer.id === category)?.code || category.slice(0, 2).toUpperCase()
}

function markerIcon(pin) {
  return L.divIcon({
    className: 'mapake-marker-wrap',
    html: `<span class="mapake-marker" style="background:${categoryColor(pin.category)}">${categoryCode(pin.category)}</span>`,
    iconSize: [38, 38],
    iconAnchor: [19, 19],
    popupAnchor: [0, -18],
  })
}

function readableAmenity(value) {
  return value
    .split('_')
    .map((part) => part.charAt(0).toUpperCase() + part.slice(1))
    .join(' ')
}

function escapeHtml(value) {
  return String(value)
    .replaceAll('&', '&amp;')
    .replaceAll('<', '&lt;')
    .replaceAll('>', '&gt;')
    .replaceAll('"', '&quot;')
    .replaceAll("'", '&#039;')
}

async function handleAuthAction() {
  if (!auth.isLoggedIn()) {
    router.push('/login')
    return
  }

  await auth.logout()
}
</script>

<style scoped>
.theme-shell {
  --bg: #f4f1ea;
  --panel: #fbfaf6;
  --panel-strong: #ffffff;
  --border: #d7d1c6;
  --text: #171b22;
  --muted: #667085;
  --muted-strong: #4b5563;
  --field: #f7f3eb;
  --accent: #f59e0b;
  --accent-soft: #fff7ed;
  --space: #111827;
  background: var(--bg);
  color: var(--text);
}

.theme-shell.dark {
  --bg: #070a12;
  --panel: #0e1421;
  --panel-strong: #121a2a;
  --border: #263143;
  --text: #f8fafc;
  --muted: #94a3b8;
  --muted-strong: #cbd5e1;
  --field: #111827;
  --accent: #f59e0b;
  --accent-soft: rgba(245, 158, 11, 0.14);
  --space: #020617;
}

.sidebar,
.app-header,
.modal {
  background: var(--panel);
  border-color: var(--border);
  color: var(--text);
}

.panel-header,
.surface,
.data-card,
.metric,
.field,
.secondary-button,
.header-button,
.tab-button,
.icon-button {
  border-color: var(--border);
}

.surface,
.data-card,
.metric,
.modal {
  background: var(--panel-strong);
}

.field {
  background: var(--field);
}

.field input,
.field textarea,
textarea.field {
  color: var(--text);
}

.muted {
  color: var(--muted);
}

.muted-strong {
  color: var(--muted-strong);
}

.eyebrow {
  color: var(--accent);
  font-size: 0.72rem;
  font-weight: 900;
  letter-spacing: 0.16em;
  text-transform: uppercase;
}

.accent-text {
  color: var(--accent);
}

.glass-panel {
  background: color-mix(in srgb, var(--panel) 88%, transparent);
  border-color: var(--border);
  color: var(--text);
  backdrop-filter: blur(14px);
}

.space-panel {
  background: var(--space);
  border-color: var(--border);
  color: #f8fafc;
}

.space-panel .muted,
.space-panel p {
  color: #dbe4f0;
}

.primary-button,
.danger-button,
.secondary-button,
.header-button,
.segmented,
.tab-button,
.icon-button {
  border-radius: 0.375rem;
  font-weight: 900;
  transition: border-color 0.16s ease, background 0.16s ease, color 0.16s ease;
}

.primary-button {
  background: var(--accent);
  color: #111827;
  padding: 0.5rem 0.75rem;
}

.danger-button {
  background: #dc2626;
  color: white;
  padding: 0.5rem 0.9rem;
}

.secondary-button,
.header-button,
.segmented,
.tab-button,
.icon-button {
  background: var(--panel-strong);
  border: 1px solid var(--border);
  color: var(--text);
}

.secondary-button,
.header-button {
  padding: 0.5rem 0.75rem;
}

.segmented {
  padding: 0.65rem 0.7rem;
  text-align: center;
}

.segmented.is-active,
.tab-button.is-active {
  background: var(--text);
  border-color: var(--text);
  color: var(--bg);
}

.icon-button {
  display: grid;
  height: 2.35rem;
  place-items: center;
  width: 2.35rem;
}

.layer-row {
  background: var(--panel-strong);
  border-color: var(--border);
}

.layer-row.is-enabled {
  background: var(--accent-soft);
  border-color: var(--accent);
}

.switch {
  background: #64748b;
  border-radius: 999px;
  display: block;
  height: 1.25rem;
  padding: 0.125rem;
  width: 2.25rem;
}

.switch span {
  background: white;
  border-radius: 999px;
  display: block;
  height: 1rem;
  transition: transform 0.16s ease;
  width: 1rem;
}

.switch.is-on {
  background: var(--accent);
}

.switch.is-on span {
  transform: translateX(1rem);
}

.pill {
  background: var(--text);
  border-radius: 999px;
  color: var(--bg);
  display: inline-flex;
  font-size: 0.72rem;
  font-weight: 900;
  padding: 0.25rem 0.7rem;
  white-space: nowrap;
}

.data-card:hover {
  border-color: var(--accent);
}

.progress-track {
  background: color-mix(in srgb, var(--muted) 28%, transparent);
}

:deep(.leaflet-container) {
  height: 100%;
  width: 100%;
  background: #070a12;
  font-family: inherit;
}

:deep(.leaflet-control-attribution) {
  font-size: 10px;
}

:deep(.leaflet-popup-content-wrapper),
:deep(.leaflet-popup-tip) {
  background: var(--panel);
  color: var(--text);
}

:deep(.mapake-marker-wrap) {
  background: transparent;
  border: 0;
}

:deep(.mapake-marker) {
  align-items: center;
  border: 3px solid white;
  border-radius: 999px;
  box-shadow: 0 10px 24px rgba(0, 0, 0, 0.35);
  color: white;
  display: flex;
  font-size: 10px;
  font-weight: 900;
  height: 38px;
  justify-content: center;
  width: 38px;
}
</style>
