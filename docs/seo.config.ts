import type { Metadata } from 'next'

const defineMetadata = <T extends Metadata>(metadata: T) => metadata

const seoConfig = defineMetadata({
  metadataBase: new URL('https://dev.arc3dia.com'),
  title: {
    template: '%s - Arc3dia | For Developers',
    default:
      'Arc3dia | For Developers - Add cross-platform multiplayer to your game, fast'
  },
  description: 'Add cross-platform multiplayer to your game, fast',
  themeColor: '#000',
  icons: [
    { rel: 'icon', url: '/favicon.ico' },
    { rel: 'mask-icon', url: '/favicon.ico' },
    { rel: 'image/x-icon', url: '/favicon.ico' }
  ],
  twitter: {
    site: '@arc3dia',
    creator: '@wanjohiryan'
  }
})

export default seoConfig