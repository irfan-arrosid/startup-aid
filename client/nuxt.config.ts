// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  pages: true,
  devtools: { enabled: true },
  css: ['~/assets/css/main.css'],
  app: {
    head: {
      title: 'StartUp Aid',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { hid: 'description', name: 'description', content: ''},
      ],
      link: [
        {
          rel: 'stylesheet',
          href: 'https://fonts.googleapis.com/css2?family=Poppins:wght@100;200;300;400;500;600;700;800;900&display=swap',
        },
      ],
    },
  },
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  modules: ['@sidebase/nuxt-auth'],
  auth: {
    baseURL: 'http://localhost:8080/api/v1',
    provider: {
      type: 'local',
      endpoints: {
        signIn: { path: '/sessions', method: 'post' },
        // signOut: false,
        signUp: { path: '/users', method: 'post' },
        getSession: { path: '/users/fetch', method: 'get' }
      }
    },
  }
})