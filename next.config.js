/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    appDir: true,
  },
  images:{
    formats: ['image/avif', 'image/webp'],
    remotePatterns:[
      {
        protocol: 'https',
        hostname: 'images.dog.ceo',
        port:'',
        pathname:'/**'
      }
    ]
  }
}

module.exports = nextConfig
