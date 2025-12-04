// https://nuxt.com/docs/api/configuration/nuxt-config
type MaybeProcessEnv = Record<string, string | undefined>

declare const defineNuxtConfig: (config: Record<string, unknown>) => unknown

const env =
  (globalThis as typeof globalThis & { process?: { env?: MaybeProcessEnv } })
    .process?.env ?? {}

const serverHost = env.NUXT_GQL_HOST ?? 'http://app:5050/query'
const browserHost = env.NUXT_PUBLIC_GQL_HOST ?? 'http://localhost:5050/query'

export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  modules: [
    'nuxt-graphql-client',
  ],
  css: [
    '~/assets/css/reset.css'
  ],
  'graphql-client': {
    tokenStorage: {
      mode: 'localStorage'
    },
    clients: {
      default: {
        host: serverHost,
      }
    }
  },
  runtimeConfig: {
    gqlHostServer: serverHost,
    public: {
      GQL_HOST: browserHost,
    }
  }
})
