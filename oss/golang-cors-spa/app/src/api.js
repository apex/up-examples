
import axios from 'axios'

/**
 * API endpoint.
 */

const url = process.env.API_URL || 'http://localhost:3000'

/**
 * API client.
 */

const api = axios.create({
  baseURL: url,
  timeout: 5000
})

/**
 * Get pets.
 */

export function pets() {
  return api.get('/pets/')
}
