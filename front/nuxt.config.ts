// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  modules: [
    'nuxt-graphql-client',
  ],
  runtimeConfig: {
    public: {
      GQL_HOST: 'http://localhost:5050/query'
    }
  }
})
