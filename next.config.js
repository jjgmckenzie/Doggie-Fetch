/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    appDir: true,
    outputStandalone:true,
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
  },
  async rewrites() {
      return [
        {
          source: '/upload',
          destination: 'http://localhost:8080/upload'
        }
      ]
    }
}

module.exports = nextConfig
