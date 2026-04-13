import client from './client'
import type { Project, Bucket } from '@/types/models'

export const projectsAPI = {
  getAll: () => {
    return client.get<Project[]>('/projects')
  },

  getById: (id: number) => {
    return client.get<Project>(`/projects/${id}`)
  },

  create: (project: Partial<Project>) => {
    return client.post<Project>('/projects', project)
  },

  update: (id: number, project: Partial<Project>) => {
    return client.put<Project>(`/projects/${id}`, project)
  },

  delete: (id: number) => {
    return client.delete(`/projects/${id}`)
  },

  getBuckets: (projectId: number) => {
    return client.get<Bucket[]>('/buckets', { params: { project_id: projectId } })
  },

  createBucket: (bucket: Partial<Bucket>) => {
    return client.post<Bucket>('/buckets', bucket)
  },

  updateBucket: (id: number, bucket: Partial<Bucket>) => {
    return client.put<Bucket>(`/buckets/${id}`, bucket)
  },

  deleteBucket: (id: number) => {
    return client.delete(`/buckets/${id}`)
  }
}
