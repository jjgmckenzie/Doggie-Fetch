/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    appDir: true,
  },
  images:{
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
