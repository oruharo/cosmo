import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react()
  ],
  server: {
    host: '0.0.0.0',
    port: 3010,
    proxy: {
      '/auth': {
        target: 'http://localhost:8450',
      },
      '/cosmo-auth-server/authserver.v1alpha1': {
        target: 'http://localhost:8450',
      },
    },
  }
})
