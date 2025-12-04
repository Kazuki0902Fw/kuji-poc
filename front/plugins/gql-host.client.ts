const DEFAULT_BROWSER_HOST = 'http://localhost:5050/query'

let didOverride = false

export default defineNuxtPlugin(() => {
  const nuxtApp = useNuxtApp()

  nuxtApp.hook('app:mounted', () => {
    if (didOverride) {
      return
    }

    try {
      const config = useRuntimeConfig()
      const client = useGqlHost('default')
      if (!client || typeof client.setHost !== 'function') {
        return
      }
      const host = config.public.GQL_HOST ?? DEFAULT_BROWSER_HOST
      if (typeof host === 'string' && host.length > 0) {
        client.setHost(host)
        didOverride = true
      }
    } catch (error) {
      console.error('failed to set gql host', error)
    }
  })
})

