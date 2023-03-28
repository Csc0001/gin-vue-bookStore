import storageService from '@/service/storageService';
import axios from 'axios';

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL,
  timeout: 1000,
  headers: { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` },
});

// Add a request interceptor
service.interceptors.request.use((config) => {
  Object.assign(config.headers, { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` });
  return config;
}, (error) => Promise.reject(error));
export default service;
