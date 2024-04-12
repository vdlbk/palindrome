// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    // @ts-ignore
    nitro: {
        preset: 'vercel'
    },
    devtools: { enabled: true },
    modules: ['@nuxt/ui', "@nuxt/image"],
    // @ts-ignore
    app: {
        head: {
            title: 'Palindromix',
            charset: 'utf-8',
            viewport: 'width=device-width, initial-scale=1.0',
            link: [
                { rel: 'icon', sizes: 'any', href: "/favicon.ico" }
            ],
            meta: [
                { name: 'description', content: 'Palindromix' },
                { name: 'theme-color', content: '#e0cb13' }
            ],
        }
    },
    css: [
      '~/assets/css/styles.css'
    ]
})